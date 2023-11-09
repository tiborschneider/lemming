/*
Package definedsets is a generated package which contains definitions
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
package definedsets

import (
	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// DefinedSets_PortSetPath represents the /openconfig-defined-sets/defined-sets/port-sets/port-set YANG schema element.
type DefinedSets_PortSetPath struct {
	*ygnmi.NodePath
}

// DefinedSets_PortSetPathAny represents the wildcard version of the /openconfig-defined-sets/defined-sets/port-sets/port-set YANG schema element.
type DefinedSets_PortSetPathAny struct {
	*ygnmi.NodePath
}

// DefinedSets_PortSetPathMap represents the /openconfig-defined-sets/defined-sets/port-sets/port-set YANG schema element.
type DefinedSets_PortSetPathMap struct {
	*ygnmi.NodePath
}

// DefinedSets_PortSetPathMapAny represents the wildcard version of the /openconfig-defined-sets/defined-sets/port-sets/port-set YANG schema element.
type DefinedSets_PortSetPathMapAny struct {
	*ygnmi.NodePath
}

// Description (leaf): A user defined description for the port set
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "*/description"
//	Path from root:       "/defined-sets/port-sets/port-set/*/description"
func (n *DefinedSets_PortSetPath) Description() *DefinedSets_PortSet_DescriptionPath {
	ps := &DefinedSets_PortSet_DescriptionPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "description"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Description (leaf): A user defined description for the port set
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "*/description"
//	Path from root:       "/defined-sets/port-sets/port-set/*/description"
func (n *DefinedSets_PortSetPathAny) Description() *DefinedSets_PortSet_DescriptionPathAny {
	ps := &DefinedSets_PortSet_DescriptionPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "description"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Name (leaf): A user defined name of the port set.
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "*/name"
//	Path from root:       "/defined-sets/port-sets/port-set/*/name"
func (n *DefinedSets_PortSetPath) Name() *DefinedSets_PortSet_NamePath {
	ps := &DefinedSets_PortSet_NamePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Name (leaf): A user defined name of the port set.
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "*/name"
//	Path from root:       "/defined-sets/port-sets/port-set/*/name"
func (n *DefinedSets_PortSetPathAny) Name() *DefinedSets_PortSet_NamePathAny {
	ps := &DefinedSets_PortSet_NamePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Port (leaf-list): A user defined set of ports to be
// used in the match conditions.
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "*/port"
//	Path from root:       "/defined-sets/port-sets/port-set/*/port"
func (n *DefinedSets_PortSetPath) Port() *DefinedSets_PortSet_PortPath {
	ps := &DefinedSets_PortSet_PortPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "port"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Port (leaf-list): A user defined set of ports to be
// used in the match conditions.
//
//	Defining module:      "openconfig-defined-sets"
//	Instantiating module: "openconfig-defined-sets"
//	Path from parent:     "*/port"
//	Path from root:       "/defined-sets/port-sets/port-set/*/port"
func (n *DefinedSets_PortSetPathAny) Port() *DefinedSets_PortSet_PortPathAny {
	ps := &DefinedSets_PortSet_PortPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "port"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// State returns a Query that can be used in gNMI operations.
func (n *DefinedSets_PortSetPath) State() ygnmi.SingletonQuery[*oc.DefinedSets_PortSet] {
	return ygnmi.NewSingletonQuery[*oc.DefinedSets_PortSet](
		"DefinedSets_PortSet",
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
func (n *DefinedSets_PortSetPathAny) State() ygnmi.WildcardQuery[*oc.DefinedSets_PortSet] {
	return ygnmi.NewWildcardQuery[*oc.DefinedSets_PortSet](
		"DefinedSets_PortSet",
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

// Config returns a Query that can be used in gNMI operations.
func (n *DefinedSets_PortSetPath) Config() ygnmi.ConfigQuery[*oc.DefinedSets_PortSet] {
	return ygnmi.NewConfigQuery[*oc.DefinedSets_PortSet](
		"DefinedSets_PortSet",
		false,
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

// Config returns a Query that can be used in gNMI operations.
func (n *DefinedSets_PortSetPathAny) Config() ygnmi.WildcardQuery[*oc.DefinedSets_PortSet] {
	return ygnmi.NewWildcardQuery[*oc.DefinedSets_PortSet](
		"DefinedSets_PortSet",
		false,
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
func (n *DefinedSets_PortSetPathMap) State() ygnmi.SingletonQuery[map[string]*oc.DefinedSets_PortSet] {
	return ygnmi.NewSingletonQuery[map[string]*oc.DefinedSets_PortSet](
		"DefinedSets",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.DefinedSets_PortSet, bool) {
			ret := gs.(*oc.DefinedSets).PortSet
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.DefinedSets) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-defined-sets:port-sets"},
			PostRelPath: []string{"openconfig-defined-sets:port-set"},
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *DefinedSets_PortSetPathMapAny) State() ygnmi.WildcardQuery[map[string]*oc.DefinedSets_PortSet] {
	return ygnmi.NewWildcardQuery[map[string]*oc.DefinedSets_PortSet](
		"DefinedSets",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.DefinedSets_PortSet, bool) {
			ret := gs.(*oc.DefinedSets).PortSet
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.DefinedSets) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-defined-sets:port-sets"},
			PostRelPath: []string{"openconfig-defined-sets:port-set"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *DefinedSets_PortSetPathMap) Config() ygnmi.ConfigQuery[map[string]*oc.DefinedSets_PortSet] {
	return ygnmi.NewConfigQuery[map[string]*oc.DefinedSets_PortSet](
		"DefinedSets",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.DefinedSets_PortSet, bool) {
			ret := gs.(*oc.DefinedSets).PortSet
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.DefinedSets) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-defined-sets:port-sets"},
			PostRelPath: []string{"openconfig-defined-sets:port-set"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *DefinedSets_PortSetPathMapAny) Config() ygnmi.WildcardQuery[map[string]*oc.DefinedSets_PortSet] {
	return ygnmi.NewWildcardQuery[map[string]*oc.DefinedSets_PortSet](
		"DefinedSets",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.DefinedSets_PortSet, bool) {
			ret := gs.(*oc.DefinedSets).PortSet
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.DefinedSets) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-defined-sets:port-sets"},
			PostRelPath: []string{"openconfig-defined-sets:port-set"},
		},
	)
}
