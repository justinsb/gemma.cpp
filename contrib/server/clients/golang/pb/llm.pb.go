// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: llm.proto

package pb

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

type ConverseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ConverseRequest) Reset() {
	*x = ConverseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConverseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConverseRequest) ProtoMessage() {}

func (x *ConverseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_llm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConverseRequest.ProtoReflect.Descriptor instead.
func (*ConverseRequest) Descriptor() ([]byte, []int) {
	return file_llm_proto_rawDescGZIP(), []int{0}
}

func (x *ConverseRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ConverseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text          []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	EndOfResponse bool     `protobuf:"varint,2,opt,name=end_of_response,json=endOfResponse,proto3" json:"end_of_response,omitempty"`
}

func (x *ConverseResponse) Reset() {
	*x = ConverseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_llm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConverseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConverseResponse) ProtoMessage() {}

func (x *ConverseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_llm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConverseResponse.ProtoReflect.Descriptor instead.
func (*ConverseResponse) Descriptor() ([]byte, []int) {
	return file_llm_proto_rawDescGZIP(), []int{1}
}

func (x *ConverseResponse) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *ConverseResponse) GetEndOfResponse() bool {
	if x != nil {
		return x.EndOfResponse
	}
	return false
}

var File_llm_proto protoreflect.FileDescriptor

var file_llm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6c, 0x6c, 0x6d,
	0x22, 0x25, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4e, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x44, 0x0a, 0x03, 0x4c, 0x4c, 0x4d, 0x12, 0x3d,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x6c, 0x6c, 0x6d,
	0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x67, 0x65, 0x6d, 0x6d, 0x61, 0x2e, 0x63, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_llm_proto_rawDescOnce sync.Once
	file_llm_proto_rawDescData = file_llm_proto_rawDesc
)

func file_llm_proto_rawDescGZIP() []byte {
	file_llm_proto_rawDescOnce.Do(func() {
		file_llm_proto_rawDescData = protoimpl.X.CompressGZIP(file_llm_proto_rawDescData)
	})
	return file_llm_proto_rawDescData
}

var file_llm_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_llm_proto_goTypes = []interface{}{
	(*ConverseRequest)(nil),  // 0: llm.ConverseRequest
	(*ConverseResponse)(nil), // 1: llm.ConverseResponse
}
var file_llm_proto_depIdxs = []int32{
	0, // 0: llm.LLM.Converse:input_type -> llm.ConverseRequest
	1, // 1: llm.LLM.Converse:output_type -> llm.ConverseResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_llm_proto_init() }
func file_llm_proto_init() {
	if File_llm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_llm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConverseRequest); i {
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
		file_llm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConverseResponse); i {
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
			RawDescriptor: file_llm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_llm_proto_goTypes,
		DependencyIndexes: file_llm_proto_depIdxs,
		MessageInfos:      file_llm_proto_msgTypes,
	}.Build()
	File_llm_proto = out.File
	file_llm_proto_rawDesc = nil
	file_llm_proto_goTypes = nil
	file_llm_proto_depIdxs = nil
}
