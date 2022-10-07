/*
Package ocpath is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by ygnmi version: v0.3.1: (ygot: v0.25.2)
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

Imported modules were sourced from:
  - public/release/models/...
  - public/third_party/ietf/...
*/
package ocpath

import (
	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/lemming/gnmi/oc/acl"
	"github.com/openconfig/lemming/gnmi/oc/interfaces"
	"github.com/openconfig/lemming/gnmi/oc/keychain"
	"github.com/openconfig/lemming/gnmi/oc/lacp"
	"github.com/openconfig/lemming/gnmi/oc/lldp"
	"github.com/openconfig/lemming/gnmi/oc/networkinstance"
	"github.com/openconfig/lemming/gnmi/oc/platform"
	"github.com/openconfig/lemming/gnmi/oc/qos"
	"github.com/openconfig/lemming/gnmi/oc/routingpolicy"
	"github.com/openconfig/lemming/gnmi/oc/system"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ytypes"
)

// RootPath represents the /root YANG schema element.
type RootPath struct {
	*ygnmi.DeviceRootBase
}

// Root returns a root path object from which YANG paths can be constructed.
func Root() *RootPath {
	return &RootPath{ygnmi.NewDeviceRootBase()}
}

// Acl (container): Top level enclosing container for ACL model config
// and operational state data
//
//	Defining module:      "openconfig-acl"
//	Instantiating module: "openconfig-acl"
//	Path from parent:     "acl"
//	Path from root:       "/acl"
func (n *RootPath) Acl() *acl.AclPath {
	return &acl.AclPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"acl"},
			map[string]interface{}{},
			n,
		),
	}
}

// ComponentAny (list): List of components, keyed by component name.
//
//	Defining module:      "openconfig-platform"
//	Instantiating module: "openconfig-platform"
//	Path from parent:     "components/component"
//	Path from root:       "/components/component"
func (n *RootPath) ComponentAny() *platform.ComponentPathAny {
	return &platform.ComponentPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"components", "component"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Component (list): List of components, keyed by component name.
//
//	Defining module:      "openconfig-platform"
//	Instantiating module: "openconfig-platform"
//	Path from parent:     "components/component"
//	Path from root:       "/components/component"
//
//	Name: string
func (n *RootPath) Component(Name string) *platform.ComponentPath {
	return &platform.ComponentPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"components", "component"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// InterfaceAny (list): The list of named interfaces on the device.
//
//	Defining module:      "openconfig-interfaces"
//	Instantiating module: "openconfig-interfaces"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
func (n *RootPath) InterfaceAny() *interfaces.InterfacePathAny {
	return &interfaces.InterfacePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): The list of named interfaces on the device.
//
//	Defining module:      "openconfig-interfaces"
//	Instantiating module: "openconfig-interfaces"
//	Path from parent:     "interfaces/interface"
//	Path from root:       "/interfaces/interface"
//
//	Name: string
func (n *RootPath) Interface(Name string) *interfaces.InterfacePath {
	return &interfaces.InterfacePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// KeychainAny (list): List of defined keychains.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keychains/keychain"
//	Path from root:       "/keychains/keychain"
func (n *RootPath) KeychainAny() *keychain.KeychainPathAny {
	return &keychain.KeychainPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"keychains", "keychain"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Keychain (list): List of defined keychains.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keychains/keychain"
//	Path from root:       "/keychains/keychain"
//
//	Name: string
func (n *RootPath) Keychain(Name string) *keychain.KeychainPath {
	return &keychain.KeychainPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"keychains", "keychain"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Lacp (container): Configuration and operational state data for LACP protocol
// operation on the aggregate interface
//
//	Defining module:      "openconfig-lacp"
//	Instantiating module: "openconfig-lacp"
//	Path from parent:     "lacp"
//	Path from root:       "/lacp"
func (n *RootPath) Lacp() *lacp.LacpPath {
	return &lacp.LacpPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lacp"},
			map[string]interface{}{},
			n,
		),
	}
}

// Lldp (container): Top-level container for LLDP configuration and state data
//
//	Defining module:      "openconfig-lldp"
//	Instantiating module: "openconfig-lldp"
//	Path from parent:     "lldp"
//	Path from root:       "/lldp"
func (n *RootPath) Lldp() *lldp.LldpPath {
	return &lldp.LldpPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"lldp"},
			map[string]interface{}{},
			n,
		),
	}
}

// NetworkInstanceAny (list): Network instances configured on the local system
//
//	Defining module:      "openconfig-network-instance"
//	Instantiating module: "openconfig-network-instance"
//	Path from parent:     "network-instances/network-instance"
//	Path from root:       "/network-instances/network-instance"
func (n *RootPath) NetworkInstanceAny() *networkinstance.NetworkInstancePathAny {
	return &networkinstance.NetworkInstancePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"network-instances", "network-instance"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// NetworkInstance (list): Network instances configured on the local system
//
//	Defining module:      "openconfig-network-instance"
//	Instantiating module: "openconfig-network-instance"
//	Path from parent:     "network-instances/network-instance"
//	Path from root:       "/network-instances/network-instance"
//
//	Name: string
func (n *RootPath) NetworkInstance(Name string) *networkinstance.NetworkInstancePath {
	return &networkinstance.NetworkInstancePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"network-instances", "network-instance"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Qos (container): Top-level container for QoS data
//
//	Defining module:      "openconfig-qos"
//	Instantiating module: "openconfig-qos"
//	Path from parent:     "qos"
//	Path from root:       "/qos"
func (n *RootPath) Qos() *qos.QosPath {
	return &qos.QosPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"qos"},
			map[string]interface{}{},
			n,
		),
	}
}

// RoutingPolicy (container): Top-level container for all routing policy configuration
//
//	Defining module:      "openconfig-routing-policy"
//	Instantiating module: "openconfig-routing-policy"
//	Path from parent:     "routing-policy"
//	Path from root:       "/routing-policy"
func (n *RootPath) RoutingPolicy() *routingpolicy.RoutingPolicyPath {
	return &routingpolicy.RoutingPolicyPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"routing-policy"},
			map[string]interface{}{},
			n,
		),
	}
}

// System (container): Enclosing container for system-related configuration and
// operational state data
//
//	Defining module:      "openconfig-system"
//	Instantiating module: "openconfig-system"
//	Path from parent:     "system"
//	Path from root:       "/system"
func (n *RootPath) System() *system.SystemPath {
	return &system.SystemPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"system"},
			map[string]interface{}{},
			n,
		),
	}
}

// Batch contains a collection of paths.
// Calling State() or Config() on the batch returns a query
// that can use to Lookup, Watch, etc on multiple paths at once.
type Batch struct {
	paths []ygnmi.PathStruct
}

// AddPaths adds the paths to the batch.
func (b *Batch) AddPaths(paths ...ygnmi.PathStruct) *Batch {
	b.paths = append(b.paths, paths...)
	return b
}

// State returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) State() ygnmi.SingletonQuery[*oc.Root] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		true,
		ygnmi.NewDeviceRootBase(),
		queryPaths,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
// The returned query is immutable, adding paths does not modify existing queries.
func (b *Batch) Config() ygnmi.SingletonQuery[*oc.Root] {
	queryPaths := make([]ygnmi.PathStruct, len(b.paths))
	copy(queryPaths, b.paths)
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		false,
		ygnmi.NewDeviceRootBase(),
		queryPaths,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *RootPath) State() ygnmi.SingletonQuery[*oc.Root] {
	return ygnmi.NewNonLeafSingletonQuery[*oc.Root](
		"Root",
		true,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *RootPath) Config() ygnmi.ConfigQuery[*oc.Root] {
	return ygnmi.NewNonLeafConfigQuery[*oc.Root](
		"Root",
		false,
		n,
		nil,
		&ytypes.Schema{
			Root:       &oc.Root{},
			SchemaTree: oc.SchemaTree,
			Unmarshal:  oc.Unmarshal,
		},
	)
}
