/*
Package lacp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.9.0: (ygot: v0.29.12)
using the following YANG input files:
  - gnsi/yang/gnsi-telemetry.yang
  - public/release/models/acl/openconfig-acl.yang
  - public/release/models/acl/openconfig-packet-match.yang
  - public/release/models/aft/openconfig-aft.yang
  - public/release/models/bfd/openconfig-bfd.yang
  - public/release/models/bgp/openconfig-bgp-policy.yang
  - public/release/models/bgp/openconfig-bgp-types.yang
  - public/release/models/interfaces/openconfig-if-aggregate.yang
  - public/release/models/interfaces/openconfig-if-ethernet.yang
  - public/release/models/interfaces/openconfig-if-ip-ext.yang
  - public/release/models/interfaces/openconfig-if-ip.yang
  - public/release/models/interfaces/openconfig-interfaces.yang
  - public/release/models/isis/openconfig-isis.yang
  - public/release/models/lacp/openconfig-lacp.yang
  - public/release/models/lldp/openconfig-lldp-types.yang
  - public/release/models/lldp/openconfig-lldp.yang
  - public/release/models/local-routing/openconfig-local-routing.yang
  - public/release/models/mpls/openconfig-mpls-types.yang
  - public/release/models/multicast/openconfig-pim.yang
  - public/release/models/network-instance/openconfig-network-instance.yang
  - public/release/models/openconfig-extensions.yang
  - public/release/models/optical-transport/openconfig-transport-types.yang
  - public/release/models/ospf/openconfig-ospfv2.yang
  - public/release/models/platform/openconfig-platform-cpu.yang
  - public/release/models/platform/openconfig-platform-integrated-circuit.yang
  - public/release/models/platform/openconfig-platform-software.yang
  - public/release/models/platform/openconfig-platform-transceiver.yang
  - public/release/models/platform/openconfig-platform.yang
  - public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  - public/release/models/policy/openconfig-policy-types.yang
  - public/release/models/qos/openconfig-qos-elements.yang
  - public/release/models/qos/openconfig-qos-interfaces.yang
  - public/release/models/qos/openconfig-qos-types.yang
  - public/release/models/qos/openconfig-qos.yang
  - public/release/models/rib/openconfig-rib-bgp.yang
  - public/release/models/segment-routing/openconfig-segment-routing-types.yang
  - public/release/models/system/openconfig-system.yang
  - public/release/models/types/openconfig-inet-types.yang
  - public/release/models/types/openconfig-types.yang
  - public/release/models/types/openconfig-yang-types.yang
  - public/release/models/vlan/openconfig-vlan.yang
  - public/third_party/ietf/iana-if-type.yang
  - public/third_party/ietf/ietf-inet-types.yang
  - public/third_party/ietf/ietf-interfaces.yang
  - public/third_party/ietf/ietf-yang-types.yang
  - yang/openconfig-bgp-gue.yang

Imported modules were sourced from:
  - public/release/models/...
  - public/third_party/ietf/...
  - gnsi/...
*/
package lacp

import (
	"reflect"

	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Lacp_Interface_Member_PartnerKeyPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/partner-key YANG schema element.
type Lacp_Interface_Member_PartnerKeyPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_PartnerKeyPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/partner-key YANG schema element.
type Lacp_Interface_Member_PartnerKeyPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/partner-key"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/partner-key"
func (n *Lacp_Interface_Member_PartnerKeyPath) State() ygnmi.SingletonQuery[uint16] {
	return ygnmi.NewSingletonQuery[uint16](
		"Lacp_Interface_Member",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "partner-key"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).PartnerKey
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/partner-key"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/partner-key"
func (n *Lacp_Interface_Member_PartnerKeyPathAny) State() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp_Interface_Member",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "partner-key"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).PartnerKey
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Lacp_Interface_Member_PartnerPortNumPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/partner-port-num YANG schema element.
type Lacp_Interface_Member_PartnerPortNumPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_PartnerPortNumPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/partner-port-num YANG schema element.
type Lacp_Interface_Member_PartnerPortNumPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/partner-port-num"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/partner-port-num"
func (n *Lacp_Interface_Member_PartnerPortNumPath) State() ygnmi.SingletonQuery[uint16] {
	return ygnmi.NewSingletonQuery[uint16](
		"Lacp_Interface_Member",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "partner-port-num"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).PartnerPortNum
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/partner-port-num"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/partner-port-num"
func (n *Lacp_Interface_Member_PartnerPortNumPathAny) State() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp_Interface_Member",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "partner-port-num"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).PartnerPortNum
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Lacp_Interface_Member_PortNumPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/port-num YANG schema element.
type Lacp_Interface_Member_PortNumPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_PortNumPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/port-num YANG schema element.
type Lacp_Interface_Member_PortNumPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/port-num"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/port-num"
func (n *Lacp_Interface_Member_PortNumPath) State() ygnmi.SingletonQuery[uint16] {
	return ygnmi.NewSingletonQuery[uint16](
		"Lacp_Interface_Member",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "port-num"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).PortNum
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/port-num"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/port-num"
func (n *Lacp_Interface_Member_PortNumPathAny) State() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp_Interface_Member",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "port-num"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).PortNum
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// Lacp_Interface_Member_SynchronizationPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization YANG schema element.
type Lacp_Interface_Member_SynchronizationPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_SynchronizationPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/synchronization YANG schema element.
type Lacp_Interface_Member_SynchronizationPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/synchronization"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/synchronization"
func (n *Lacp_Interface_Member_SynchronizationPath) State() ygnmi.SingletonQuery[oc.E_Lacp_LacpSynchronizationType] {
	return ygnmi.NewSingletonQuery[oc.E_Lacp_LacpSynchronizationType](
		"Lacp_Interface_Member",
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "synchronization"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpSynchronizationType, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).Synchronization
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/synchronization"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/synchronization"
func (n *Lacp_Interface_Member_SynchronizationPathAny) State() ygnmi.WildcardQuery[oc.E_Lacp_LacpSynchronizationType] {
	return ygnmi.NewWildcardQuery[oc.E_Lacp_LacpSynchronizationType](
		"Lacp_Interface_Member",
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "synchronization"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpSynchronizationType, bool) {
			ret := gs.(*oc.Lacp_Interface_Member).Synchronization
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		nil,
	)
}
