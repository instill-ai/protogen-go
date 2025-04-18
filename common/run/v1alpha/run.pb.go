// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: common/run/v1alpha/run.proto

package runv1alpha

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
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

const file_common_run_v1alpha_run_proto_rawDesc = "" +
	"\n" +
	"\x1ccommon/run/v1alpha/run.proto\x12\x12common.run.v1alpha*\x8a\x01\n" +
	"\tRunStatus\x12\x1a\n" +
	"\x16RUN_STATUS_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15RUN_STATUS_PROCESSING\x10\x01\x12\x18\n" +
	"\x14RUN_STATUS_COMPLETED\x10\x02\x12\x15\n" +
	"\x11RUN_STATUS_FAILED\x10\x03\x12\x15\n" +
	"\x11RUN_STATUS_QUEUED\x10\x04*S\n" +
	"\tRunSource\x12\x1a\n" +
	"\x16RUN_SOURCE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12RUN_SOURCE_CONSOLE\x10\x01\x12\x12\n" +
	"\x0eRUN_SOURCE_API\x10\x02B\xcd\x01\n" +
	"\x16com.common.run.v1alphaB\bRunProtoP\x01Z?github.com/instill-ai/protogen-go/common/run/v1alpha;runv1alpha\xa2\x02\x03CRX\xaa\x02\x12Common.Run.V1alpha\xca\x02\x12Common\\Run\\V1alpha\xe2\x02\x1eCommon\\Run\\V1alpha\\GPBMetadata\xea\x02\x14Common::Run::V1alphab\x06proto3"

var (
	file_common_run_v1alpha_run_proto_rawDescOnce sync.Once
	file_common_run_v1alpha_run_proto_rawDescData []byte
)

func file_common_run_v1alpha_run_proto_rawDescGZIP() []byte {
	file_common_run_v1alpha_run_proto_rawDescOnce.Do(func() {
		file_common_run_v1alpha_run_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_common_run_v1alpha_run_proto_rawDesc), len(file_common_run_v1alpha_run_proto_rawDesc)))
	})
	return file_common_run_v1alpha_run_proto_rawDescData
}

var file_common_run_v1alpha_run_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_run_v1alpha_run_proto_goTypes = []any{
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_run_v1alpha_run_proto_rawDesc), len(file_common_run_v1alpha_run_proto_rawDesc)),
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
	file_common_run_v1alpha_run_proto_goTypes = nil
	file_common_run_v1alpha_run_proto_depIdxs = nil
}
