// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: common/run/v1alpha/run.proto

package runv1alpha

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

// RunStatus defines the status of a pipeline or model run.
type RunStatus int32

const (
	// Unspecified.
	RunStatus_RUN_STATUS_UNSPECIFIED RunStatus = 0
	// Run in progress.
	RunStatus_RUN_STATUS_PROCESSING RunStatus = 1
	// Run succeeded.
	RunStatus_RUN_STATUS_COMPLETED RunStatus = 2
	// Run failed.
	RunStatus_RUN_STATUS_FAILED RunStatus = 3
	// Run is waiting to be executed.
	RunStatus_RUN_STATUS_QUEUED RunStatus = 4
)

// Enum value maps for RunStatus.
var (
	RunStatus_name = map[int32]string{
		0: "RUN_STATUS_UNSPECIFIED",
		1: "RUN_STATUS_PROCESSING",
		2: "RUN_STATUS_COMPLETED",
		3: "RUN_STATUS_FAILED",
		4: "RUN_STATUS_QUEUED",
	}
	RunStatus_value = map[string]int32{
		"RUN_STATUS_UNSPECIFIED": 0,
		"RUN_STATUS_PROCESSING":  1,
		"RUN_STATUS_COMPLETED":   2,
		"RUN_STATUS_FAILED":      3,
		"RUN_STATUS_QUEUED":      4,
	}
)

func (x RunStatus) Enum() *RunStatus {
	p := new(RunStatus)
	*p = x
	return p
}

func (x RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_run_v1alpha_run_proto_enumTypes[0].Descriptor()
}

func (RunStatus) Type() protoreflect.EnumType {
	return &file_common_run_v1alpha_run_proto_enumTypes[0]
}

func (x RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunStatus.Descriptor instead.
func (RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_run_v1alpha_run_proto_rawDescGZIP(), []int{0}
}

// RunSource defines the source of a pipeline or model run.
type RunSource int32

const (
	// Unspecified.
	RunSource_RUN_SOURCE_UNSPECIFIED RunSource = 0
	// Run from frontend UI.
	RunSource_RUN_SOURCE_CONSOLE RunSource = 1
	// Run from API or SDK.
	RunSource_RUN_SOURCE_API RunSource = 2
)

// Enum value maps for RunSource.
var (
	RunSource_name = map[int32]string{
		0: "RUN_SOURCE_UNSPECIFIED",
		1: "RUN_SOURCE_CONSOLE",
		2: "RUN_SOURCE_API",
	}
	RunSource_value = map[string]int32{
		"RUN_SOURCE_UNSPECIFIED": 0,
		"RUN_SOURCE_CONSOLE":     1,
		"RUN_SOURCE_API":         2,
	}
)

func (x RunSource) Enum() *RunSource {
	p := new(RunSource)
	*p = x
	return p
}

func (x RunSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunSource) Descriptor() protoreflect.EnumDescriptor {
	return file_common_run_v1alpha_run_proto_enumTypes[1].Descriptor()
}

func (RunSource) Type() protoreflect.EnumType {
	return &file_common_run_v1alpha_run_proto_enumTypes[1]
}

func (x RunSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunSource.Descriptor instead.
func (RunSource) EnumDescriptor() ([]byte, []int) {
	return file_common_run_v1alpha_run_proto_rawDescGZIP(), []int{1}
}

var File_common_run_v1alpha_run_proto protoreflect.FileDescriptor

var file_common_run_v1alpha_run_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2a, 0x8a, 0x01, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1a, 0x0a, 0x16, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45,
	0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x55, 0x4e, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x55, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x04, 0x2a,
	0x53, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x16,
	0x52, 0x55, 0x4e, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x55, 0x4e, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x4f, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x41,
	0x50, 0x49, 0x10, 0x02, 0x42, 0xcd, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x08, 0x52, 0x75, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d,
	0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x72, 0x75, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x43,
	0x52, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5c, 0x52, 0x75, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x52, 0x75, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x52, 0x75, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_run_v1alpha_run_proto_rawDescOnce sync.Once
	file_common_run_v1alpha_run_proto_rawDescData = file_common_run_v1alpha_run_proto_rawDesc
)

func file_common_run_v1alpha_run_proto_rawDescGZIP() []byte {
	file_common_run_v1alpha_run_proto_rawDescOnce.Do(func() {
		file_common_run_v1alpha_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_run_v1alpha_run_proto_rawDescData)
	})
	return file_common_run_v1alpha_run_proto_rawDescData
}

var file_common_run_v1alpha_run_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_run_v1alpha_run_proto_goTypes = []interface{}{
	(RunStatus)(0), // 0: common.run.v1alpha.RunStatus
	(RunSource)(0), // 1: common.run.v1alpha.RunSource
}
var file_common_run_v1alpha_run_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_run_v1alpha_run_proto_init() }
func file_common_run_v1alpha_run_proto_init() {
	if File_common_run_v1alpha_run_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_run_v1alpha_run_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_run_v1alpha_run_proto_goTypes,
		DependencyIndexes: file_common_run_v1alpha_run_proto_depIdxs,
		EnumInfos:         file_common_run_v1alpha_run_proto_enumTypes,
	}.Build()
	File_common_run_v1alpha_run_proto = out.File
	file_common_run_v1alpha_run_proto_rawDesc = nil
	file_common_run_v1alpha_run_proto_goTypes = nil
	file_common_run_v1alpha_run_proto_depIdxs = nil
}