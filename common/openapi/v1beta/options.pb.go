// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: common/openapi/v1beta/options.proto

package openapiv1beta

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_common_openapi_v1beta_options_proto protoreflect.FileDescriptor

var file_common_openapi_v1beta_options_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x93, 0x06, 0x92,
	0x41, 0xa9, 0x04, 0x12, 0xfa, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x20,
	0x41, 0x49, 0x20, 0x41, 0x50, 0x49, 0x12, 0x44, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x20, 0x41, 0x49,
	0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x20, 0x65, 0x78, 0x70, 0x6f, 0x73,
	0x65, 0x64, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x48, 0x54, 0x54, 0x50, 0x22, 0x3f, 0x0a, 0x0a,
	0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x20, 0x41, 0x49, 0x12, 0x1d, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x1a, 0x12, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x40, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2a, 0x57, 0x0a,
	0x1a, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65,
	0x20, 0x32, 0x2e, 0x30, 0x20, 0x28, 0x45, 0x4c, 0x76, 0x32, 0x29, 0x12, 0x39, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c,
	0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x08, 0x30, 0x2e, 0x34, 0x2d, 0x62, 0x65, 0x74, 0x61,
	0x1a, 0x18, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x52, 0x3c, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x35, 0x0a, 0x33, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x20, 0x61, 0x72, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x2e,
	0x5a, 0x64, 0x0a, 0x62, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x58, 0x08, 0x02,
	0x12, 0x43, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x60, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x20, 0x60, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2c, 0x20, 0x65, 0x2e, 0x67,
	0x2e, 0x20, 0x60, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x61, 0x62, 0x63, 0x64, 0x65, 0x31,
	0x32, 0x33, 0x34, 0x35, 0x60, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x12, 0x00, 0x72, 0x36, 0x0a, 0x15, 0x4d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x20, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x20, 0x41, 0x49, 0x12, 0x1d, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64, 0x6f, 0x63, 0x73, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x3b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0xa2, 0x02,
	0x03, 0x43, 0x4f, 0x58, 0xaa, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xca, 0x02, 0x15, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0xe2, 0x02, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x4f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x3a, 0x3a, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_common_openapi_v1beta_options_proto_goTypes = []interface{}{}
var file_common_openapi_v1beta_options_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_openapi_v1beta_options_proto_init() }
func file_common_openapi_v1beta_options_proto_init() {
	if File_common_openapi_v1beta_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_openapi_v1beta_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_openapi_v1beta_options_proto_goTypes,
		DependencyIndexes: file_common_openapi_v1beta_options_proto_depIdxs,
	}.Build()
	File_common_openapi_v1beta_options_proto = out.File
	file_common_openapi_v1beta_options_proto_rawDesc = nil
	file_common_openapi_v1beta_options_proto_goTypes = nil
	file_common_openapi_v1beta_options_proto_depIdxs = nil
}