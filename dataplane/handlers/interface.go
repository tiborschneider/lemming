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

// Package handlers contains gNMI task handlers.
package handlers

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/openconfig/lemming/dataplane/internal/engine"
	"github.com/openconfig/lemming/dataplane/internal/kernel"
	"github.com/openconfig/lemming/gnmi/gnmiclient"
	"github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/lemming/gnmi/oc/ocpath"
	"github.com/openconfig/lemming/gnmi/reconciler"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"

	log "github.com/golang/glog"
	fwdpb "github.com/openconfig/lemming/proto/forwarding"
)

// Interface handles config updates to the /interfaces/... paths.
type Interface struct {
	c *ygnmi.Client
	// closers functions should all be invoked when the interface handler stops running.
	closers []func()
	fwd     fwdpb.ServiceClient
	stateMu sync.RWMutex
	// state keeps track of the applied state of the device's interfaces so that we do not issue duplicate configuration commands to the device's interfaces.
	state     map[string]*oc.Interface
	idxToName map[int]string
}

// NewInterface creates a new interface handler.
func NewInterface(fwd fwdpb.ServiceClient) *reconciler.BuiltReconciler {
	i := &Interface{
		fwd:       fwd,
		idxToName: map[int]string{},
		state:     map[string]*oc.Interface{},
	}
	return reconciler.NewBuilder("interface").WithStart(i.start).WithStop(i.stop).Build()
}

// Start starts running the handler, watching the cache and the kernel interfaces.
func (ni *Interface) start(ctx context.Context, client *ygnmi.Client) error {
	log.Info("starting interface handler")
	b := &ocpath.Batch{}
	ni.c = client

	if err := ni.setupPorts(ctx); err != nil {
		return fmt.Errorf("failed to setup ports: %v", err)
	}

	b.AddPaths(
		ocpath.Root().InterfaceAny().Name().Config().PathStruct(),
		ocpath.Root().InterfaceAny().Ethernet().MacAddress().Config().PathStruct(),
		ocpath.Root().InterfaceAny().Subinterface(0).Enabled().Config().PathStruct(), // TODO: Do we need enable at root interface level?
		ocpath.Root().InterfaceAny().Subinterface(0).Ipv4().AddressAny().Ip().Config().PathStruct(),
		ocpath.Root().InterfaceAny().Subinterface(0).Ipv4().AddressAny().PrefixLength().Config().PathStruct(),
		ocpath.Root().InterfaceAny().Subinterface(0).Ipv6().AddressAny().Ip().Config().PathStruct(),
		ocpath.Root().InterfaceAny().Subinterface(0).Ipv6().AddressAny().PrefixLength().Config().PathStruct(),
	)
	cancelCtx, cancelFn := context.WithCancel(ctx)

	watcher := ygnmi.Watch(cancelCtx, ni.c, b.Config(), func(val *ygnmi.Value[*oc.Root]) error {
		log.V(2).Info("reconciling interfaces")
		root, ok := val.Val()
		if !ok || root.Interface == nil {
			return ygnmi.Continue
		}
		for _, i := range root.Interface {
			ni.reconcile(i)
		}
		return ygnmi.Continue
	})

	linkDoneCh := make(chan struct{})
	linkUpdateCh := make(chan netlink.LinkUpdate)
	addrDoneCh := make(chan struct{})
	addrUpdateCh := make(chan netlink.AddrUpdate)
	neighDoneCh := make(chan struct{})
	neighUpdateCh := make(chan netlink.NeighUpdate)
	ni.closers = append(ni.closers, func() {
		close(linkDoneCh)
		close(addrDoneCh)
		close(neighDoneCh)
	}, cancelFn)

	if err := netlink.LinkSubscribe(linkUpdateCh, linkDoneCh); err != nil {
		return fmt.Errorf("failed to sub to link: %v", err)
	}
	if err := netlink.AddrSubscribe(addrUpdateCh, addrDoneCh); err != nil {
		return fmt.Errorf("failed to sub to addr: %v", err)
	}
	if err := netlink.NeighSubscribe(neighUpdateCh, addrDoneCh); err != nil {
		return fmt.Errorf("failed to sub to neighbor: %v", err)
	}

	go func() {
		for {
			select {
			case up := <-linkUpdateCh:
				ni.handleLinkUpdate(ctx, &up)
			case up := <-addrUpdateCh:
				ni.handleAddrUpdate(ctx, &up)
			case up := <-neighUpdateCh:
				ni.handleNeighborUpdate(ctx, &up)
			}
		}
	}()

	go func() {
		// TODO: handle error
		if _, err := watcher.Await(); err != nil {
			log.Warningf("interface watch err: %v", err)
		}
	}()

	ni.startCounterUpdates(ctx)

	return nil
}

// Stop stops all watchers.
func (ni *Interface) stop(context.Context) error {
	// TODO: prevent stopping more than once.
	for _, closeFn := range ni.closers {
		closeFn()
	}
	return nil
}

// startCounterUpdates starts a goroutine for updating counters for configured
// interfaces.
func (ni *Interface) startCounterUpdates(ctx context.Context) {
	tick := time.NewTicker(time.Second)
	ni.closers = append(ni.closers, tick.Stop)
	go func() {
		// Design comment:
		// This polling can be eliminated if either the forwarding
		// service supported streaming the counters, or if somehow the
		// gnmi cache were able to forward queries to prompt the data
		// producer to populate the leaf.
		//
		// However, given counters are likely frequently-updated values
		// anyways, it may be fine for counter values to be polled.
		for range tick.C {
			ni.stateMu.RLock()
			var intfNames []string
			for intfName := range ni.state {
				// TODO(wenbli): Support interface state deletion when interface is deleted.
				intfNames = append(intfNames, intfName)
			}
			ni.stateMu.RUnlock()
			for _, intfName := range intfNames {
				countersReply, err := ni.fwd.ObjectCounters(ctx, &fwdpb.ObjectCountersRequest{
					ObjectId:  &fwdpb.ObjectId{Id: intfName},
					ContextId: &fwdpb.ContextId{Id: engine.DefaultContextID},
				})
				log.V(1).Infof("querying counters for interface %q, got %v", intfName, countersReply)
				if err != nil {
					log.Errorf("interface handler: could not retrieve counter for interface %q", intfName)
					continue
				}
				for _, counter := range countersReply.Counters {
					switch counter.Id {
					case fwdpb.CounterId_COUNTER_ID_RX_PACKETS:
						// TODO(wenbli): Perhaps should make a logging version of ygnmi.
						if _, err := gnmiclient.Replace(ctx, ni.c, ocpath.Root().Interface(intfName).Counters().InPkts().State(), counter.Value); err != nil {
							log.Errorf("interface handler: %v", err)
						}
					case fwdpb.CounterId_COUNTER_ID_TX_PACKETS:
						if _, err := gnmiclient.Replace(ctx, ni.c, ocpath.Root().Interface(intfName).Counters().OutPkts().State(), counter.Value); err != nil {
							log.Errorf("interface handler: %v", err)
						}
					}
				}
			}
		}
	}()
}

// reconcile compares the interface config with state and modifies state to match config.
func (ni *Interface) reconcile(config *oc.Interface) {
	ni.stateMu.RLock()
	defer ni.stateMu.RUnlock()

	tapName := engine.IntfNameToTapName(config.GetName())
	state := ni.getOrCreateInterface(config.GetName())

	if config.GetOrCreateEthernet().MacAddress != nil {
		if config.GetEthernet().GetMacAddress() != state.GetEthernet().GetMacAddress() {
			log.V(1).Infof("setting interface %s hw-addr %q", tapName, config.GetEthernet().GetMacAddress())
			if err := kernel.SetInterfaceHWAddr(config.GetName(), config.GetEthernet().GetMacAddress()); err != nil {
				log.Warningf("Failed to set mac address of port: %v", err)
			}
		}
	} else {
		// Deleting the configured MAC address means it should be the system-assigned MAC address, as detailed in the OpenConfig schema.
		// https://openconfig.net/projects/models/schemadocs/yangdoc/openconfig-interfaces.html#interfaces-interface-ethernet-state-mac-address
		if state.GetEthernet().GetHwMacAddress() != state.GetEthernet().GetMacAddress() {
			log.V(1).Infof("resetting interface %s hw-addr %q", tapName, state.GetEthernet().GetHwMacAddress())
			if err := kernel.SetInterfaceHWAddr(config.GetName(), state.GetEthernet().GetHwMacAddress()); err != nil {
				log.Warningf("Failed to set mac address of port: %v", err)
			}
		}
	}

	if config.GetOrCreateSubinterface(0).Enabled != nil {
		if state.GetOrCreateSubinterface(0).Enabled == nil || config.GetSubinterface(0).GetEnabled() != state.GetSubinterface(0).GetEnabled() {
			log.V(1).Infof("setting interface %s enabled %t", tapName, config.GetSubinterface(0).GetEnabled())
			if err := kernel.SetInterfaceState(tapName, config.GetSubinterface(0).GetEnabled()); err != nil {
				log.Warningf("Failed to set state address of port: %v", err)
			}
		}
	}

	type prefixPair struct {
		cfgIP, stateIP *string
		cfgPL, statePL *uint8
	}

	// Get all state IPs and their corresponding config IPs (if they exist).
	var interfacePairs []*prefixPair
	for _, addr := range state.GetOrCreateSubinterface(0).GetOrCreateIpv4().Address {
		pair := &prefixPair{
			stateIP: addr.Ip,
			statePL: addr.PrefixLength,
		}
		if pairAddr := config.GetSubinterface(0).GetIpv4().GetAddress(addr.GetIp()); pairAddr != nil {
			pair.cfgIP = pairAddr.Ip
			pair.cfgPL = pairAddr.PrefixLength
		}
		interfacePairs = append(interfacePairs, pair)
	}
	for _, addr := range state.GetOrCreateSubinterface(0).GetOrCreateIpv6().Address {
		pair := &prefixPair{
			stateIP: addr.Ip,
			statePL: addr.PrefixLength,
		}
		if pairAddr := config.GetSubinterface(0).GetIpv6().GetAddress(addr.GetIp()); pairAddr != nil {
			pair.cfgIP = pairAddr.Ip
			pair.cfgPL = pairAddr.PrefixLength
		}
		interfacePairs = append(interfacePairs, pair)
	}

	// Get all config IPs and their corresponding state IPs (if they exist).
	for _, addr := range config.GetOrCreateSubinterface(0).GetOrCreateIpv4().Address {
		pair := &prefixPair{
			cfgIP: addr.Ip,
			cfgPL: addr.PrefixLength,
		}
		if pairAddr := state.GetSubinterface(0).GetIpv4().GetAddress(addr.GetIp()); pairAddr != nil {
			pair.stateIP = pairAddr.Ip
			pair.statePL = pairAddr.PrefixLength
		}
		interfacePairs = append(interfacePairs, pair)
	}
	for _, addr := range config.GetOrCreateSubinterface(0).GetOrCreateIpv6().Address {
		pair := &prefixPair{
			cfgIP: addr.Ip,
			cfgPL: addr.PrefixLength,
		}
		if pairAddr := state.GetSubinterface(0).GetIpv6().GetAddress(addr.GetIp()); pairAddr != nil {
			pair.stateIP = pairAddr.Ip
			pair.statePL = pairAddr.PrefixLength
		}
		interfacePairs = append(interfacePairs, pair)
	}

	for _, pair := range interfacePairs {
		// If an IP exists in state, but not in config, remove the IP.
		if (pair.stateIP != nil && pair.statePL != nil) && (pair.cfgIP == nil && pair.cfgPL == nil) {
			log.V(1).Infof("Delete Config IP: %v, Config PL: %v. State IP: %v, State PL: %v", pair.cfgIP, pair.cfgPL, *pair.stateIP, *pair.statePL)
			log.V(2).Infof("deleting interface %s ip %s/%d", tapName, *pair.stateIP, *pair.statePL)
			if err := kernel.DeleteInterfaceIP(tapName, *pair.stateIP, int(*pair.statePL)); err != nil {
				log.Warningf("Failed to set ip address of port: %v", err)
			}
		}
		// If an IP exists in config, but not in state (or state is different) add the IP.
		if (pair.cfgIP != nil && pair.cfgPL != nil) && (pair.stateIP == nil || *pair.statePL != *pair.cfgPL) {
			log.V(1).Infof("Set Config IP: %v, Config PL: %v. State IP: %v, State PL: %v", *pair.cfgIP, *pair.cfgPL, pair.stateIP, pair.statePL)
			log.V(2).Infof("setting interface %s ip %s/%d", tapName, *pair.cfgIP, *pair.cfgPL)
			if err := kernel.SetInterfaceIP(tapName, *pair.cfgIP, int(*pair.cfgPL)); err != nil {
				log.Warningf("Failed to set ip address of port: %v", err)
			}
		}
	}
}

// getOrCreateInterface returns the state interface from the cache.
func (ni *Interface) getOrCreateInterface(iface string) *oc.Interface {
	if _, ok := ni.state[iface]; !ok {
		ni.state[iface] = &oc.Interface{
			Name: &iface,
		}
	}
	return ni.state[iface]
}

// handleLinkUpdate modifies the state based on changes to link state.
func (ni *Interface) handleLinkUpdate(ctx context.Context, lu *netlink.LinkUpdate) {
	ni.stateMu.Lock()
	defer ni.stateMu.Unlock()

	log.V(1).Infof("handling link update for %s", lu.Attrs().Name)

	if !engine.IsTap(lu.Attrs().Name) {
		return
	}

	modelName := engine.TapNameToIntfName(lu.Attrs().Name)
	iface := ni.getOrCreateInterface(modelName)
	if err := engine.UpdatePortSrcMAC(ctx, ni.fwd, modelName, lu.Attrs().HardwareAddr); err != nil {
		log.Warningf("failed to update src mac: %v", err)
	}
	iface.GetOrCreateEthernet().MacAddress = ygot.String(lu.Attrs().HardwareAddr.String())

	iface.Ifindex = ygot.Uint32(uint32(lu.Attrs().Index))
	iface.Enabled = ygot.Bool(lu.Attrs().Flags&net.FlagUp != 0)
	iface.AdminStatus = oc.Interface_AdminStatus_DOWN
	if *iface.Enabled {
		iface.AdminStatus = oc.Interface_AdminStatus_UP
	}

	// TODO: handle other states.
	var operStatus oc.E_Interface_OperStatus
	switch lu.Attrs().OperState {
	case netlink.OperDown:
		operStatus = oc.Interface_OperStatus_DOWN
	case netlink.OperUp, netlink.OperUnknown: // TAP interface may be unknown state because the dataplane doesn't bind to its fd, so treat unknown as up.
		operStatus = oc.Interface_OperStatus_UP
	}
	iface.OperStatus = operStatus

	sb := &ygnmi.SetBatch{}

	gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Ethernet().MacAddress().State(), *iface.Ethernet.MacAddress)
	gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Ifindex().State(), *iface.Ifindex)
	gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Enabled().State(), *iface.Enabled)
	gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).OperStatus().State(), iface.OperStatus)
	if _, err := sb.Set(ctx, ni.c); err != nil {
		log.Warningf("failed to set link status: %v", err)
	}
}

// handleAddrUpdate modifies the state based on changes to addresses.
func (ni *Interface) handleAddrUpdate(ctx context.Context, au *netlink.AddrUpdate) {
	ni.stateMu.Lock()
	defer ni.stateMu.Unlock()
	name := ni.idxToName[au.LinkIndex]
	if !engine.IsTap(name) || name == "" {
		return
	}

	sb := &ygnmi.SetBatch{}
	modelName := engine.TapNameToIntfName(name)
	sub := ni.getOrCreateInterface(modelName).GetOrCreateSubinterface(0)

	ip := au.LinkAddress.IP.String()
	pl, _ := au.LinkAddress.Mask.Size()
	isV4 := au.LinkAddress.IP.To4() != nil
	log.V(1).Infof("handling addr update for %s ip %v pl %v", name, ip, pl)
	if au.NewAddr {
		var ipBytes []byte
		if isV4 {
			ipBytes = au.LinkAddress.IP.To4()
			sub.GetOrCreateIpv4().GetOrCreateAddress(ip).PrefixLength = ygot.Uint8(uint8(pl))
			gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Address(ip).Ip().State(), au.LinkAddress.IP.String())
			gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Address(ip).PrefixLength().State(), uint8(pl))
		} else {
			ipBytes = au.LinkAddress.IP.To16()
			sub.GetOrCreateIpv6().GetOrCreateAddress(ip).PrefixLength = ygot.Uint8(uint8(pl))
			gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Address(ip).Ip().State(), au.LinkAddress.IP.String())
			gnmiclient.BatchUpdate(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Address(ip).PrefixLength().State(), uint8(pl))
		}
		// Forward all packets destined to this interface to the corresponding TAP interface.
		if err := engine.AddLayer3PuntRule(ctx, ni.fwd, modelName, ipBytes); err != nil {
			log.Warningf("failed to add layer3 punt rule: %v", err)
		}
	} else {
		if isV4 {
			sub.GetOrCreateIpv4().DeleteAddress(ip)
			gnmiclient.BatchDelete(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Address(ip).State())
		} else {
			sub.GetOrCreateIpv6().DeleteAddress(ip)
			gnmiclient.BatchDelete(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Address(ip).State())
		}
	}
	if _, err := sb.Set(ctx, ni.c); err != nil {
		log.Warningf("failed to set link status: %v", err)
	}
}

// handleNeighborUpdate modifies the state based on changes to the neighbor.
func (ni *Interface) handleNeighborUpdate(ctx context.Context, nu *netlink.NeighUpdate) {
	ni.stateMu.Lock()
	defer ni.stateMu.Unlock()
	log.V(1).Infof("handling neighbor update for %s on %d", nu.IP.String(), nu.LinkIndex)

	name := ni.idxToName[nu.LinkIndex]
	if name == "" {
		return
	}
	sb := &ygnmi.SetBatch{}
	modelName := engine.TapNameToIntfName(name)
	sub := ni.getOrCreateInterface(modelName).GetOrCreateSubinterface(0)

	switch nu.Type {
	case unix.RTM_DELNEIGH:
		if err := engine.RemoveNeighbor(ctx, ni.fwd, ipToBytes(nu.IP)); err != nil {
			log.Warningf("failed to add neighbor to dataplane: %v", err)
			return
		}
		if nu.Family == unix.AF_INET6 {
			sub.GetOrCreateIpv6().DeleteNeighbor(nu.IP.String())
			gnmiclient.BatchDelete(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Neighbor(nu.IP.String()).State())
		} else {
			sub.GetOrCreateIpv4().DeleteNeighbor(nu.IP.String())
			gnmiclient.BatchDelete(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Neighbor(nu.IP.String()).State())
		}
	case unix.RTM_NEWNEIGH:
		if len(nu.HardwareAddr) == 0 {
			log.Info("skipping neighbor update with no hwaddr")
			return
		}
		if err := engine.AddNeighbor(ctx, ni.fwd, ipToBytes(nu.IP), nu.HardwareAddr); err != nil {
			log.Warningf("failed to add neighbor to dataplane: %v", err)
			return
		}
		if nu.Family == unix.AF_INET6 {
			neigh := sub.GetOrCreateIpv6().GetOrCreateNeighbor(nu.IP.String())
			neigh.LinkLayerAddress = ygot.String(nu.HardwareAddr.String())
			if nu.Flags&unix.NUD_PERMANENT != 0 {
				neigh.Origin = oc.IfIp_NeighborOrigin_STATIC
			} else {
				neigh.Origin = oc.IfIp_NeighborOrigin_DYNAMIC
			}
			gnmiclient.BatchReplace(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Neighbor(nu.IP.String()).Ip().State(), neigh.GetIp())
			gnmiclient.BatchReplace(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Neighbor(nu.IP.String()).LinkLayerAddress().State(), neigh.GetLinkLayerAddress())
			gnmiclient.BatchReplace(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv6().Neighbor(nu.IP.String()).Origin().State(), neigh.GetOrigin())
		} else {
			neigh := sub.GetOrCreateIpv4().GetOrCreateNeighbor(nu.IP.String())
			neigh.LinkLayerAddress = ygot.String(nu.HardwareAddr.String())
			if nu.Flags&unix.NUD_PERMANENT != 0 {
				neigh.Origin = oc.IfIp_NeighborOrigin_STATIC
			} else {
				neigh.Origin = oc.IfIp_NeighborOrigin_DYNAMIC
			}
			gnmiclient.BatchReplace(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Neighbor(nu.IP.String()).Ip().State(), neigh.GetIp())
			gnmiclient.BatchReplace(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Neighbor(nu.IP.String()).LinkLayerAddress().State(), neigh.GetLinkLayerAddress())
			gnmiclient.BatchReplace(sb, ocpath.Root().Interface(modelName).Subinterface(0).Ipv4().Neighbor(nu.IP.String()).Origin().State(), neigh.GetOrigin())
		}
	default:
		log.Warningf("unknown neigh update type: %v", nu.Type)
	}

	if _, err := sb.Set(ctx, ni.c); err != nil {
		log.Warningf("failed to set link status: %v", err)
	}
}

// setupPorts creates the dataplane ports and TAP interfaces for all interfaces on the device.
func (ni *Interface) setupPorts(ctx context.Context) error {
	ifs, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, i := range ifs {
		// Skip loopback, k8s pod interface, and tap interfaces.
		if i.Name == "lo" || i.Name == "eth0" || engine.IsTap(i.Name) {
			continue
		}
		fd, err := kernel.CreateTAP(engine.IntfNameToTapName(i.Name))
		if err != nil {
			return fmt.Errorf("failed to create tap port %q: %w", engine.IntfNameToTapName(i.Name), err)
		}
		tap, err := net.InterfaceByName(engine.IntfNameToTapName(i.Name))
		if err != nil {
			return fmt.Errorf("failed to find tap interface %q: %w", engine.IntfNameToTapName(i.Name), err)
		}
		ni.idxToName[i.Index] = i.Name
		ni.idxToName[tap.Index] = tap.Name
		if err := engine.CreateLocalPort(ctx, ni.fwd, tap.Name, fd); err != nil {
			return fmt.Errorf("failed to create internal port %q: %w", tap.Name, err)
		}
		if err := engine.CreateExternalPort(ctx, ni.fwd, i.Name); err != nil {
			return fmt.Errorf("failed to create external port %q: %w", i.Name, err)
		}
		// Make port only reply to IPs it have
		if err := os.WriteFile(fmt.Sprintf("/proc/sys/net/ipv4/conf/%s/arp_ignore", i.Name), []byte("2"), 0600); err != nil {
			return fmt.Errorf("failed to set arp_ignore to 2: %v", err)
		}
		if err := engine.UpdatePortSrcMAC(ctx, ni.fwd, i.Name, tap.HardwareAddr); err != nil {
			return fmt.Errorf("failed to update MAC address for port %q: %w", i.Name, err)
		}
		ni.getOrCreateInterface(i.Name).GetOrCreateEthernet().SetHwMacAddress(tap.HardwareAddr.String())
		ni.getOrCreateInterface(i.Name).GetOrCreateEthernet().SetMacAddress(tap.HardwareAddr.String())
		if _, err := gnmiclient.Update(ctx, ni.c, ocpath.Root().Interface(i.Name).Ethernet().HwMacAddress().State(), tap.HardwareAddr.String()); err != nil {
			return fmt.Errorf("failed to set hw addr of interface %q: %v", tap.Name, err)
		}
		if _, err := gnmiclient.Update(ctx, ni.c, ocpath.Root().Interface(i.Name).Ethernet().MacAddress().State(), tap.HardwareAddr.String()); err != nil {
			return fmt.Errorf("failed to set hw addr of interface %q: %v", tap.Name, err)
		}
	}
	return nil
}

// ipToBytes converts a net.IP to a slice of bytes of the correct length (4 for IPv4, 16 for IPv6).
func ipToBytes(ip net.IP) []byte {
	if ip.To4() != nil {
		return ip.To4()
	}
	return ip.To16()
}
