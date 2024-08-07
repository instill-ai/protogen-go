// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: artifact/artifact/v1alpha/qa.proto

package artifactv1alpha

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

// QuestionAnsweringRequest
type QuestionAnsweringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of the namespace
	NamespaceId string `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	// id of the catalog
	CatalogId string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	// question to be answered
	Question string `protobuf:"bytes,3,opt,name=question,proto3" json:"question,omitempty"`
	// top k default to 5
	TopK int32 `protobuf:"varint,4,opt,name=top_k,json=topK,proto3" json:"top_k,omitempty"`
}

func (x *QuestionAnsweringRequest) Reset() {
	*x = QuestionAnsweringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_qa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionAnsweringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionAnsweringRequest) ProtoMessage() {}

func (x *QuestionAnsweringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_qa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionAnsweringRequest.ProtoReflect.Descriptor instead.
func (*QuestionAnsweringRequest) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_qa_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionAnsweringRequest) GetNamespaceId() string {
	if x != nil {
		return x.NamespaceId
	}
	return ""
}

func (x *QuestionAnsweringRequest) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

func (x *QuestionAnsweringRequest) GetQuestion() string {
	if x != nil {
		return x.Question
	}
	return ""
}

func (x *QuestionAnsweringRequest) GetTopK() int32 {
	if x != nil {
		return x.TopK
	}
	return 0
}

// QuestionAnsweringResponse
type QuestionAnsweringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// answer to the question
	Answer string `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	// chunks
	SimilarChunks []*SimilarityChunk `protobuf:"bytes,2,rep,name=similar_chunks,json=similarChunks,proto3" json:"similar_chunks,omitempty"`
}

func (x *QuestionAnsweringResponse) Reset() {
	*x = QuestionAnsweringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifact_v1alpha_qa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionAnsweringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionAnsweringResponse) ProtoMessage() {}

func (x *QuestionAnsweringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifact_v1alpha_qa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionAnsweringResponse.ProtoReflect.Descriptor instead.
func (*QuestionAnsweringResponse) Descriptor() ([]byte, []int) {
	return file_artifact_artifact_v1alpha_qa_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionAnsweringResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *QuestionAnsweringResponse) GetSimilarChunks() []*SimilarityChunk {
	if x != nil {
		return x.SimilarChunks
	}
	return nil
}

var File_artifact_artifact_v1alpha_qa_proto protoreflect.FileDescriptor

var file_artifact_artifact_v1alpha_qa_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x71, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x25, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x22, 0x86, 0x01, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0e,
	0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x52, 0x0d, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x42,
	0xfb, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x42, 0x07, 0x51, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c,
	0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa,
	0x02, 0x19, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x19, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x25, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1b, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_artifact_v1alpha_qa_proto_rawDescOnce sync.Once
	file_artifact_artifact_v1alpha_qa_proto_rawDescData = file_artifact_artifact_v1alpha_qa_proto_rawDesc
)

func file_artifact_artifact_v1alpha_qa_proto_rawDescGZIP() []byte {
	file_artifact_artifact_v1alpha_qa_proto_rawDescOnce.Do(func() {
		file_artifact_artifact_v1alpha_qa_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_artifact_v1alpha_qa_proto_rawDescData)
	})
	return file_artifact_artifact_v1alpha_qa_proto_rawDescData
}

var file_artifact_artifact_v1alpha_qa_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artifact_artifact_v1alpha_qa_proto_goTypes = []interface{}{
	(*QuestionAnsweringRequest)(nil),  // 0: artifact.artifact.v1alpha.QuestionAnsweringRequest
	(*QuestionAnsweringResponse)(nil), // 1: artifact.artifact.v1alpha.QuestionAnsweringResponse
	(*SimilarityChunk)(nil),           // 2: artifact.artifact.v1alpha.SimilarityChunk
}
var file_artifact_artifact_v1alpha_qa_proto_depIdxs = []int32{
	2, // 0: artifact.artifact.v1alpha.QuestionAnsweringResponse.similar_chunks:type_name -> artifact.artifact.v1alpha.SimilarityChunk
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_artifact_artifact_v1alpha_qa_proto_init() }
func file_artifact_artifact_v1alpha_qa_proto_init() {
	if File_artifact_artifact_v1alpha_qa_proto != nil {
		return
	}
	file_artifact_artifact_v1alpha_chunk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_artifact_artifact_v1alpha_qa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionAnsweringRequest); i {
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
		file_artifact_artifact_v1alpha_qa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionAnsweringResponse); i {
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
			RawDescriptor: file_artifact_artifact_v1alpha_qa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_artifact_v1alpha_qa_proto_goTypes,
		DependencyIndexes: file_artifact_artifact_v1alpha_qa_proto_depIdxs,
		MessageInfos:      file_artifact_artifact_v1alpha_qa_proto_msgTypes,
	}.Build()
	File_artifact_artifact_v1alpha_qa_proto = out.File
	file_artifact_artifact_v1alpha_qa_proto_rawDesc = nil
	file_artifact_artifact_v1alpha_qa_proto_goTypes = nil
	file_artifact_artifact_v1alpha_qa_proto_depIdxs = nil
}
