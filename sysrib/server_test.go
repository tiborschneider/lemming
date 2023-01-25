// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sysrib

import (
	"context"
	"fmt"
	"net"
	"net/netip"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gribigo/afthelper"
	"github.com/openconfig/lemming/gnmi"
	"github.com/openconfig/lemming/gnmi/gnmiclient"
	"github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/lemming/gnmi/oc/ocpath"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"google.golang.org/grpc"

	pb "github.com/openconfig/lemming/proto/sysrib"
)

const (
	// Each quantum is 100 ms
	maxGNMIWaitQuanta = 200 // 20s
)

type AddIntfAction struct {
	name    string
	ifindex int32
	enabled bool
	prefix  string
	niName  string
}

type SetRouteRequestAction struct {
	Desc     string
	RouteReq *pb.SetRouteRequest
}

type FakeDataplane struct {
	mu             sync.Mutex
	incomingRoutes []*ResolvedRoute

	// failRoutes are routes that will fail to program.
	failRoutes []*ResolvedRoute
}

func (dp *FakeDataplane) ProgramRoute(r *ResolvedRoute) error {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	for _, failroute := range dp.failRoutes {
		if reflect.DeepEqual(r, failroute) {
			return fmt.Errorf("route failed to program: %v", r)
		}
	}
	dp.incomingRoutes = append(dp.incomingRoutes, r)
	return nil
}

// SetupFailRoutes sets routes that will fail to program.
func (dp *FakeDataplane) SetupFailRoutes(failRoutes []*ResolvedRoute) {
	dp.failRoutes = failRoutes
}

func (dp *FakeDataplane) GetRoutes() []*ResolvedRoute {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	return dp.incomingRoutes
}

func (dp *FakeDataplane) ClearQueue() {
	dp.mu.Lock()
	defer dp.mu.Unlock()
	dp.incomingRoutes = []*ResolvedRoute{}
}

func NewFakeDataplane() *FakeDataplane {
	return &FakeDataplane{}
}

// routeSliceToMap converts a slice of ResolvedRoute to a map keyed by their
// RouteKeys. It returns an error if any of the routes were nil or if there is
// a duplicate.
func routeSliceToMap(rs []*ResolvedRoute) (map[RouteKey]*ResolvedRoute, error) {
	ret := map[RouteKey]*ResolvedRoute{}
	for _, rr := range rs {
		if rr == nil {
			return nil, fmt.Errorf("Got nil route in ResolvedRoute slice")
		}
		if existing, ok := ret[rr.RouteKey]; ok {
			return nil, fmt.Errorf("Got duplicate route key:\nFirst: %+v\nDuplicate: %+v", existing, rr)
		}
		ret[rr.RouteKey] = rr
	}
	return ret, nil
}

func checkResolvedRoutesEqual(got, want []*ResolvedRoute) error {
	gotRoutes, err := routeSliceToMap(got)
	if err != nil {
		return err
	}
	wantRoutes, err := routeSliceToMap(want)
	if err != nil {
		return err
	}

	if diff := cmp.Diff(gotRoutes, wantRoutes); diff != "" {
		return fmt.Errorf("Resolved routes are not equal: (-got, +want):\n%s", diff)
	}
	return nil
}

func configureInterface(t *testing.T, intf *AddIntfAction, yclient *ygnmi.Client) {
	t.Helper()

	ocintf := &oc.Interface{}
	ocintf.Name = ygot.String(intf.name)
	ocintf.Enabled = ygot.Bool(intf.enabled)
	ocintf.Ifindex = ygot.Uint32(uint32(intf.ifindex))
	prefix, err := netip.ParsePrefix(intf.prefix)
	if err != nil {
		t.Fatalf("Invalid prefix: %q", intf.prefix)
	}
	switch {
	case prefix.Addr().Is4():
		ocaddr := ocintf.GetOrCreateSubinterface(0).GetOrCreateIpv4().GetOrCreateAddress(prefix.Addr().String())
		ocaddr.PrefixLength = ygot.Uint8(uint8(prefix.Bits()))
	case prefix.Addr().Is6():
		ocaddr := ocintf.GetOrCreateSubinterface(0).GetOrCreateIpv6().GetOrCreateAddress(prefix.Addr().String())
		ocaddr.PrefixLength = ygot.Uint8(uint8(prefix.Bits()))
	default:
		t.Fatalf("Prefix is neither IPv4 nor IPv6: %q", intf.prefix)
	}

	if _, err := gnmiclient.Replace(context.Background(), yclient, ocpath.Root().Interface(intf.name).State(), ocintf); err != nil {
		t.Fatalf("Cannot configure interface: %v", err)
	}
}

var (
	// v4v6Mapper provides a mapping of IPv4 addresses to IPv6 addresses
	// for running the same tests, but using v6 addresses.
	v4v6Mapper = map[string]string{
		"192.168.1.1/24": "2001::1/49",
		"192.168.2.1/24": "2002::1/49",
		"192.168.3.1/24": "2003::1/49",
		"192.168.4.1/24": "2004::1/49",
		"192.168.5.1/24": "2005::1/49",
		"192.168.1.0/24": "2001::/49",
		"192.168.2.0/24": "2002::/49",
		"192.168.3.0/24": "2003::/49",
		"192.168.4.0/24": "2004::/49",
		"192.168.5.0/24": "2005::/49",
		"192.168.1.42":   "2001::42",
		"192.168.2.42":   "2002::42",
		"192.168.3.42":   "2003::42",
		"192.168.4.42":   "2004::42",
		"192.168.5.42":   "2005::42",
		"10.0.0.0/8":     "1000::/40",
		"20.0.0.0/8":     "2100::/40",
		"30.0.0.0/8":     "3000::/40",
		"40.0.0.0/8":     "4000::/40",
		"10.0.0.0":       "1000::",
		"20.0.0.0":       "2100::",
		"30.0.0.0":       "3000::",
		"40.0.0.0":       "4000::",
		"10.10.10.10":    "1000::10",
		"20.10.10.10":    "2100::10",
		"15.0.0.0":       "1500::",
		"11.10.10.10":    "1100::10",
		"192.0.2.1/30":   "2001:1::1/49",
		"192.0.2.5/30":   "2001:2::1/49",
		"192.0.2.9/30":   "2001:3::1/49",
		"192.0.2.0/30":   "2001:1::/49",
		"192.0.2.4/30":   "2001:2::/49",
		"192.0.2.8/30":   "2001:3::/49",
		// GUE only
		"10.10.20.20": "1000:0:20::20",
	}

	v4v6PolicyMap = map[GUEHeaders]GUEHeaders{
		{
			GUEPolicy: GUEPolicy{
				dstPortv4: 8,
				srcIP4:    [4]byte{42, 42, 42, 42},
				isV6:      false,
			},
			dstIP4: [4]byte{192, 168, 1, 42},
		}: {
			GUEPolicy: GUEPolicy{
				dstPortv6: 8,
				srcIP6:    [16]byte{42, 42, 42, 42, 42},
				isV6:      true,
			},
			dstIP6: [16]byte{0x20, 0x1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x42},
		},
		{
			GUEPolicy: GUEPolicy{
				dstPortv4: 9,
				srcIP4:    [4]byte{43, 43, 43, 43},
				isV6:      false,
			},
			dstIP4: [4]byte{10, 10, 10, 10},
		}: {
			GUEPolicy: GUEPolicy{
				dstPortv6: 9,
				srcIP6:    [16]byte{43, 43, 43, 43, 43},
				isV6:      true,
			},
			// 10::10
			dstIP6: [16]byte{0x10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10},
		},
		{
			GUEPolicy: GUEPolicy{
				dstPortv4: 8,
				srcIP4:    [4]byte{8, 8, 8, 8},
				isV6:      false,
			},
			dstIP4: [4]byte{10, 10, 10, 10},
		}: {
			GUEPolicy: GUEPolicy{
				dstPortv6: 8,
				srcIP6:    [16]byte{8, 8, 8, 8, 8},
				isV6:      true,
			},
			// 10::10
			dstIP6: [16]byte{0x10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x10},
		},
		{
			GUEPolicy: GUEPolicy{
				dstPortv4: 8,
				srcIP4:    [4]byte{8, 8, 8, 8},
				isV6:      false,
			},
			dstIP4: [4]byte{10, 10, 20, 20},
		}: {
			GUEPolicy: GUEPolicy{
				dstPortv6: 8,
				srcIP6:    [16]byte{8, 8, 8, 8, 8},
				isV6:      true,
			},
			dstIP6: [16]byte{0x10, 0, 0, 0, 0, 0x20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x20},
		},
		{
			GUEPolicy: GUEPolicy{
				dstPortv4: 16,
				srcIP4:    [4]byte{16, 16, 16, 16},
				isV6:      false,
			},
			dstIP4: [4]byte{10, 10, 20, 20},
		}: {
			GUEPolicy: GUEPolicy{
				dstPortv6: 16,
				srcIP6:    [16]byte{16, 16, 16, 16, 16},
				isV6:      true,
			},
			// 1000:20::20
			dstIP6: [16]byte{0x10, 0, 0, 0, 0, 0x20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x20},
		},
	}
)

func mapPolicy(t *testing.T, h *GUEHeaders) {
	zero := GUEHeaders{}
	if *h == zero {
		return
	}
	v6, ok := v4v6PolicyMap[*h]
	if !ok {
		t.Fatalf("v4 policy not in v4v6PolicyMap: %v", *h)
	}
	*h = v6
}

func mapAddress(t *testing.T, dst *string) {
	if *dst == "" {
		return
	}
	v6, ok := v4v6Mapper[*dst]
	if !ok {
		t.Fatalf("v4 address not in v4v6Mapper: %q", *dst)
	}
	*dst = v6
}

func mapResolvedRoute(t *testing.T, route *ResolvedRoute) {
	mapAddress(t, &route.Prefix)
	// Since this is not a pointer need to overwrite.
	nexthops := map[ResolvedNexthop]bool{}
	for nh, v := range route.Nexthops {
		mapAddress(t, &nh.Address)
		mapPolicy(t, &nh.GUEHeaders)
		nexthops[nh] = v
	}
	route.Nexthops = nexthops
}

func mapPrefix(t *testing.T, prefix *pb.Prefix) {
	if prefix.Family == pb.Prefix_FAMILY_IPV4 {
		prefix.Family = pb.Prefix_FAMILY_IPV6
	}
	mapAddress(t, &prefix.Address)
	if prefix.MaskLength == 8 {
		prefix.MaskLength = 40
	}
}

func TestServer(t *testing.T) {
	// TODO(wenbli): This test should be refactored such that the wantResolvedRoutes is inlined with the inSetRouteRequests for easier reading.
	tests := []struct {
		desc                       string
		inInterfaces               []*AddIntfAction
		wantInitialConnectedRoutes []*ResolvedRoute
		inSetRouteRequests         []*SetRouteRequestAction
		inFailRoutes               []*ResolvedRoute
		wantResolvedRoutes         [][]*ResolvedRoute
	}{{
		desc: "Route Additions", // TODO(wenbli): test route deletion in this test case once it's implemented.
		inInterfaces: []*AddIntfAction{{
			name:    "eth0",
			ifindex: 0,
			enabled: true,
			prefix:  "192.168.1.1/24",
			niName:  "DEFAULT",
		}, {
			name:    "eth1",
			ifindex: 1,
			enabled: true,
			prefix:  "192.168.2.1/24",
			niName:  "DEFAULT",
		}, {
			name:    "eth2",
			ifindex: 2,
			enabled: true,
			prefix:  "192.168.3.1/24",
			niName:  "DEFAULT",
		}, {
			name:    "eth3",
			ifindex: 3,
			enabled: true,
			prefix:  "192.168.4.1/24",
			niName:  "DEFAULT",
		}, {
			name:    "eth4",
			ifindex: 4,
			enabled: true,
			prefix:  "192.168.5.1/24",
			niName:  "DEFAULT",
		}},
		wantInitialConnectedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "192.168.1.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "192.168.2.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth1",
						Index: 1,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "192.168.3.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 2,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "192.168.4.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth3",
						Index: 3,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "192.168.5.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth4",
						Index: 4,
					},
				}: true,
			},
		}},
		inSetRouteRequests: []*SetRouteRequestAction{{
			Desc: "1st level indirect route",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 10,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.1.42",
					// TODO(wenbli): Implement WCMP, for all route requests in this test.
					Weight: 1,
				}},
			},
		}, {
			Desc: "2nd level indirect route",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 10,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "20.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "10.10.10.10",
				}},
			},
		}, {
			Desc: "3rd level indirect route",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 10,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "30.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "20.10.10.10",
				}},
			},
		}, {
			Desc: "secondary 1st level indirect route that has higher admin distance",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 20,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.2.42",
				}},
			},
		}, {
			Desc: "secondary 1st level indirect route that has lower admin distance",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 5,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.3.42",
				}},
			},
		}, {
			Desc: "secondary 1st level indirect route that has higher metric",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 5,
				Metric:        999,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.4.42",
				}},
			},
		}, {
			Desc: "secondary 1st level indirect route that has lower metric",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 5,
				Metric:        5,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.5.42",
				}},
			},
		}},
		wantResolvedRoutes: [][]*ResolvedRoute{{{
			RouteKey: RouteKey{
				Prefix: "10.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}}, {{
			RouteKey: RouteKey{
				Prefix: "20.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}}, {{
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}}, {}, {{
			RouteKey: RouteKey{
				Prefix: "10.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.3.42",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 2,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "20.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.3.42",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 2,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.3.42",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 2,
					},
				}: true,
			},
		}}, {}, {{
			RouteKey: RouteKey{
				Prefix: "10.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.5.42",
					},
					Port: Interface{
						Name:  "eth4",
						Index: 4,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "20.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.5.42",
					},
					Port: Interface{
						Name:  "eth4",
						Index: 4,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.5.42",
					},
					Port: Interface{
						Name:  "eth4",
						Index: 4,
					},
				}: true,
			},
		}}},
	}, {
		desc: "Unresolvable and ECMP",
		inInterfaces: []*AddIntfAction{{
			name:    "eth0",
			ifindex: 0,
			enabled: false,
			prefix:  "192.168.1.1/24",
			niName:  "DEFAULT",
		}, {
			name:    "eth1",
			ifindex: 1,
			enabled: true,
			prefix:  "192.168.2.1/24",
			niName:  "DEFAULT",
		}, {
			name:    "eth2",
			ifindex: 2,
			enabled: true,
			prefix:  "192.168.3.1/24",
			niName:  "DEFAULT",
		}},
		wantInitialConnectedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "192.168.2.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth1",
						Index: 1,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "192.168.3.0/24",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 2,
					},
				}: true,
			},
		}},
		inSetRouteRequests: []*SetRouteRequestAction{{
			Desc: "unresolvable route",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 10,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "15.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "11.10.10.10",
				}},
			},
		}, {
			Desc: "route that resolves over down interface",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 10,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.1.42",
				}},
			},
		}, {
			Desc: "same route that resolves over up interface with higher admin distance",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 20,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.2.42",
				}},
			},
		}, {
			Desc: "ECMP",
			RouteReq: &pb.SetRouteRequest{
				AdminDistance: 20,
				Metric:        10,
				Prefix: &pb.Prefix{
					Family:     pb.Prefix_FAMILY_IPV4,
					Address:    "10.0.0.0",
					MaskLength: 8,
				},
				Nexthops: []*pb.Nexthop{{
					Type:    pb.Nexthop_TYPE_IPV4,
					Address: "192.168.3.42",
				}},
			},
		}},
		wantResolvedRoutes: [][]*ResolvedRoute{{}, {}, {{
			RouteKey: RouteKey{
				Prefix: "10.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.2.42",
					},
					Port: Interface{
						Name:  "eth1",
						Index: 1,
					},
				}: true,
			},
		}}, {{
			RouteKey: RouteKey{
				Prefix: "10.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.2.42",
					},
					Port: Interface{
						Name:  "eth1",
						Index: 1,
					},
				}: true,
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.3.42",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 2,
					},
				}: true,
			},
		}}},
	}, {
		desc: "test route program failures",
		inInterfaces: []*AddIntfAction{{
			name:    "eth1",
			ifindex: 3,
			enabled: true,
			prefix:  "192.0.2.1/30",
			niName:  "DEFAULT",
		}, {
			name:    "eth2",
			ifindex: 4,
			enabled: true,
			prefix:  "192.0.2.5/30",
			niName:  "DEFAULT",
		}, {
			name:    "eth3",
			ifindex: 5,
			enabled: true,
			prefix:  "192.0.2.9/30",
			niName:  "DEFAULT",
		}},
		inFailRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "192.0.2.0/30",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth1",
						Index: 3,
					},
				}: true,
			},
		}},
		wantInitialConnectedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "192.0.2.4/30",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth2",
						Index: 4,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "192.0.2.8/30",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
					},
					Port: Interface{
						Name:  "eth3",
						Index: 5,
					},
				}: true,
			},
		}},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			for _, v4 := range []bool{true, false} {
				desc := "v4"
				if !v4 {
					desc = "v6"
					// Convert all v4 addresses to v6.
					for _, intf := range tt.inInterfaces {
						mapAddress(t, &intf.prefix)
					}
					for _, req := range tt.inSetRouteRequests {
						mapPrefix(t, req.RouteReq.Prefix)
						for _, nh := range req.RouteReq.Nexthops {
							if nh.Type == pb.Nexthop_TYPE_IPV4 {
								nh.Type = pb.Nexthop_TYPE_IPV6
							}
							mapAddress(t, &nh.Address)
						}
					}
					for _, route := range tt.wantInitialConnectedRoutes {
						mapResolvedRoute(t, route)
					}
					for _, route := range tt.inFailRoutes {
						mapResolvedRoute(t, route)
					}
					for _, routes := range tt.wantResolvedRoutes {
						for _, route := range routes {
							mapResolvedRoute(t, route)
						}
					}
				}

				t.Run(desc, func(t *testing.T) {
					// TODO(wenbli): Don't re-create gNMI server, simply erase it and then reconnect to it afterwards.
					grpcServer := grpc.NewServer()
					gnmiServer, err := gnmi.New(grpcServer, "local")
					if err != nil {
						t.Fatal(err)
					}
					lis, err := net.Listen("tcp", "localhost:0")
					if err != nil {
						t.Fatalf("Failed to start listener: %v", err)
					}
					go func() {
						grpcServer.Serve(lis)
					}()

					dp := NewFakeDataplane()
					dp.SetupFailRoutes(tt.inFailRoutes)
					s, err := New()
					if err != nil {
						t.Fatal(err)
					}

					// Update the interface configuration on the gNMI server.
					client := gnmiServer.LocalClient()
					if err := s.Start(client, "local", ""); err != nil {
						t.Fatalf("cannot start sysrib server, %v", err)
					}
					s.dataplane = dp
					defer s.Stop()

					c, err := ygnmi.NewClient(client, ygnmi.WithTarget("local"))
					if err != nil {
						t.Fatalf("cannot create ygnmi client: %v", err)
					}
					for _, intf := range tt.inInterfaces {
						configureInterface(t, intf, c)
					}

					// Wait for Sysrib to pick up the connected prefixes.
					for i := 0; i != maxGNMIWaitQuanta; i++ {
						if err = checkResolvedRoutesEqual(dp.GetRoutes(), tt.wantInitialConnectedRoutes); err == nil {
							break
						}
						time.Sleep(100 * time.Millisecond)
					}
					if err != nil {
						t.Fatalf("After initial interface operations: %v", err)
					}
					dp.ClearQueue()

					for i, routeReq := range tt.inSetRouteRequests {
						// TODO(wenbli): Test SetRouteResponse
						if _, err := s.SetRoute(context.Background(), routeReq.RouteReq); err != nil {
							t.Fatalf("%s: Got unexpected error during call to SetRoute: %v", routeReq.Desc, err)
						}
						if err := checkResolvedRoutesEqual(dp.GetRoutes(), tt.wantResolvedRoutes[i]); err != nil {
							t.Fatalf("%s: %v", routeReq.Desc, err)
						}
						dp.ClearQueue()
					}
				})
			}
		})
	}
}

func TestBGPGUEPolicy(t *testing.T) {
	intfs := []*AddIntfAction{{
		name:    "eth0",
		ifindex: 0,
		enabled: true,
		prefix:  "192.168.1.1/24",
		niName:  "DEFAULT",
	}, {
		name:    "eth1",
		ifindex: 1,
		enabled: true,
		prefix:  "192.168.2.1/24",
		niName:  "DEFAULT",
	}, {
		name:    "eth2",
		ifindex: 2,
		enabled: true,
		prefix:  "192.168.3.1/24",
		niName:  "DEFAULT",
	}, {
		name:    "eth3",
		ifindex: 3,
		enabled: true,
		prefix:  "192.168.4.1/24",
		niName:  "DEFAULT",
	}, {
		name:    "eth4",
		ifindex: 4,
		enabled: true,
		prefix:  "192.168.5.1/24",
		niName:  "DEFAULT",
	}}

	wantConnectedRoutes := []*ResolvedRoute{{
		RouteKey: RouteKey{
			Prefix: "192.168.1.0/24",
			NIName: "DEFAULT",
		},
		Nexthops: map[ResolvedNexthop]bool{
			{
				NextHopSummary: afthelper.NextHopSummary{
					NetworkInstance: "DEFAULT",
				},
				Port: Interface{
					Name:  "eth0",
					Index: 0,
				},
			}: true,
		},
	}, {
		RouteKey: RouteKey{
			Prefix: "192.168.2.0/24",
			NIName: "DEFAULT",
		},
		Nexthops: map[ResolvedNexthop]bool{
			{
				NextHopSummary: afthelper.NextHopSummary{
					NetworkInstance: "DEFAULT",
				},
				Port: Interface{
					Name:  "eth1",
					Index: 1,
				},
			}: true,
		},
	}, {
		RouteKey: RouteKey{
			Prefix: "192.168.3.0/24",
			NIName: "DEFAULT",
		},
		Nexthops: map[ResolvedNexthop]bool{
			{
				NextHopSummary: afthelper.NextHopSummary{
					NetworkInstance: "DEFAULT",
				},
				Port: Interface{
					Name:  "eth2",
					Index: 2,
				},
			}: true,
		},
	}, {
		RouteKey: RouteKey{
			Prefix: "192.168.4.0/24",
			NIName: "DEFAULT",
		},
		Nexthops: map[ResolvedNexthop]bool{
			{
				NextHopSummary: afthelper.NextHopSummary{
					NetworkInstance: "DEFAULT",
				},
				Port: Interface{
					Name:  "eth3",
					Index: 3,
				},
			}: true,
		},
	}, {
		RouteKey: RouteKey{
			Prefix: "192.168.5.0/24",
			NIName: "DEFAULT",
		},
		Nexthops: map[ResolvedNexthop]bool{
			{
				NextHopSummary: afthelper.NextHopSummary{
					NetworkInstance: "DEFAULT",
				},
				Port: Interface{
					Name:  "eth4",
					Index: 4,
				},
			}: true,
		},
	}}

	// Note: This is a sequential test -- each test case depends on the previous one.
	tests := []struct {
		desc               string
		inSetRouteRequests []*pb.SetRouteRequest
		inAddPolicies      map[string]GUEPolicy
		inAddPoliciesV6    map[string]GUEPolicy
		inDeletePolicies   []string
		inDeletePoliciesV6 []string
		wantResolvedRoutes []*ResolvedRoute
	}{{
		desc: "Add static and BGP routes",
		inSetRouteRequests: []*pb.SetRouteRequest{{
			AdminDistance: 10, // not BGP
			Metric:        10,
			Prefix: &pb.Prefix{
				Family:     pb.Prefix_FAMILY_IPV4,
				Address:    "10.0.0.0",
				MaskLength: 8,
			},
			Nexthops: []*pb.Nexthop{{
				Type:    pb.Nexthop_TYPE_IPV4,
				Address: "192.168.1.42",
			}},
		}, {
			AdminDistance: 20, // EBGP
			Metric:        10,
			Prefix: &pb.Prefix{
				Family:     pb.Prefix_FAMILY_IPV4,
				Address:    "20.0.0.0",
				MaskLength: 8,
			},
			Nexthops: []*pb.Nexthop{{
				Type:    pb.Nexthop_TYPE_IPV4,
				Address: "192.168.1.42",
			}},
		}},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "10.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "20.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}, {
		desc: "Add Policy",
		inAddPolicies: map[string]GUEPolicy{
			"192.168.0.0/16": {
				dstPortv4: 8,
				srcIP4:    [4]byte{42, 42, 42, 42},
				isV6:      false,
			},
		},
		inAddPoliciesV6: map[string]GUEPolicy{
			"2001::/16": {
				dstPortv6: 8,
				srcIP6:    [16]byte{42, 42, 42, 42, 42},
				isV6:      true,
			},
		},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "20.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
					GUEHeaders: GUEHeaders{
						GUEPolicy: GUEPolicy{
							dstPortv4: 8,
							srcIP4:    [4]byte{42, 42, 42, 42},
							isV6:      false,
						},
						dstIP4: [4]byte{192, 168, 1, 42},
					},
				}: true,
			},
		}},
	}, {
		desc:               "Remove Policy",
		inDeletePolicies:   []string{"192.168.0.0/16"},
		inDeletePoliciesV6: []string{"2001::/16"},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "20.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}, {
		desc: "Add BGP route that resolves over the static route",
		inSetRouteRequests: []*pb.SetRouteRequest{{
			AdminDistance: 20, // EBGP
			Metric:        10,
			Prefix: &pb.Prefix{
				Family:     pb.Prefix_FAMILY_IPV4,
				Address:    "30.0.0.0",
				MaskLength: 8,
			},
			Nexthops: []*pb.Nexthop{{
				Type:    pb.Nexthop_TYPE_IPV4,
				Address: "10.10.10.10",
			}},
		}},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}, {
		desc: "Add Policy for second BGP route",
		inAddPolicies: map[string]GUEPolicy{
			"10.10.0.0/16": {
				dstPortv4: 9,
				srcIP4:    [4]byte{43, 43, 43, 43},
				isV6:      false,
			},
		},
		inAddPoliciesV6: map[string]GUEPolicy{
			"1000::/16": {
				dstPortv6: 9,
				srcIP6:    [16]byte{43, 43, 43, 43, 43},
				isV6:      true,
			},
		},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
					GUEHeaders: GUEHeaders{
						GUEPolicy: GUEPolicy{
							dstPortv4: 9,
							srcIP4:    [4]byte{43, 43, 43, 43},
							isV6:      false,
						},
						dstIP4: [4]byte{10, 10, 10, 10},
					},
				}: true,
			},
		}},
	}, {
		desc:               "Remove Policy for second BGP route",
		inDeletePolicies:   []string{"10.10.0.0/16"},
		inDeletePoliciesV6: []string{"1000::/16"},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}, {
		desc: "Add another BGP route that resolves over the static route",
		inSetRouteRequests: []*pb.SetRouteRequest{{
			AdminDistance: 20, // EBGP
			Metric:        10,
			Prefix: &pb.Prefix{
				Family:     pb.Prefix_FAMILY_IPV4,
				Address:    "40.0.0.0",
				MaskLength: 8,
			},
			Nexthops: []*pb.Nexthop{{
				Type:    pb.Nexthop_TYPE_IPV4,
				Address: "10.10.20.20",
			}},
		}},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "40.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}, {
		desc: "Add a policy that applies to two BGP routes",
		inAddPolicies: map[string]GUEPolicy{
			"10.0.0.0/8": {
				dstPortv4: 8,
				srcIP4:    [4]byte{8, 8, 8, 8},
				isV6:      false,
			},
		},
		inAddPoliciesV6: map[string]GUEPolicy{
			"1000::/8": {
				dstPortv6: 8,
				srcIP6:    [16]byte{8, 8, 8, 8, 8},
				isV6:      true,
			},
		},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
					GUEHeaders: GUEHeaders{
						GUEPolicy: GUEPolicy{
							dstPortv4: 8,
							srcIP4:    [4]byte{8, 8, 8, 8},
							isV6:      false,
						},
						dstIP4: [4]byte{10, 10, 10, 10},
					},
				}: true,
			},
		}, {
			RouteKey: RouteKey{
				Prefix: "40.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
					GUEHeaders: GUEHeaders{
						GUEPolicy: GUEPolicy{
							dstPortv4: 8,
							srcIP4:    [4]byte{8, 8, 8, 8},
							isV6:      false,
						},
						dstIP4: [4]byte{10, 10, 20, 20},
					},
				}: true,
			},
		}},
	}, {
		desc: "Add a more specific policy that applies to a BGP route",
		inAddPolicies: map[string]GUEPolicy{
			"10.10.20.0/24": {
				dstPortv4: 16,
				srcIP4:    [4]byte{16, 16, 16, 16},
				isV6:      false,
			},
		},
		inAddPoliciesV6: map[string]GUEPolicy{
			"1000:0:20::/48": {
				dstPortv6: 16,
				srcIP6:    [16]byte{16, 16, 16, 16, 16},
				isV6:      true,
			},
		},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "40.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
					GUEHeaders: GUEHeaders{
						GUEPolicy: GUEPolicy{
							dstPortv4: 16,
							srcIP4:    [4]byte{16, 16, 16, 16},
							isV6:      false,
						},
						dstIP4: [4]byte{10, 10, 20, 20},
					},
				}: true,
			},
		}},
	}, {
		desc:               "Remove the less-specific policy",
		inDeletePolicies:   []string{"10.0.0.0/8"},
		inDeletePoliciesV6: []string{"1000::/8"},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "30.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}, {
		desc:               "Remove the more-specific policy",
		inDeletePolicies:   []string{"10.10.20.0/24"},
		inDeletePoliciesV6: []string{"1000:0:20::/48"},
		wantResolvedRoutes: []*ResolvedRoute{{
			RouteKey: RouteKey{
				Prefix: "40.0.0.0/8",
				NIName: "DEFAULT",
			},
			Nexthops: map[ResolvedNexthop]bool{
				{
					NextHopSummary: afthelper.NextHopSummary{
						NetworkInstance: "DEFAULT",
						Address:         "192.168.1.42",
					},
					Port: Interface{
						Name:  "eth0",
						Index: 0,
					},
				}: true,
			},
		}},
	}}

	for _, v4 := range []bool{true, false} {
		desc := "v4"
		if !v4 {
			desc = "v6"
			// Convert all v4 addresses to v6.
			for _, intf := range intfs {
				mapAddress(t, &intf.prefix)
			}
			for _, route := range wantConnectedRoutes {
				mapResolvedRoute(t, route)
			}
		}

		t.Run(desc, func(t *testing.T) {
			grpcServer := grpc.NewServer()
			gnmiServer, err := gnmi.New(grpcServer, "local")
			if err != nil {
				t.Fatal(err)
			}
			lis, err := net.Listen("tcp", "localhost:0")
			if err != nil {
				t.Fatalf("Failed to start listener: %v", err)
			}
			go func() {
				grpcServer.Serve(lis)
			}()

			dp := NewFakeDataplane()
			s, err := New()
			if err != nil {
				t.Fatal(err)
			}

			// Update the interface configuration on the gNMI server.
			client := gnmiServer.LocalClient()
			if err := s.Start(client, "local", ""); err != nil {
				t.Fatalf("cannot start sysrib server, %v", err)
			}
			s.dataplane = dp
			defer s.Stop()

			c, err := ygnmi.NewClient(client, ygnmi.WithTarget("local"))
			if err != nil {
				t.Fatalf("cannot create ygnmi client: %v", err)
			}

			for _, intf := range intfs {
				configureInterface(t, intf, c)
			}
			// Wait for Sysrib to pick up the connected prefixes.
			for i := 0; i != maxGNMIWaitQuanta; i++ {
				if err = checkResolvedRoutesEqual(dp.GetRoutes(), wantConnectedRoutes); err == nil {
					break
				}
				time.Sleep(100 * time.Millisecond)
			}
			if err != nil {
				t.Fatalf("After initial interface operations: %v", err)
			}
			dp.ClearQueue()

			for _, tt := range tests {
				inAddPolicies := tt.inAddPolicies
				inDeletePolicies := tt.inDeletePolicies
				if !v4 { // Convert v4 to v6.
					for _, req := range tt.inSetRouteRequests {
						mapPrefix(t, req.Prefix)
						for _, nh := range req.Nexthops {
							if nh.Type == pb.Nexthop_TYPE_IPV4 {
								nh.Type = pb.Nexthop_TYPE_IPV6
							}
							mapAddress(t, &nh.Address)
						}
					}
					for _, route := range tt.wantResolvedRoutes {
						mapResolvedRoute(t, route)
					}
					inAddPolicies = tt.inAddPoliciesV6
					inDeletePolicies = tt.inDeletePoliciesV6
				}

				t.Run(tt.desc, func(t *testing.T) {
					for _, routeReq := range tt.inSetRouteRequests {
						if _, err := s.SetRoute(context.Background(), routeReq); err != nil {
							t.Fatalf("Got unexpected error during call to SetRoute: %v", err)
						}
					}
					for prefix, policy := range inAddPolicies {
						s.setGUEPolicy(prefix, policy)
					}
					for _, prefix := range inDeletePolicies {
						s.deleteGUEPolicy(prefix)
					}
					if err := checkResolvedRoutesEqual(dp.GetRoutes(), tt.wantResolvedRoutes); err != nil {
						fmt.Println("debug", tt.desc)
						s.rib.PrintRIB()
						t.Fatalf("%v", err)
					}
					dp.ClearQueue()
				})
			}
		})
	}
}
