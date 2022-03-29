// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: proto/streaming.proto

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

type KeywordsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword []string `protobuf:"bytes,1,rep,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *KeywordsData) Reset() {
	*x = KeywordsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordsData) ProtoMessage() {}

func (x *KeywordsData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordsData.ProtoReflect.Descriptor instead.
func (*KeywordsData) Descriptor() ([]byte, []int) {
	return file_proto_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordsData) GetKeyword() []string {
	if x != nil {
		return x.Keyword
	}
	return nil
}

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_streaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_proto_streaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_proto_streaming_proto_rawDescGZIP(), []int{1}
}

var File_proto_streaming_proto protoreflect.FileDescriptor

var file_proto_streaming_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x07, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x32, 0x39, 0x0a, 0x08, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x1a, 0x0d, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_streaming_proto_rawDescOnce sync.Once
	file_proto_streaming_proto_rawDescData = file_proto_streaming_proto_rawDesc
)

func file_proto_streaming_proto_rawDescGZIP() []byte {
	file_proto_streaming_proto_rawDescOnce.Do(func() {
		file_proto_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_streaming_proto_rawDescData)
	})
	return file_proto_streaming_proto_rawDescData
}

var file_proto_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_streaming_proto_goTypes = []interface{}{
	(*KeywordsData)(nil), // 0: KeywordsData
	(*Input)(nil),        // 1: Input
}
var file_proto_streaming_proto_depIdxs = []int32{
	1, // 0: Keywords.StreamingKeyword:input_type -> Input
	0, // 1: Keywords.StreamingKeyword:output_type -> KeywordsData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_streaming_proto_init() }
func file_proto_streaming_proto_init() {
	if File_proto_streaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_streaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordsData); i {
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
		file_proto_streaming_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
			RawDescriptor: file_proto_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_streaming_proto_goTypes,
		DependencyIndexes: file_proto_streaming_proto_depIdxs,
		MessageInfos:      file_proto_streaming_proto_msgTypes,
	}.Build()
	File_proto_streaming_proto = out.File
	file_proto_streaming_proto_rawDesc = nil
	file_proto_streaming_proto_goTypes = nil
	file_proto_streaming_proto_depIdxs = nil
}
