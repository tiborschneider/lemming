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

// Interface for consuming packets.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/forwarding/forwarding_packetsink.proto

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

// PacketInjectRequest is a request to inject frames into the port.
type PacketInjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortId       *PortId             `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`               // Port used to process the packet
	ContextId    *ContextId          `protobuf:"bytes,2,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`      // Context containing the port
	Bytes        []byte              `protobuf:"bytes,3,opt,name=bytes,proto3" json:"bytes,omitempty"`                               // Bytes in the L2 frame
	Action       PortAction          `protobuf:"varint,4,opt,name=action,proto3,enum=forwarding.PortAction" json:"action,omitempty"` // Select a set of actions on the port
	Preprocesses []*ActionDesc       `protobuf:"bytes,7,rep,name=preprocesses,proto3" json:"preprocesses,omitempty"`                 // Actions used to preprocess the packet
	Ingress      *PortId             `protobuf:"bytes,8,opt,name=ingress,proto3" json:"ingress,omitempty"`                           // Identifies the ingress of the packet
	Egress       *PortId             `protobuf:"bytes,9,opt,name=egress,proto3" json:"egress,omitempty"`                             // Identifies the egress of the packet
	StartHeader  PacketHeaderId      `protobuf:"varint,10,opt,name=start_header,json=startHeader,proto3,enum=forwarding.PacketHeaderId" json:"start_header,omitempty"`
	ParsedFields []*PacketFieldBytes `protobuf:"bytes,11,rep,name=parsed_fields,json=parsedFields,proto3" json:"parsed_fields,omitempty"` // List of fields describing the L2 frame
	Debug        bool                `protobuf:"varint,12,opt,name=debug,proto3" json:"debug,omitempty"`                                  // Whether to turn on debug messages.
}

func (x *PacketInjectRequest) Reset() {
	*x = PacketInjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_packetsink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketInjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketInjectRequest) ProtoMessage() {}

func (x *PacketInjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_packetsink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketInjectRequest.ProtoReflect.Descriptor instead.
func (*PacketInjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_packetsink_proto_rawDescGZIP(), []int{0}
}

func (x *PacketInjectRequest) GetPortId() *PortId {
	if x != nil {
		return x.PortId
	}
	return nil
}

func (x *PacketInjectRequest) GetContextId() *ContextId {
	if x != nil {
		return x.ContextId
	}
	return nil
}

func (x *PacketInjectRequest) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *PacketInjectRequest) GetAction() PortAction {
	if x != nil {
		return x.Action
	}
	return PortAction_PORT_ACTION_UNSPECIFIED
}

func (x *PacketInjectRequest) GetPreprocesses() []*ActionDesc {
	if x != nil {
		return x.Preprocesses
	}
	return nil
}

func (x *PacketInjectRequest) GetIngress() *PortId {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *PacketInjectRequest) GetEgress() *PortId {
	if x != nil {
		return x.Egress
	}
	return nil
}

func (x *PacketInjectRequest) GetStartHeader() PacketHeaderId {
	if x != nil {
		return x.StartHeader
	}
	return PacketHeaderId_PACKET_HEADER_ID_UNSPECIFIED
}

func (x *PacketInjectRequest) GetParsedFields() []*PacketFieldBytes {
	if x != nil {
		return x.ParsedFields
	}
	return nil
}

func (x *PacketInjectRequest) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type PacketInjectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PacketInjectReply) Reset() {
	*x = PacketInjectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_packetsink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketInjectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketInjectReply) ProtoMessage() {}

func (x *PacketInjectReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_packetsink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketInjectReply.ProtoReflect.Descriptor instead.
func (*PacketInjectReply) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_packetsink_proto_rawDescGZIP(), []int{1}
}

var File_proto_forwarding_forwarding_packetsink_proto protoreflect.FileDescriptor

var file_proto_forwarding_forwarding_packetsink_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec,
	0x03, 0x0a, 0x13, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x52, 0x06, 0x70, 0x6f, 0x72,
	0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x72,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3a, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x52, 0x0c, 0x70,
	0x72, 0x65, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64,
	0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x52, 0x06, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x64, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65,
	0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x22, 0x13, 0x0a,
	0x11, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_forwarding_forwarding_packetsink_proto_rawDescOnce sync.Once
	file_proto_forwarding_forwarding_packetsink_proto_rawDescData = file_proto_forwarding_forwarding_packetsink_proto_rawDesc
)

func file_proto_forwarding_forwarding_packetsink_proto_rawDescGZIP() []byte {
	file_proto_forwarding_forwarding_packetsink_proto_rawDescOnce.Do(func() {
		file_proto_forwarding_forwarding_packetsink_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_forwarding_forwarding_packetsink_proto_rawDescData)
	})
	return file_proto_forwarding_forwarding_packetsink_proto_rawDescData
}

var file_proto_forwarding_forwarding_packetsink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_forwarding_forwarding_packetsink_proto_goTypes = []interface{}{
	(*PacketInjectRequest)(nil), // 0: forwarding.PacketInjectRequest
	(*PacketInjectReply)(nil),   // 1: forwarding.PacketInjectReply
	(*PortId)(nil),              // 2: forwarding.PortId
	(*ContextId)(nil),           // 3: forwarding.ContextId
	(PortAction)(0),             // 4: forwarding.PortAction
	(*ActionDesc)(nil),          // 5: forwarding.ActionDesc
	(PacketHeaderId)(0),         // 6: forwarding.PacketHeaderId
	(*PacketFieldBytes)(nil),    // 7: forwarding.PacketFieldBytes
}
var file_proto_forwarding_forwarding_packetsink_proto_depIdxs = []int32{
	2, // 0: forwarding.PacketInjectRequest.port_id:type_name -> forwarding.PortId
	3, // 1: forwarding.PacketInjectRequest.context_id:type_name -> forwarding.ContextId
	4, // 2: forwarding.PacketInjectRequest.action:type_name -> forwarding.PortAction
	5, // 3: forwarding.PacketInjectRequest.preprocesses:type_name -> forwarding.ActionDesc
	2, // 4: forwarding.PacketInjectRequest.ingress:type_name -> forwarding.PortId
	2, // 5: forwarding.PacketInjectRequest.egress:type_name -> forwarding.PortId
	6, // 6: forwarding.PacketInjectRequest.start_header:type_name -> forwarding.PacketHeaderId
	7, // 7: forwarding.PacketInjectRequest.parsed_fields:type_name -> forwarding.PacketFieldBytes
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_forwarding_forwarding_packetsink_proto_init() }
func file_proto_forwarding_forwarding_packetsink_proto_init() {
	if File_proto_forwarding_forwarding_packetsink_proto != nil {
		return
	}
	file_proto_forwarding_forwarding_action_proto_init()
	file_proto_forwarding_forwarding_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_forwarding_forwarding_packetsink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketInjectRequest); i {
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
		file_proto_forwarding_forwarding_packetsink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketInjectReply); i {
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
			RawDescriptor: file_proto_forwarding_forwarding_packetsink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_forwarding_forwarding_packetsink_proto_goTypes,
		DependencyIndexes: file_proto_forwarding_forwarding_packetsink_proto_depIdxs,
		MessageInfos:      file_proto_forwarding_forwarding_packetsink_proto_msgTypes,
	}.Build()
	File_proto_forwarding_forwarding_packetsink_proto = out.File
	file_proto_forwarding_forwarding_packetsink_proto_rawDesc = nil
	file_proto_forwarding_forwarding_packetsink_proto_goTypes = nil
	file_proto_forwarding_forwarding_packetsink_proto_depIdxs = nil
}
