/*
Package lldp is a generated package which contains definitions
of structs which generate gNMI paths for a YANG schema. The generated paths are
based on a compressed form of the schema.

This package was generated by /usr/local/google/home/robjs/go/pkg/mod/github.com/openconfig/ygot@v0.24.2/genutil/names.go
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
package lldp

import (
	"github.com/openconfig/ygot/ygot"
)

// LldpPath represents the /openconfig-lldp/lldp YANG schema element.
type LldpPath struct {
	*ygot.NodePath
}

// LldpPathAny represents the wildcard version of the /openconfig-lldp/lldp YANG schema element.
type LldpPathAny struct {
	*ygot.NodePath
}

// Lldp_ChassisIdPath represents the /openconfig-lldp/lldp/state/chassis-id YANG schema element.
type Lldp_ChassisIdPath struct {
	*ygot.NodePath
}

// Lldp_ChassisIdPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/chassis-id YANG schema element.
type Lldp_ChassisIdPathAny struct {
	*ygot.NodePath
}

// Lldp_ChassisIdTypePath represents the /openconfig-lldp/lldp/state/chassis-id-type YANG schema element.
type Lldp_ChassisIdTypePath struct {
	*ygot.NodePath
}

// Lldp_ChassisIdTypePathAny represents the wildcard version of the /openconfig-lldp/lldp/state/chassis-id-type YANG schema element.
type Lldp_ChassisIdTypePathAny struct {
	*ygot.NodePath
}

// Lldp_EnabledPath represents the /openconfig-lldp/lldp/state/enabled YANG schema element.
type Lldp_EnabledPath struct {
	*ygot.NodePath
}

// Lldp_EnabledPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/enabled YANG schema element.
type Lldp_EnabledPathAny struct {
	*ygot.NodePath
}

// Lldp_HelloTimerPath represents the /openconfig-lldp/lldp/state/hello-timer YANG schema element.
type Lldp_HelloTimerPath struct {
	*ygot.NodePath
}

// Lldp_HelloTimerPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/hello-timer YANG schema element.
type Lldp_HelloTimerPathAny struct {
	*ygot.NodePath
}

// Lldp_SuppressTlvAdvertisementPath represents the /openconfig-lldp/lldp/state/suppress-tlv-advertisement YANG schema element.
type Lldp_SuppressTlvAdvertisementPath struct {
	*ygot.NodePath
}

// Lldp_SuppressTlvAdvertisementPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/suppress-tlv-advertisement YANG schema element.
type Lldp_SuppressTlvAdvertisementPathAny struct {
	*ygot.NodePath
}

// Lldp_SystemDescriptionPath represents the /openconfig-lldp/lldp/state/system-description YANG schema element.
type Lldp_SystemDescriptionPath struct {
	*ygot.NodePath
}

// Lldp_SystemDescriptionPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/system-description YANG schema element.
type Lldp_SystemDescriptionPathAny struct {
	*ygot.NodePath
}

// Lldp_SystemNamePath represents the /openconfig-lldp/lldp/state/system-name YANG schema element.
type Lldp_SystemNamePath struct {
	*ygot.NodePath
}

// Lldp_SystemNamePathAny represents the wildcard version of the /openconfig-lldp/lldp/state/system-name YANG schema element.
type Lldp_SystemNamePathAny struct {
	*ygot.NodePath
}

// ChassisId (leaf): The Chassis ID is a mandatory TLV which identifies the
// chassis component of the endpoint identifier associated with
// the transmitting LLDP agent
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/chassis-id"
// Path from root: "/lldp/state/chassis-id"
func (n *LldpPath) ChassisId() *Lldp_ChassisIdPath {
	return &Lldp_ChassisIdPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "chassis-id"},
			map[string]interface{}{},
			n,
		),
	}
}

// ChassisId (leaf): The Chassis ID is a mandatory TLV which identifies the
// chassis component of the endpoint identifier associated with
// the transmitting LLDP agent
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/chassis-id"
// Path from root: "/lldp/state/chassis-id"
func (n *LldpPathAny) ChassisId() *Lldp_ChassisIdPathAny {
	return &Lldp_ChassisIdPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "chassis-id"},
			map[string]interface{}{},
			n,
		),
	}
}

// ChassisIdType (leaf): This field identifies the format and source of the chassis
// identifier string. It is an enumerator defined by the
// LldpChassisIdSubtype object from IEEE 802.1AB MIB.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/chassis-id-type"
// Path from root: "/lldp/state/chassis-id-type"
func (n *LldpPath) ChassisIdType() *Lldp_ChassisIdTypePath {
	return &Lldp_ChassisIdTypePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "chassis-id-type"},
			map[string]interface{}{},
			n,
		),
	}
}

// ChassisIdType (leaf): This field identifies the format and source of the chassis
// identifier string. It is an enumerator defined by the
// LldpChassisIdSubtype object from IEEE 802.1AB MIB.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/chassis-id-type"
// Path from root: "/lldp/state/chassis-id-type"
func (n *LldpPathAny) ChassisIdType() *Lldp_ChassisIdTypePathAny {
	return &Lldp_ChassisIdTypePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "chassis-id-type"},
			map[string]interface{}{},
			n,
		),
	}
}

// Counters (container): Global LLDP counters
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/counters"
// Path from root: "/lldp/state/counters"
func (n *LldpPath) Counters() *Lldp_CountersPath {
	return &Lldp_CountersPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// Counters (container): Global LLDP counters
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/counters"
// Path from root: "/lldp/state/counters"
func (n *LldpPathAny) Counters() *Lldp_CountersPathAny {
	return &Lldp_CountersPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "counters"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): System level state of the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/enabled"
// Path from root: "/lldp/state/enabled"
func (n *LldpPath) Enabled() *Lldp_EnabledPath {
	return &Lldp_EnabledPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// Enabled (leaf): System level state of the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/enabled"
// Path from root: "/lldp/state/enabled"
func (n *LldpPathAny) Enabled() *Lldp_EnabledPathAny {
	return &Lldp_EnabledPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "enabled"},
			map[string]interface{}{},
			n,
		),
	}
}

// HelloTimer (leaf): System level hello timer for the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/hello-timer"
// Path from root: "/lldp/state/hello-timer"
func (n *LldpPath) HelloTimer() *Lldp_HelloTimerPath {
	return &Lldp_HelloTimerPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "hello-timer"},
			map[string]interface{}{},
			n,
		),
	}
}

// HelloTimer (leaf): System level hello timer for the LLDP protocol.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/hello-timer"
// Path from root: "/lldp/state/hello-timer"
func (n *LldpPathAny) HelloTimer() *Lldp_HelloTimerPathAny {
	return &Lldp_HelloTimerPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "hello-timer"},
			map[string]interface{}{},
			n,
		),
	}
}

// InterfaceAny (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name (wildcarded): string
func (n *LldpPath) InterfaceAny() *Lldp_InterfacePathAny {
	return &Lldp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// InterfaceAny (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name (wildcarded): string
func (n *LldpPathAny) InterfaceAny() *Lldp_InterfacePathAny {
	return &Lldp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": "*"},
			n,
		),
	}
}

// Interface (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name: string
func (n *LldpPath) Interface(Name string) *Lldp_InterfacePath {
	return &Lldp_InterfacePath{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// Interface (list): List of interfaces on which LLDP is enabled / available
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "interfaces/interface"
// Path from root: "/lldp/interfaces/interface"
// Name: string
func (n *LldpPathAny) Interface(Name string) *Lldp_InterfacePathAny {
	return &Lldp_InterfacePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"interfaces", "interface"},
			map[string]interface{}{"name": Name},
			n,
		),
	}
}

// SuppressTlvAdvertisement (leaf-list): Indicates whether the local system should suppress the
// advertisement of particular TLVs with the LLDP PDUs that it
// transmits. Where a TLV type is specified within this list, it
// should not be included in any LLDP PDU transmitted by the
// local agent.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/suppress-tlv-advertisement"
// Path from root: "/lldp/state/suppress-tlv-advertisement"
func (n *LldpPath) SuppressTlvAdvertisement() *Lldp_SuppressTlvAdvertisementPath {
	return &Lldp_SuppressTlvAdvertisementPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "suppress-tlv-advertisement"},
			map[string]interface{}{},
			n,
		),
	}
}

// SuppressTlvAdvertisement (leaf-list): Indicates whether the local system should suppress the
// advertisement of particular TLVs with the LLDP PDUs that it
// transmits. Where a TLV type is specified within this list, it
// should not be included in any LLDP PDU transmitted by the
// local agent.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/suppress-tlv-advertisement"
// Path from root: "/lldp/state/suppress-tlv-advertisement"
func (n *LldpPathAny) SuppressTlvAdvertisement() *Lldp_SuppressTlvAdvertisementPathAny {
	return &Lldp_SuppressTlvAdvertisementPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "suppress-tlv-advertisement"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemDescription (leaf): The system description field shall contain an alpha-numeric
// string that is the textual description of the network entity.
// The system description should include the full name and
// version identification of the system's hardware type,
// software operating system, and networking software. If
// implementations support IETF RFC 3418, the sysDescr object
// should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/system-description"
// Path from root: "/lldp/state/system-description"
func (n *LldpPath) SystemDescription() *Lldp_SystemDescriptionPath {
	return &Lldp_SystemDescriptionPath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "system-description"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemDescription (leaf): The system description field shall contain an alpha-numeric
// string that is the textual description of the network entity.
// The system description should include the full name and
// version identification of the system's hardware type,
// software operating system, and networking software. If
// implementations support IETF RFC 3418, the sysDescr object
// should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/system-description"
// Path from root: "/lldp/state/system-description"
func (n *LldpPathAny) SystemDescription() *Lldp_SystemDescriptionPathAny {
	return &Lldp_SystemDescriptionPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "system-description"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemName (leaf): The system name field shall contain an alpha-numeric string
// that indicates the system's administratively assigned name.
// The system name should be the system's fully qualified domain
// name. If implementations support IETF RFC 3418, the sysName
// object should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/system-name"
// Path from root: "/lldp/state/system-name"
func (n *LldpPath) SystemName() *Lldp_SystemNamePath {
	return &Lldp_SystemNamePath{
		NodePath: ygot.NewNodePath(
			[]string{"state", "system-name"},
			map[string]interface{}{},
			n,
		),
	}
}

// SystemName (leaf): The system name field shall contain an alpha-numeric string
// that indicates the system's administratively assigned name.
// The system name should be the system's fully qualified domain
// name. If implementations support IETF RFC 3418, the sysName
// object should be used for this field.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "state/system-name"
// Path from root: "/lldp/state/system-name"
func (n *LldpPathAny) SystemName() *Lldp_SystemNamePathAny {
	return &Lldp_SystemNamePathAny{
		NodePath: ygot.NewNodePath(
			[]string{"state", "system-name"},
			map[string]interface{}{},
			n,
		),
	}
}

// Lldp_CountersPath represents the /openconfig-lldp/lldp/state/counters YANG schema element.
type Lldp_CountersPath struct {
	*ygot.NodePath
}

// Lldp_CountersPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters YANG schema element.
type Lldp_CountersPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_EntriesAgedOutPath represents the /openconfig-lldp/lldp/state/counters/entries-aged-out YANG schema element.
type Lldp_Counters_EntriesAgedOutPath struct {
	*ygot.NodePath
}

// Lldp_Counters_EntriesAgedOutPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/entries-aged-out YANG schema element.
type Lldp_Counters_EntriesAgedOutPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameDiscardPath represents the /openconfig-lldp/lldp/state/counters/frame-discard YANG schema element.
type Lldp_Counters_FrameDiscardPath struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameDiscardPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/frame-discard YANG schema element.
type Lldp_Counters_FrameDiscardPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameErrorInPath represents the /openconfig-lldp/lldp/state/counters/frame-error-in YANG schema element.
type Lldp_Counters_FrameErrorInPath struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameErrorInPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/frame-error-in YANG schema element.
type Lldp_Counters_FrameErrorInPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameInPath represents the /openconfig-lldp/lldp/state/counters/frame-in YANG schema element.
type Lldp_Counters_FrameInPath struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameInPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/frame-in YANG schema element.
type Lldp_Counters_FrameInPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameOutPath represents the /openconfig-lldp/lldp/state/counters/frame-out YANG schema element.
type Lldp_Counters_FrameOutPath struct {
	*ygot.NodePath
}

// Lldp_Counters_FrameOutPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/frame-out YANG schema element.
type Lldp_Counters_FrameOutPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_LastClearPath represents the /openconfig-lldp/lldp/state/counters/last-clear YANG schema element.
type Lldp_Counters_LastClearPath struct {
	*ygot.NodePath
}

// Lldp_Counters_LastClearPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/last-clear YANG schema element.
type Lldp_Counters_LastClearPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_TlvAcceptedPath represents the /openconfig-lldp/lldp/state/counters/tlv-accepted YANG schema element.
type Lldp_Counters_TlvAcceptedPath struct {
	*ygot.NodePath
}

// Lldp_Counters_TlvAcceptedPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/tlv-accepted YANG schema element.
type Lldp_Counters_TlvAcceptedPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_TlvDiscardPath represents the /openconfig-lldp/lldp/state/counters/tlv-discard YANG schema element.
type Lldp_Counters_TlvDiscardPath struct {
	*ygot.NodePath
}

// Lldp_Counters_TlvDiscardPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/tlv-discard YANG schema element.
type Lldp_Counters_TlvDiscardPathAny struct {
	*ygot.NodePath
}

// Lldp_Counters_TlvUnknownPath represents the /openconfig-lldp/lldp/state/counters/tlv-unknown YANG schema element.
type Lldp_Counters_TlvUnknownPath struct {
	*ygot.NodePath
}

// Lldp_Counters_TlvUnknownPathAny represents the wildcard version of the /openconfig-lldp/lldp/state/counters/tlv-unknown YANG schema element.
type Lldp_Counters_TlvUnknownPathAny struct {
	*ygot.NodePath
}

// EntriesAgedOut (leaf): The number of entries aged out due to timeout.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "entries-aged-out"
// Path from root: "/lldp/state/counters/entries-aged-out"
func (n *Lldp_CountersPath) EntriesAgedOut() *Lldp_Counters_EntriesAgedOutPath {
	return &Lldp_Counters_EntriesAgedOutPath{
		NodePath: ygot.NewNodePath(
			[]string{"entries-aged-out"},
			map[string]interface{}{},
			n,
		),
	}
}

// EntriesAgedOut (leaf): The number of entries aged out due to timeout.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "entries-aged-out"
// Path from root: "/lldp/state/counters/entries-aged-out"
func (n *Lldp_CountersPathAny) EntriesAgedOut() *Lldp_Counters_EntriesAgedOutPathAny {
	return &Lldp_Counters_EntriesAgedOutPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"entries-aged-out"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameDiscard (leaf): The number of LLDP frames received and discarded.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-discard"
// Path from root: "/lldp/state/counters/frame-discard"
func (n *Lldp_CountersPath) FrameDiscard() *Lldp_Counters_FrameDiscardPath {
	return &Lldp_Counters_FrameDiscardPath{
		NodePath: ygot.NewNodePath(
			[]string{"frame-discard"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameDiscard (leaf): The number of LLDP frames received and discarded.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-discard"
// Path from root: "/lldp/state/counters/frame-discard"
func (n *Lldp_CountersPathAny) FrameDiscard() *Lldp_Counters_FrameDiscardPathAny {
	return &Lldp_Counters_FrameDiscardPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"frame-discard"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameErrorIn (leaf): The number of LLDP frames received with errors.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-error-in"
// Path from root: "/lldp/state/counters/frame-error-in"
func (n *Lldp_CountersPath) FrameErrorIn() *Lldp_Counters_FrameErrorInPath {
	return &Lldp_Counters_FrameErrorInPath{
		NodePath: ygot.NewNodePath(
			[]string{"frame-error-in"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameErrorIn (leaf): The number of LLDP frames received with errors.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-error-in"
// Path from root: "/lldp/state/counters/frame-error-in"
func (n *Lldp_CountersPathAny) FrameErrorIn() *Lldp_Counters_FrameErrorInPathAny {
	return &Lldp_Counters_FrameErrorInPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"frame-error-in"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameIn (leaf): The number of lldp frames received.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-in"
// Path from root: "/lldp/state/counters/frame-in"
func (n *Lldp_CountersPath) FrameIn() *Lldp_Counters_FrameInPath {
	return &Lldp_Counters_FrameInPath{
		NodePath: ygot.NewNodePath(
			[]string{"frame-in"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameIn (leaf): The number of lldp frames received.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-in"
// Path from root: "/lldp/state/counters/frame-in"
func (n *Lldp_CountersPathAny) FrameIn() *Lldp_Counters_FrameInPathAny {
	return &Lldp_Counters_FrameInPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"frame-in"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameOut (leaf): The number of frames transmitted out.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-out"
// Path from root: "/lldp/state/counters/frame-out"
func (n *Lldp_CountersPath) FrameOut() *Lldp_Counters_FrameOutPath {
	return &Lldp_Counters_FrameOutPath{
		NodePath: ygot.NewNodePath(
			[]string{"frame-out"},
			map[string]interface{}{},
			n,
		),
	}
}

// FrameOut (leaf): The number of frames transmitted out.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "frame-out"
// Path from root: "/lldp/state/counters/frame-out"
func (n *Lldp_CountersPathAny) FrameOut() *Lldp_Counters_FrameOutPathAny {
	return &Lldp_Counters_FrameOutPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"frame-out"},
			map[string]interface{}{},
			n,
		),
	}
}

// LastClear (leaf): Indicates the last time the counters were
// cleared.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "last-clear"
// Path from root: "/lldp/state/counters/last-clear"
func (n *Lldp_CountersPath) LastClear() *Lldp_Counters_LastClearPath {
	return &Lldp_Counters_LastClearPath{
		NodePath: ygot.NewNodePath(
			[]string{"last-clear"},
			map[string]interface{}{},
			n,
		),
	}
}

// LastClear (leaf): Indicates the last time the counters were
// cleared.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "last-clear"
// Path from root: "/lldp/state/counters/last-clear"
func (n *Lldp_CountersPathAny) LastClear() *Lldp_Counters_LastClearPathAny {
	return &Lldp_Counters_LastClearPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"last-clear"},
			map[string]interface{}{},
			n,
		),
	}
}

// TlvAccepted (leaf): The number of valid TLVs received.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "tlv-accepted"
// Path from root: "/lldp/state/counters/tlv-accepted"
func (n *Lldp_CountersPath) TlvAccepted() *Lldp_Counters_TlvAcceptedPath {
	return &Lldp_Counters_TlvAcceptedPath{
		NodePath: ygot.NewNodePath(
			[]string{"tlv-accepted"},
			map[string]interface{}{},
			n,
		),
	}
}

// TlvAccepted (leaf): The number of valid TLVs received.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "tlv-accepted"
// Path from root: "/lldp/state/counters/tlv-accepted"
func (n *Lldp_CountersPathAny) TlvAccepted() *Lldp_Counters_TlvAcceptedPathAny {
	return &Lldp_Counters_TlvAcceptedPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"tlv-accepted"},
			map[string]interface{}{},
			n,
		),
	}
}

// TlvDiscard (leaf): The number of TLV frames received and discarded.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "tlv-discard"
// Path from root: "/lldp/state/counters/tlv-discard"
func (n *Lldp_CountersPath) TlvDiscard() *Lldp_Counters_TlvDiscardPath {
	return &Lldp_Counters_TlvDiscardPath{
		NodePath: ygot.NewNodePath(
			[]string{"tlv-discard"},
			map[string]interface{}{},
			n,
		),
	}
}

// TlvDiscard (leaf): The number of TLV frames received and discarded.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "tlv-discard"
// Path from root: "/lldp/state/counters/tlv-discard"
func (n *Lldp_CountersPathAny) TlvDiscard() *Lldp_Counters_TlvDiscardPathAny {
	return &Lldp_Counters_TlvDiscardPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"tlv-discard"},
			map[string]interface{}{},
			n,
		),
	}
}

// TlvUnknown (leaf): The number of frames received with unknown TLV.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "tlv-unknown"
// Path from root: "/lldp/state/counters/tlv-unknown"
func (n *Lldp_CountersPath) TlvUnknown() *Lldp_Counters_TlvUnknownPath {
	return &Lldp_Counters_TlvUnknownPath{
		NodePath: ygot.NewNodePath(
			[]string{"tlv-unknown"},
			map[string]interface{}{},
			n,
		),
	}
}

// TlvUnknown (leaf): The number of frames received with unknown TLV.
// ----------------------------------------
// Defining module: "openconfig-lldp"
// Instantiating module: "openconfig-lldp"
// Path from parent: "tlv-unknown"
// Path from root: "/lldp/state/counters/tlv-unknown"
func (n *Lldp_CountersPathAny) TlvUnknown() *Lldp_Counters_TlvUnknownPathAny {
	return &Lldp_Counters_TlvUnknownPathAny{
		NodePath: ygot.NewNodePath(
			[]string{"tlv-unknown"},
			map[string]interface{}{},
			n,
		),
	}
}
