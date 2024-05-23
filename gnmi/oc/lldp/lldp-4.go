/*
Package lldp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.11.1: (ygot: v0.29.18)
using the following YANG input files:
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
*/
package lldp

import (
	"reflect"

	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Lldp_Interface_Counters_TlvUnknownPath represents the /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown YANG schema element.
type Lldp_Interface_Counters_TlvUnknownPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_Interface_Counters_TlvUnknownPathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/state/counters/tlv-unknown YANG schema element.
type Lldp_Interface_Counters_TlvUnknownPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "tlv-unknown"
//	Path from root:       "/lldp/interfaces/interface/state/counters/tlv-unknown"
func (n *Lldp_Interface_Counters_TlvUnknownPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lldp_Interface_Counters",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"tlv-unknown"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp_Interface_Counters).TlvUnknown
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Counters) },
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
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "tlv-unknown"
//	Path from root:       "/lldp/interfaces/interface/state/counters/tlv-unknown"
func (n *Lldp_Interface_Counters_TlvUnknownPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lldp_Interface_Counters",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"tlv-unknown"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp_Interface_Counters).TlvUnknown
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Counters) },
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

// Lldp_Interface_CountersPath represents the /openconfig-lldp/lldp/interfaces/interface/state/counters YANG schema element.
type Lldp_Interface_CountersPath struct {
	*ygnmi.NodePath
}

// Lldp_Interface_CountersPathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/state/counters YANG schema element.
type Lldp_Interface_CountersPathAny struct {
	*ygnmi.NodePath
}

// FrameDiscard (leaf): The number of LLDP frames received and discarded.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-discard"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-discard"
func (n *Lldp_Interface_CountersPath) FrameDiscard() *Lldp_Interface_Counters_FrameDiscardPath {
	ps := &Lldp_Interface_Counters_FrameDiscardPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-discard"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameDiscard (leaf): The number of LLDP frames received and discarded.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-discard"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-discard"
func (n *Lldp_Interface_CountersPathAny) FrameDiscard() *Lldp_Interface_Counters_FrameDiscardPathAny {
	ps := &Lldp_Interface_Counters_FrameDiscardPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-discard"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameErrorIn (leaf): The number of LLDP frames received with errors.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-error-in"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-error-in"
func (n *Lldp_Interface_CountersPath) FrameErrorIn() *Lldp_Interface_Counters_FrameErrorInPath {
	ps := &Lldp_Interface_Counters_FrameErrorInPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-error-in"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameErrorIn (leaf): The number of LLDP frames received with errors.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-error-in"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-error-in"
func (n *Lldp_Interface_CountersPathAny) FrameErrorIn() *Lldp_Interface_Counters_FrameErrorInPathAny {
	ps := &Lldp_Interface_Counters_FrameErrorInPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-error-in"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameErrorOut (leaf): The number of frame transmit errors on the
// interface.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-error-out"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-error-out"
func (n *Lldp_Interface_CountersPath) FrameErrorOut() *Lldp_Interface_Counters_FrameErrorOutPath {
	ps := &Lldp_Interface_Counters_FrameErrorOutPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-error-out"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameErrorOut (leaf): The number of frame transmit errors on the
// interface.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-error-out"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-error-out"
func (n *Lldp_Interface_CountersPathAny) FrameErrorOut() *Lldp_Interface_Counters_FrameErrorOutPathAny {
	ps := &Lldp_Interface_Counters_FrameErrorOutPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-error-out"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameIn (leaf): The number of lldp frames received.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-in"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-in"
func (n *Lldp_Interface_CountersPath) FrameIn() *Lldp_Interface_Counters_FrameInPath {
	ps := &Lldp_Interface_Counters_FrameInPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-in"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameIn (leaf): The number of lldp frames received.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-in"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-in"
func (n *Lldp_Interface_CountersPathAny) FrameIn() *Lldp_Interface_Counters_FrameInPathAny {
	ps := &Lldp_Interface_Counters_FrameInPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-in"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameOut (leaf): The number of frames transmitted out.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-out"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-out"
func (n *Lldp_Interface_CountersPath) FrameOut() *Lldp_Interface_Counters_FrameOutPath {
	ps := &Lldp_Interface_Counters_FrameOutPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-out"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// FrameOut (leaf): The number of frames transmitted out.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "frame-out"
//	Path from root:       "/lldp/interfaces/interface/state/counters/frame-out"
func (n *Lldp_Interface_CountersPathAny) FrameOut() *Lldp_Interface_Counters_FrameOutPathAny {
	ps := &Lldp_Interface_Counters_FrameOutPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"frame-out"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LastClear (leaf): Indicates the last time the counters were
// cleared.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "last-clear"
//	Path from root:       "/lldp/interfaces/interface/state/counters/last-clear"
func (n *Lldp_Interface_CountersPath) LastClear() *Lldp_Interface_Counters_LastClearPath {
	ps := &Lldp_Interface_Counters_LastClearPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"last-clear"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// LastClear (leaf): Indicates the last time the counters were
// cleared.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "last-clear"
//	Path from root:       "/lldp/interfaces/interface/state/counters/last-clear"
func (n *Lldp_Interface_CountersPathAny) LastClear() *Lldp_Interface_Counters_LastClearPathAny {
	ps := &Lldp_Interface_Counters_LastClearPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"last-clear"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// TlvDiscard (leaf): The number of TLV frames received and discarded.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "tlv-discard"
//	Path from root:       "/lldp/interfaces/interface/state/counters/tlv-discard"
func (n *Lldp_Interface_CountersPath) TlvDiscard() *Lldp_Interface_Counters_TlvDiscardPath {
	ps := &Lldp_Interface_Counters_TlvDiscardPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"tlv-discard"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// TlvDiscard (leaf): The number of TLV frames received and discarded.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "tlv-discard"
//	Path from root:       "/lldp/interfaces/interface/state/counters/tlv-discard"
func (n *Lldp_Interface_CountersPathAny) TlvDiscard() *Lldp_Interface_Counters_TlvDiscardPathAny {
	ps := &Lldp_Interface_Counters_TlvDiscardPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"tlv-discard"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// TlvUnknown (leaf): The number of frames received with unknown TLV.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "tlv-unknown"
//	Path from root:       "/lldp/interfaces/interface/state/counters/tlv-unknown"
func (n *Lldp_Interface_CountersPath) TlvUnknown() *Lldp_Interface_Counters_TlvUnknownPath {
	ps := &Lldp_Interface_Counters_TlvUnknownPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"tlv-unknown"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// TlvUnknown (leaf): The number of frames received with unknown TLV.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "tlv-unknown"
//	Path from root:       "/lldp/interfaces/interface/state/counters/tlv-unknown"
func (n *Lldp_Interface_CountersPathAny) TlvUnknown() *Lldp_Interface_Counters_TlvUnknownPathAny {
	ps := &Lldp_Interface_Counters_TlvUnknownPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"tlv-unknown"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// State returns a Query that can be used in gNMI operations.
func (n *Lldp_Interface_CountersPath) State() ygnmi.SingletonQuery[*oc.Lldp_Interface_Counters] {
	return ygnmi.NewSingletonQuery[*oc.Lldp_Interface_Counters](
		"Lldp_Interface_Counters",
		true,
		false,
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
func (n *Lldp_Interface_CountersPathAny) State() ygnmi.WildcardQuery[*oc.Lldp_Interface_Counters] {
	return ygnmi.NewWildcardQuery[*oc.Lldp_Interface_Counters](
		"Lldp_Interface_Counters",
		true,
		false,
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

// Lldp_Interface_Neighbor_AgePath represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age YANG schema element.
type Lldp_Interface_Neighbor_AgePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_Interface_Neighbor_AgePathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/age YANG schema element.
type Lldp_Interface_Neighbor_AgePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/age"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/age"
func (n *Lldp_Interface_Neighbor_AgePath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "age"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).Age
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/age"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/age"
func (n *Lldp_Interface_Neighbor_AgePathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "age"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).Age
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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

// Lldp_Interface_Neighbor_ChassisIdPath represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id YANG schema element.
type Lldp_Interface_Neighbor_ChassisIdPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_Interface_Neighbor_ChassisIdPathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id YANG schema element.
type Lldp_Interface_Neighbor_ChassisIdPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/chassis-id"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id"
func (n *Lldp_Interface_Neighbor_ChassisIdPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "chassis-id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).ChassisId
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/chassis-id"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id"
func (n *Lldp_Interface_Neighbor_ChassisIdPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "chassis-id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).ChassisId
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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

// Lldp_Interface_Neighbor_ChassisIdTypePath represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type YANG schema element.
type Lldp_Interface_Neighbor_ChassisIdTypePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_Interface_Neighbor_ChassisIdTypePathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type YANG schema element.
type Lldp_Interface_Neighbor_ChassisIdTypePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/chassis-id-type"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type"
func (n *Lldp_Interface_Neighbor_ChassisIdTypePath) State() ygnmi.SingletonQuery[oc.E_LldpTypes_ChassisIdType] {
	return ygnmi.NewSingletonQuery[oc.E_LldpTypes_ChassisIdType](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "chassis-id-type"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_LldpTypes_ChassisIdType, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).ChassisIdType
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/chassis-id-type"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/chassis-id-type"
func (n *Lldp_Interface_Neighbor_ChassisIdTypePathAny) State() ygnmi.WildcardQuery[oc.E_LldpTypes_ChassisIdType] {
	return ygnmi.NewWildcardQuery[oc.E_LldpTypes_ChassisIdType](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "chassis-id-type"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_LldpTypes_ChassisIdType, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).ChassisIdType
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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

// Lldp_Interface_Neighbor_IdPath represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id YANG schema element.
type Lldp_Interface_Neighbor_IdPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_Interface_Neighbor_IdPathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/id YANG schema element.
type Lldp_Interface_Neighbor_IdPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/id"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/id"
func (n *Lldp_Interface_Neighbor_IdPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).Id
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/id"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/id"
func (n *Lldp_Interface_Neighbor_IdPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).Id
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "id"
//	Path from root:       ""
func (n *Lldp_Interface_Neighbor_IdPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewConfigQuery[string](
		"Lldp_Interface_Neighbor",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).Id
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "id"
//	Path from root:       ""
func (n *Lldp_Interface_Neighbor_IdPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp_Interface_Neighbor",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).Id
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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

// Lldp_Interface_Neighbor_LastUpdatePath represents the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update YANG schema element.
type Lldp_Interface_Neighbor_LastUpdatePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_Interface_Neighbor_LastUpdatePathAny represents the wildcard version of the /openconfig-lldp/lldp/interfaces/interface/neighbors/neighbor/state/last-update YANG schema element.
type Lldp_Interface_Neighbor_LastUpdatePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/last-update"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/last-update"
func (n *Lldp_Interface_Neighbor_LastUpdatePath) State() ygnmi.SingletonQuery[int64] {
	return ygnmi.NewSingletonQuery[int64](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "last-update"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (int64, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).LastUpdate
			if ret == nil {
				var zero int64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/last-update"
//	Path from root:       "/lldp/interfaces/interface/neighbors/neighbor/state/last-update"
func (n *Lldp_Interface_Neighbor_LastUpdatePathAny) State() ygnmi.WildcardQuery[int64] {
	return ygnmi.NewWildcardQuery[int64](
		"Lldp_Interface_Neighbor",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "last-update"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (int64, bool) {
			ret := gs.(*oc.Lldp_Interface_Neighbor).LastUpdate
			if ret == nil {
				var zero int64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp_Interface_Neighbor) },
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
