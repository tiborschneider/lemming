/*
Package lacp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.8.7: (ygot: v0.29.10)
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

// Lacp_SystemPriorityPath represents the /openconfig-lacp/lacp/state/system-priority YANG schema element.
type Lacp_SystemPriorityPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_SystemPriorityPathAny represents the wildcard version of the /openconfig-lacp/lacp/state/system-priority YANG schema element.
type Lacp_SystemPriorityPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/system-priority"
//	Path from root:       "/lacp/state/system-priority"
func (n *Lacp_SystemPriorityPath) State() ygnmi.SingletonQuery[uint16] {
	return ygnmi.NewSingletonQuery[uint16](
		"Lacp",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
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
//	Path from parent:     "state/system-priority"
//	Path from root:       "/lacp/state/system-priority"
func (n *Lacp_SystemPriorityPathAny) State() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp",
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
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

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/system-priority"
//	Path from root:       "/lacp/config/system-priority"
func (n *Lacp_SystemPriorityPath) Config() ygnmi.ConfigQuery[uint16] {
	return ygnmi.NewConfigQuery[uint16](
		"Lacp",
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
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
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/system-priority"
//	Path from root:       "/lacp/config/system-priority"
func (n *Lacp_SystemPriorityPathAny) Config() ygnmi.WildcardQuery[uint16] {
	return ygnmi.NewWildcardQuery[uint16](
		"Lacp",
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-priority"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint16, bool) {
			ret := gs.(*oc.Lacp).SystemPriority
			if ret == nil {
				var zero uint16
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp) },
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

// LacpPath represents the /openconfig-lacp/lacp YANG schema element.
type LacpPath struct {
	*ygnmi.NodePath
}

// LacpPathAny represents the wildcard version of the /openconfig-lacp/lacp YANG schema element.
type LacpPathAny struct {
	*ygnmi.NodePath
}

// InterfaceAny (list): List of aggregate interfaces managed by LACP
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/lacp/interfaces/interface"
func (n *LacpPath) InterfaceAny() *Lacp_InterfacePathAny {
	ps := &Lacp_InterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// InterfaceAny (list): List of aggregate interfaces managed by LACP
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/lacp/interfaces/interface"
func (n *LacpPathAny) InterfaceAny() *Lacp_InterfacePathAny {
	ps := &Lacp_InterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
	return ps
}

// Interface (list): List of aggregate interfaces managed by LACP
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/lacp/interfaces/interface"
//
//	Name: string
func (n *LacpPath) Interface(Name string) *Lacp_InterfacePath {
	ps := &Lacp_InterfacePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// Interface (list): List of aggregate interfaces managed by LACP
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/lacp/interfaces/interface"
//
//	Name: string
func (n *LacpPathAny) Interface(Name string) *Lacp_InterfacePathAny {
	ps := &Lacp_InterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
	return ps
}

// InterfaceMap (list): List of aggregate interfaces managed by LACP
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/lacp/interfaces/interface"
func (n *LacpPath) InterfaceMap() *Lacp_InterfacePathMap {
	ps := &Lacp_InterfacePathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// InterfaceMap (list): List of aggregate interfaces managed by LACP
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/lacp/interfaces/interface"
func (n *LacpPathAny) InterfaceMap() *Lacp_InterfacePathMapAny {
	ps := &Lacp_InterfacePathMapAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/system-priority"
//	Path from root:       "/lacp/*/system-priority"
func (n *LacpPath) SystemPriority() *Lacp_SystemPriorityPath {
	ps := &Lacp_SystemPriorityPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "system-priority"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// SystemPriority (leaf): Sytem priority used by the node on this LAG interface.
// Lower value is higher priority for determining which node
// is the controlling system.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "*/system-priority"
//	Path from root:       "/lacp/*/system-priority"
func (n *LacpPathAny) SystemPriority() *Lacp_SystemPriorityPathAny {
	ps := &Lacp_SystemPriorityPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "system-priority"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

func binarySliceToFloatSlice(in []oc.Binary) []float32 {
	converted := make([]float32, 0, len(in))
	for _, binary := range in {
		converted = append(converted, ygot.BinaryToFloat32(binary))
	}
	return converted
}

// State returns a Query that can be used in gNMI operations.
func (n *LacpPath) State() ygnmi.SingletonQuery[*oc.Lacp] {
	return ygnmi.NewSingletonQuery[*oc.Lacp](
		"Lacp",
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
func (n *LacpPathAny) State() ygnmi.WildcardQuery[*oc.Lacp] {
	return ygnmi.NewWildcardQuery[*oc.Lacp](
		"Lacp",
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

// Config returns a Query that can be used in gNMI operations.
func (n *LacpPath) Config() ygnmi.ConfigQuery[*oc.Lacp] {
	return ygnmi.NewConfigQuery[*oc.Lacp](
		"Lacp",
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

// Config returns a Query that can be used in gNMI operations.
func (n *LacpPathAny) Config() ygnmi.WildcardQuery[*oc.Lacp] {
	return ygnmi.NewWildcardQuery[*oc.Lacp](
		"Lacp",
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
	)
}

// Lacp_Interface_IntervalPath represents the /openconfig-lacp/lacp/interfaces/interface/state/interval YANG schema element.
type Lacp_Interface_IntervalPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_IntervalPathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/state/interval YANG schema element.
type Lacp_Interface_IntervalPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/interval"
//	Path from root:       "/lacp/interfaces/interface/state/interval"
func (n *Lacp_Interface_IntervalPath) State() ygnmi.SingletonQuery[oc.E_Lacp_LacpPeriodType] {
	return ygnmi.NewSingletonQuery[oc.E_Lacp_LacpPeriodType](
		"Lacp_Interface",
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "interval"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpPeriodType, bool) {
			ret := gs.(*oc.Lacp_Interface).Interval
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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
//	Path from parent:     "state/interval"
//	Path from root:       "/lacp/interfaces/interface/state/interval"
func (n *Lacp_Interface_IntervalPathAny) State() ygnmi.WildcardQuery[oc.E_Lacp_LacpPeriodType] {
	return ygnmi.NewWildcardQuery[oc.E_Lacp_LacpPeriodType](
		"Lacp_Interface",
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "interval"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpPeriodType, bool) {
			ret := gs.(*oc.Lacp_Interface).Interval
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/interval"
//	Path from root:       "/lacp/interfaces/interface/config/interval"
func (n *Lacp_Interface_IntervalPath) Config() ygnmi.ConfigQuery[oc.E_Lacp_LacpPeriodType] {
	return ygnmi.NewConfigQuery[oc.E_Lacp_LacpPeriodType](
		"Lacp_Interface",
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "interval"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpPeriodType, bool) {
			ret := gs.(*oc.Lacp_Interface).Interval
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/interval"
//	Path from root:       "/lacp/interfaces/interface/config/interval"
func (n *Lacp_Interface_IntervalPathAny) Config() ygnmi.WildcardQuery[oc.E_Lacp_LacpPeriodType] {
	return ygnmi.NewWildcardQuery[oc.E_Lacp_LacpPeriodType](
		"Lacp_Interface",
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "interval"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpPeriodType, bool) {
			ret := gs.(*oc.Lacp_Interface).Interval
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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

// Lacp_Interface_LacpModePath represents the /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode YANG schema element.
type Lacp_Interface_LacpModePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lacp_Interface_LacpModePathAny represents the wildcard version of the /openconfig-lacp/lacp/interfaces/interface/state/lacp-mode YANG schema element.
type Lacp_Interface_LacpModePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "state/lacp-mode"
//	Path from root:       "/lacp/interfaces/interface/state/lacp-mode"
func (n *Lacp_Interface_LacpModePath) State() ygnmi.SingletonQuery[oc.E_Lacp_LacpActivityType] {
	return ygnmi.NewSingletonQuery[oc.E_Lacp_LacpActivityType](
		"Lacp_Interface",
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "lacp-mode"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpActivityType, bool) {
			ret := gs.(*oc.Lacp_Interface).LacpMode
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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
//	Path from parent:     "state/lacp-mode"
//	Path from root:       "/lacp/interfaces/interface/state/lacp-mode"
func (n *Lacp_Interface_LacpModePathAny) State() ygnmi.WildcardQuery[oc.E_Lacp_LacpActivityType] {
	return ygnmi.NewWildcardQuery[oc.E_Lacp_LacpActivityType](
		"Lacp_Interface",
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "lacp-mode"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpActivityType, bool) {
			ret := gs.(*oc.Lacp_Interface).LacpMode
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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

// Config returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/lacp-mode"
//	Path from root:       "/lacp/interfaces/interface/config/lacp-mode"
func (n *Lacp_Interface_LacpModePath) Config() ygnmi.ConfigQuery[oc.E_Lacp_LacpActivityType] {
	return ygnmi.NewConfigQuery[oc.E_Lacp_LacpActivityType](
		"Lacp_Interface",
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "lacp-mode"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpActivityType, bool) {
			ret := gs.(*oc.Lacp_Interface).LacpMode
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "config/lacp-mode"
//	Path from root:       "/lacp/interfaces/interface/config/lacp-mode"
func (n *Lacp_Interface_LacpModePathAny) Config() ygnmi.WildcardQuery[oc.E_Lacp_LacpActivityType] {
	return ygnmi.NewWildcardQuery[oc.E_Lacp_LacpActivityType](
		"Lacp_Interface",
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "lacp-mode"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_Lacp_LacpActivityType, bool) {
			ret := gs.(*oc.Lacp_Interface).LacpMode
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lacp_Interface) },
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
