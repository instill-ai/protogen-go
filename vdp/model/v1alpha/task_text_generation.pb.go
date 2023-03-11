// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vdp/model/v1alpha/task_text_generation.proto

package modelv1alpha

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// TextGenerationInput represents the input of text generation task
type TextGenerationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prompt text
	Prompt string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// The maximum number of tokens for model to generate
	OutputLen *int64 `protobuf:"varint,2,opt,name=output_len,json=outputLen,proto3,oneof" json:"output_len,omitempty"`
	// The words not to avoid being generated by the model
	BadWordsList *string `protobuf:"bytes,3,opt,name=bad_words_list,json=badWordsList,proto3,oneof" json:"bad_words_list,omitempty"`
	// The trigger words to stop generation
	StopWordsList *string `protobuf:"bytes,4,opt,name=stop_words_list,json=stopWordsList,proto3,oneof" json:"stop_words_list,omitempty"`
	// Top k for sampling
	Topk *int64 `protobuf:"varint,5,opt,name=topk,proto3,oneof" json:"topk,omitempty"`
	// The seed
	Seed *int64 `protobuf:"varint,6,opt,name=seed,proto3,oneof" json:"seed,omitempty"`
}

func (x *TextGenerationInput) Reset() {
	*x = TextGenerationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextGenerationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextGenerationInput) ProtoMessage() {}

func (x *TextGenerationInput) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextGenerationInput.ProtoReflect.Descriptor instead.
func (*TextGenerationInput) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_task_text_generation_proto_rawDescGZIP(), []int{0}
}

func (x *TextGenerationInput) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *TextGenerationInput) GetOutputLen() int64 {
	if x != nil && x.OutputLen != nil {
		return *x.OutputLen
	}
	return 0
}

func (x *TextGenerationInput) GetBadWordsList() string {
	if x != nil && x.BadWordsList != nil {
		return *x.BadWordsList
	}
	return ""
}

func (x *TextGenerationInput) GetStopWordsList() string {
	if x != nil && x.StopWordsList != nil {
		return *x.StopWordsList
	}
	return ""
}

func (x *TextGenerationInput) GetTopk() int64 {
	if x != nil && x.Topk != nil {
		return *x.Topk
	}
	return 0
}

func (x *TextGenerationInput) GetSeed() int64 {
	if x != nil && x.Seed != nil {
		return *x.Seed
	}
	return 0
}

// TextGenerationOutput represents the output of text generation task
type TextGenerationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The text generated by the model
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextGenerationOutput) Reset() {
	*x = TextGenerationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextGenerationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextGenerationOutput) ProtoMessage() {}

func (x *TextGenerationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextGenerationOutput.ProtoReflect.Descriptor instead.
func (*TextGenerationOutput) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_task_text_generation_proto_rawDescGZIP(), []int{1}
}

func (x *TextGenerationOutput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_vdp_model_v1alpha_task_text_generation_proto protoreflect.FileDescriptor

var file_vdp_model_v1alpha_task_text_generation_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x13, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x48, 0x00, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4c, 0x65, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x2e, 0x0a, 0x0e, 0x62, 0x61, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x01, 0x52,
	0x0c, 0x62, 0x61, 0x64, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x30, 0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x02,
	0x52, 0x0d, 0x73, 0x74, 0x6f, 0x70, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x6f, 0x70, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x6b, 0x88, 0x01, 0x01,
	0x12, 0x1c, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x48, 0x04, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x62, 0x61, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x6f, 0x70, 0x6b, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x73, 0x65, 0x65, 0x64, 0x22, 0x2f, 0x0a, 0x14, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0xd8, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x42, 0x17, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c,
	0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02,
	0x03, 0x56, 0x4d, 0x58, 0xaa, 0x02, 0x11, 0x56, 0x64, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x11, 0x56, 0x64, 0x70, 0x5c, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1d, 0x56,
	0x64, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x56,
	0x64, 0x70, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vdp_model_v1alpha_task_text_generation_proto_rawDescOnce sync.Once
	file_vdp_model_v1alpha_task_text_generation_proto_rawDescData = file_vdp_model_v1alpha_task_text_generation_proto_rawDesc
)

func file_vdp_model_v1alpha_task_text_generation_proto_rawDescGZIP() []byte {
	file_vdp_model_v1alpha_task_text_generation_proto_rawDescOnce.Do(func() {
		file_vdp_model_v1alpha_task_text_generation_proto_rawDescData = protoimpl.X.CompressGZIP(file_vdp_model_v1alpha_task_text_generation_proto_rawDescData)
	})
	return file_vdp_model_v1alpha_task_text_generation_proto_rawDescData
}

var file_vdp_model_v1alpha_task_text_generation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vdp_model_v1alpha_task_text_generation_proto_goTypes = []interface{}{
	(*TextGenerationInput)(nil),  // 0: vdp.model.v1alpha.TextGenerationInput
	(*TextGenerationOutput)(nil), // 1: vdp.model.v1alpha.TextGenerationOutput
}
var file_vdp_model_v1alpha_task_text_generation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vdp_model_v1alpha_task_text_generation_proto_init() }
func file_vdp_model_v1alpha_task_text_generation_proto_init() {
	if File_vdp_model_v1alpha_task_text_generation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextGenerationInput); i {
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
		file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextGenerationOutput); i {
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
	file_vdp_model_v1alpha_task_text_generation_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_model_v1alpha_task_text_generation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vdp_model_v1alpha_task_text_generation_proto_goTypes,
		DependencyIndexes: file_vdp_model_v1alpha_task_text_generation_proto_depIdxs,
		MessageInfos:      file_vdp_model_v1alpha_task_text_generation_proto_msgTypes,
	}.Build()
	File_vdp_model_v1alpha_task_text_generation_proto = out.File
	file_vdp_model_v1alpha_task_text_generation_proto_rawDesc = nil
	file_vdp_model_v1alpha_task_text_generation_proto_goTypes = nil
	file_vdp_model_v1alpha_task_text_generation_proto_depIdxs = nil
}