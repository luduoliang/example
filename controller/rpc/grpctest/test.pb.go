package grpctest

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestFirst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestFirst) Reset() {
	*x = RequestFirst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFirst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFirst) ProtoMessage() {}

func (x *RequestFirst) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFirst.ProtoReflect.Descriptor instead.
func (*RequestFirst) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *RequestFirst) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ResponseFirst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseFirst) Reset() {
	*x = ResponseFirst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseFirst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFirst) ProtoMessage() {}

func (x *ResponseFirst) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFirst.ProtoReflect.Descriptor instead.
func (*ResponseFirst) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseFirst) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseFirst) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResponseFirst) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x51, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x7d, 0x0a, 0x09, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*RequestFirst)(nil),  // 0: test.RequestFirst
	(*ResponseFirst)(nil), // 1: test.ResponseFirst
}
var file_test_proto_depIdxs = []int32{
	0, // 0: test.TestFirst.GetData:input_type -> test.RequestFirst
	0, // 1: test.TestFirst.Communite:input_type -> test.RequestFirst
	1, // 2: test.TestFirst.GetData:output_type -> test.ResponseFirst
	1, // 3: test.TestFirst.Communite:output_type -> test.ResponseFirst
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFirst); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseFirst); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestFirstClient is the client API for TestFirst service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestFirstClient interface {
	GetData(ctx context.Context, in *RequestFirst, opts ...grpc.CallOption) (*ResponseFirst, error)
	Communite(ctx context.Context, opts ...grpc.CallOption) (TestFirst_CommuniteClient, error)
}

type testFirstClient struct {
	cc grpc.ClientConnInterface
}

func NewTestFirstClient(cc grpc.ClientConnInterface) TestFirstClient {
	return &testFirstClient{cc}
}

func (c *testFirstClient) GetData(ctx context.Context, in *RequestFirst, opts ...grpc.CallOption) (*ResponseFirst, error) {
	out := new(ResponseFirst)
	err := c.cc.Invoke(ctx, "/test.TestFirst/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testFirstClient) Communite(ctx context.Context, opts ...grpc.CallOption) (TestFirst_CommuniteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestFirst_serviceDesc.Streams[0], "/test.TestFirst/Communite", opts...)
	if err != nil {
		return nil, err
	}
	x := &testFirstCommuniteClient{stream}
	return x, nil
}

type TestFirst_CommuniteClient interface {
	Send(*RequestFirst) error
	Recv() (*ResponseFirst, error)
	grpc.ClientStream
}

type testFirstCommuniteClient struct {
	grpc.ClientStream
}

func (x *testFirstCommuniteClient) Send(m *RequestFirst) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testFirstCommuniteClient) Recv() (*ResponseFirst, error) {
	m := new(ResponseFirst)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestFirstServer is the server API for TestFirst service.
type TestFirstServer interface {
	GetData(context.Context, *RequestFirst) (*ResponseFirst, error)
	Communite(TestFirst_CommuniteServer) error
}

// UnimplementedTestFirstServer can be embedded to have forward compatible implementations.
type UnimplementedTestFirstServer struct {
}

func (*UnimplementedTestFirstServer) GetData(context.Context, *RequestFirst) (*ResponseFirst, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (*UnimplementedTestFirstServer) Communite(TestFirst_CommuniteServer) error {
	return status.Errorf(codes.Unimplemented, "method Communite not implemented")
}

func RegisterTestFirstServer(s *grpc.Server, srv TestFirstServer) {
	s.RegisterService(&_TestFirst_serviceDesc, srv)
}

func _TestFirst_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFirst)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestFirstServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.TestFirst/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestFirstServer).GetData(ctx, req.(*RequestFirst))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestFirst_Communite_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestFirstServer).Communite(&testFirstCommuniteServer{stream})
}

type TestFirst_CommuniteServer interface {
	Send(*ResponseFirst) error
	Recv() (*RequestFirst, error)
	grpc.ServerStream
}

type testFirstCommuniteServer struct {
	grpc.ServerStream
}

func (x *testFirstCommuniteServer) Send(m *ResponseFirst) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testFirstCommuniteServer) Recv() (*RequestFirst, error) {
	m := new(RequestFirst)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestFirst_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.TestFirst",
	HandlerType: (*TestFirstServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _TestFirst_GetData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Communite",
			Handler:       _TestFirst_Communite_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}
