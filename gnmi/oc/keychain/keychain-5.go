/*
Package keychain is a generated package which contains definitions
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
package keychain

import (
	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// Keychain_Key_SendLifetime_EndTimePath represents the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time YANG schema element.
type Keychain_Key_SendLifetime_EndTimePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Keychain_Key_SendLifetime_EndTimePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/end-time YANG schema element.
type Keychain_Key_SendLifetime_EndTimePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "state/end-time"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/state/end-time"
func (n *Keychain_Key_SendLifetime_EndTimePath) State() ygnmi.SingletonQuery[uint64] {
	return ygnmi.NewSingletonQuery[uint64](
		"Keychain_Key_SendLifetime",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "end-time"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).EndTime
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "state/end-time"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/state/end-time"
func (n *Keychain_Key_SendLifetime_EndTimePathAny) State() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Keychain_Key_SendLifetime",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "end-time"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).EndTime
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "config/end-time"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/config/end-time"
func (n *Keychain_Key_SendLifetime_EndTimePath) Config() ygnmi.ConfigQuery[uint64] {
	return ygnmi.NewConfigQuery[uint64](
		"Keychain_Key_SendLifetime",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "end-time"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).EndTime
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "config/end-time"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/config/end-time"
func (n *Keychain_Key_SendLifetime_EndTimePathAny) Config() ygnmi.WildcardQuery[uint64] {
	return ygnmi.NewWildcardQuery[uint64](
		"Keychain_Key_SendLifetime",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "end-time"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (uint64, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).EndTime
			if ret == nil {
				var zero uint64
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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

// Keychain_Key_SendLifetime_SendAndReceivePath represents the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive YANG schema element.
type Keychain_Key_SendLifetime_SendAndReceivePath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Keychain_Key_SendLifetime_SendAndReceivePathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/send-lifetime/state/send-and-receive YANG schema element.
type Keychain_Key_SendLifetime_SendAndReceivePathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "state/send-and-receive"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/state/send-and-receive"
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) State() ygnmi.SingletonQuery[bool] {
	return ygnmi.NewSingletonQuery[bool](
		"Keychain_Key_SendLifetime",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "send-and-receive"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).SendAndReceive
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "state/send-and-receive"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/state/send-and-receive"
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) State() ygnmi.WildcardQuery[bool] {
	return ygnmi.NewWildcardQuery[bool](
		"Keychain_Key_SendLifetime",
		true,
		false,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "send-and-receive"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).SendAndReceive
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "config/send-and-receive"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/config/send-and-receive"
func (n *Keychain_Key_SendLifetime_SendAndReceivePath) Config() ygnmi.ConfigQuery[bool] {
	return ygnmi.NewConfigQuery[bool](
		"Keychain_Key_SendLifetime",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "send-and-receive"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).SendAndReceive
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "config/send-and-receive"
//	Path from root:       "/keychains/keychain/keys/key/send-lifetime/config/send-and-receive"
func (n *Keychain_Key_SendLifetime_SendAndReceivePathAny) Config() ygnmi.WildcardQuery[bool] {
	return ygnmi.NewWildcardQuery[bool](
		"Keychain_Key_SendLifetime",
		false,
		true,
		true,
		true,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "send-and-receive"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (bool, bool) {
			ret := gs.(*oc.Keychain_Key_SendLifetime).SendAndReceive
			if ret == nil {
				var zero bool
				return zero, false
			}
			return *ret, true
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key_SendLifetime) },
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
