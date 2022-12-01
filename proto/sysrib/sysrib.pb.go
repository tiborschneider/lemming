// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: sysrib.proto

package sysrib

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetRouteRequest_Safi int32

const (
	SetRouteRequest_SAFI_UNSPECIFIED SetRouteRequest_Safi = 0
	SetRouteRequest_SAFI_UNICAST     SetRouteRequest_Safi = 1 //MULTICAST = 1;
)

// Enum value maps for SetRouteRequest_Safi.
var (
	SetRouteRequest_Safi_name = map[int32]string{
		0: "SAFI_UNSPECIFIED",
		1: "SAFI_UNICAST",
	}
	SetRouteRequest_Safi_value = map[string]int32{
		"SAFI_UNSPECIFIED": 0,
		"SAFI_UNICAST":     1,
	}
)

func (x SetRouteRequest_Safi) Enum() *SetRouteRequest_Safi {
	p := new(SetRouteRequest_Safi)
	*p = x
	return p
}

func (x SetRouteRequest_Safi) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetRouteRequest_Safi) Descriptor() protoreflect.EnumDescriptor {
	return file_sysrib_proto_enumTypes[0].Descriptor()
}

func (SetRouteRequest_Safi) Type() protoreflect.EnumType {
	return &file_sysrib_proto_enumTypes[0]
}

func (x SetRouteRequest_Safi) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetRouteRequest_Safi.Descriptor instead.
func (SetRouteRequest_Safi) EnumDescriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{0, 0}
}

type Prefix_Family int32

const (
	Prefix_FAMILY_UNSPECIFIED Prefix_Family = 0
	Prefix_FAMILY_IPV4        Prefix_Family = 1 //IPv6 = 1;
)

// Enum value maps for Prefix_Family.
var (
	Prefix_Family_name = map[int32]string{
		0: "FAMILY_UNSPECIFIED",
		1: "FAMILY_IPV4",
	}
	Prefix_Family_value = map[string]int32{
		"FAMILY_UNSPECIFIED": 0,
		"FAMILY_IPV4":        1,
	}
)

func (x Prefix_Family) Enum() *Prefix_Family {
	p := new(Prefix_Family)
	*p = x
	return p
}

func (x Prefix_Family) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prefix_Family) Descriptor() protoreflect.EnumDescriptor {
	return file_sysrib_proto_enumTypes[1].Descriptor()
}

func (Prefix_Family) Type() protoreflect.EnumType {
	return &file_sysrib_proto_enumTypes[1]
}

func (x Prefix_Family) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Prefix_Family.Descriptor instead.
func (Prefix_Family) EnumDescriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{1, 0}
}

type Nexthop_Type int32

const (
	Nexthop_TYPE_UNSPECIFIED Nexthop_Type = 0
	Nexthop_TYPE_IPV4        Nexthop_Type = 1
)

// Enum value maps for Nexthop_Type.
var (
	Nexthop_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_IPV4",
	}
	Nexthop_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_IPV4":        1,
	}
)

func (x Nexthop_Type) Enum() *Nexthop_Type {
	p := new(Nexthop_Type)
	*p = x
	return p
}

func (x Nexthop_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nexthop_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sysrib_proto_enumTypes[2].Descriptor()
}

func (Nexthop_Type) Type() protoreflect.EnumType {
	return &file_sysrib_proto_enumTypes[2]
}

func (x Nexthop_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nexthop_Type.Descriptor instead.
func (Nexthop_Type) EnumDescriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{2, 0}
}

type SetRouteResponse_Status int32

const (
	SetRouteResponse_STATUS_UNSPECIFIED SetRouteResponse_Status = 0
	SetRouteResponse_STATUS_SUCCESS     SetRouteResponse_Status = 1
	SetRouteResponse_STATUS_FAIL        SetRouteResponse_Status = 2
)

// Enum value maps for SetRouteResponse_Status.
var (
	SetRouteResponse_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_SUCCESS",
		2: "STATUS_FAIL",
	}
	SetRouteResponse_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_SUCCESS":     1,
		"STATUS_FAIL":        2,
	}
)

func (x SetRouteResponse_Status) Enum() *SetRouteResponse_Status {
	p := new(SetRouteResponse_Status)
	*p = x
	return p
}

func (x SetRouteResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetRouteResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_sysrib_proto_enumTypes[3].Descriptor()
}

func (SetRouteResponse_Status) Type() protoreflect.EnumType {
	return &file_sysrib_proto_enumTypes[3]
}

func (x SetRouteResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetRouteResponse_Status.Descriptor instead.
func (SetRouteResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{3, 0}
}

// SetRouteRequest and its dependent messages are derived from
// ZEBRA_ROUTE_ADD/ZEBRA_ROUTE_DELETE.
type SetRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delete         bool                 `protobuf:"varint,1,opt,name=delete,proto3" json:"delete,omitempty"`
	VrfId          uint32               `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	AdminDistance  uint32               `protobuf:"varint,3,opt,name=admin_distance,json=adminDistance,proto3" json:"admin_distance,omitempty"`
	ProtocolName   string               `protobuf:"bytes,4,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	Safi           SetRouteRequest_Safi `protobuf:"varint,5,opt,name=safi,proto3,enum=sysrib.SetRouteRequest_Safi" json:"safi,omitempty"`
	Prefix         *Prefix              `protobuf:"bytes,6,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Metric         uint32               `protobuf:"varint,7,opt,name=metric,proto3" json:"metric,omitempty"`
	Nexthops       []*Nexthop           `protobuf:"bytes,8,rep,name=nexthops,proto3" json:"nexthops,omitempty"`
	BackupNexthops []*Nexthop           `protobuf:"bytes,9,rep,name=backup_nexthops,json=backupNexthops,proto3" json:"backup_nexthops,omitempty"`
}

func (x *SetRouteRequest) Reset() {
	*x = SetRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysrib_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRouteRequest) ProtoMessage() {}

func (x *SetRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sysrib_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRouteRequest.ProtoReflect.Descriptor instead.
func (*SetRouteRequest) Descriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{0}
}

func (x *SetRouteRequest) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *SetRouteRequest) GetVrfId() uint32 {
	if x != nil {
		return x.VrfId
	}
	return 0
}

func (x *SetRouteRequest) GetAdminDistance() uint32 {
	if x != nil {
		return x.AdminDistance
	}
	return 0
}

func (x *SetRouteRequest) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *SetRouteRequest) GetSafi() SetRouteRequest_Safi {
	if x != nil {
		return x.Safi
	}
	return SetRouteRequest_SAFI_UNSPECIFIED
}

func (x *SetRouteRequest) GetPrefix() *Prefix {
	if x != nil {
		return x.Prefix
	}
	return nil
}

func (x *SetRouteRequest) GetMetric() uint32 {
	if x != nil {
		return x.Metric
	}
	return 0
}

func (x *SetRouteRequest) GetNexthops() []*Nexthop {
	if x != nil {
		return x.Nexthops
	}
	return nil
}

func (x *SetRouteRequest) GetBackupNexthops() []*Nexthop {
	if x != nil {
		return x.BackupNexthops
	}
	return nil
}

// TODO(wenbli): This probably goes in some common proto file.
type Prefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family     Prefix_Family `protobuf:"varint,1,opt,name=family,proto3,enum=sysrib.Prefix_Family" json:"family,omitempty"`
	Address    string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	MaskLength uint32        `protobuf:"varint,3,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
}

func (x *Prefix) Reset() {
	*x = Prefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysrib_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prefix) ProtoMessage() {}

func (x *Prefix) ProtoReflect() protoreflect.Message {
	mi := &file_sysrib_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prefix.ProtoReflect.Descriptor instead.
func (*Prefix) Descriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{1}
}

func (x *Prefix) GetFamily() Prefix_Family {
	if x != nil {
		return x.Family
	}
	return Prefix_FAMILY_UNSPECIFIED
}

func (x *Prefix) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Prefix) GetMaskLength() uint32 {
	if x != nil {
		return x.MaskLength
	}
	return 0
}

type Nexthop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VrfId uint32       `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Type  Nexthop_Type `protobuf:"varint,2,opt,name=type,proto3,enum=sysrib.Nexthop_Type" json:"type,omitempty"`
	// int32 ifindex = 3;
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Weight  uint64 `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Nexthop) Reset() {
	*x = Nexthop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysrib_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nexthop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nexthop) ProtoMessage() {}

func (x *Nexthop) ProtoReflect() protoreflect.Message {
	mi := &file_sysrib_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nexthop.ProtoReflect.Descriptor instead.
func (*Nexthop) Descriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{2}
}

func (x *Nexthop) GetVrfId() uint32 {
	if x != nil {
		return x.VrfId
	}
	return 0
}

func (x *Nexthop) GetType() Nexthop_Type {
	if x != nil {
		return x.Type
	}
	return Nexthop_TYPE_UNSPECIFIED
}

func (x *Nexthop) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Nexthop) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type SetRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SetRouteResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=sysrib.SetRouteResponse_Status" json:"status,omitempty"` // tableid
}

func (x *SetRouteResponse) Reset() {
	*x = SetRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysrib_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRouteResponse) ProtoMessage() {}

func (x *SetRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sysrib_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRouteResponse.ProtoReflect.Descriptor instead.
func (*SetRouteResponse) Descriptor() ([]byte, []int) {
	return file_sysrib_proto_rawDescGZIP(), []int{3}
}

func (x *SetRouteResponse) GetStatus() SetRouteResponse_Status {
	if x != nil {
		return x.Status
	}
	return SetRouteResponse_STATUS_UNSPECIFIED
}

var File_sysrib_proto protoreflect.FileDescriptor

var file_sysrib_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x22, 0x95, 0x03, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x72, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x72, 0x66, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x61, 0x66, 0x69, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x2e, 0x53, 0x65, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x61, 0x66,
	0x69, 0x52, 0x04, 0x73, 0x61, 0x66, 0x69, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62,
	0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x68,
	0x6f, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x79, 0x73, 0x72,
	0x69, 0x62, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74,
	0x68, 0x6f, 0x70, 0x73, 0x12, 0x38, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e,
	0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x52, 0x0e,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x73, 0x22, 0x2e,
	0x0a, 0x04, 0x53, 0x61, 0x66, 0x69, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x41, 0x46, 0x49, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x41, 0x46, 0x49, 0x5f, 0x55, 0x4e, 0x49, 0x43, 0x41, 0x53, 0x54, 0x10, 0x01, 0x22, 0xa5,
	0x01, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x79, 0x73, 0x72,
	0x69, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x73, 0x6b, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x31, 0x0a, 0x06, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x12, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f,
	0x49, 0x50, 0x56, 0x34, 0x10, 0x01, 0x22, 0xa9, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x78, 0x74, 0x68,
	0x6f, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x72, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x72, 0x66, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62,
	0x2e, 0x4e, 0x65, 0x78, 0x74, 0x68, 0x6f, 0x70, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x2b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x50, 0x56, 0x34,
	0x10, 0x01, 0x22, 0x92, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62,
	0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x45, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x32, 0x47, 0x0a, 0x06, 0x53, 0x79, 0x73, 0x72, 0x69,
	0x62, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x2e,
	0x53, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x79, 0x73, 0x72, 0x69, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sysrib_proto_rawDescOnce sync.Once
	file_sysrib_proto_rawDescData = file_sysrib_proto_rawDesc
)

func file_sysrib_proto_rawDescGZIP() []byte {
	file_sysrib_proto_rawDescOnce.Do(func() {
		file_sysrib_proto_rawDescData = protoimpl.X.CompressGZIP(file_sysrib_proto_rawDescData)
	})
	return file_sysrib_proto_rawDescData
}

var file_sysrib_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_sysrib_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sysrib_proto_goTypes = []interface{}{
	(SetRouteRequest_Safi)(0),    // 0: sysrib.SetRouteRequest.Safi
	(Prefix_Family)(0),           // 1: sysrib.Prefix.Family
	(Nexthop_Type)(0),            // 2: sysrib.Nexthop.Type
	(SetRouteResponse_Status)(0), // 3: sysrib.SetRouteResponse.Status
	(*SetRouteRequest)(nil),      // 4: sysrib.SetRouteRequest
	(*Prefix)(nil),               // 5: sysrib.Prefix
	(*Nexthop)(nil),              // 6: sysrib.Nexthop
	(*SetRouteResponse)(nil),     // 7: sysrib.SetRouteResponse
}
var file_sysrib_proto_depIdxs = []int32{
	0, // 0: sysrib.SetRouteRequest.safi:type_name -> sysrib.SetRouteRequest.Safi
	5, // 1: sysrib.SetRouteRequest.prefix:type_name -> sysrib.Prefix
	6, // 2: sysrib.SetRouteRequest.nexthops:type_name -> sysrib.Nexthop
	6, // 3: sysrib.SetRouteRequest.backup_nexthops:type_name -> sysrib.Nexthop
	1, // 4: sysrib.Prefix.family:type_name -> sysrib.Prefix.Family
	2, // 5: sysrib.Nexthop.type:type_name -> sysrib.Nexthop.Type
	3, // 6: sysrib.SetRouteResponse.status:type_name -> sysrib.SetRouteResponse.Status
	4, // 7: sysrib.Sysrib.SetRoute:input_type -> sysrib.SetRouteRequest
	7, // 8: sysrib.Sysrib.SetRoute:output_type -> sysrib.SetRouteResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_sysrib_proto_init() }
func file_sysrib_proto_init() {
	if File_sysrib_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sysrib_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sysrib_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prefix); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sysrib_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nexthop); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sysrib_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sysrib_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sysrib_proto_goTypes,
		DependencyIndexes: file_sysrib_proto_depIdxs,
		EnumInfos:         file_sysrib_proto_enumTypes,
		MessageInfos:      file_sysrib_proto_msgTypes,
	}.Build()
	File_sysrib_proto = out.File
	file_sysrib_proto_rawDesc = nil
	file_sysrib_proto_goTypes = nil
	file_sysrib_proto_depIdxs = nil
}
