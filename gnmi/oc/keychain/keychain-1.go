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
	"reflect"

	oc "github.com/openconfig/lemming/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
)

// KeychainPath represents the /openconfig-keychain/keychains/keychain YANG schema element.
type KeychainPath struct {
	*ygnmi.NodePath
}

// KeychainPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain YANG schema element.
type KeychainPathAny struct {
	*ygnmi.NodePath
}

// KeychainPathMap represents the /openconfig-keychain/keychains/keychain YANG schema element.
type KeychainPathMap struct {
	*ygnmi.NodePath
}

// KeychainPathMapAny represents the wildcard version of the /openconfig-keychain/keychains/keychain YANG schema element.
type KeychainPathMapAny struct {
	*ygnmi.NodePath
}

// KeyAny (list): List of configured keys for the keychain.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keys/key"
//	Path from root:       "/keychains/keychain/keys/key"
func (n *KeychainPath) KeyAny() *Keychain_KeyPathAny {
	ps := &Keychain_KeyPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": "*"},
			n,
		),
	}
	return ps
}

// KeyAny (list): List of configured keys for the keychain.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keys/key"
//	Path from root:       "/keychains/keychain/keys/key"
func (n *KeychainPathAny) KeyAny() *Keychain_KeyPathAny {
	ps := &Keychain_KeyPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": "*"},
			n,
		),
	}
	return ps
}

// Key (list): List of configured keys for the keychain.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keys/key"
//	Path from root:       "/keychains/keychain/keys/key"
//
//	KeyId: [oc.UnionString, oc.UnionUint64]
func (n *KeychainPath) Key(KeyId oc.Keychain_Key_KeyId_Union) *Keychain_KeyPath {
	ps := &Keychain_KeyPath{
		NodePath: ygnmi.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": KeyId},
			n,
		),
	}
	return ps
}

// Key (list): List of configured keys for the keychain.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keys/key"
//	Path from root:       "/keychains/keychain/keys/key"
//
//	KeyId: [oc.UnionString, oc.UnionUint64]
func (n *KeychainPathAny) Key(KeyId oc.Keychain_Key_KeyId_Union) *Keychain_KeyPathAny {
	ps := &Keychain_KeyPathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"keys", "key"},
			map[string]interface{}{"key-id": KeyId},
			n,
		),
	}
	return ps
}

// KeyMap (list): List of configured keys for the keychain.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keys/key"
//	Path from root:       "/keychains/keychain/keys/key"
func (n *KeychainPath) KeyMap() *Keychain_KeyPathMap {
	ps := &Keychain_KeyPathMap{
		NodePath: ygnmi.NewNodePath(
			[]string{"keys"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// KeyMap (list): List of configured keys for the keychain.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "keys/key"
//	Path from root:       "/keychains/keychain/keys/key"
func (n *KeychainPathAny) KeyMap() *Keychain_KeyPathMapAny {
	ps := &Keychain_KeyPathMapAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"keys"},
			map[string]interface{}{},
			n,
		),
	}
	return ps
}

// Name (leaf): Keychain name.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "*/name"
//	Path from root:       "/keychains/keychain/*/name"
func (n *KeychainPath) Name() *Keychain_NamePath {
	ps := &Keychain_NamePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Name (leaf): Keychain name.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "*/name"
//	Path from root:       "/keychains/keychain/*/name"
func (n *KeychainPathAny) Name() *Keychain_NamePathAny {
	ps := &Keychain_NamePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "name"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Tolerance (leaf): Tolerance (overlap time) that a receive key should be accepted.  May be
// expressed as range in seconds, or using the FOREVER value to indicate
// that the key does not expire.  The default value should be 0, i.e., no
// tolerance.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "*/tolerance"
//	Path from root:       "/keychains/keychain/*/tolerance"
func (n *KeychainPath) Tolerance() *Keychain_TolerancePath {
	ps := &Keychain_TolerancePath{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "tolerance"},
			map[string]interface{}{},
			n,
		),
		parent: n,
	}
	return ps
}

// Tolerance (leaf): Tolerance (overlap time) that a receive key should be accepted.  May be
// expressed as range in seconds, or using the FOREVER value to indicate
// that the key does not expire.  The default value should be 0, i.e., no
// tolerance.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "*/tolerance"
//	Path from root:       "/keychains/keychain/*/tolerance"
func (n *KeychainPathAny) Tolerance() *Keychain_TolerancePathAny {
	ps := &Keychain_TolerancePathAny{
		NodePath: ygnmi.NewNodePath(
			[]string{"*", "tolerance"},
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
func (n *KeychainPath) State() ygnmi.SingletonQuery[*oc.Keychain] {
	return ygnmi.NewSingletonQuery[*oc.Keychain](
		"Keychain",
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
func (n *KeychainPathAny) State() ygnmi.WildcardQuery[*oc.Keychain] {
	return ygnmi.NewWildcardQuery[*oc.Keychain](
		"Keychain",
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
func (n *KeychainPath) Config() ygnmi.ConfigQuery[*oc.Keychain] {
	return ygnmi.NewConfigQuery[*oc.Keychain](
		"Keychain",
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
func (n *KeychainPathAny) Config() ygnmi.WildcardQuery[*oc.Keychain] {
	return ygnmi.NewWildcardQuery[*oc.Keychain](
		"Keychain",
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
func (n *KeychainPathMap) State() ygnmi.SingletonQuery[map[string]*oc.Keychain] {
	return ygnmi.NewSingletonQuery[map[string]*oc.Keychain](
		"Root",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Keychain, bool) {
			ret := gs.(*oc.Root).Keychain
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Root) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-keychain:keychains"},
			PostRelPath: []string{"openconfig-keychain:keychain"},
		},
	)
}

// State returns a Query that can be used in gNMI operations.
func (n *KeychainPathMapAny) State() ygnmi.WildcardQuery[map[string]*oc.Keychain] {
	return ygnmi.NewWildcardQuery[map[string]*oc.Keychain](
		"Root",
		true,
		false,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Keychain, bool) {
			ret := gs.(*oc.Root).Keychain
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Root) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-keychain:keychains"},
			PostRelPath: []string{"openconfig-keychain:keychain"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *KeychainPathMap) Config() ygnmi.ConfigQuery[map[string]*oc.Keychain] {
	return ygnmi.NewConfigQuery[map[string]*oc.Keychain](
		"Root",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Keychain, bool) {
			ret := gs.(*oc.Root).Keychain
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Root) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-keychain:keychains"},
			PostRelPath: []string{"openconfig-keychain:keychain"},
		},
	)
}

// Config returns a Query that can be used in gNMI operations.
func (n *KeychainPathMapAny) Config() ygnmi.WildcardQuery[map[string]*oc.Keychain] {
	return ygnmi.NewWildcardQuery[map[string]*oc.Keychain](
		"Root",
		false,
		true,
		false,
		false,
		true,
		true,
		n,
		func(gs ygot.ValidatedGoStruct) (map[string]*oc.Keychain, bool) {
			ret := gs.(*oc.Root).Keychain
			return ret, ret != nil
		},
		func() ygot.ValidatedGoStruct { return new(oc.Root) },
		func() *ytypes.Schema {
			return &ytypes.Schema{
				Root:       &oc.Root{},
				SchemaTree: oc.SchemaTree,
				Unmarshal:  oc.Unmarshal,
			}
		},
		nil,
		&ygnmi.CompressionInfo{
			PreRelPath:  []string{"openconfig-keychain:keychains"},
			PostRelPath: []string{"openconfig-keychain:keychain"},
		},
	)
}

// Keychain_Key_CryptoAlgorithmPath represents the /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm YANG schema element.
type Keychain_Key_CryptoAlgorithmPath struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// Keychain_Key_CryptoAlgorithmPathAny represents the wildcard version of the /openconfig-keychain/keychains/keychain/keys/key/state/crypto-algorithm YANG schema element.
type Keychain_Key_CryptoAlgorithmPathAny struct {
	*ygnmi.NodePath
	parent ygnmi.PathStruct
}

// State returns a Query that can be used in gNMI operations.
//
//	Defining module:      "openconfig-keychain"
//	Instantiating module: "openconfig-keychain"
//	Path from parent:     "state/crypto-algorithm"
//	Path from root:       "/keychains/keychain/keys/key/state/crypto-algorithm"
func (n *Keychain_Key_CryptoAlgorithmPath) State() ygnmi.SingletonQuery[oc.E_KeychainTypes_CRYPTO_TYPE] {
	return ygnmi.NewSingletonQuery[oc.E_KeychainTypes_CRYPTO_TYPE](
		"Keychain_Key",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "crypto-algorithm"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_KeychainTypes_CRYPTO_TYPE, bool) {
			ret := gs.(*oc.Keychain_Key).CryptoAlgorithm
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key) },
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
//	Path from parent:     "state/crypto-algorithm"
//	Path from root:       "/keychains/keychain/keys/key/state/crypto-algorithm"
func (n *Keychain_Key_CryptoAlgorithmPathAny) State() ygnmi.WildcardQuery[oc.E_KeychainTypes_CRYPTO_TYPE] {
	return ygnmi.NewWildcardQuery[oc.E_KeychainTypes_CRYPTO_TYPE](
		"Keychain_Key",
		true,
		false,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"state", "crypto-algorithm"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_KeychainTypes_CRYPTO_TYPE, bool) {
			ret := gs.(*oc.Keychain_Key).CryptoAlgorithm
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key) },
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
//	Path from parent:     "config/crypto-algorithm"
//	Path from root:       "/keychains/keychain/keys/key/config/crypto-algorithm"
func (n *Keychain_Key_CryptoAlgorithmPath) Config() ygnmi.ConfigQuery[oc.E_KeychainTypes_CRYPTO_TYPE] {
	return ygnmi.NewConfigQuery[oc.E_KeychainTypes_CRYPTO_TYPE](
		"Keychain_Key",
		false,
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "crypto-algorithm"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_KeychainTypes_CRYPTO_TYPE, bool) {
			ret := gs.(*oc.Keychain_Key).CryptoAlgorithm
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key) },
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
//	Path from parent:     "config/crypto-algorithm"
//	Path from root:       "/keychains/keychain/keys/key/config/crypto-algorithm"
func (n *Keychain_Key_CryptoAlgorithmPathAny) Config() ygnmi.WildcardQuery[oc.E_KeychainTypes_CRYPTO_TYPE] {
	return ygnmi.NewWildcardQuery[oc.E_KeychainTypes_CRYPTO_TYPE](
		"Keychain_Key",
		false,
		true,
		true,
		false,
		true,
		false,
		ygnmi.NewNodePath(
			[]string{"config", "crypto-algorithm"},
			nil,
			n.parent,
		),
		func(gs ygot.ValidatedGoStruct) (oc.E_KeychainTypes_CRYPTO_TYPE, bool) {
			ret := gs.(*oc.Keychain_Key).CryptoAlgorithm
			return ret, !reflect.ValueOf(ret).IsZero()
		},
		func() ygot.ValidatedGoStruct { return new(oc.Keychain_Key) },
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
