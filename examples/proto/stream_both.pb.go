// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: stream_both.proto

package proto

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

// 流式请求：和普通请求没有区别
type StreamBothRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamBothRequest) Reset() {
	*x = StreamBothRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_both_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamBothRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBothRequest) ProtoMessage() {}

func (x *StreamBothRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stream_both_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBothRequest.ProtoReflect.Descriptor instead.
func (*StreamBothRequest) Descriptor() ([]byte, []int) {
	return file_stream_both_proto_rawDescGZIP(), []int{0}
}

func (x *StreamBothRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 流式响应：和普通响应没有区别
type StreamBothResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamBothResponse) Reset() {
	*x = StreamBothResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stream_both_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamBothResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBothResponse) ProtoMessage() {}

func (x *StreamBothResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stream_both_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBothResponse.ProtoReflect.Descriptor instead.
func (*StreamBothResponse) Descriptor() ([]byte, []int) {
	return file_stream_both_proto_rawDescGZIP(), []int{1}
}

func (x *StreamBothResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_stream_both_proto protoreflect.FileDescriptor

var file_stream_both_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x62, 0x6f, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x6f, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x12,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x6f, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4d, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x42, 0x6f, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x42, 0x6f,
	0x74, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x42, 0x6f, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x6f, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stream_both_proto_rawDescOnce sync.Once
	file_stream_both_proto_rawDescData = file_stream_both_proto_rawDesc
)

func file_stream_both_proto_rawDescGZIP() []byte {
	file_stream_both_proto_rawDescOnce.Do(func() {
		file_stream_both_proto_rawDescData = protoimpl.X.CompressGZIP(file_stream_both_proto_rawDescData)
	})
	return file_stream_both_proto_rawDescData
}

var file_stream_both_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_stream_both_proto_goTypes = []interface{}{
	(*StreamBothRequest)(nil),  // 0: StreamBothRequest
	(*StreamBothResponse)(nil), // 1: StreamBothResponse
}
var file_stream_both_proto_depIdxs = []int32{
	0, // 0: StreamBothServer.BothStream:input_type -> StreamBothRequest
	1, // 1: StreamBothServer.BothStream:output_type -> StreamBothResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stream_both_proto_init() }
func file_stream_both_proto_init() {
	if File_stream_both_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stream_both_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamBothRequest); i {
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
		file_stream_both_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamBothResponse); i {
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
			RawDescriptor: file_stream_both_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stream_both_proto_goTypes,
		DependencyIndexes: file_stream_both_proto_depIdxs,
		MessageInfos:      file_stream_both_proto_msgTypes,
	}.Build()
	File_stream_both_proto = out.File
	file_stream_both_proto_rawDesc = nil
	file_stream_both_proto_goTypes = nil
	file_stream_both_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StreamBothServerClient is the client API for StreamBothServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamBothServerClient interface {
	// 双向端流模式
	BothStream(ctx context.Context, opts ...grpc.CallOption) (StreamBothServer_BothStreamClient, error)
}

type streamBothServerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamBothServerClient(cc grpc.ClientConnInterface) StreamBothServerClient {
	return &streamBothServerClient{cc}
}

func (c *streamBothServerClient) BothStream(ctx context.Context, opts ...grpc.CallOption) (StreamBothServer_BothStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamBothServer_serviceDesc.Streams[0], "/StreamBothServer/BothStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamBothServerBothStreamClient{stream}
	return x, nil
}

type StreamBothServer_BothStreamClient interface {
	Send(*StreamBothRequest) error
	Recv() (*StreamBothResponse, error)
	grpc.ClientStream
}

type streamBothServerBothStreamClient struct {
	grpc.ClientStream
}

func (x *streamBothServerBothStreamClient) Send(m *StreamBothRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamBothServerBothStreamClient) Recv() (*StreamBothResponse, error) {
	m := new(StreamBothResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamBothServerServer is the server API for StreamBothServer service.
type StreamBothServerServer interface {
	// 双向端流模式
	BothStream(StreamBothServer_BothStreamServer) error
}

// UnimplementedStreamBothServerServer can be embedded to have forward compatible implementations.
type UnimplementedStreamBothServerServer struct {
}

func (*UnimplementedStreamBothServerServer) BothStream(StreamBothServer_BothStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BothStream not implemented")
}

func RegisterStreamBothServerServer(s *grpc.Server, srv StreamBothServerServer) {
	s.RegisterService(&_StreamBothServer_serviceDesc, srv)
}

func _StreamBothServer_BothStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamBothServerServer).BothStream(&streamBothServerBothStreamServer{stream})
}

type StreamBothServer_BothStreamServer interface {
	Send(*StreamBothResponse) error
	Recv() (*StreamBothRequest, error)
	grpc.ServerStream
}

type streamBothServerBothStreamServer struct {
	grpc.ServerStream
}

func (x *streamBothServerBothStreamServer) Send(m *StreamBothResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamBothServerBothStreamServer) Recv() (*StreamBothRequest, error) {
	m := new(StreamBothRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamBothServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StreamBothServer",
	HandlerType: (*StreamBothServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BothStream",
			Handler:       _StreamBothServer_BothStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream_both.proto",
}