// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: artifact/artifact/v1alpha/artifact_public_service.proto

package artifactv1alpha

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ArtifactPublicService_Liveness_FullMethodName                  = "/artifact.artifact.v1alpha.ArtifactPublicService/Liveness"
	ArtifactPublicService_Readiness_FullMethodName                 = "/artifact.artifact.v1alpha.ArtifactPublicService/Readiness"
	ArtifactPublicService_CreateKnowledgeBase_FullMethodName       = "/artifact.artifact.v1alpha.ArtifactPublicService/CreateKnowledgeBase"
	ArtifactPublicService_ListKnowledgeBases_FullMethodName        = "/artifact.artifact.v1alpha.ArtifactPublicService/ListKnowledgeBases"
	ArtifactPublicService_UpdateKnowledgeBase_FullMethodName       = "/artifact.artifact.v1alpha.ArtifactPublicService/UpdateKnowledgeBase"
	ArtifactPublicService_DeleteKnowledgeBase_FullMethodName       = "/artifact.artifact.v1alpha.ArtifactPublicService/DeleteKnowledgeBase"
	ArtifactPublicService_UploadKnowledgeBaseFile_FullMethodName   = "/artifact.artifact.v1alpha.ArtifactPublicService/UploadKnowledgeBaseFile"
	ArtifactPublicService_DeleteKnowledgeBaseFile_FullMethodName   = "/artifact.artifact.v1alpha.ArtifactPublicService/DeleteKnowledgeBaseFile"
	ArtifactPublicService_ProcessKnowledgeBaseFiles_FullMethodName = "/artifact.artifact.v1alpha.ArtifactPublicService/ProcessKnowledgeBaseFiles"
	ArtifactPublicService_ListKnowledgeBaseFiles_FullMethodName    = "/artifact.artifact.v1alpha.ArtifactPublicService/ListKnowledgeBaseFiles"
)

// ArtifactPublicServiceClient is the client API for ArtifactPublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArtifactPublicServiceClient interface {
	// Check if the artifact server is alive
	//
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md.
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// Check if the artifact server is ready
	//
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// Create a knowledge base
	CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*CreateKnowledgeBaseResponse, error)
	// Get all knowledge bases info
	ListKnowledgeBases(ctx context.Context, in *ListKnowledgeBasesRequest, opts ...grpc.CallOption) (*ListKnowledgeBasesResponse, error)
	// Update a knowledge base info
	UpdateKnowledgeBase(ctx context.Context, in *UpdateKnowledgeBaseRequest, opts ...grpc.CallOption) (*UpdateKnowledgeBaseResponse, error)
	// Delete a knowledge base
	DeleteKnowledgeBase(ctx context.Context, in *DeleteKnowledgeBaseRequest, opts ...grpc.CallOption) (*DeleteKnowledgeBaseResponse, error)
	// Create a file
	UploadKnowledgeBaseFile(ctx context.Context, in *UploadKnowledgeBaseFileRequest, opts ...grpc.CallOption) (*UploadKnowledgeBaseFileResponse, error)
	// Delete a file
	DeleteKnowledgeBaseFile(ctx context.Context, in *DeleteKnowledgeBaseFileRequest, opts ...grpc.CallOption) (*DeleteKnowledgeBaseFileResponse, error)
	// process file
	ProcessKnowledgeBaseFiles(ctx context.Context, in *ProcessKnowledgeBaseFilesRequest, opts ...grpc.CallOption) (*ProcessKnowledgeBaseFilesResponse, error)
	// list files
	ListKnowledgeBaseFiles(ctx context.Context, in *ListKnowledgeBaseFilesRequest, opts ...grpc.CallOption) (*ListKnowledgeBaseFilesResponse, error)
}

type artifactPublicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactPublicServiceClient(cc grpc.ClientConnInterface) ArtifactPublicServiceClient {
	return &artifactPublicServiceClient{cc}
}

func (c *artifactPublicServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_Liveness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_Readiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*CreateKnowledgeBaseResponse, error) {
	out := new(CreateKnowledgeBaseResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_CreateKnowledgeBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) ListKnowledgeBases(ctx context.Context, in *ListKnowledgeBasesRequest, opts ...grpc.CallOption) (*ListKnowledgeBasesResponse, error) {
	out := new(ListKnowledgeBasesResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_ListKnowledgeBases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) UpdateKnowledgeBase(ctx context.Context, in *UpdateKnowledgeBaseRequest, opts ...grpc.CallOption) (*UpdateKnowledgeBaseResponse, error) {
	out := new(UpdateKnowledgeBaseResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_UpdateKnowledgeBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) DeleteKnowledgeBase(ctx context.Context, in *DeleteKnowledgeBaseRequest, opts ...grpc.CallOption) (*DeleteKnowledgeBaseResponse, error) {
	out := new(DeleteKnowledgeBaseResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_DeleteKnowledgeBase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) UploadKnowledgeBaseFile(ctx context.Context, in *UploadKnowledgeBaseFileRequest, opts ...grpc.CallOption) (*UploadKnowledgeBaseFileResponse, error) {
	out := new(UploadKnowledgeBaseFileResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_UploadKnowledgeBaseFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) DeleteKnowledgeBaseFile(ctx context.Context, in *DeleteKnowledgeBaseFileRequest, opts ...grpc.CallOption) (*DeleteKnowledgeBaseFileResponse, error) {
	out := new(DeleteKnowledgeBaseFileResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_DeleteKnowledgeBaseFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) ProcessKnowledgeBaseFiles(ctx context.Context, in *ProcessKnowledgeBaseFilesRequest, opts ...grpc.CallOption) (*ProcessKnowledgeBaseFilesResponse, error) {
	out := new(ProcessKnowledgeBaseFilesResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_ProcessKnowledgeBaseFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactPublicServiceClient) ListKnowledgeBaseFiles(ctx context.Context, in *ListKnowledgeBaseFilesRequest, opts ...grpc.CallOption) (*ListKnowledgeBaseFilesResponse, error) {
	out := new(ListKnowledgeBaseFilesResponse)
	err := c.cc.Invoke(ctx, ArtifactPublicService_ListKnowledgeBaseFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactPublicServiceServer is the server API for ArtifactPublicService service.
// All implementations should embed UnimplementedArtifactPublicServiceServer
// for forward compatibility
type ArtifactPublicServiceServer interface {
	// Check if the artifact server is alive
	//
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md.
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// Check if the artifact server is ready
	//
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// Create a knowledge base
	CreateKnowledgeBase(context.Context, *CreateKnowledgeBaseRequest) (*CreateKnowledgeBaseResponse, error)
	// Get all knowledge bases info
	ListKnowledgeBases(context.Context, *ListKnowledgeBasesRequest) (*ListKnowledgeBasesResponse, error)
	// Update a knowledge base info
	UpdateKnowledgeBase(context.Context, *UpdateKnowledgeBaseRequest) (*UpdateKnowledgeBaseResponse, error)
	// Delete a knowledge base
	DeleteKnowledgeBase(context.Context, *DeleteKnowledgeBaseRequest) (*DeleteKnowledgeBaseResponse, error)
	// Create a file
	UploadKnowledgeBaseFile(context.Context, *UploadKnowledgeBaseFileRequest) (*UploadKnowledgeBaseFileResponse, error)
	// Delete a file
	DeleteKnowledgeBaseFile(context.Context, *DeleteKnowledgeBaseFileRequest) (*DeleteKnowledgeBaseFileResponse, error)
	// process file
	ProcessKnowledgeBaseFiles(context.Context, *ProcessKnowledgeBaseFilesRequest) (*ProcessKnowledgeBaseFilesResponse, error)
	// list files
	ListKnowledgeBaseFiles(context.Context, *ListKnowledgeBaseFilesRequest) (*ListKnowledgeBaseFilesResponse, error)
}

// UnimplementedArtifactPublicServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArtifactPublicServiceServer struct {
}

func (UnimplementedArtifactPublicServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedArtifactPublicServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (UnimplementedArtifactPublicServiceServer) CreateKnowledgeBase(context.Context, *CreateKnowledgeBaseRequest) (*CreateKnowledgeBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKnowledgeBase not implemented")
}
func (UnimplementedArtifactPublicServiceServer) ListKnowledgeBases(context.Context, *ListKnowledgeBasesRequest) (*ListKnowledgeBasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKnowledgeBases not implemented")
}
func (UnimplementedArtifactPublicServiceServer) UpdateKnowledgeBase(context.Context, *UpdateKnowledgeBaseRequest) (*UpdateKnowledgeBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKnowledgeBase not implemented")
}
func (UnimplementedArtifactPublicServiceServer) DeleteKnowledgeBase(context.Context, *DeleteKnowledgeBaseRequest) (*DeleteKnowledgeBaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKnowledgeBase not implemented")
}
func (UnimplementedArtifactPublicServiceServer) UploadKnowledgeBaseFile(context.Context, *UploadKnowledgeBaseFileRequest) (*UploadKnowledgeBaseFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadKnowledgeBaseFile not implemented")
}
func (UnimplementedArtifactPublicServiceServer) DeleteKnowledgeBaseFile(context.Context, *DeleteKnowledgeBaseFileRequest) (*DeleteKnowledgeBaseFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKnowledgeBaseFile not implemented")
}
func (UnimplementedArtifactPublicServiceServer) ProcessKnowledgeBaseFiles(context.Context, *ProcessKnowledgeBaseFilesRequest) (*ProcessKnowledgeBaseFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessKnowledgeBaseFiles not implemented")
}
func (UnimplementedArtifactPublicServiceServer) ListKnowledgeBaseFiles(context.Context, *ListKnowledgeBaseFilesRequest) (*ListKnowledgeBaseFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKnowledgeBaseFiles not implemented")
}

// UnsafeArtifactPublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactPublicServiceServer will
// result in compilation errors.
type UnsafeArtifactPublicServiceServer interface {
	mustEmbedUnimplementedArtifactPublicServiceServer()
}

func RegisterArtifactPublicServiceServer(s grpc.ServiceRegistrar, srv ArtifactPublicServiceServer) {
	s.RegisterService(&ArtifactPublicService_ServiceDesc, srv)
}

func _ArtifactPublicService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_Liveness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_Readiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_CreateKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).CreateKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_CreateKnowledgeBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).CreateKnowledgeBase(ctx, req.(*CreateKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_ListKnowledgeBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKnowledgeBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).ListKnowledgeBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_ListKnowledgeBases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).ListKnowledgeBases(ctx, req.(*ListKnowledgeBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_UpdateKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).UpdateKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_UpdateKnowledgeBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).UpdateKnowledgeBase(ctx, req.(*UpdateKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_DeleteKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).DeleteKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_DeleteKnowledgeBase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).DeleteKnowledgeBase(ctx, req.(*DeleteKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_UploadKnowledgeBaseFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadKnowledgeBaseFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).UploadKnowledgeBaseFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_UploadKnowledgeBaseFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).UploadKnowledgeBaseFile(ctx, req.(*UploadKnowledgeBaseFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_DeleteKnowledgeBaseFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKnowledgeBaseFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).DeleteKnowledgeBaseFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_DeleteKnowledgeBaseFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).DeleteKnowledgeBaseFile(ctx, req.(*DeleteKnowledgeBaseFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_ProcessKnowledgeBaseFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessKnowledgeBaseFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).ProcessKnowledgeBaseFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_ProcessKnowledgeBaseFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).ProcessKnowledgeBaseFiles(ctx, req.(*ProcessKnowledgeBaseFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactPublicService_ListKnowledgeBaseFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKnowledgeBaseFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactPublicServiceServer).ListKnowledgeBaseFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactPublicService_ListKnowledgeBaseFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactPublicServiceServer).ListKnowledgeBaseFiles(ctx, req.(*ListKnowledgeBaseFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactPublicService_ServiceDesc is the grpc.ServiceDesc for ArtifactPublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactPublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "artifact.artifact.v1alpha.ArtifactPublicService",
	HandlerType: (*ArtifactPublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _ArtifactPublicService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _ArtifactPublicService_Readiness_Handler,
		},
		{
			MethodName: "CreateKnowledgeBase",
			Handler:    _ArtifactPublicService_CreateKnowledgeBase_Handler,
		},
		{
			MethodName: "ListKnowledgeBases",
			Handler:    _ArtifactPublicService_ListKnowledgeBases_Handler,
		},
		{
			MethodName: "UpdateKnowledgeBase",
			Handler:    _ArtifactPublicService_UpdateKnowledgeBase_Handler,
		},
		{
			MethodName: "DeleteKnowledgeBase",
			Handler:    _ArtifactPublicService_DeleteKnowledgeBase_Handler,
		},
		{
			MethodName: "UploadKnowledgeBaseFile",
			Handler:    _ArtifactPublicService_UploadKnowledgeBaseFile_Handler,
		},
		{
			MethodName: "DeleteKnowledgeBaseFile",
			Handler:    _ArtifactPublicService_DeleteKnowledgeBaseFile_Handler,
		},
		{
			MethodName: "ProcessKnowledgeBaseFiles",
			Handler:    _ArtifactPublicService_ProcessKnowledgeBaseFiles_Handler,
		},
		{
			MethodName: "ListKnowledgeBaseFiles",
			Handler:    _ArtifactPublicService_ListKnowledgeBaseFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "artifact/artifact/v1alpha/artifact_public_service.proto",
}
