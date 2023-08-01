// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/model/v1alpha/task_instance_segmentation.proto

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

// InstanceSegmentationObject corresponding to a instance segmentation object
type InstanceSegmentationObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instance RLE segmentation mask
	Rle string `protobuf:"bytes,1,opt,name=rle,proto3" json:"rle,omitempty"`
	// Instance category
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// Instance score
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	// Instance bounding box
	BoundingBox *BoundingBox `protobuf:"bytes,4,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
}

func (x *InstanceSegmentationObject) Reset() {
	*x = InstanceSegmentationObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceSegmentationObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceSegmentationObject) ProtoMessage() {}

func (x *InstanceSegmentationObject) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceSegmentationObject.ProtoReflect.Descriptor instead.
func (*InstanceSegmentationObject) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_instance_segmentation_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceSegmentationObject) GetRle() string {
	if x != nil {
		return x.Rle
	}
	return ""
}

func (x *InstanceSegmentationObject) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *InstanceSegmentationObject) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *InstanceSegmentationObject) GetBoundingBox() *BoundingBox {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

// InstanceSegmentationInput represents the input of instance segmentation task
type InstanceSegmentationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input type
	//
	// Types that are assignable to Type:
	//
	//	*InstanceSegmentationInput_ImageUrl
	//	*InstanceSegmentationInput_ImageBase64
	Type isInstanceSegmentationInput_Type `protobuf_oneof:"type"`
}

func (x *InstanceSegmentationInput) Reset() {
	*x = InstanceSegmentationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceSegmentationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceSegmentationInput) ProtoMessage() {}

func (x *InstanceSegmentationInput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceSegmentationInput.ProtoReflect.Descriptor instead.
func (*InstanceSegmentationInput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_instance_segmentation_proto_rawDescGZIP(), []int{1}
}

func (m *InstanceSegmentationInput) GetType() isInstanceSegmentationInput_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *InstanceSegmentationInput) GetImageUrl() string {
	if x, ok := x.GetType().(*InstanceSegmentationInput_ImageUrl); ok {
		return x.ImageUrl
	}
	return ""
}

func (x *InstanceSegmentationInput) GetImageBase64() string {
	if x, ok := x.GetType().(*InstanceSegmentationInput_ImageBase64); ok {
		return x.ImageBase64
	}
	return ""
}

type isInstanceSegmentationInput_Type interface {
	isInstanceSegmentationInput_Type()
}

type InstanceSegmentationInput_ImageUrl struct {
	// Image type URL
	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3,oneof"`
}

type InstanceSegmentationInput_ImageBase64 struct {
	// Image type base64
	ImageBase64 string `protobuf:"bytes,2,opt,name=image_base64,json=imageBase64,proto3,oneof"`
}

func (*InstanceSegmentationInput_ImageUrl) isInstanceSegmentationInput_Type() {}

func (*InstanceSegmentationInput_ImageBase64) isInstanceSegmentationInput_Type() {}

// InstanceSegmentationInputStream represents the input of instance segmentation
// task when using stream method
type InstanceSegmentationInputStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of file length for each uploaded binary file
	FileLengths []uint64 `protobuf:"varint,1,rep,packed,name=file_lengths,json=fileLengths,proto3" json:"file_lengths,omitempty"`
	// Content of images in bytes
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *InstanceSegmentationInputStream) Reset() {
	*x = InstanceSegmentationInputStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceSegmentationInputStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceSegmentationInputStream) ProtoMessage() {}

func (x *InstanceSegmentationInputStream) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceSegmentationInputStream.ProtoReflect.Descriptor instead.
func (*InstanceSegmentationInputStream) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_instance_segmentation_proto_rawDescGZIP(), []int{2}
}

func (x *InstanceSegmentationInputStream) GetFileLengths() []uint64 {
	if x != nil {
		return x.FileLengths
	}
	return nil
}

func (x *InstanceSegmentationInputStream) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// InstanceSegmentationOutput represents the output of instance segmentation
// task
type InstanceSegmentationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of instance segmentation objects
	Objects []*InstanceSegmentationObject `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
}

func (x *InstanceSegmentationOutput) Reset() {
	*x = InstanceSegmentationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceSegmentationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceSegmentationOutput) ProtoMessage() {}

func (x *InstanceSegmentationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceSegmentationOutput.ProtoReflect.Descriptor instead.
func (*InstanceSegmentationOutput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_instance_segmentation_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceSegmentationOutput) GetObjects() []*InstanceSegmentationObject {
	if x != nil {
		return x.Objects
	}
	return nil
}

var File_model_model_v1alpha_task_instance_segmentation_proto protoreflect.FileDescriptor

var file_model_model_v1alpha_task_instance_segmentation_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9,
	0x01, 0x0a, 0x1a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x15, 0x0a,
	0x03, 0x72, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x03, 0x72, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x48, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x22, 0x67, 0x0a, 0x19, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x1f, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x26, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x12, 0x1d,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a,
	0x1a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x4e, 0x0a, 0x07, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x07, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x42, 0xea, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x1d, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x02, 0x03, 0x4d,
	0x4d, 0x58, 0xaa, 0x02, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02,
	0x1f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_model_v1alpha_task_instance_segmentation_proto_rawDescOnce sync.Once
	file_model_model_v1alpha_task_instance_segmentation_proto_rawDescData = file_model_model_v1alpha_task_instance_segmentation_proto_rawDesc
)

func file_model_model_v1alpha_task_instance_segmentation_proto_rawDescGZIP() []byte {
	file_model_model_v1alpha_task_instance_segmentation_proto_rawDescOnce.Do(func() {
		file_model_model_v1alpha_task_instance_segmentation_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_model_v1alpha_task_instance_segmentation_proto_rawDescData)
	})
	return file_model_model_v1alpha_task_instance_segmentation_proto_rawDescData
}

var file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_model_v1alpha_task_instance_segmentation_proto_goTypes = []interface{}{
	(*InstanceSegmentationObject)(nil),      // 0: model.model.v1alpha.InstanceSegmentationObject
	(*InstanceSegmentationInput)(nil),       // 1: model.model.v1alpha.InstanceSegmentationInput
	(*InstanceSegmentationInputStream)(nil), // 2: model.model.v1alpha.InstanceSegmentationInputStream
	(*InstanceSegmentationOutput)(nil),      // 3: model.model.v1alpha.InstanceSegmentationOutput
	(*BoundingBox)(nil),                     // 4: model.model.v1alpha.BoundingBox
}
var file_model_model_v1alpha_task_instance_segmentation_proto_depIdxs = []int32{
	4, // 0: model.model.v1alpha.InstanceSegmentationObject.bounding_box:type_name -> model.model.v1alpha.BoundingBox
	0, // 1: model.model.v1alpha.InstanceSegmentationOutput.objects:type_name -> model.model.v1alpha.InstanceSegmentationObject
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_model_v1alpha_task_instance_segmentation_proto_init() }
func file_model_model_v1alpha_task_instance_segmentation_proto_init() {
	if File_model_model_v1alpha_task_instance_segmentation_proto != nil {
		return
	}
	file_model_model_v1alpha_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceSegmentationObject); i {
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
		file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceSegmentationInput); i {
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
		file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceSegmentationInputStream); i {
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
		file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceSegmentationOutput); i {
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
	file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InstanceSegmentationInput_ImageUrl)(nil),
		(*InstanceSegmentationInput_ImageBase64)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_model_v1alpha_task_instance_segmentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_v1alpha_task_instance_segmentation_proto_goTypes,
		DependencyIndexes: file_model_model_v1alpha_task_instance_segmentation_proto_depIdxs,
		MessageInfos:      file_model_model_v1alpha_task_instance_segmentation_proto_msgTypes,
	}.Build()
	File_model_model_v1alpha_task_instance_segmentation_proto = out.File
	file_model_model_v1alpha_task_instance_segmentation_proto_rawDesc = nil
	file_model_model_v1alpha_task_instance_segmentation_proto_goTypes = nil
	file_model_model_v1alpha_task_instance_segmentation_proto_depIdxs = nil
}