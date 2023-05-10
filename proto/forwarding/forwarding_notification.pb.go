// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proto/forwarding/forwarding_notification.proto

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

type Event int32

const (
	Event_EVENT_UNSPECIFIED Event = 0
	Event_EVENT_PORT        Event = 1
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_PORT",
	}
	Event_value = map[string]int32{
		"EVENT_UNSPECIFIED": 0,
		"EVENT_PORT":        1,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_forwarding_forwarding_notification_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_proto_forwarding_forwarding_notification_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_notification_proto_rawDescGZIP(), []int{0}
}

type EventDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event          Event  `protobuf:"varint,1,opt,name=event,proto3,enum=forwarding.Event" json:"event,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	// Types that are assignable to Desc:
	//
	//	*EventDesc_Port
	Desc isEventDesc_Desc `protobuf_oneof:"desc"`
}

func (x *EventDesc) Reset() {
	*x = EventDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDesc) ProtoMessage() {}

func (x *EventDesc) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDesc.ProtoReflect.Descriptor instead.
func (*EventDesc) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_notification_proto_rawDescGZIP(), []int{0}
}

func (x *EventDesc) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_EVENT_UNSPECIFIED
}

func (x *EventDesc) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (m *EventDesc) GetDesc() isEventDesc_Desc {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (x *EventDesc) GetPort() *PortEventDesc {
	if x, ok := x.GetDesc().(*EventDesc_Port); ok {
		return x.Port
	}
	return nil
}

type isEventDesc_Desc interface {
	isEventDesc_Desc()
}

type EventDesc_Port struct {
	Port *PortEventDesc `protobuf:"bytes,3,opt,name=port,proto3,oneof"`
}

func (*EventDesc_Port) isEventDesc_Desc() {}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_notification_proto_rawDescGZIP(), []int{1}
}

type PortEventDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context   *ContextId     `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	PortId    *PortId        `protobuf:"bytes,5,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	LinkEvent *LinkStateDesc `protobuf:"bytes,6,opt,name=link_event,json=linkEvent,proto3" json:"link_event,omitempty"`
}

func (x *PortEventDesc) Reset() {
	*x = PortEventDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_forwarding_forwarding_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortEventDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortEventDesc) ProtoMessage() {}

func (x *PortEventDesc) ProtoReflect() protoreflect.Message {
	mi := &file_proto_forwarding_forwarding_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortEventDesc.ProtoReflect.Descriptor instead.
func (*PortEventDesc) Descriptor() ([]byte, []int) {
	return file_proto_forwarding_forwarding_notification_proto_rawDescGZIP(), []int{2}
}

func (x *PortEventDesc) GetContext() *ContextId {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *PortEventDesc) GetPortId() *PortId {
	if x != nil {
		return x.PortId
	}
	return nil
}

func (x *PortEventDesc) GetLinkEvent() *LinkStateDesc {
	if x != nil {
		return x.LinkEvent
	}
	return nil
}

var File_proto_forwarding_forwarding_notification_proto protoreflect.FileDescriptor

var file_proto_forwarding_forwarding_notification_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x28, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96,
	0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x27, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2f,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x52, 0x06, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63, 0x52,
	0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x2a, 0x2e, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_forwarding_forwarding_notification_proto_rawDescOnce sync.Once
	file_proto_forwarding_forwarding_notification_proto_rawDescData = file_proto_forwarding_forwarding_notification_proto_rawDesc
)

func file_proto_forwarding_forwarding_notification_proto_rawDescGZIP() []byte {
	file_proto_forwarding_forwarding_notification_proto_rawDescOnce.Do(func() {
		file_proto_forwarding_forwarding_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_forwarding_forwarding_notification_proto_rawDescData)
	})
	return file_proto_forwarding_forwarding_notification_proto_rawDescData
}

var file_proto_forwarding_forwarding_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_forwarding_forwarding_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_forwarding_forwarding_notification_proto_goTypes = []interface{}{
	(Event)(0),            // 0: forwarding.Event
	(*EventDesc)(nil),     // 1: forwarding.EventDesc
	(*Empty)(nil),         // 2: forwarding.Empty
	(*PortEventDesc)(nil), // 3: forwarding.PortEventDesc
	(*ContextId)(nil),     // 4: forwarding.ContextId
	(*PortId)(nil),        // 5: forwarding.PortId
	(*LinkStateDesc)(nil), // 6: forwarding.LinkStateDesc
}
var file_proto_forwarding_forwarding_notification_proto_depIdxs = []int32{
	0, // 0: forwarding.EventDesc.event:type_name -> forwarding.Event
	3, // 1: forwarding.EventDesc.port:type_name -> forwarding.PortEventDesc
	4, // 2: forwarding.PortEventDesc.context:type_name -> forwarding.ContextId
	5, // 3: forwarding.PortEventDesc.port_id:type_name -> forwarding.PortId
	6, // 4: forwarding.PortEventDesc.link_event:type_name -> forwarding.LinkStateDesc
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_forwarding_forwarding_notification_proto_init() }
func file_proto_forwarding_forwarding_notification_proto_init() {
	if File_proto_forwarding_forwarding_notification_proto != nil {
		return
	}
	file_proto_forwarding_forwarding_common_proto_init()
	file_proto_forwarding_forwarding_port_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_forwarding_forwarding_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventDesc); i {
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
		file_proto_forwarding_forwarding_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_forwarding_forwarding_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortEventDesc); i {
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
	file_proto_forwarding_forwarding_notification_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EventDesc_Port)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_forwarding_forwarding_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_forwarding_forwarding_notification_proto_goTypes,
		DependencyIndexes: file_proto_forwarding_forwarding_notification_proto_depIdxs,
		EnumInfos:         file_proto_forwarding_forwarding_notification_proto_enumTypes,
		MessageInfos:      file_proto_forwarding_forwarding_notification_proto_msgTypes,
	}.Build()
	File_proto_forwarding_forwarding_notification_proto = out.File
	file_proto_forwarding_forwarding_notification_proto_rawDesc = nil
	file_proto_forwarding_forwarding_notification_proto_goTypes = nil
	file_proto_forwarding_forwarding_notification_proto_depIdxs = nil
}
