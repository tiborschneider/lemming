/*
Package lacp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.8.7: (ygot: v0.29.11)
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
	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Lacp_Interface_Member_Counters_LacpInPktsPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts YANG schema element.
type Lacp_Interface_Member_Counters_LacpInPktsPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_Counters_LacpInPktsPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts YANG schema element.
type Lacp_Interface_Member_Counters_LacpInPktsPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-in-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts"
func (n *Lacp_Interface_Member_Counters_LacpInPktsPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-in-pkts"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpInPkts
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
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
//	Path from parent:     "lacp-in-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts"
func (n *Lacp_Interface_Member_Counters_LacpInPktsPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-in-pkts"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpInPkts
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// Lacp_Interface_Member_Counters_LacpOutPktsPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts YANG schema element.
type Lacp_Interface_Member_Counters_LacpOutPktsPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_Counters_LacpOutPktsPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts YANG schema element.
type Lacp_Interface_Member_Counters_LacpOutPktsPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-out-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts"
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-out-pkts"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpOutPkts
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
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
//	Path from parent:     "lacp-out-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts"
func (n *Lacp_Interface_Member_Counters_LacpOutPktsPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-out-pkts"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpOutPkts
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// Lacp_Interface_Member_Counters_LacpRxErrorsPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors YANG schema element.
type Lacp_Interface_Member_Counters_LacpRxErrorsPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_Counters_LacpRxErrorsPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors YANG schema element.
type Lacp_Interface_Member_Counters_LacpRxErrorsPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-rx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors"
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-rx-errors"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpRxErrors
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
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
//	Path from parent:     "lacp-rx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors"
func (n *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-rx-errors"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpRxErrors
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}

// Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions YANG schema element.
type Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions YANG schema element.
type Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-timeout-transitions"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions"
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-timeout-transitions"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpTimeoutTransitions
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
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
//	Path from parent:     "lacp-timeout-transitions"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions"
func (n *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-timeout-transitions"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpTimeoutTransitions
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface_Member_Counters) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
	)
}
