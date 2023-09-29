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

// Lacp_Interface_Member_Counters_LacpTxErrorsPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors YANG schema element.
type Lacp_Interface_Member_Counters_LacpTxErrorsPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_Counters_LacpTxErrorsPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors YANG schema element.
type Lacp_Interface_Member_Counters_LacpTxErrorsPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-tx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors"
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-tx-errors"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpTxErrors
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
//	Path from parent:     "lacp-tx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors"
func (n *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-tx-errors"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpTxErrors
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

// Lacp_Interface_Member_Counters_LacpUnknownErrorsPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors YANG schema element.
type Lacp_Interface_Member_Counters_LacpUnknownErrorsPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors YANG schema element.
type Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-unknown-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors"
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-unknown-errors"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpUnknownErrors
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
//	Path from parent:     "lacp-unknown-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors"
func (n *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lacp_Interface_Member_Counters",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"lacp-unknown-errors"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lacp_Interface_Member_Counters).LacpUnknownErrors
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

// Lacp_Interface_Member_CountersPath represents the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters YANG schema element.
type Lacp_Interface_Member_CountersPath struct {
	*ygnmi.NodePath
}

// Lacp_Interface_Member_CountersPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/members/member/state/counters YANG schema element.
type Lacp_Interface_Member_CountersPathAny struct {
	*ygnmi.NodePath
}

// LacpErrors (leaf): Number of LACPDU illegal packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-errors"
func (n *Lacp_Interface_Member_CountersPath) LacpErrors() *Lacp_Interface_Member_Counters_LacpErrorsPath {
	ps := &Lacp_Interface_Member_Counters_LacpErrorsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpErrors (leaf): Number of LACPDU illegal packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-errors"
func (n *Lacp_Interface_Member_CountersPathAny) LacpErrors() *Lacp_Interface_Member_Counters_LacpErrorsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpErrorsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpInPkts (leaf): Number of LACPDUs received
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-in-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts"
func (n *Lacp_Interface_Member_CountersPath) LacpInPkts() *Lacp_Interface_Member_Counters_LacpInPktsPath {
	ps := &Lacp_Interface_Member_Counters_LacpInPktsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-in-pkts"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpInPkts (leaf): Number of LACPDUs received
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-in-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-in-pkts"
func (n *Lacp_Interface_Member_CountersPathAny) LacpInPkts() *Lacp_Interface_Member_Counters_LacpInPktsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpInPktsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-in-pkts"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpOutPkts (leaf): Number of LACPDUs transmitted
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-out-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts"
func (n *Lacp_Interface_Member_CountersPath) LacpOutPkts() *Lacp_Interface_Member_Counters_LacpOutPktsPath {
	ps := &Lacp_Interface_Member_Counters_LacpOutPktsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-out-pkts"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpOutPkts (leaf): Number of LACPDUs transmitted
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-out-pkts"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-out-pkts"
func (n *Lacp_Interface_Member_CountersPathAny) LacpOutPkts() *Lacp_Interface_Member_Counters_LacpOutPktsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpOutPktsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-out-pkts"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpRxErrors (leaf): Number of LACPDU receive packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-rx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors"
func (n *Lacp_Interface_Member_CountersPath) LacpRxErrors() *Lacp_Interface_Member_Counters_LacpRxErrorsPath {
	ps := &Lacp_Interface_Member_Counters_LacpRxErrorsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-rx-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpRxErrors (leaf): Number of LACPDU receive packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-rx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-rx-errors"
func (n *Lacp_Interface_Member_CountersPathAny) LacpRxErrors() *Lacp_Interface_Member_Counters_LacpRxErrorsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpRxErrorsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-rx-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpTimeoutTransitions (leaf): Number of times the LACP state has transitioned
// with a timeout since the time the device restarted
// or the interface was brought up, whichever is most
// recent. The last state change of the LACP timeout
// is defined as what is reported as the operating state
// to the system. The state change is both a timeout
// event and when the timeout event is no longer active.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-timeout-transitions"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions"
func (n *Lacp_Interface_Member_CountersPath) LacpTimeoutTransitions() *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath {
	ps := &Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-timeout-transitions"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpTimeoutTransitions (leaf): Number of times the LACP state has transitioned
// with a timeout since the time the device restarted
// or the interface was brought up, whichever is most
// recent. The last state change of the LACP timeout
// is defined as what is reported as the operating state
// to the system. The state change is both a timeout
// event and when the timeout event is no longer active.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-timeout-transitions"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-timeout-transitions"
func (n *Lacp_Interface_Member_CountersPathAny) LacpTimeoutTransitions() *Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpTimeoutTransitionsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-timeout-transitions"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpTxErrors (leaf): Number of LACPDU transmit packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-tx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors"
func (n *Lacp_Interface_Member_CountersPath) LacpTxErrors() *Lacp_Interface_Member_Counters_LacpTxErrorsPath {
	ps := &Lacp_Interface_Member_Counters_LacpTxErrorsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-tx-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpTxErrors (leaf): Number of LACPDU transmit packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-tx-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-tx-errors"
func (n *Lacp_Interface_Member_CountersPathAny) LacpTxErrors() *Lacp_Interface_Member_Counters_LacpTxErrorsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpTxErrorsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-tx-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpUnknownErrors (leaf): Number of LACPDU unknown packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-unknown-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors"
func (n *Lacp_Interface_Member_CountersPath) LacpUnknownErrors() *Lacp_Interface_Member_Counters_LacpUnknownErrorsPath {
	ps := &Lacp_Interface_Member_Counters_LacpUnknownErrorsPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-unknown-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LacpUnknownErrors (leaf): Number of LACPDU unknown packet errors
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp-unknown-errors"
//	Path from root:       "/lacp/interfaces/interface/members/member/state/counters/lacp-unknown-errors"
func (n *Lacp_Interface_Member_CountersPathAny) LacpUnknownErrors() *Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny {
	ps := &Lacp_Interface_Member_Counters_LacpUnknownErrorsPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp-unknown-errors"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// State returns a Query that can be used in gNMI operations.
func (n *Lacp_Interface_Member_CountersPath) State() ygnmi.SingletonQuery[*oc.Lacp_Interface_Member_Counters] {
	return ygnmi.NewSingletonQuery[*oc.Lacp_Interface_Member_Counters](
		"Lacp_Interface_Member_Counters",
		true,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
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
func (n *Lacp_Interface_Member_CountersPathAny) State() ygnmi.WildcardQuery[*oc.Lacp_Interface_Member_Counters] {
	return ygnmi.NewWildcardQuery[*oc.Lacp_Interface_Member_Counters](
		"Lacp_Interface_Member_Counters",
		true,
		false,
		false,
		true,
		false,
		n,
		nil,
		nil,
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
