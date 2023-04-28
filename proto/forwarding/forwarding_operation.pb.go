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

// Protocol buffers to wrap all streaming compatible programmings
// so that they can be serialized within a single stream to provide
// the necessary FIFO guarantees.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/forwarding/forwarding_operation.proto

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

// Union of all possible streaming message requests.
type OperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*OperationRequest_TableEntryAdd
	//	*OperationRequest_TableEntryRemove
	Request isOperationRequest_Request `protobuf_oneof:"request"`
}

func (x *OperationRequest) Reset() {
	*x = OperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationRequest) ProtoMessage() {}

func (x *OperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationRequest.ProtoReflect.Descriptor instead.
func (*OperationRequest) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_operation_proto_rawDescGZIP(), []int{0}
}

func (m *OperationRequest) GetRequest() isOperationRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *OperationRequest) GetTableEntryAdd() *TableEntryAddRequest {
	if x, ok := x.GetRequest().(*OperationRequest_TableEntryAdd); ok {
		return x.TableEntryAdd
	}
	return nil
}

func (x *OperationRequest) GetTableEntryRemove() *TableEntryRemoveRequest {
	if x, ok := x.GetRequest().(*OperationRequest_TableEntryRemove); ok {
		return x.TableEntryRemove
	}
	return nil
}

type isOperationRequest_Request interface {
	isOperationRequest_Request()
}

type OperationRequest_TableEntryAdd struct {
	TableEntryAdd *TableEntryAddRequest `protobuf:"bytes,1,opt,name=table_entry_add,json=tableEntryAdd,proto3,oneof"`
}

type OperationRequest_TableEntryRemove struct {
	TableEntryRemove *TableEntryRemoveRequest `protobuf:"bytes,2,opt,name=table_entry_remove,json=tableEntryRemove,proto3,oneof"`
}

func (*OperationRequest_TableEntryAdd) isOperationRequest_Request() {}

func (*OperationRequest_TableEntryRemove) isOperationRequest_Request() {}

// Union of all possible streaming message replies.
type OperationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reply:
	//
	//	*OperationReply_TableEntryAdd
	//	*OperationReply_TableEntryRemove
	Reply isOperationReply_Reply `protobuf_oneof:"reply"`
}

func (x *OperationReply) Reset() {
	*x = OperationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationReply) ProtoMessage() {}

func (x *OperationReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationReply.ProtoReflect.Descriptor instead.
func (*OperationReply) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_operation_proto_rawDescGZIP(), []int{1}
}

func (m *OperationReply) GetReply() isOperationReply_Reply {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (x *OperationReply) GetTableEntryAdd() *TableEntryAddReply {
	if x, ok := x.GetReply().(*OperationReply_TableEntryAdd); ok {
		return x.TableEntryAdd
	}
	return nil
}

func (x *OperationReply) GetTableEntryRemove() *TableEntryRemoveReply {
	if x, ok := x.GetReply().(*OperationReply_TableEntryRemove); ok {
		return x.TableEntryRemove
	}
	return nil
}

type isOperationReply_Reply interface {
	isOperationReply_Reply()
}

type OperationReply_TableEntryAdd struct {
	TableEntryAdd *TableEntryAddReply `protobuf:"bytes,1,opt,name=table_entry_add,json=tableEntryAdd,proto3,oneof"`
}

type OperationReply_TableEntryRemove struct {
	TableEntryRemove *TableEntryRemoveReply `protobuf:"bytes,2,opt,name=table_entry_remove,json=tableEntryRemove,proto3,oneof"`
}

func (*OperationReply_TableEntryAdd) isOperationReply_Reply() {}

func (*OperationReply_TableEntryRemove) isOperationReply_Reply() {}

var File_proto_forwarding_forwarding_operation_proto protoreflect.FileDescriptor

var file_proto_forwarding_forwarding_operation_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x10, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x41, 0x64, 0x64, 0x12, 0x53, 0x0a, 0x12, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x10, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x48, 0x0a, 0x0f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48,
	0x00, 0x52, 0x0d, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x64, 0x64,
	0x12, 0x51, 0x0a, 0x12, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48,
	0x00, 0x52, 0x10, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_forwarding_forwarding_operation_proto_rawDescOnce sync.Once
	file_proto_forwarding_forwarding_operation_proto_rawDescData = file_proto_forwarding_forwarding_operation_proto_rawDesc
)

func file_proto_forwarding_forwarding_operation_proto_rawDescGZIP() []byte {
	file_proto_forwarding_forwarding_operation_proto_rawDescOnce.Do(func() {
		file_proto_forwarding_forwarding_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_forwarding_forwarding_operation_proto_rawDescData)
	})
	return file_proto_forwarding_forwarding_operation_proto_rawDescData
}

var file_proto_forwarding_forwarding_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_forwarding_forwarding_operation_proto_goTypes = []interface{}{
	(*OperationRequest)(nil),        // 0: forwarding.OperationRequest
	(*OperationReply)(nil),          // 1: forwarding.OperationReply
	(*TableEntryAddRequest)(nil),    // 2: forwarding.TableEntryAddRequest
	(*TableEntryRemoveRequest)(nil), // 3: forwarding.TableEntryRemoveRequest
	(*TableEntryAddReply)(nil),      // 4: forwarding.TableEntryAddReply
	(*TableEntryRemoveReply)(nil),   // 5: forwarding.TableEntryRemoveReply
}
var file_proto_forwarding_forwarding_operation_proto_depIdxs = []int32{
	2, // 0: forwarding.OperationRequest.table_entry_add:type_name -> forwarding.TableEntryAddRequest
	3, // 1: forwarding.OperationRequest.table_entry_remove:type_name -> forwarding.TableEntryRemoveRequest
	4, // 2: forwarding.OperationReply.table_entry_add:type_name -> forwarding.TableEntryAddReply
	5, // 3: forwarding.OperationReply.table_entry_remove:type_name -> forwarding.TableEntryRemoveReply
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_forwarding_forwarding_operation_proto_init() }
func file_proto_forwarding_forwarding_operation_proto_init() {
	if File_proto_forwarding_forwarding_operation_proto != nil {
		return
	}
	file_proto_forwarding_forwarding_table_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_forwarding_forwarding_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationRequest); i {
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
		file_proto_forwarding_forwarding_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationReply); i {
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
	file_proto_forwarding_forwarding_operation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OperationRequest_TableEntryAdd)(nil),
		(*OperationRequest_TableEntryRemove)(nil),
	}
	file_proto_forwarding_forwarding_operation_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*OperationReply_TableEntryAdd)(nil),
		(*OperationReply_TableEntryRemove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_forwarding_forwarding_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_forwarding_forwarding_operation_proto_goTypes,
		DependencyIndexes: file_proto_forwarding_forwarding_operation_proto_depIdxs,
		MessageInfos:      file_proto_forwarding_forwarding_operation_proto_msgTypes,
	}.Build()
	File_proto_forwarding_forwarding_operation_proto = out.File
	file_proto_forwarding_forwarding_operation_proto_rawDesc = nil
	file_proto_forwarding_forwarding_operation_proto_goTypes = nil
	file_proto_forwarding_forwarding_operation_proto_depIdxs = nil
}
