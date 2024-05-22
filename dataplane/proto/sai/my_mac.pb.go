// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: dataplane/proto/sai/my_mac.proto

package sai

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MyMacAttr int32

const (
	MyMacAttr_MY_MAC_ATTR_UNSPECIFIED      MyMacAttr = 0
	MyMacAttr_MY_MAC_ATTR_PRIORITY         MyMacAttr = 1
	MyMacAttr_MY_MAC_ATTR_PORT_ID          MyMacAttr = 2
	MyMacAttr_MY_MAC_ATTR_VLAN_ID          MyMacAttr = 3
	MyMacAttr_MY_MAC_ATTR_MAC_ADDRESS      MyMacAttr = 4
	MyMacAttr_MY_MAC_ATTR_MAC_ADDRESS_MASK MyMacAttr = 5
)

// Enum value maps for MyMacAttr.
var (
	MyMacAttr_name = map[int32]string{
		0: "MY_MAC_ATTR_UNSPECIFIED",
		1: "MY_MAC_ATTR_PRIORITY",
		2: "MY_MAC_ATTR_PORT_ID",
		3: "MY_MAC_ATTR_VLAN_ID",
		4: "MY_MAC_ATTR_MAC_ADDRESS",
		5: "MY_MAC_ATTR_MAC_ADDRESS_MASK",
	}
	MyMacAttr_value = map[string]int32{
		"MY_MAC_ATTR_UNSPECIFIED":      0,
		"MY_MAC_ATTR_PRIORITY":         1,
		"MY_MAC_ATTR_PORT_ID":          2,
		"MY_MAC_ATTR_VLAN_ID":          3,
		"MY_MAC_ATTR_MAC_ADDRESS":      4,
		"MY_MAC_ATTR_MAC_ADDRESS_MASK": 5,
	}
)

func (x MyMacAttr) Enum() *MyMacAttr {
	p := new(MyMacAttr)
	*p = x
	return p
}

func (x MyMacAttr) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MyMacAttr) Descriptor() protoreflect.EnumDescriptor {
	return file_dataplane_proto_sai_my_mac_proto_enumTypes[0].Descriptor()
}

func (MyMacAttr) Type() protoreflect.EnumType {
	return &file_dataplane_proto_sai_my_mac_proto_enumTypes[0]
}

func (x MyMacAttr) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MyMacAttr.Descriptor instead.
func (MyMacAttr) EnumDescriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{0}
}

type CreateMyMacRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Switch         uint64  `protobuf:"varint,1,opt,name=switch,proto3" json:"switch,omitempty"`
	Priority       *uint32 `protobuf:"varint,2,opt,name=priority,proto3,oneof" json:"priority,omitempty"`
	PortId         *uint64 `protobuf:"varint,3,opt,name=port_id,json=portId,proto3,oneof" json:"port_id,omitempty"`
	VlanId         *uint32 `protobuf:"varint,4,opt,name=vlan_id,json=vlanId,proto3,oneof" json:"vlan_id,omitempty"`
	MacAddress     []byte  `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3,oneof" json:"mac_address,omitempty"`
	MacAddressMask []byte  `protobuf:"bytes,6,opt,name=mac_address_mask,json=macAddressMask,proto3,oneof" json:"mac_address_mask,omitempty"`
}

func (x *CreateMyMacRequest) Reset() {
	*x = CreateMyMacRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMyMacRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMyMacRequest) ProtoMessage() {}

func (x *CreateMyMacRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMyMacRequest.ProtoReflect.Descriptor instead.
func (*CreateMyMacRequest) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMyMacRequest) GetSwitch() uint64 {
	if x != nil {
		return x.Switch
	}
	return 0
}

func (x *CreateMyMacRequest) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

func (x *CreateMyMacRequest) GetPortId() uint64 {
	if x != nil && x.PortId != nil {
		return *x.PortId
	}
	return 0
}

func (x *CreateMyMacRequest) GetVlanId() uint32 {
	if x != nil && x.VlanId != nil {
		return *x.VlanId
	}
	return 0
}

func (x *CreateMyMacRequest) GetMacAddress() []byte {
	if x != nil {
		return x.MacAddress
	}
	return nil
}

func (x *CreateMyMacRequest) GetMacAddressMask() []byte {
	if x != nil {
		return x.MacAddressMask
	}
	return nil
}

type CreateMyMacResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid uint64 `protobuf:"varint,1,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *CreateMyMacResponse) Reset() {
	*x = CreateMyMacResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMyMacResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMyMacResponse) ProtoMessage() {}

func (x *CreateMyMacResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMyMacResponse.ProtoReflect.Descriptor instead.
func (*CreateMyMacResponse) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMyMacResponse) GetOid() uint64 {
	if x != nil {
		return x.Oid
	}
	return 0
}

type RemoveMyMacRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid uint64 `protobuf:"varint,1,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *RemoveMyMacRequest) Reset() {
	*x = RemoveMyMacRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMyMacRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMyMacRequest) ProtoMessage() {}

func (x *RemoveMyMacRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMyMacRequest.ProtoReflect.Descriptor instead.
func (*RemoveMyMacRequest) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveMyMacRequest) GetOid() uint64 {
	if x != nil {
		return x.Oid
	}
	return 0
}

type RemoveMyMacResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveMyMacResponse) Reset() {
	*x = RemoveMyMacResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMyMacResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMyMacResponse) ProtoMessage() {}

func (x *RemoveMyMacResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMyMacResponse.ProtoReflect.Descriptor instead.
func (*RemoveMyMacResponse) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{3}
}

type SetMyMacAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid      uint64  `protobuf:"varint,1,opt,name=oid,proto3" json:"oid,omitempty"`
	Priority *uint32 `protobuf:"varint,2,opt,name=priority,proto3,oneof" json:"priority,omitempty"`
}

func (x *SetMyMacAttributeRequest) Reset() {
	*x = SetMyMacAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMyMacAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMyMacAttributeRequest) ProtoMessage() {}

func (x *SetMyMacAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMyMacAttributeRequest.ProtoReflect.Descriptor instead.
func (*SetMyMacAttributeRequest) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{4}
}

func (x *SetMyMacAttributeRequest) GetOid() uint64 {
	if x != nil {
		return x.Oid
	}
	return 0
}

func (x *SetMyMacAttributeRequest) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

type SetMyMacAttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetMyMacAttributeResponse) Reset() {
	*x = SetMyMacAttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMyMacAttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMyMacAttributeResponse) ProtoMessage() {}

func (x *SetMyMacAttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMyMacAttributeResponse.ProtoReflect.Descriptor instead.
func (*SetMyMacAttributeResponse) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{5}
}

type GetMyMacAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid      uint64      `protobuf:"varint,1,opt,name=oid,proto3" json:"oid,omitempty"`
	AttrType []MyMacAttr `protobuf:"varint,2,rep,packed,name=attr_type,json=attrType,proto3,enum=lemming.dataplane.sai.MyMacAttr" json:"attr_type,omitempty"`
}

func (x *GetMyMacAttributeRequest) Reset() {
	*x = GetMyMacAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyMacAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyMacAttributeRequest) ProtoMessage() {}

func (x *GetMyMacAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyMacAttributeRequest.ProtoReflect.Descriptor instead.
func (*GetMyMacAttributeRequest) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{6}
}

func (x *GetMyMacAttributeRequest) GetOid() uint64 {
	if x != nil {
		return x.Oid
	}
	return 0
}

func (x *GetMyMacAttributeRequest) GetAttrType() []MyMacAttr {
	if x != nil {
		return x.AttrType
	}
	return nil
}

type GetMyMacAttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attr *MyMacAttribute `protobuf:"bytes,1,opt,name=attr,proto3" json:"attr,omitempty"`
}

func (x *GetMyMacAttributeResponse) Reset() {
	*x = GetMyMacAttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyMacAttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyMacAttributeResponse) ProtoMessage() {}

func (x *GetMyMacAttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dataplane_proto_sai_my_mac_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyMacAttributeResponse.ProtoReflect.Descriptor instead.
func (*GetMyMacAttributeResponse) Descriptor() ([]byte, []int) {
	return file_dataplane_proto_sai_my_mac_proto_rawDescGZIP(), []int{7}
}

func (x *GetMyMacAttributeResponse) GetAttr() *MyMacAttribute {
	if x != nil {
		return x.Attr
	}
	return nil
}

var File_dataplane_proto_sai_my_mac_proto protoreflect.FileDescriptor

var file_dataplane_proto_sai_my_mac_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x61, 0x69, 0x2f, 0x6d, 0x79, 0x5f, 0x6d, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x1a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xf0, 0xdc,
	0x93, 0xad, 0x0f, 0x01, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x06, 0xf0, 0xdc, 0x93, 0xad, 0x0f, 0x02, 0x48, 0x01, 0x52, 0x06,
	0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x07, 0x76, 0x6c, 0x61,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xf0, 0xdc, 0x93, 0xad,
	0x0f, 0x03, 0x48, 0x02, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x2c, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xf0, 0xdc, 0x93, 0xad, 0x0f, 0x04, 0x48, 0x03, 0x52, 0x0a,
	0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a,
	0x10, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xf0, 0xdc, 0x93, 0xad, 0x0f, 0x05, 0x48,
	0x04, 0x52, 0x0e, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x73,
	0x6b, 0x88, 0x01, 0x01, 0x3a, 0x06, 0xa0, 0xa9, 0x90, 0xad, 0x0f, 0x62, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x27, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x22,
	0x26, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62,
	0x0a, 0x18, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06,
	0xf0, 0xdc, 0x93, 0xad, 0x0f, 0x01, 0x48, 0x00, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x6b, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x12, 0x3d, 0x0a,
	0x09, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74,
	0x74, 0x72, 0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x56, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x61, 0x74, 0x74,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e,
	0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x2a, 0xb3, 0x01, 0x0a, 0x09, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74,
	0x74, 0x72, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x54, 0x54,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x50,
	0x52, 0x49, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x59, 0x5f,
	0x4d, 0x41, 0x43, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x49, 0x44,
	0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x59, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x54, 0x54,
	0x52, 0x5f, 0x56, 0x4c, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x4d,
	0x59, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41,
	0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x59, 0x5f, 0x4d,
	0x41, 0x43, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x52,
	0x45, 0x53, 0x53, 0x5f, 0x4d, 0x41, 0x53, 0x4b, 0x10, 0x05, 0x32, 0xcb, 0x03, 0x0a, 0x05, 0x4d,
	0x79, 0x4d, 0x61, 0x63, 0x12, 0x66, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x79,
	0x4d, 0x61, 0x63, 0x12, 0x29, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x79, 0x4d,
	0x61, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0b,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x12, 0x29, 0x2e, 0x6c, 0x65,
	0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x73, 0x61, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6c, 0x65, 0x6d, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61,
	0x69, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6c, 0x65, 0x6d,
	0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73,
	0x61, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x73, 0x61, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x79, 0x4d, 0x61, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dataplane_proto_sai_my_mac_proto_rawDescOnce sync.Once
	file_dataplane_proto_sai_my_mac_proto_rawDescData = file_dataplane_proto_sai_my_mac_proto_rawDesc
)

func file_dataplane_proto_sai_my_mac_proto_rawDescGZIP() []byte {
	file_dataplane_proto_sai_my_mac_proto_rawDescOnce.Do(func() {
		file_dataplane_proto_sai_my_mac_proto_rawDescData = protoimpl.X.CompressGZIP(file_dataplane_proto_sai_my_mac_proto_rawDescData)
	})
	return file_dataplane_proto_sai_my_mac_proto_rawDescData
}

var file_dataplane_proto_sai_my_mac_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dataplane_proto_sai_my_mac_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dataplane_proto_sai_my_mac_proto_goTypes = []interface{}{
	(MyMacAttr)(0),                    // 0: lemming.dataplane.sai.MyMacAttr
	(*CreateMyMacRequest)(nil),        // 1: lemming.dataplane.sai.CreateMyMacRequest
	(*CreateMyMacResponse)(nil),       // 2: lemming.dataplane.sai.CreateMyMacResponse
	(*RemoveMyMacRequest)(nil),        // 3: lemming.dataplane.sai.RemoveMyMacRequest
	(*RemoveMyMacResponse)(nil),       // 4: lemming.dataplane.sai.RemoveMyMacResponse
	(*SetMyMacAttributeRequest)(nil),  // 5: lemming.dataplane.sai.SetMyMacAttributeRequest
	(*SetMyMacAttributeResponse)(nil), // 6: lemming.dataplane.sai.SetMyMacAttributeResponse
	(*GetMyMacAttributeRequest)(nil),  // 7: lemming.dataplane.sai.GetMyMacAttributeRequest
	(*GetMyMacAttributeResponse)(nil), // 8: lemming.dataplane.sai.GetMyMacAttributeResponse
	(*MyMacAttribute)(nil),            // 9: lemming.dataplane.sai.MyMacAttribute
}
var file_dataplane_proto_sai_my_mac_proto_depIdxs = []int32{
	0, // 0: lemming.dataplane.sai.GetMyMacAttributeRequest.attr_type:type_name -> lemming.dataplane.sai.MyMacAttr
	9, // 1: lemming.dataplane.sai.GetMyMacAttributeResponse.attr:type_name -> lemming.dataplane.sai.MyMacAttribute
	1, // 2: lemming.dataplane.sai.MyMac.CreateMyMac:input_type -> lemming.dataplane.sai.CreateMyMacRequest
	3, // 3: lemming.dataplane.sai.MyMac.RemoveMyMac:input_type -> lemming.dataplane.sai.RemoveMyMacRequest
	5, // 4: lemming.dataplane.sai.MyMac.SetMyMacAttribute:input_type -> lemming.dataplane.sai.SetMyMacAttributeRequest
	7, // 5: lemming.dataplane.sai.MyMac.GetMyMacAttribute:input_type -> lemming.dataplane.sai.GetMyMacAttributeRequest
	2, // 6: lemming.dataplane.sai.MyMac.CreateMyMac:output_type -> lemming.dataplane.sai.CreateMyMacResponse
	4, // 7: lemming.dataplane.sai.MyMac.RemoveMyMac:output_type -> lemming.dataplane.sai.RemoveMyMacResponse
	6, // 8: lemming.dataplane.sai.MyMac.SetMyMacAttribute:output_type -> lemming.dataplane.sai.SetMyMacAttributeResponse
	8, // 9: lemming.dataplane.sai.MyMac.GetMyMacAttribute:output_type -> lemming.dataplane.sai.GetMyMacAttributeResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dataplane_proto_sai_my_mac_proto_init() }
func file_dataplane_proto_sai_my_mac_proto_init() {
	if File_dataplane_proto_sai_my_mac_proto != nil {
		return
	}
	file_dataplane_proto_sai_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dataplane_proto_sai_my_mac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMyMacRequest); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMyMacResponse); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMyMacRequest); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMyMacResponse); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMyMacAttributeRequest); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMyMacAttributeResponse); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyMacAttributeRequest); i {
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
		file_dataplane_proto_sai_my_mac_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyMacAttributeResponse); i {
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
	file_dataplane_proto_sai_my_mac_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_dataplane_proto_sai_my_mac_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dataplane_proto_sai_my_mac_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dataplane_proto_sai_my_mac_proto_goTypes,
		DependencyIndexes: file_dataplane_proto_sai_my_mac_proto_depIdxs,
		EnumInfos:         file_dataplane_proto_sai_my_mac_proto_enumTypes,
		MessageInfos:      file_dataplane_proto_sai_my_mac_proto_msgTypes,
	}.Build()
	File_dataplane_proto_sai_my_mac_proto = out.File
	file_dataplane_proto_sai_my_mac_proto_rawDesc = nil
	file_dataplane_proto_sai_my_mac_proto_goTypes = nil
	file_dataplane_proto_sai_my_mac_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MyMacClient is the client API for MyMac service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyMacClient interface {
	CreateMyMac(ctx context.Context, in *CreateMyMacRequest, opts ...grpc.CallOption) (*CreateMyMacResponse, error)
	RemoveMyMac(ctx context.Context, in *RemoveMyMacRequest, opts ...grpc.CallOption) (*RemoveMyMacResponse, error)
	SetMyMacAttribute(ctx context.Context, in *SetMyMacAttributeRequest, opts ...grpc.CallOption) (*SetMyMacAttributeResponse, error)
	GetMyMacAttribute(ctx context.Context, in *GetMyMacAttributeRequest, opts ...grpc.CallOption) (*GetMyMacAttributeResponse, error)
}

type myMacClient struct {
	cc grpc.ClientConnInterface
}

func NewMyMacClient(cc grpc.ClientConnInterface) MyMacClient {
	return &myMacClient{cc}
}

func (c *myMacClient) CreateMyMac(ctx context.Context, in *CreateMyMacRequest, opts ...grpc.CallOption) (*CreateMyMacResponse, error) {
	out := new(CreateMyMacResponse)
	err := c.cc.Invoke(ctx, "/lemming.dataplane.sai.MyMac/CreateMyMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMacClient) RemoveMyMac(ctx context.Context, in *RemoveMyMacRequest, opts ...grpc.CallOption) (*RemoveMyMacResponse, error) {
	out := new(RemoveMyMacResponse)
	err := c.cc.Invoke(ctx, "/lemming.dataplane.sai.MyMac/RemoveMyMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMacClient) SetMyMacAttribute(ctx context.Context, in *SetMyMacAttributeRequest, opts ...grpc.CallOption) (*SetMyMacAttributeResponse, error) {
	out := new(SetMyMacAttributeResponse)
	err := c.cc.Invoke(ctx, "/lemming.dataplane.sai.MyMac/SetMyMacAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMacClient) GetMyMacAttribute(ctx context.Context, in *GetMyMacAttributeRequest, opts ...grpc.CallOption) (*GetMyMacAttributeResponse, error) {
	out := new(GetMyMacAttributeResponse)
	err := c.cc.Invoke(ctx, "/lemming.dataplane.sai.MyMac/GetMyMacAttribute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyMacServer is the server API for MyMac service.
type MyMacServer interface {
	CreateMyMac(context.Context, *CreateMyMacRequest) (*CreateMyMacResponse, error)
	RemoveMyMac(context.Context, *RemoveMyMacRequest) (*RemoveMyMacResponse, error)
	SetMyMacAttribute(context.Context, *SetMyMacAttributeRequest) (*SetMyMacAttributeResponse, error)
	GetMyMacAttribute(context.Context, *GetMyMacAttributeRequest) (*GetMyMacAttributeResponse, error)
}

// UnimplementedMyMacServer can be embedded to have forward compatible implementations.
type UnimplementedMyMacServer struct {
}

func (*UnimplementedMyMacServer) CreateMyMac(context.Context, *CreateMyMacRequest) (*CreateMyMacResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMyMac not implemented")
}
func (*UnimplementedMyMacServer) RemoveMyMac(context.Context, *RemoveMyMacRequest) (*RemoveMyMacResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMyMac not implemented")
}
func (*UnimplementedMyMacServer) SetMyMacAttribute(context.Context, *SetMyMacAttributeRequest) (*SetMyMacAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMyMacAttribute not implemented")
}
func (*UnimplementedMyMacServer) GetMyMacAttribute(context.Context, *GetMyMacAttributeRequest) (*GetMyMacAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyMacAttribute not implemented")
}

func RegisterMyMacServer(s *grpc.Server, srv MyMacServer) {
	s.RegisterService(&_MyMac_serviceDesc, srv)
}

func _MyMac_CreateMyMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMyMacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMacServer).CreateMyMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemming.dataplane.sai.MyMac/CreateMyMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMacServer).CreateMyMac(ctx, req.(*CreateMyMacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMac_RemoveMyMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMyMacRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMacServer).RemoveMyMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemming.dataplane.sai.MyMac/RemoveMyMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMacServer).RemoveMyMac(ctx, req.(*RemoveMyMacRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMac_SetMyMacAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMyMacAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMacServer).SetMyMacAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemming.dataplane.sai.MyMac/SetMyMacAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMacServer).SetMyMacAttribute(ctx, req.(*SetMyMacAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMac_GetMyMacAttribute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyMacAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMacServer).GetMyMacAttribute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lemming.dataplane.sai.MyMac/GetMyMacAttribute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMacServer).GetMyMacAttribute(ctx, req.(*GetMyMacAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyMac_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lemming.dataplane.sai.MyMac",
	HandlerType: (*MyMacServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMyMac",
			Handler:    _MyMac_CreateMyMac_Handler,
		},
		{
			MethodName: "RemoveMyMac",
			Handler:    _MyMac_RemoveMyMac_Handler,
		},
		{
			MethodName: "SetMyMacAttribute",
			Handler:    _MyMac_SetMyMacAttribute_Handler,
		},
		{
			MethodName: "GetMyMacAttribute",
			Handler:    _MyMac_GetMyMacAttribute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplane/proto/sai/my_mac.proto",
}
