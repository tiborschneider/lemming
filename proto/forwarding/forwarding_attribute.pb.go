// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: proto/forwarding/forwarding_attribute.proto

package forwarding

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

type AttributeDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Help string `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
}

func (x *AttributeDesc) Reset() {
	*x = AttributeDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeDesc) ProtoMessage() {}

func (x *AttributeDesc) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeDesc.ProtoReflect.Descriptor instead.
func (*AttributeDesc) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeDesc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AttributeDesc) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

type AttributeListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AttributeListRequest) Reset() {
	*x = AttributeListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeListRequest) ProtoMessage() {}

func (x *AttributeListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeListRequest.ProtoReflect.Descriptor instead.
func (*AttributeListRequest) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{1}
}

type AttributeListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attrs []*AttributeDesc `protobuf:"bytes,1,rep,name=attrs,proto3" json:"attrs,omitempty"`
}

func (x *AttributeListReply) Reset() {
	*x = AttributeListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeListReply) ProtoMessage() {}

func (x *AttributeListReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeListReply.ProtoReflect.Descriptor instead.
func (*AttributeListReply) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{2}
}

func (x *AttributeListReply) GetAttrs() []*AttributeDesc {
	if x != nil {
		return x.Attrs
	}
	return nil
}

type AttributeUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextId *ContextId `protobuf:"bytes,1,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	ObjectId  *ObjectId  `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	AttrId    string     `protobuf:"bytes,3,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	AttrValue string     `protobuf:"bytes,4,opt,name=attr_value,json=attrValue,proto3" json:"attr_value,omitempty"`
}

func (x *AttributeUpdateRequest) Reset() {
	*x = AttributeUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeUpdateRequest) ProtoMessage() {}

func (x *AttributeUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeUpdateRequest.ProtoReflect.Descriptor instead.
func (*AttributeUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{3}
}

func (x *AttributeUpdateRequest) GetContextId() *ContextId {
	if x != nil {
		return x.ContextId
	}
	return nil
}

func (x *AttributeUpdateRequest) GetObjectId() *ObjectId {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

func (x *AttributeUpdateRequest) GetAttrId() string {
	if x != nil {
		return x.AttrId
	}
	return ""
}

func (x *AttributeUpdateRequest) GetAttrValue() string {
	if x != nil {
		return x.AttrValue
	}
	return ""
}

type AttributeUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AttributeUpdateReply) Reset() {
	*x = AttributeUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeUpdateReply) ProtoMessage() {}

func (x *AttributeUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeUpdateReply.ProtoReflect.Descriptor instead.
func (*AttributeUpdateReply) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{4}
}

type AttributeQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextId *ContextId `protobuf:"bytes,1,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	ObjectId  *ObjectId  `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	AttrId    string     `protobuf:"bytes,3,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
}

func (x *AttributeQueryRequest) Reset() {
	*x = AttributeQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeQueryRequest) ProtoMessage() {}

func (x *AttributeQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeQueryRequest.ProtoReflect.Descriptor instead.
func (*AttributeQueryRequest) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{5}
}

func (x *AttributeQueryRequest) GetContextId() *ContextId {
	if x != nil {
		return x.ContextId
	}
	return nil
}

func (x *AttributeQueryRequest) GetObjectId() *ObjectId {
	if x != nil {
		return x.ObjectId
	}
	return nil
}

func (x *AttributeQueryRequest) GetAttrId() string {
	if x != nil {
		return x.AttrId
	}
	return ""
}

type AttributeQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrValue string `protobuf:"bytes,1,opt,name=attr_value,json=attrValue,proto3" json:"attr_value,omitempty"`
}

func (x *AttributeQueryReply) Reset() {
	*x = AttributeQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeQueryReply) ProtoMessage() {}

func (x *AttributeQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_attribute_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeQueryReply.ProtoReflect.Descriptor instead.
func (*AttributeQueryReply) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP(), []int{6}
}

func (x *AttributeQueryReply) GetAttrValue() string {
	if x != nil {
		return x.AttrValue
	}
	return ""
}

var File_proto_forwarding_forwarding_attribute_proto protoreflect.FileDescriptor

var file_proto_forwarding_forwarding_attribute_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x22, 0x16, 0x0a, 0x14,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x12, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x61, 0x74,
	0x74, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x16,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49,
	0x64, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74,
	0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x99, 0x01, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x49, 0x64, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_forwarding_forwarding_attribute_proto_rawDescOnce sync.Once
	file_proto_forwarding_forwarding_attribute_proto_rawDescData = file_proto_forwarding_forwarding_attribute_proto_rawDesc
)

func file_proto_forwarding_forwarding_attribute_proto_rawDescGZIP() []byte {
	file_proto_forwarding_forwarding_attribute_proto_rawDescOnce.Do(func() {
		file_proto_forwarding_forwarding_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_forwarding_forwarding_attribute_proto_rawDescData)
	})
	return file_proto_forwarding_forwarding_attribute_proto_rawDescData
}

var file_proto_forwarding_forwarding_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_forwarding_forwarding_attribute_proto_goTypes = []interface{}{
	(*AttributeDesc)(nil),          // 0: forwarding.AttributeDesc
	(*AttributeListRequest)(nil),   // 1: forwarding.AttributeListRequest
	(*AttributeListReply)(nil),     // 2: forwarding.AttributeListReply
	(*AttributeUpdateRequest)(nil), // 3: forwarding.AttributeUpdateRequest
	(*AttributeUpdateReply)(nil),   // 4: forwarding.AttributeUpdateReply
	(*AttributeQueryRequest)(nil),  // 5: forwarding.AttributeQueryRequest
	(*AttributeQueryReply)(nil),    // 6: forwarding.AttributeQueryReply
	(*ContextId)(nil),              // 7: forwarding.ContextId
	(*ObjectId)(nil),               // 8: forwarding.ObjectId
}
var file_proto_forwarding_forwarding_attribute_proto_depIdxs = []int32{
	0, // 0: forwarding.AttributeListReply.attrs:type_name -> forwarding.AttributeDesc
	7, // 1: forwarding.AttributeUpdateRequest.context_id:type_name -> forwarding.ContextId
	8, // 2: forwarding.AttributeUpdateRequest.object_id:type_name -> forwarding.ObjectId
	7, // 3: forwarding.AttributeQueryRequest.context_id:type_name -> forwarding.ContextId
	8, // 4: forwarding.AttributeQueryRequest.object_id:type_name -> forwarding.ObjectId
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_forwarding_forwarding_attribute_proto_init() }
func file_proto_forwarding_forwarding_attribute_proto_init() {
	if File_proto_forwarding_forwarding_attribute_proto != nil {
		return
	}
	file_proto_forwarding_forwarding_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeDesc); i {
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
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeListRequest); i {
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
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeListReply); i {
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
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeUpdateRequest); i {
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
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeUpdateReply); i {
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
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeQueryRequest); i {
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
		file_proto_forwarding_forwarding_attribute_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeQueryReply); i {
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
			RawDescriptor: file_proto_forwarding_forwarding_attribute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_forwarding_forwarding_attribute_proto_goTypes,
		DependencyIndexes: file_proto_forwarding_forwarding_attribute_proto_depIdxs,
		MessageInfos:      file_proto_forwarding_forwarding_attribute_proto_msgTypes,
	}.Build()
	File_proto_forwarding_forwarding_attribute_proto = out.File
	file_proto_forwarding_forwarding_attribute_proto_rawDesc = nil
	file_proto_forwarding_forwarding_attribute_proto_goTypes = nil
	file_proto_forwarding_forwarding_attribute_proto_depIdxs = nil
}
