// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: core/mgmt/v1beta/mgmt_private_service.proto

package mgmtv1beta

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

var File_core_mgmt_v1beta_mgmt_private_service_proto protoreflect.FileDescriptor

var file_core_mgmt_v1beta_mgmt_private_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a,
	0x1b, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xac, 0x0e, 0x0a, 0x12, 0x4d, 0x67, 0x6d, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xa4, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x6f, 0x6b,
	0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c,
	0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3c, 0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x12, 0xa0,
	0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0xaa, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0xda, 0x41, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xc4,
	0x01, 0x0a, 0x17, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x6f,
	0x6f, 0x6b, 0x55, 0x70, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x44, 0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c,
	0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x12, 0xbf, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0xda, 0x41, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xdf, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x39, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x44, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x83, 0x01, 0x0a, 0x13, 0x53, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f,
	0xda, 0x41, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x88, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x30, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x08, 0xda, 0x41, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a, 0x13, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x2c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0c, 0xda, 0x41, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x10, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x42,
	0xd0, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x17, 0x4d, 0x67, 0x6d, 0x74, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x67, 0x6d, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x6f, 0x72, 0x65, 0x2e,
	0x4d, 0x67, 0x6d, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xca, 0x02, 0x10, 0x43, 0x6f,
	0x72, 0x65, 0x5c, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xe2, 0x02,
	0x1c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12,
	0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_core_mgmt_v1beta_mgmt_private_service_proto_goTypes = []interface{}{
	(*ListUsersAdminRequest)(nil),                    // 0: core.mgmt.v1beta.ListUsersAdminRequest
	(*GetUserAdminRequest)(nil),                      // 1: core.mgmt.v1beta.GetUserAdminRequest
	(*LookUpUserAdminRequest)(nil),                   // 2: core.mgmt.v1beta.LookUpUserAdminRequest
	(*ListOrganizationsAdminRequest)(nil),            // 3: core.mgmt.v1beta.ListOrganizationsAdminRequest
	(*GetOrganizationAdminRequest)(nil),              // 4: core.mgmt.v1beta.GetOrganizationAdminRequest
	(*LookUpOrganizationAdminRequest)(nil),           // 5: core.mgmt.v1beta.LookUpOrganizationAdminRequest
	(*GetUserSubscriptionAdminRequest)(nil),          // 6: core.mgmt.v1beta.GetUserSubscriptionAdminRequest
	(*GetOrganizationSubscriptionAdminRequest)(nil),  // 7: core.mgmt.v1beta.GetOrganizationSubscriptionAdminRequest
	(*SubtractCreditAdminRequest)(nil),               // 8: core.mgmt.v1beta.SubtractCreditAdminRequest
	(*GetRemainingCreditAdminRequest)(nil),           // 9: core.mgmt.v1beta.GetRemainingCreditAdminRequest
	(*CheckNamespaceAdminRequest)(nil),               // 10: core.mgmt.v1beta.CheckNamespaceAdminRequest
	(*ListUsersAdminResponse)(nil),                   // 11: core.mgmt.v1beta.ListUsersAdminResponse
	(*GetUserAdminResponse)(nil),                     // 12: core.mgmt.v1beta.GetUserAdminResponse
	(*LookUpUserAdminResponse)(nil),                  // 13: core.mgmt.v1beta.LookUpUserAdminResponse
	(*ListOrganizationsAdminResponse)(nil),           // 14: core.mgmt.v1beta.ListOrganizationsAdminResponse
	(*GetOrganizationAdminResponse)(nil),             // 15: core.mgmt.v1beta.GetOrganizationAdminResponse
	(*LookUpOrganizationAdminResponse)(nil),          // 16: core.mgmt.v1beta.LookUpOrganizationAdminResponse
	(*GetUserSubscriptionAdminResponse)(nil),         // 17: core.mgmt.v1beta.GetUserSubscriptionAdminResponse
	(*GetOrganizationSubscriptionAdminResponse)(nil), // 18: core.mgmt.v1beta.GetOrganizationSubscriptionAdminResponse
	(*SubtractCreditAdminResponse)(nil),              // 19: core.mgmt.v1beta.SubtractCreditAdminResponse
	(*GetRemainingCreditAdminResponse)(nil),          // 20: core.mgmt.v1beta.GetRemainingCreditAdminResponse
	(*CheckNamespaceAdminResponse)(nil),              // 21: core.mgmt.v1beta.CheckNamespaceAdminResponse
}
var file_core_mgmt_v1beta_mgmt_private_service_proto_depIdxs = []int32{
	0,  // 0: core.mgmt.v1beta.MgmtPrivateService.ListUsersAdmin:input_type -> core.mgmt.v1beta.ListUsersAdminRequest
	1,  // 1: core.mgmt.v1beta.MgmtPrivateService.GetUserAdmin:input_type -> core.mgmt.v1beta.GetUserAdminRequest
	2,  // 2: core.mgmt.v1beta.MgmtPrivateService.LookUpUserAdmin:input_type -> core.mgmt.v1beta.LookUpUserAdminRequest
	3,  // 3: core.mgmt.v1beta.MgmtPrivateService.ListOrganizationsAdmin:input_type -> core.mgmt.v1beta.ListOrganizationsAdminRequest
	4,  // 4: core.mgmt.v1beta.MgmtPrivateService.GetOrganizationAdmin:input_type -> core.mgmt.v1beta.GetOrganizationAdminRequest
	5,  // 5: core.mgmt.v1beta.MgmtPrivateService.LookUpOrganizationAdmin:input_type -> core.mgmt.v1beta.LookUpOrganizationAdminRequest
	6,  // 6: core.mgmt.v1beta.MgmtPrivateService.GetUserSubscriptionAdmin:input_type -> core.mgmt.v1beta.GetUserSubscriptionAdminRequest
	7,  // 7: core.mgmt.v1beta.MgmtPrivateService.GetOrganizationSubscriptionAdmin:input_type -> core.mgmt.v1beta.GetOrganizationSubscriptionAdminRequest
	8,  // 8: core.mgmt.v1beta.MgmtPrivateService.SubtractCreditAdmin:input_type -> core.mgmt.v1beta.SubtractCreditAdminRequest
	9,  // 9: core.mgmt.v1beta.MgmtPrivateService.GetRemainingCreditAdmin:input_type -> core.mgmt.v1beta.GetRemainingCreditAdminRequest
	10, // 10: core.mgmt.v1beta.MgmtPrivateService.CheckNamespaceAdmin:input_type -> core.mgmt.v1beta.CheckNamespaceAdminRequest
	11, // 11: core.mgmt.v1beta.MgmtPrivateService.ListUsersAdmin:output_type -> core.mgmt.v1beta.ListUsersAdminResponse
	12, // 12: core.mgmt.v1beta.MgmtPrivateService.GetUserAdmin:output_type -> core.mgmt.v1beta.GetUserAdminResponse
	13, // 13: core.mgmt.v1beta.MgmtPrivateService.LookUpUserAdmin:output_type -> core.mgmt.v1beta.LookUpUserAdminResponse
	14, // 14: core.mgmt.v1beta.MgmtPrivateService.ListOrganizationsAdmin:output_type -> core.mgmt.v1beta.ListOrganizationsAdminResponse
	15, // 15: core.mgmt.v1beta.MgmtPrivateService.GetOrganizationAdmin:output_type -> core.mgmt.v1beta.GetOrganizationAdminResponse
	16, // 16: core.mgmt.v1beta.MgmtPrivateService.LookUpOrganizationAdmin:output_type -> core.mgmt.v1beta.LookUpOrganizationAdminResponse
	17, // 17: core.mgmt.v1beta.MgmtPrivateService.GetUserSubscriptionAdmin:output_type -> core.mgmt.v1beta.GetUserSubscriptionAdminResponse
	18, // 18: core.mgmt.v1beta.MgmtPrivateService.GetOrganizationSubscriptionAdmin:output_type -> core.mgmt.v1beta.GetOrganizationSubscriptionAdminResponse
	19, // 19: core.mgmt.v1beta.MgmtPrivateService.SubtractCreditAdmin:output_type -> core.mgmt.v1beta.SubtractCreditAdminResponse
	20, // 20: core.mgmt.v1beta.MgmtPrivateService.GetRemainingCreditAdmin:output_type -> core.mgmt.v1beta.GetRemainingCreditAdminResponse
	21, // 21: core.mgmt.v1beta.MgmtPrivateService.CheckNamespaceAdmin:output_type -> core.mgmt.v1beta.CheckNamespaceAdminResponse
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_core_mgmt_v1beta_mgmt_private_service_proto_init() }
func file_core_mgmt_v1beta_mgmt_private_service_proto_init() {
	if File_core_mgmt_v1beta_mgmt_private_service_proto != nil {
		return
	}
	file_core_mgmt_v1beta_mgmt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_mgmt_v1beta_mgmt_private_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_mgmt_v1beta_mgmt_private_service_proto_goTypes,
		DependencyIndexes: file_core_mgmt_v1beta_mgmt_private_service_proto_depIdxs,
	}.Build()
	File_core_mgmt_v1beta_mgmt_private_service_proto = out.File
	file_core_mgmt_v1beta_mgmt_private_service_proto_rawDesc = nil
	file_core_mgmt_v1beta_mgmt_private_service_proto_goTypes = nil
	file_core_mgmt_v1beta_mgmt_private_service_proto_depIdxs = nil
}
