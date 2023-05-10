// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/dataplane/dataplane.proto

package dataplane

import (
	forwarding "github.com/openconfig/lemming/proto/forwarding"
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

type NextHop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port               string                   `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Ip                 string                   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Weight             uint64                   `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	PreTransmitActions []*forwarding.ActionDesc `protobuf:"bytes,4,rep,name=pre_transmit_actions,json=preTransmitActions,proto3" json:"pre_transmit_actions,omitempty"`
}

func (x *NextHop) Reset() {
	*x = NextHop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dataplane_dataplane_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextHop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextHop) ProtoMessage() {}

func (x *NextHop) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataplane_dataplane_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextHop.ProtoReflect.Descriptor instead.
func (*NextHop) Descriptor() ([]byte, []int) {
	return file_proto_dataplane_dataplane_proto_rawDescGZIP(), []int{0}
}

func (x *NextHop) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *NextHop) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NextHop) GetWeight() uint64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *NextHop) GetPreTransmitActions() []*forwarding.ActionDesc {
	if x != nil {
		return x.PreTransmitActions
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vrf      uint64     `protobuf:"varint,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	Prefix   string     `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	NextHops []*NextHop `protobuf:"bytes,3,rep,name=next_hops,json=nextHops,proto3" json:"next_hops,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dataplane_dataplane_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dataplane_dataplane_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_proto_dataplane_dataplane_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetVrf() uint64 {
	if x != nil {
		return x.Vrf
	}
	return 0
}

func (x *Route) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Route) GetNextHops() []*NextHop {
	if x != nil {
		return x.NextHops
	}
	return nil
}

var File_proto_dataplane_dataplane_proto protoreflect.FileDescriptor

var file_proto_dataplane_dataplane_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f,
	0x01, 0x0a, 0x07, 0x4e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x48, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x52, 0x12, 0x70, 0x72,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x6a, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x72, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x76, 0x72, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x37, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x48,
	0x6f, 0x70, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x48, 0x6f, 0x70, 0x73, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dataplane_dataplane_proto_rawDescOnce sync.Once
	file_proto_dataplane_dataplane_proto_rawDescData = file_proto_dataplane_dataplane_proto_rawDesc
)

func file_proto_dataplane_dataplane_proto_rawDescGZIP() []byte {
	file_proto_dataplane_dataplane_proto_rawDescOnce.Do(func() {
		file_proto_dataplane_dataplane_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dataplane_dataplane_proto_rawDescData)
	})
	return file_proto_dataplane_dataplane_proto_rawDescData
}

var file_proto_dataplane_dataplane_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_dataplane_dataplane_proto_goTypes = []interface{}{
	(*NextHop)(nil),               // 0: lemming.dataplane.NextHop
	(*Route)(nil),                 // 1: lemming.dataplane.Route
	(*forwarding.ActionDesc)(nil), // 2: forwarding.ActionDesc
}
var file_proto_dataplane_dataplane_proto_depIdxs = []int32{
	2, // 0: lemming.dataplane.NextHop.pre_transmit_actions:type_name -> forwarding.ActionDesc
	0, // 1: lemming.dataplane.Route.next_hops:type_name -> lemming.dataplane.NextHop
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_dataplane_dataplane_proto_init() }
func file_proto_dataplane_dataplane_proto_init() {
	if File_proto_dataplane_dataplane_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dataplane_dataplane_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextHop); i {
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
		file_proto_dataplane_dataplane_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
			RawDescriptor: file_proto_dataplane_dataplane_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_dataplane_dataplane_proto_goTypes,
		DependencyIndexes: file_proto_dataplane_dataplane_proto_depIdxs,
		MessageInfos:      file_proto_dataplane_dataplane_proto_msgTypes,
	}.Build()
	File_proto_dataplane_dataplane_proto = out.File
	file_proto_dataplane_dataplane_proto_rawDesc = nil
	file_proto_dataplane_dataplane_proto_goTypes = nil
	file_proto_dataplane_dataplane_proto_depIdxs = nil
}
