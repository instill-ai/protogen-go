// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vdp/usage/v1alpha/usage_service.proto

package usagev1alpha

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_vdp_usage_v1alpha_usage_service_proto protoreflect.FileDescriptor

var file_vdp_usage_v1alpha_usage_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x76, 0x64, 0x70, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x76, 0x64, 0x70, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x76, 0x64, 0x70, 0x2f, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbf, 0x04, 0x0a, 0x0c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x6e,
	0x65, 0x73, 0x73, 0x12, 0x22, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x76, 0x65,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2e, 0x5a, 0x17, 0x12, 0x15, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x5f, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x6e, 0x65,
	0x73, 0x73, 0x12, 0x74, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x23, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x5f, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0xda,
	0x41, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a,
	0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x99, 0x01, 0x0a, 0x11,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x2b, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0xda, 0x41,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x06, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x42, 0xd2, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x42, 0x11, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x64, 0x70, 0x2f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x56, 0x55, 0x58, 0xaa, 0x02,
	0x11, 0x56, 0x64, 0x70, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xca, 0x02, 0x11, 0x56, 0x64, 0x70, 0x5c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1d, 0x56, 0x64, 0x70, 0x5c, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x56, 0x64, 0x70, 0x3a, 0x3a, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_vdp_usage_v1alpha_usage_service_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),           // 0: vdp.usage.v1alpha.LivenessRequest
	(*ReadinessRequest)(nil),          // 1: vdp.usage.v1alpha.ReadinessRequest
	(*CreateSessionRequest)(nil),      // 2: vdp.usage.v1alpha.CreateSessionRequest
	(*SendSessionReportRequest)(nil),  // 3: vdp.usage.v1alpha.SendSessionReportRequest
	(*LivenessResponse)(nil),          // 4: vdp.usage.v1alpha.LivenessResponse
	(*ReadinessResponse)(nil),         // 5: vdp.usage.v1alpha.ReadinessResponse
	(*CreateSessionResponse)(nil),     // 6: vdp.usage.v1alpha.CreateSessionResponse
	(*SendSessionReportResponse)(nil), // 7: vdp.usage.v1alpha.SendSessionReportResponse
}
var file_vdp_usage_v1alpha_usage_service_proto_depIdxs = []int32{
	0, // 0: vdp.usage.v1alpha.UsageService.Liveness:input_type -> vdp.usage.v1alpha.LivenessRequest
	1, // 1: vdp.usage.v1alpha.UsageService.Readiness:input_type -> vdp.usage.v1alpha.ReadinessRequest
	2, // 2: vdp.usage.v1alpha.UsageService.CreateSession:input_type -> vdp.usage.v1alpha.CreateSessionRequest
	3, // 3: vdp.usage.v1alpha.UsageService.SendSessionReport:input_type -> vdp.usage.v1alpha.SendSessionReportRequest
	4, // 4: vdp.usage.v1alpha.UsageService.Liveness:output_type -> vdp.usage.v1alpha.LivenessResponse
	5, // 5: vdp.usage.v1alpha.UsageService.Readiness:output_type -> vdp.usage.v1alpha.ReadinessResponse
	6, // 6: vdp.usage.v1alpha.UsageService.CreateSession:output_type -> vdp.usage.v1alpha.CreateSessionResponse
	7, // 7: vdp.usage.v1alpha.UsageService.SendSessionReport:output_type -> vdp.usage.v1alpha.SendSessionReportResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vdp_usage_v1alpha_usage_service_proto_init() }
func file_vdp_usage_v1alpha_usage_service_proto_init() {
	if File_vdp_usage_v1alpha_usage_service_proto != nil {
		return
	}
	file_vdp_usage_v1alpha_healthcheck_proto_init()
	file_vdp_usage_v1alpha_usage_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_usage_v1alpha_usage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vdp_usage_v1alpha_usage_service_proto_goTypes,
		DependencyIndexes: file_vdp_usage_v1alpha_usage_service_proto_depIdxs,
	}.Build()
	File_vdp_usage_v1alpha_usage_service_proto = out.File
	file_vdp_usage_v1alpha_usage_service_proto_rawDesc = nil
	file_vdp_usage_v1alpha_usage_service_proto_goTypes = nil
	file_vdp_usage_v1alpha_usage_service_proto_depIdxs = nil
}
