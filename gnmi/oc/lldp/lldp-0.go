/*
Package lldp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema.

This package was generated by ygnmi version: v0.10.0: (ygot: v0.29.12)
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
package lldp

import (
	"reflect"

	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Lldp_ChassisIdPath represents the /openconfig-lldp/lldp/state/chassis-id YANG schema element.
type Lldp_ChassisIdPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_ChassisIdPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/chassis-id YANG schema element.
type Lldp_ChassisIdPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/chassis-id"
//	Path from root:       "/lldp/state/chassis-id"
func (n *Lldp_ChassisIdPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lldp",
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
			ret := gs.(*oc.Lldp).ChassisId
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from root:       "/lldp/state/chassis-id"
func (n *Lldp_ChassisIdPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp",
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
			ret := gs.(*oc.Lldp).ChassisId
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/chassis-id"
//	Path from root:       "/lldp/config/chassis-id"
func (n *Lldp_ChassisIdPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewConfigQuery[string](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "chassis-id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).ChassisId
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/chassis-id"
//	Path from root:       "/lldp/config/chassis-id"
func (n *Lldp_ChassisIdPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "chassis-id"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).ChassisId
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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

// Lldp_ChassisIdTypePath represents the /openconfig-lldp/lldp/state/chassis-id-type YANG schema element.
type Lldp_ChassisIdTypePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_ChassisIdTypePathAny represents the wildcard version of the /openconfig-lldp/lldp/state/chassis-id-type YANG schema element.
type Lldp_ChassisIdTypePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/chassis-id-type"
//	Path from root:       "/lldp/state/chassis-id-type"
func (n *Lldp_ChassisIdTypePath) State() ygnmi.SingletonQuery[oc.E_LldpTypes_ChassisIdType] {
	return ygnmi.NewSingletonQuery[oc.E_LldpTypes_ChassisIdType](
		"Lldp",
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
			ret := gs.(*oc.Lldp).ChassisIdType
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from root:       "/lldp/state/chassis-id-type"
func (n *Lldp_ChassisIdTypePathAny) State() ygnmi.WildcardQuery[oc.E_LldpTypes_ChassisIdType] {
	return ygnmi.NewWildcardQuery[oc.E_LldpTypes_ChassisIdType](
		"Lldp",
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
			ret := gs.(*oc.Lldp).ChassisIdType
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/chassis-id-type"
//	Path from root:       "/lldp/config/chassis-id-type"
func (n *Lldp_ChassisIdTypePath) Config() ygnmi.ConfigQuery[oc.E_LldpTypes_ChassisIdType] {
	return ygnmi.NewConfigQuery[oc.E_LldpTypes_ChassisIdType](
		"Lldp",
		false,
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "chassis-id-type"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_LldpTypes_ChassisIdType, bool) {
			ret := gs.(*oc.Lldp).ChassisIdType
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/chassis-id-type"
//	Path from root:       "/lldp/config/chassis-id-type"
func (n *Lldp_ChassisIdTypePathAny) Config() ygnmi.WildcardQuery[oc.E_LldpTypes_ChassisIdType] {
	return ygnmi.NewWildcardQuery[oc.E_LldpTypes_ChassisIdType](
		"Lldp",
		false,
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "chassis-id-type"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_LldpTypes_ChassisIdType, bool) {
			ret := gs.(*oc.Lldp).ChassisIdType
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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

// Lldp_EnabledPath represents the /openconfig-lldp/lldp/state/enabled YANG schema element.
type Lldp_EnabledPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_EnabledPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/enabled YANG schema element.
type Lldp_EnabledPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/enabled"
//	Path from root:       "/lldp/state/enabled"
func (n *Lldp_EnabledPath) State() ygnmi.SingletonQuery[bool] {
	return ygnmi.NewSingletonQuery[bool](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Lldp).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "state/enabled"
//	Path from root:       "/lldp/state/enabled"
func (n *Lldp_EnabledPathAny) State() ygnmi.WildcardQuery[bool] {
	return ygnmi.NewWildcardQuery[bool](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Lldp).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/enabled"
//	Path from root:       "/lldp/config/enabled"
func (n *Lldp_EnabledPath) Config() ygnmi.ConfigQuery[bool] {
	return ygnmi.NewConfigQuery[bool](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Lldp).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/enabled"
//	Path from root:       "/lldp/config/enabled"
func (n *Lldp_EnabledPathAny) Config() ygnmi.WildcardQuery[bool] {
	return ygnmi.NewWildcardQuery[bool](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "enabled"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Lldp).Enabled
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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

// Lldp_HelloTimerPath represents the /openconfig-lldp/lldp/state/hello-timer YANG schema element.
type Lldp_HelloTimerPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_HelloTimerPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/hello-timer YANG schema element.
type Lldp_HelloTimerPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/hello-timer"
//	Path from root:       "/lldp/state/hello-timer"
func (n *Lldp_HelloTimerPath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "hello-timer"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp).HelloTimer
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "state/hello-timer"
//	Path from root:       "/lldp/state/hello-timer"
func (n *Lldp_HelloTimerPathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "hello-timer"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp).HelloTimer
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/hello-timer"
//	Path from root:       "/lldp/config/hello-timer"
func (n *Lldp_HelloTimerPath) Config() ygnmi.ConfigQuery[uint64] {
	return ygnmi.NewConfigQuery[uint64](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "hello-timer"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp).HelloTimer
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/hello-timer"
//	Path from root:       "/lldp/config/hello-timer"
func (n *Lldp_HelloTimerPathAny) Config() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "hello-timer"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Lldp).HelloTimer
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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

// Lldp_SuppressTlvAdvertisementPath represents the /openconfig-lldp/lldp/state/suppress-tlv-advertisement YANG schema element.
type Lldp_SuppressTlvAdvertisementPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_SuppressTlvAdvertisementPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/suppress-tlv-advertisement YANG schema element.
type Lldp_SuppressTlvAdvertisementPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/suppress-tlv-advertisement"
//	Path from root:       "/lldp/state/suppress-tlv-advertisement"
func (n *Lldp_SuppressTlvAdvertisementPath) State() ygnmi.SingletonQuery[[]oc.E_LldpTypes_LLDP_TLV] {
	return ygnmi.NewSingletonQuery[[]oc.E_LldpTypes_LLDP_TLV](
		"Lldp",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "suppress-tlv-advertisement"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]oc.E_LldpTypes_LLDP_TLV, bool) {
			ret := gs.(*oc.Lldp).SuppressTlvAdvertisement
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "state/suppress-tlv-advertisement"
//	Path from root:       "/lldp/state/suppress-tlv-advertisement"
func (n *Lldp_SuppressTlvAdvertisementPathAny) State() ygnmi.WildcardQuery[[]oc.E_LldpTypes_LLDP_TLV] {
	return ygnmi.NewWildcardQuery[[]oc.E_LldpTypes_LLDP_TLV](
		"Lldp",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "suppress-tlv-advertisement"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]oc.E_LldpTypes_LLDP_TLV, bool) {
			ret := gs.(*oc.Lldp).SuppressTlvAdvertisement
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/suppress-tlv-advertisement"
//	Path from root:       "/lldp/config/suppress-tlv-advertisement"
func (n *Lldp_SuppressTlvAdvertisementPath) Config() ygnmi.ConfigQuery[[]oc.E_LldpTypes_LLDP_TLV] {
	return ygnmi.NewConfigQuery[[]oc.E_LldpTypes_LLDP_TLV](
		"Lldp",
		false,
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "suppress-tlv-advertisement"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]oc.E_LldpTypes_LLDP_TLV, bool) {
			ret := gs.(*oc.Lldp).SuppressTlvAdvertisement
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/suppress-tlv-advertisement"
//	Path from root:       "/lldp/config/suppress-tlv-advertisement"
func (n *Lldp_SuppressTlvAdvertisementPathAny) Config() ygnmi.WildcardQuery[[]oc.E_LldpTypes_LLDP_TLV] {
	return ygnmi.NewWildcardQuery[[]oc.E_LldpTypes_LLDP_TLV](
		"Lldp",
		false,
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "suppress-tlv-advertisement"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) ([]oc.E_LldpTypes_LLDP_TLV, bool) {
			ret := gs.(*oc.Lldp).SuppressTlvAdvertisement
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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

// Lldp_SystemDescriptionPath represents the /openconfig-lldp/lldp/state/system-description YANG schema element.
type Lldp_SystemDescriptionPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_SystemDescriptionPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/system-description YANG schema element.
type Lldp_SystemDescriptionPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/system-description"
//	Path from root:       "/lldp/state/system-description"
func (n *Lldp_SystemDescriptionPath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-description"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemDescription
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "state/system-description"
//	Path from root:       "/lldp/state/system-description"
func (n *Lldp_SystemDescriptionPathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-description"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemDescription
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/system-description"
//	Path from root:       "/lldp/config/system-description"
func (n *Lldp_SystemDescriptionPath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewConfigQuery[string](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-description"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemDescription
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/system-description"
//	Path from root:       "/lldp/config/system-description"
func (n *Lldp_SystemDescriptionPathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-description"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemDescription
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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

// Lldp_SystemNamePath represents the /openconfig-lldp/lldp/state/system-name YANG schema element.
type Lldp_SystemNamePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Lldp_SystemNamePathAny represents the wildcard version of the /openconfig-lldp/lldp/state/system-name YANG schema element.
type Lldp_SystemNamePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "state/system-name"
//	Path from root:       "/lldp/state/system-name"
func (n *Lldp_SystemNamePath) State() ygnmi.SingletonQuery[string] {
	return ygnmi.NewSingletonQuery[string](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemName
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "state/system-name"
//	Path from root:       "/lldp/state/system-name"
func (n *Lldp_SystemNamePathAny) State() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "system-name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemName
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/system-name"
//	Path from root:       "/lldp/config/system-name"
func (n *Lldp_SystemNamePath) Config() ygnmi.ConfigQuery[string] {
	return ygnmi.NewConfigQuery[string](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemName
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
//	Path from parent:     "config/system-name"
//	Path from root:       "/lldp/config/system-name"
func (n *Lldp_SystemNamePathAny) Config() ygnmi.WildcardQuery[string] {
	return ygnmi.NewWildcardQuery[string](
		"Lldp",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "system-name"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (string, bool) {
			ret := gs.(*oc.Lldp).SystemName
			if ret == nil {
				var zero string
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Lldp) },
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
