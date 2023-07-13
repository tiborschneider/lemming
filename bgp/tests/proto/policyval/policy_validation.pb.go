// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: bgp/tests/proto/policyval/policy_validation.proto

package policyval

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

type RouteTestResult int32

const (
	RouteTestResult_ROUTE_TEST_RESULT_UNSPECIFIED RouteTestResult = 0
	RouteTestResult_ROUTE_TEST_RESULT_ACCEPT      RouteTestResult = 1
	RouteTestResult_ROUTE_TEST_RESULT_DISCARD     RouteTestResult = 2
)

// Enum value maps for RouteTestResult.
var (
	RouteTestResult_name = map[int32]string{
		0: "ROUTE_TEST_RESULT_UNSPECIFIED",
		1: "ROUTE_TEST_RESULT_ACCEPT",
		2: "ROUTE_TEST_RESULT_DISCARD",
	}
	RouteTestResult_value = map[string]int32{
		"ROUTE_TEST_RESULT_UNSPECIFIED": 0,
		"ROUTE_TEST_RESULT_ACCEPT":      1,
		"ROUTE_TEST_RESULT_DISCARD":     2,
	}
)

func (x RouteTestResult) Enum() *RouteTestResult {
	p := new(RouteTestResult)
	*p = x
	return p
}

func (x RouteTestResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RouteTestResult) Descriptor() protoreflect.EnumDescriptor {
	return file_bgp_tests_proto_policyval_policy_validation_proto_enumTypes[0].Descriptor()
}

func (RouteTestResult) Type() protoreflect.EnumType {
	return &file_bgp_tests_proto_policyval_policy_validation_proto_enumTypes[0]
}

func (x RouteTestResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RouteTestResult.Descriptor instead.
func (RouteTestResult) EnumDescriptor() ([]byte, []int) {
	return file_bgp_tests_proto_policyval_policy_validation_proto_rawDescGZIP(), []int{0}
}

type PolicyTestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string           `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	RouteTests  []*RouteTestCase `protobuf:"bytes,4,rep,name=route_tests,json=routeTests,proto3" json:"route_tests,omitempty"`
}

func (x *PolicyTestCase) Reset() {
	*x = PolicyTestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyTestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyTestCase) ProtoMessage() {}

func (x *PolicyTestCase) ProtoReflect() protoreflect.Message {
	mi := &file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyTestCase.ProtoReflect.Descriptor instead.
func (*PolicyTestCase) Descriptor() ([]byte, []int) {
	return file_bgp_tests_proto_policyval_policy_validation_proto_rawDescGZIP(), []int{0}
}

func (x *PolicyTestCase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PolicyTestCase) GetRouteTests() []*RouteTestCase {
	if x != nil {
		return x.RouteTests
	}
	return nil
}

type TestRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReachPrefix string `protobuf:"bytes,1,opt,name=reach_prefix,json=reachPrefix,proto3" json:"reach_prefix,omitempty"`
}

func (x *TestRoute) Reset() {
	*x = TestRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRoute) ProtoMessage() {}

func (x *TestRoute) ProtoReflect() protoreflect.Message {
	mi := &file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRoute.ProtoReflect.Descriptor instead.
func (*TestRoute) Descriptor() ([]byte, []int) {
	return file_bgp_tests_proto_policyval_policy_validation_proto_rawDescGZIP(), []int{1}
}

func (x *TestRoute) GetReachPrefix() string {
	if x != nil {
		return x.ReachPrefix
	}
	return ""
}

type RouteTestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description    string          `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Input          *TestRoute      `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	ExpectedResult RouteTestResult `protobuf:"varint,3,opt,name=expected_result,json=expectedResult,proto3,enum=policy_validation.RouteTestResult" json:"expected_result,omitempty"`
}

func (x *RouteTestCase) Reset() {
	*x = RouteTestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteTestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteTestCase) ProtoMessage() {}

func (x *RouteTestCase) ProtoReflect() protoreflect.Message {
	mi := &file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteTestCase.ProtoReflect.Descriptor instead.
func (*RouteTestCase) Descriptor() ([]byte, []int) {
	return file_bgp_tests_proto_policyval_policy_validation_proto_rawDescGZIP(), []int{2}
}

func (x *RouteTestCase) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RouteTestCase) GetInput() *TestRoute {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *RouteTestCase) GetExpectedResult() RouteTestResult {
	if x != nil {
		return x.ExpectedResult
	}
	return RouteTestResult_ROUTE_TEST_RESULT_UNSPECIFIED
}

var File_bgp_tests_proto_policyval_policy_validation_proto protoreflect.FileDescriptor

var file_bgp_tests_proto_policyval_policy_validation_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x67, 0x70, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x76, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x0e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x73, 0x22, 0x2e, 0x0a,
	0x09, 0x54, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x61, 0x63, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x61, 0x63, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0xb2, 0x01,
	0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4b, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2a, 0x71, 0x0a, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f, 0x54,
	0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x4f, 0x55, 0x54,
	0x45, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x41, 0x43,
	0x43, 0x45, 0x50, 0x54, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x43,
	0x41, 0x52, 0x44, 0x10, 0x02, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x65, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x67, 0x70, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x76, 0x61, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bgp_tests_proto_policyval_policy_validation_proto_rawDescOnce sync.Once
	file_bgp_tests_proto_policyval_policy_validation_proto_rawDescData = file_bgp_tests_proto_policyval_policy_validation_proto_rawDesc
)

func file_bgp_tests_proto_policyval_policy_validation_proto_rawDescGZIP() []byte {
	file_bgp_tests_proto_policyval_policy_validation_proto_rawDescOnce.Do(func() {
		file_bgp_tests_proto_policyval_policy_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgp_tests_proto_policyval_policy_validation_proto_rawDescData)
	})
	return file_bgp_tests_proto_policyval_policy_validation_proto_rawDescData
}

var file_bgp_tests_proto_policyval_policy_validation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bgp_tests_proto_policyval_policy_validation_proto_goTypes = []interface{}{
	(RouteTestResult)(0),   // 0: policy_validation.RouteTestResult
	(*PolicyTestCase)(nil), // 1: policy_validation.PolicyTestCase
	(*TestRoute)(nil),      // 2: policy_validation.TestRoute
	(*RouteTestCase)(nil),  // 3: policy_validation.RouteTestCase
}
var file_bgp_tests_proto_policyval_policy_validation_proto_depIdxs = []int32{
	3, // 0: policy_validation.PolicyTestCase.route_tests:type_name -> policy_validation.RouteTestCase
	2, // 1: policy_validation.RouteTestCase.input:type_name -> policy_validation.TestRoute
	0, // 2: policy_validation.RouteTestCase.expected_result:type_name -> policy_validation.RouteTestResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bgp_tests_proto_policyval_policy_validation_proto_init() }
func file_bgp_tests_proto_policyval_policy_validation_proto_init() {
	if File_bgp_tests_proto_policyval_policy_validation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyTestCase); i {
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
		file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRoute); i {
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
		file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteTestCase); i {
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
			RawDescriptor: file_bgp_tests_proto_policyval_policy_validation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bgp_tests_proto_policyval_policy_validation_proto_goTypes,
		DependencyIndexes: file_bgp_tests_proto_policyval_policy_validation_proto_depIdxs,
		EnumInfos:         file_bgp_tests_proto_policyval_policy_validation_proto_enumTypes,
		MessageInfos:      file_bgp_tests_proto_policyval_policy_validation_proto_msgTypes,
	}.Build()
	File_bgp_tests_proto_policyval_policy_validation_proto = out.File
	file_bgp_tests_proto_policyval_policy_validation_proto_rawDesc = nil
	file_bgp_tests_proto_policyval_policy_validation_proto_goTypes = nil
	file_bgp_tests_proto_policyval_policy_validation_proto_depIdxs = nil
}
