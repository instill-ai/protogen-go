// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vdp/mgmt/v1alpha/mgmt_private_service.proto

package mgmtv1alpha

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

var File_vdp_mgmt_v1alpha_mgmt_private_service_proto protoreflect.FileDescriptor

var file_vdp_mgmt_v1alpha_mgmt_private_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76,
	0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x67, 0x6d, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xbc, 0x07, 0x0a, 0x12, 0x4d, 0x67, 0x6d, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x27, 0x2e,
	0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x91,
	0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76,
	0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0xda, 0x41, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x14, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x25, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2c, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d,
	0x12, 0xab, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0xda, 0x41, 0x10, 0x75, 0x73,
	0x65, 0x72, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x22, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x94,
	0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76,
	0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xa5, 0x01, 0x0a, 0x0f, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x76, 0x64, 0x70, 0x2e,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f,
	0x6b, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d,
	0xda, 0x41, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x7b, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x6e, 0x6b, 0x3d, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x6f, 0x6f, 0x6b, 0x55, 0x70, 0x1a, 0x13, 0xca,
	0x41, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65,
	0x63, 0x68, 0x42, 0xd1, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x17, 0x4d, 0x67, 0x6d,
	0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x67, 0x6d, 0x74, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x56, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x56,
	0x64, 0x70, 0x2e, 0x4d, 0x67, 0x6d, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca,
	0x02, 0x10, 0x56, 0x64, 0x70, 0x5c, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xe2, 0x02, 0x1c, 0x56, 0x64, 0x70, 0x5c, 0x4d, 0x67, 0x6d, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x56, 0x64, 0x70, 0x3a, 0x3a, 0x4d, 0x67, 0x6d, 0x74, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_vdp_mgmt_v1alpha_mgmt_private_service_proto_goTypes = []interface{}{
	(*ListUsersAdminRequest)(nil),   // 0: vdp.mgmt.v1alpha.ListUsersAdminRequest
	(*CreateUserAdminRequest)(nil),  // 1: vdp.mgmt.v1alpha.CreateUserAdminRequest
	(*GetUserAdminRequest)(nil),     // 2: vdp.mgmt.v1alpha.GetUserAdminRequest
	(*UpdateUserAdminRequest)(nil),  // 3: vdp.mgmt.v1alpha.UpdateUserAdminRequest
	(*DeleteUserAdminRequest)(nil),  // 4: vdp.mgmt.v1alpha.DeleteUserAdminRequest
	(*LookUpUserAdminRequest)(nil),  // 5: vdp.mgmt.v1alpha.LookUpUserAdminRequest
	(*ListUsersAdminResponse)(nil),  // 6: vdp.mgmt.v1alpha.ListUsersAdminResponse
	(*CreateUserAdminResponse)(nil), // 7: vdp.mgmt.v1alpha.CreateUserAdminResponse
	(*GetUserAdminResponse)(nil),    // 8: vdp.mgmt.v1alpha.GetUserAdminResponse
	(*UpdateUserAdminResponse)(nil), // 9: vdp.mgmt.v1alpha.UpdateUserAdminResponse
	(*DeleteUserAdminResponse)(nil), // 10: vdp.mgmt.v1alpha.DeleteUserAdminResponse
	(*LookUpUserAdminResponse)(nil), // 11: vdp.mgmt.v1alpha.LookUpUserAdminResponse
}
var file_vdp_mgmt_v1alpha_mgmt_private_service_proto_depIdxs = []int32{
	0,  // 0: vdp.mgmt.v1alpha.MgmtPrivateService.ListUsersAdmin:input_type -> vdp.mgmt.v1alpha.ListUsersAdminRequest
	1,  // 1: vdp.mgmt.v1alpha.MgmtPrivateService.CreateUserAdmin:input_type -> vdp.mgmt.v1alpha.CreateUserAdminRequest
	2,  // 2: vdp.mgmt.v1alpha.MgmtPrivateService.GetUserAdmin:input_type -> vdp.mgmt.v1alpha.GetUserAdminRequest
	3,  // 3: vdp.mgmt.v1alpha.MgmtPrivateService.UpdateUserAdmin:input_type -> vdp.mgmt.v1alpha.UpdateUserAdminRequest
	4,  // 4: vdp.mgmt.v1alpha.MgmtPrivateService.DeleteUserAdmin:input_type -> vdp.mgmt.v1alpha.DeleteUserAdminRequest
	5,  // 5: vdp.mgmt.v1alpha.MgmtPrivateService.LookUpUserAdmin:input_type -> vdp.mgmt.v1alpha.LookUpUserAdminRequest
	6,  // 6: vdp.mgmt.v1alpha.MgmtPrivateService.ListUsersAdmin:output_type -> vdp.mgmt.v1alpha.ListUsersAdminResponse
	7,  // 7: vdp.mgmt.v1alpha.MgmtPrivateService.CreateUserAdmin:output_type -> vdp.mgmt.v1alpha.CreateUserAdminResponse
	8,  // 8: vdp.mgmt.v1alpha.MgmtPrivateService.GetUserAdmin:output_type -> vdp.mgmt.v1alpha.GetUserAdminResponse
	9,  // 9: vdp.mgmt.v1alpha.MgmtPrivateService.UpdateUserAdmin:output_type -> vdp.mgmt.v1alpha.UpdateUserAdminResponse
	10, // 10: vdp.mgmt.v1alpha.MgmtPrivateService.DeleteUserAdmin:output_type -> vdp.mgmt.v1alpha.DeleteUserAdminResponse
	11, // 11: vdp.mgmt.v1alpha.MgmtPrivateService.LookUpUserAdmin:output_type -> vdp.mgmt.v1alpha.LookUpUserAdminResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_vdp_mgmt_v1alpha_mgmt_private_service_proto_init() }
func file_vdp_mgmt_v1alpha_mgmt_private_service_proto_init() {
	if File_vdp_mgmt_v1alpha_mgmt_private_service_proto != nil {
		return
	}
	file_vdp_mgmt_v1alpha_mgmt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_mgmt_v1alpha_mgmt_private_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vdp_mgmt_v1alpha_mgmt_private_service_proto_goTypes,
		DependencyIndexes: file_vdp_mgmt_v1alpha_mgmt_private_service_proto_depIdxs,
	}.Build()
	File_vdp_mgmt_v1alpha_mgmt_private_service_proto = out.File
	file_vdp_mgmt_v1alpha_mgmt_private_service_proto_rawDesc = nil
	file_vdp_mgmt_v1alpha_mgmt_private_service_proto_goTypes = nil
	file_vdp_mgmt_v1alpha_mgmt_private_service_proto_depIdxs = nil
}