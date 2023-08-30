// Copyright 2023 Google LLC
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

package local_test

import (
	"fmt"
	"testing"

	"github.com/openconfig/lemming/bgp"
	"github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/lemming/gnmi/oc/ocpath"
	"github.com/openconfig/ygot/ygot"

	valpb "github.com/openconfig/lemming/bgp/tests/proto/policyval"
)

func TestCommunitySet(t *testing.T) {
	installRejectPolicy := func(t *testing.T, dut1, dut2, _, _, _ *Device) {
		// Policy to reject routes with the given community set conditions
		policyName := "community-sets"
		policy := &oc.RoutingPolicy_PolicyDefinition_Statement_OrderedMap{}

		// Create ANY community set
		anyCommSetName := "ANY-community-set"
		anyCommMemberPath := ocpath.Root().RoutingPolicy().DefinedSets().BgpDefinedSets().CommunitySet(anyCommSetName).CommunityMember()
		Replace(t, dut2, anyCommMemberPath.Config(), []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union{
			oc.UnionString("11111:11111"),
			oc.UnionString("22222:22222"),
		})
		Replace(t, dut2, ocpath.Root().RoutingPolicy().DefinedSets().BgpDefinedSets().CommunitySet(anyCommSetName).MatchSetOptions().Config(), oc.RoutingPolicy_MatchSetOptionsType_ANY)

		// Create ALL community set
		allCommSetName := "ALL-community-set"
		allCommMemberPath := ocpath.Root().RoutingPolicy().DefinedSets().BgpDefinedSets().CommunitySet(allCommSetName).CommunityMember()
		Replace(t, dut2, allCommMemberPath.Config(), []oc.RoutingPolicy_DefinedSets_BgpDefinedSets_CommunitySet_CommunityMember_Union{
			oc.UnionString("33333:33333"),
			oc.UnionString("44444:44444"),
		})
		Replace(t, dut2, ocpath.Root().RoutingPolicy().DefinedSets().BgpDefinedSets().CommunitySet(allCommSetName).MatchSetOptions().Config(), oc.RoutingPolicy_MatchSetOptionsType_ALL)

		// Match on given list of community set members and reject them.
		stmt, err := policy.AppendNew("reject-any-community-sets")
		if err != nil {
			t.Fatalf("Cannot append new BGP policy statement: %v", err)
		}
		stmt.GetOrCreateConditions().GetOrCreateBgpConditions().SetCommunitySet(anyCommSetName)
		stmt.GetOrCreateActions().SetPolicyResult(oc.RoutingPolicy_PolicyResultType_REJECT_ROUTE)

		stmt, err = policy.AppendNew("reject-all-community-sets")
		if err != nil {
			t.Fatalf("Cannot append new BGP policy statement: %v", err)
		}
		stmt.GetOrCreateConditions().GetOrCreateBgpConditions().SetCommunitySet(allCommSetName)
		stmt.GetOrCreateActions().SetPolicyResult(oc.RoutingPolicy_PolicyResultType_REJECT_ROUTE)

		// Install policy
		Replace(t, dut2, ocpath.Root().RoutingPolicy().PolicyDefinition(policyName).Config(), &oc.RoutingPolicy_PolicyDefinition{Statement: policy})
		Replace(t, dut2, bgp.BGPPath.Neighbor(dut1.RouterID).ApplyPolicy().ImportPolicy().Config(), []string{policyName})
	}

	routeUnderTestList := []string{
		"10.0.0.0/16",
		"10.1.0.0/16",
		"10.2.0.0/16",
		"10.3.0.0/16",
		"10.4.0.0/16",
		"10.5.0.0/16",
	}

	installSetPolicy := func(t *testing.T, dut1, dut2, _, _, _ *Device) {
		if debug {
			fmt.Println("Installing test policies")
		}

		policyName := "set-communities"
		policy := &oc.RoutingPolicy_PolicyDefinition_Statement_OrderedMap{}

		for i, route := range routeUnderTestList {
			// Create prefix set
			prefixSetName := "accept-" + route
			prefixPath := ocpath.Root().RoutingPolicy().DefinedSets().PrefixSet(prefixSetName).Prefix(route, "exact").IpPrefix()
			Replace(t, dut1, prefixPath.Config(), route)

			stmt, err := policy.AppendNew(fmt.Sprintf("stmt%d", i))
			if err != nil {
				t.Fatalf("Cannot append new BGP policy statement: %v", err)
			}
			stmt.GetOrCreateConditions().GetOrCreateMatchPrefixSet().SetPrefixSet(prefixSetName)
			stmt.GetOrCreateConditions().GetOrCreateMatchPrefixSet().SetMatchSetOptions(oc.RoutingPolicy_MatchSetOptionsRestrictedType_ANY)

			switch i {
			case 0:
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_REPLACE)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("10000:10000"),
					},
				)
			case 1:
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_REPLACE)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("11111:11111"),
					},
				)
			case 2:
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_REPLACE)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("33333:33333"),
					},
				)
			case 3:
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_REPLACE)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("33333:33333"),
						oc.UnionString("44444:44444"),
					},
				)
			case 4:
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_REPLACE)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("55555:55555"),
						oc.UnionString("44444:44444"),
						oc.UnionString("33333:33333"),
					},
				)
			case 5:
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_REPLACE)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("33333:33333"),
					},
				)

				stmt, err = policy.AppendNew(fmt.Sprintf("stmt%d-2", i))
				if err != nil {
					t.Fatalf("Cannot append new BGP policy statement: %v", err)
				}
				stmt.GetOrCreateConditions().GetOrCreateMatchPrefixSet().SetPrefixSet(prefixSetName)
				stmt.GetOrCreateConditions().GetOrCreateMatchPrefixSet().SetMatchSetOptions(oc.RoutingPolicy_MatchSetOptionsRestrictedType_ANY)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().SetOptions(oc.BgpPolicy_BgpSetCommunityOptionType_ADD)
				stmt.GetOrCreateActions().GetOrCreateBgpActions().GetOrCreateSetCommunity().GetOrCreateInline().SetCommunities(
					[]oc.RoutingPolicy_PolicyDefinition_Statement_Actions_BgpActions_SetCommunity_Inline_Communities_Union{
						oc.UnionString("44444:44444"),
					},
				)
				// TODO(wenbli): Test REMOVE, it's possible GoBGP does not support it.
			}
			stmt.GetOrCreateActions().SetPolicyResult(oc.RoutingPolicy_PolicyResultType_ACCEPT_ROUTE)
		}
		// Install policy
		Replace(t, dut1, ocpath.Root().RoutingPolicy().PolicyDefinition(policyName).Config(), &oc.RoutingPolicy_PolicyDefinition{Name: ygot.String(policyName), Statement: policy})
		Replace(t, dut1, bgp.BGPPath.Neighbor(dut2.RouterID).ApplyPolicy().ExportPolicy().Config(), []string{policyName})
	}

	testPolicy(t, PolicyTestCase{
		spec: &valpb.PolicyTestCase{
			Description: "Test community set ANY and ALL",
			RouteTests: []*valpb.RouteTestCase{{
				Description: "No match",
				Input: &valpb.TestRoute{
					ReachPrefix: routeUnderTestList[0],
				},
				ExpectedResultBeforePolicy: valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
				ExpectedResult:             valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
			}, {
				Description: "Matches ANY",
				Input: &valpb.TestRoute{
					ReachPrefix: routeUnderTestList[1],
				},
				ExpectedResultBeforePolicy: valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
				ExpectedResult:             valpb.RouteTestResult_ROUTE_TEST_RESULT_DISCARD,
			}, {
				Description: "Partially matches ALL",
				Input: &valpb.TestRoute{
					ReachPrefix: routeUnderTestList[2],
				},
				ExpectedResultBeforePolicy: valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
				ExpectedResult:             valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
			}, {
				Description: "Matches ALL",
				Input: &valpb.TestRoute{
					ReachPrefix: routeUnderTestList[3],
				},
				ExpectedResultBeforePolicy: valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
				ExpectedResult:             valpb.RouteTestResult_ROUTE_TEST_RESULT_DISCARD,
			}, {
				Description: "Matches ALL reversed and with extra community",
				Input: &valpb.TestRoute{
					ReachPrefix: routeUnderTestList[4],
				},
				ExpectedResultBeforePolicy: valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
				ExpectedResult:             valpb.RouteTestResult_ROUTE_TEST_RESULT_DISCARD,
			}, {
				Description: "Matches ALL after ADD",
				Input: &valpb.TestRoute{
					ReachPrefix: routeUnderTestList[5],
				},
				ExpectedResultBeforePolicy: valpb.RouteTestResult_ROUTE_TEST_RESULT_ACCEPT,
				ExpectedResult:             valpb.RouteTestResult_ROUTE_TEST_RESULT_DISCARD,
			}},
		},
		installPolicies: func(t *testing.T, dut1, dut2, dut3, dut4, dut5 *Device) {
			installSetPolicy(t, dut1, dut2, dut3, dut4, dut5)
			installRejectPolicy(t, dut1, dut2, dut3, dut4, dut5)
		},
	})
}
