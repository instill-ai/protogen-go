// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vdp/model/v1alpha/model_definition.proto

package modelv1alpha

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// View represents a view of any resource. The resource view is implemented by
// adding a parameter to the method request which allows the client to specify
// which view of the resource it wants to receive in the response.
type View int32

const (
	// View: UNSPECIFIED, equivalent to BASIC.
	View_VIEW_UNSPECIFIED View = 0
	// View: BASIC, server response only include basic information of the resource
	View_VIEW_BASIC View = 1
	// View: FULL, full representation of the resource
	View_VIEW_FULL View = 2
)

// Enum value maps for View.
var (
	View_name = map[int32]string{
		0: "VIEW_UNSPECIFIED",
		1: "VIEW_BASIC",
		2: "VIEW_FULL",
	}
	View_value = map[string]int32{
		"VIEW_UNSPECIFIED": 0,
		"VIEW_BASIC":       1,
		"VIEW_FULL":        2,
	}
)

func (x View) Enum() *View {
	p := new(View)
	*p = x
	return p
}

func (x View) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (View) Descriptor() protoreflect.EnumDescriptor {
	return file_vdp_model_v1alpha_model_definition_proto_enumTypes[0].Descriptor()
}

func (View) Type() protoreflect.EnumType {
	return &file_vdp_model_v1alpha_model_definition_proto_enumTypes[0]
}

func (x View) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use View.Descriptor instead.
func (View) EnumDescriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP(), []int{0}
}

///////////////////////////////////////////////////////////////////
// ModelDefinition represents the definition of a model
type ModelDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ModelDefinition resource name. It must have the format of
	// "model-definitions/{model-definition}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ModelDefinition ID in UUIDv4
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// ModelDefinition resource ID (the last segment of the resource name) used to
	// construct the resource name.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// ModelDefinition display official title
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// ModelDefinition documentation url
	DocumentationUrl string `protobuf:"bytes,5,opt,name=documentation_url,json=documentationUrl,proto3" json:"documentation_url,omitempty"`
	// ModelDefinition icon
	Icon string `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	// ModelDefinition model specification represents the JSON schema used to
	// validate the JSON configurations of a model created from a specific model
	// source. Must be a valid JSON that includes what fields are needed to
	// create/display a model.
	ModelSpec *structpb.Struct `protobuf:"bytes,7,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// ModelDefinition model instance specification represents the JSON schema
	// used to validate the JSON configurations of a model instance from a
	// specific model source. Must be a valid JSON that includes what fields are
	// needed to display a model instance.
	ModelInstanceSpec *structpb.Struct `protobuf:"bytes,8,opt,name=model_instance_spec,json=modelInstanceSpec,proto3" json:"model_instance_spec,omitempty"`
	// ModelDefinition create time
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// ModelDefinition update time
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *ModelDefinition) Reset() {
	*x = ModelDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelDefinition) ProtoMessage() {}

func (x *ModelDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelDefinition.ProtoReflect.Descriptor instead.
func (*ModelDefinition) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP(), []int{0}
}

func (x *ModelDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelDefinition) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ModelDefinition) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModelDefinition) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ModelDefinition) GetDocumentationUrl() string {
	if x != nil {
		return x.DocumentationUrl
	}
	return ""
}

func (x *ModelDefinition) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *ModelDefinition) GetModelSpec() *structpb.Struct {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

func (x *ModelDefinition) GetModelInstanceSpec() *structpb.Struct {
	if x != nil {
		return x.ModelInstanceSpec
	}
	return nil
}

func (x *ModelDefinition) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ModelDefinition) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// ListModelDefinitionRequest represents a request to list all supported model
// definitions
type ListModelDefinitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page size: the maximum number of resources to return. The service may
	// return fewer than this value. If unspecified, at most 10 ModelDefinitions
	// will be returned. The maximum value is 100; values above 100 will be
	// coereced to 100.
	PageSize *int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"`
	// Page token
	PageToken *string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3,oneof" json:"page_token,omitempty"`
	// Definition view (default is VIEW_BASIC)
	// VIEW_UNSPECIFIED/VIEW_BASIC: omit `ModelDefinition.model_spec` and
	// `ModelDefinition.model_instance_spec`
	// VIEW_FULL: show full information
	View *View `protobuf:"varint,3,opt,name=view,proto3,enum=vdp.model.v1alpha.View,oneof" json:"view,omitempty"`
}

func (x *ListModelDefinitionRequest) Reset() {
	*x = ListModelDefinitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelDefinitionRequest) ProtoMessage() {}

func (x *ListModelDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelDefinitionRequest.ProtoReflect.Descriptor instead.
func (*ListModelDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP(), []int{1}
}

func (x *ListModelDefinitionRequest) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *ListModelDefinitionRequest) GetPageToken() string {
	if x != nil && x.PageToken != nil {
		return *x.PageToken
	}
	return ""
}

func (x *ListModelDefinitionRequest) GetView() View {
	if x != nil && x.View != nil {
		return *x.View
	}
	return View_VIEW_UNSPECIFIED
}

// ListModelDefinitionResponse represents a response to list all supported model
// definitions
type ListModelDefinitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a list of ModelDefinition instances
	ModelDefinitions []*ModelDefinition `protobuf:"bytes,1,rep,name=model_definitions,json=modelDefinitions,proto3" json:"model_definitions,omitempty"`
	// Next page token
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Total count of model definitions
	TotalSize int64 `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListModelDefinitionResponse) Reset() {
	*x = ListModelDefinitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelDefinitionResponse) ProtoMessage() {}

func (x *ListModelDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelDefinitionResponse.ProtoReflect.Descriptor instead.
func (*ListModelDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP(), []int{2}
}

func (x *ListModelDefinitionResponse) GetModelDefinitions() []*ModelDefinition {
	if x != nil {
		return x.ModelDefinitions
	}
	return nil
}

func (x *ListModelDefinitionResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListModelDefinitionResponse) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

// GetModelDefinitionRequest represents a request to query a model definition
type GetModelDefinitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name of the model definition.
	// For example "model-definitions/{uuid}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Definition view (default is VIEW_BASIC)
	// VIEW_UNSPECIFIED/VIEW_BASIC: omit `ModelDefinition.model_spec` and
	// `ModelDefinition.model_instance_spec`
	// VIEW_FULL: show full information
	View *View `protobuf:"varint,2,opt,name=view,proto3,enum=vdp.model.v1alpha.View,oneof" json:"view,omitempty"`
}

func (x *GetModelDefinitionRequest) Reset() {
	*x = GetModelDefinitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelDefinitionRequest) ProtoMessage() {}

func (x *GetModelDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelDefinitionRequest.ProtoReflect.Descriptor instead.
func (*GetModelDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP(), []int{3}
}

func (x *GetModelDefinitionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetModelDefinitionRequest) GetView() View {
	if x != nil && x.View != nil {
		return *x.View
	}
	return View_VIEW_UNSPECIFIED
}

// GetModelDefinitionResponse represents a response for a model definition
type GetModelDefinitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A model definition instance
	ModelDefinition *ModelDefinition `protobuf:"bytes,1,opt,name=model_definition,json=modelDefinition,proto3" json:"model_definition,omitempty"`
}

func (x *GetModelDefinitionResponse) Reset() {
	*x = GetModelDefinitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelDefinitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelDefinitionResponse) ProtoMessage() {}

func (x *GetModelDefinitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vdp_model_v1alpha_model_definition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelDefinitionResponse.ProtoReflect.Descriptor instead.
func (*GetModelDefinitionResponse) Descriptor() ([]byte, []int) {
	return file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP(), []int{4}
}

func (x *GetModelDefinitionResponse) GetModelDefinition() *ModelDefinition {
	if x != nil {
		return x.ModelDefinition
	}
	return nil
}

var File_vdp_model_v1alpha_model_definition_proto protoreflect.FileDescriptor

var file_vdp_model_v1alpha_model_definition_proto_rawDesc = []byte{
	0x0a, 0x28, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x76, 0x64, 0x70, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x04, 0x0a, 0x0f, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x11,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x17,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x4c, 0x0a, 0x13, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x4b, 0xea, 0x41, 0x48, 0x0a, 0x20, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x22, 0xc9, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x48, 0x01, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x35, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x02,
	0x52, 0x04, 0x76, 0x69, 0x65, 0x77, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x22,
	0xb5, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x64, 0x70,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x46, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c, 0x6c, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x92, 0x41, 0x1b,
	0xca, 0x3e, 0x18, 0xfa, 0x02, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x6b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x64, 0x70, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x3b,
	0x0a, 0x04, 0x56, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x56, 0x49, 0x45, 0x57, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x43, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x56, 0x49, 0x45, 0x57, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0xd5, 0x01, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x64, 0x70, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x14, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x6c,
	0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x76, 0x64, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2,
	0x02, 0x03, 0x56, 0x4d, 0x58, 0xaa, 0x02, 0x11, 0x56, 0x64, 0x70, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x11, 0x56, 0x64, 0x70, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1d,
	0x56, 0x64, 0x70, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13,
	0x56, 0x64, 0x70, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vdp_model_v1alpha_model_definition_proto_rawDescOnce sync.Once
	file_vdp_model_v1alpha_model_definition_proto_rawDescData = file_vdp_model_v1alpha_model_definition_proto_rawDesc
)

func file_vdp_model_v1alpha_model_definition_proto_rawDescGZIP() []byte {
	file_vdp_model_v1alpha_model_definition_proto_rawDescOnce.Do(func() {
		file_vdp_model_v1alpha_model_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_vdp_model_v1alpha_model_definition_proto_rawDescData)
	})
	return file_vdp_model_v1alpha_model_definition_proto_rawDescData
}

var file_vdp_model_v1alpha_model_definition_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vdp_model_v1alpha_model_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vdp_model_v1alpha_model_definition_proto_goTypes = []interface{}{
	(View)(0),                           // 0: vdp.model.v1alpha.View
	(*ModelDefinition)(nil),             // 1: vdp.model.v1alpha.ModelDefinition
	(*ListModelDefinitionRequest)(nil),  // 2: vdp.model.v1alpha.ListModelDefinitionRequest
	(*ListModelDefinitionResponse)(nil), // 3: vdp.model.v1alpha.ListModelDefinitionResponse
	(*GetModelDefinitionRequest)(nil),   // 4: vdp.model.v1alpha.GetModelDefinitionRequest
	(*GetModelDefinitionResponse)(nil),  // 5: vdp.model.v1alpha.GetModelDefinitionResponse
	(*structpb.Struct)(nil),             // 6: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),       // 7: google.protobuf.Timestamp
}
var file_vdp_model_v1alpha_model_definition_proto_depIdxs = []int32{
	6, // 0: vdp.model.v1alpha.ModelDefinition.model_spec:type_name -> google.protobuf.Struct
	6, // 1: vdp.model.v1alpha.ModelDefinition.model_instance_spec:type_name -> google.protobuf.Struct
	7, // 2: vdp.model.v1alpha.ModelDefinition.create_time:type_name -> google.protobuf.Timestamp
	7, // 3: vdp.model.v1alpha.ModelDefinition.update_time:type_name -> google.protobuf.Timestamp
	0, // 4: vdp.model.v1alpha.ListModelDefinitionRequest.view:type_name -> vdp.model.v1alpha.View
	1, // 5: vdp.model.v1alpha.ListModelDefinitionResponse.model_definitions:type_name -> vdp.model.v1alpha.ModelDefinition
	0, // 6: vdp.model.v1alpha.GetModelDefinitionRequest.view:type_name -> vdp.model.v1alpha.View
	1, // 7: vdp.model.v1alpha.GetModelDefinitionResponse.model_definition:type_name -> vdp.model.v1alpha.ModelDefinition
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_vdp_model_v1alpha_model_definition_proto_init() }
func file_vdp_model_v1alpha_model_definition_proto_init() {
	if File_vdp_model_v1alpha_model_definition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vdp_model_v1alpha_model_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelDefinition); i {
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
		file_vdp_model_v1alpha_model_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelDefinitionRequest); i {
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
		file_vdp_model_v1alpha_model_definition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelDefinitionResponse); i {
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
		file_vdp_model_v1alpha_model_definition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelDefinitionRequest); i {
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
		file_vdp_model_v1alpha_model_definition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelDefinitionResponse); i {
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
	file_vdp_model_v1alpha_model_definition_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_vdp_model_v1alpha_model_definition_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vdp_model_v1alpha_model_definition_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vdp_model_v1alpha_model_definition_proto_goTypes,
		DependencyIndexes: file_vdp_model_v1alpha_model_definition_proto_depIdxs,
		EnumInfos:         file_vdp_model_v1alpha_model_definition_proto_enumTypes,
		MessageInfos:      file_vdp_model_v1alpha_model_definition_proto_msgTypes,
	}.Build()
	File_vdp_model_v1alpha_model_definition_proto = out.File
	file_vdp_model_v1alpha_model_definition_proto_rawDesc = nil
	file_vdp_model_v1alpha_model_definition_proto_goTypes = nil
	file_vdp_model_v1alpha_model_definition_proto_depIdxs = nil
}