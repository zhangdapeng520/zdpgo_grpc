// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: proto/stream_response.proto

package proto

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

// 流式请求：和普通请求没有区别
type StreamResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamResponseRequest) Reset() {
	*x = StreamResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stream_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponseRequest) ProtoMessage() {}

func (x *StreamResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponseRequest.ProtoReflect.Descriptor instead.
func (*StreamResponseRequest) Descriptor() ([]byte, []int) {
	return file_proto_stream_response_proto_rawDescGZIP(), []int{0}
}

func (x *StreamResponseRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 流式响应：和普通响应没有区别
type StreamResponseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamResponseResponse) Reset() {
	*x = StreamResponseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_stream_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponseResponse) ProtoMessage() {}

func (x *StreamResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_stream_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponseResponse.ProtoReflect.Descriptor instead.
func (*StreamResponseResponse) Descriptor() ([]byte, []int) {
	return file_proto_stream_response_proto_rawDescGZIP(), []int{1}
}

func (x *StreamResponseResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_stream_response_proto protoreflect.FileDescriptor

var file_proto_stream_response_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a,
	0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x16, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x59, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x41, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x16, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_stream_response_proto_rawDescOnce sync.Once
	file_proto_stream_response_proto_rawDescData = file_proto_stream_response_proto_rawDesc
)

func file_proto_stream_response_proto_rawDescGZIP() []byte {
	file_proto_stream_response_proto_rawDescOnce.Do(func() {
		file_proto_stream_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_stream_response_proto_rawDescData)
	})
	return file_proto_stream_response_proto_rawDescData
}

var file_proto_stream_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_stream_response_proto_goTypes = []interface{}{
	(*StreamResponseRequest)(nil),  // 0: StreamResponseRequest
	(*StreamResponseResponse)(nil), // 1: StreamResponseResponse
}
var file_proto_stream_response_proto_depIdxs = []int32{
	0, // 0: StreamResponseServer.ServerStream:input_type -> StreamResponseRequest
	1, // 1: StreamResponseServer.ServerStream:output_type -> StreamResponseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_stream_response_proto_init() }
func file_proto_stream_response_proto_init() {
	if File_proto_stream_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_stream_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponseRequest); i {
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
		file_proto_stream_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponseResponse); i {
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
			RawDescriptor: file_proto_stream_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_stream_response_proto_goTypes,
		DependencyIndexes: file_proto_stream_response_proto_depIdxs,
		MessageInfos:      file_proto_stream_response_proto_msgTypes,
	}.Build()
	File_proto_stream_response_proto = out.File
	file_proto_stream_response_proto_rawDesc = nil
	file_proto_stream_response_proto_goTypes = nil
	file_proto_stream_response_proto_depIdxs = nil
}
