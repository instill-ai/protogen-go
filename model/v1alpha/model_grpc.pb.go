// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package modelv1alpha

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

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServiceClient interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// CreateModelByGitHub method receives a
	// CreateModelByGitHubRequest message and returns a
	// CreateModelByGitHubResponse message.
	CreateModelByGitHub(ctx context.Context, in *CreateModelByGitHubRequest, opts ...grpc.CallOption) (*CreateModelByGitHubResponse, error)
	// CreateModelBinaryFileUpload method receives a
	// CreateModelBinaryFileUploadRequest message and returns a
	// CreateModelBinaryFileUploadResponse message.
	CreateModelBinaryFileUpload(ctx context.Context, opts ...grpc.CallOption) (ModelService_CreateModelBinaryFileUploadClient, error)
	// CreateModel method receives a CreateModelRequest message and returns
	// a CreateModelResponse message.
	ListModel(ctx context.Context, in *ListModelRequest, opts ...grpc.CallOption) (*ListModelResponse, error)
	// GetModel method receives a GetModelRequest message and returns
	// a GetModelResponse message.
	GetModel(ctx context.Context, in *GetModelRequest, opts ...grpc.CallOption) (*GetModelResponse, error)
	// UpdateModelInstance method receives a UpdateModelInstanceRequest message and
	// returns a UpdateModelInstanceResponse message.
	UpdateModelInstance(ctx context.Context, in *UpdateModelInstanceRequest, opts ...grpc.CallOption) (*UpdateModelInstanceResponse, error)
	// DeleteModel method receives a DeleteModelRequest message and returns
	// a DeleteModelResponse message.
	DeleteModel(ctx context.Context, in *DeleteModelRequest, opts ...grpc.CallOption) (*DeleteModelResponse, error)
	// DeleteModelInstance method receives a DeleteModelInstanceRequest message and
	// returns a DeleteModelInstanceResponse message.
	DeleteModelInstance(ctx context.Context, in *DeleteModelInstanceRequest, opts ...grpc.CallOption) (*DeleteModelInstanceResponse, error)
	// TriggerModel method receives a TriggerModelRequest message and
	// returns a TriggerModelResponse message.
	TriggerModel(ctx context.Context, in *TriggerModelRequest, opts ...grpc.CallOption) (*TriggerModelResponse, error)
	// TriggerModelBinaryFileUpload method receives a
	// TriggerModelBinaryFileUploadRequest message and returns a
	// TriggerModelBinaryFileUploadResponse message.
	TriggerModelBinaryFileUpload(ctx context.Context, opts ...grpc.CallOption) (ModelService_TriggerModelBinaryFileUploadClient, error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/Readiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) CreateModelByGitHub(ctx context.Context, in *CreateModelByGitHubRequest, opts ...grpc.CallOption) (*CreateModelByGitHubResponse, error) {
	out := new(CreateModelByGitHubResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/CreateModelByGitHub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) CreateModelBinaryFileUpload(ctx context.Context, opts ...grpc.CallOption) (ModelService_CreateModelBinaryFileUploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ModelService_ServiceDesc.Streams[0], "/instill.model.v1alpha.ModelService/CreateModelBinaryFileUpload", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelServiceCreateModelBinaryFileUploadClient{stream}
	return x, nil
}

type ModelService_CreateModelBinaryFileUploadClient interface {
	Send(*CreateModelBinaryFileUploadRequest) error
	CloseAndRecv() (*CreateModelBinaryFileUploadResponse, error)
	grpc.ClientStream
}

type modelServiceCreateModelBinaryFileUploadClient struct {
	grpc.ClientStream
}

func (x *modelServiceCreateModelBinaryFileUploadClient) Send(m *CreateModelBinaryFileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *modelServiceCreateModelBinaryFileUploadClient) CloseAndRecv() (*CreateModelBinaryFileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateModelBinaryFileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *modelServiceClient) ListModel(ctx context.Context, in *ListModelRequest, opts ...grpc.CallOption) (*ListModelResponse, error) {
	out := new(ListModelResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/ListModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) GetModel(ctx context.Context, in *GetModelRequest, opts ...grpc.CallOption) (*GetModelResponse, error) {
	out := new(GetModelResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/GetModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) UpdateModelInstance(ctx context.Context, in *UpdateModelInstanceRequest, opts ...grpc.CallOption) (*UpdateModelInstanceResponse, error) {
	out := new(UpdateModelInstanceResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/UpdateModelInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) DeleteModel(ctx context.Context, in *DeleteModelRequest, opts ...grpc.CallOption) (*DeleteModelResponse, error) {
	out := new(DeleteModelResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/DeleteModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) DeleteModelInstance(ctx context.Context, in *DeleteModelInstanceRequest, opts ...grpc.CallOption) (*DeleteModelInstanceResponse, error) {
	out := new(DeleteModelInstanceResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/DeleteModelInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) TriggerModel(ctx context.Context, in *TriggerModelRequest, opts ...grpc.CallOption) (*TriggerModelResponse, error) {
	out := new(TriggerModelResponse)
	err := c.cc.Invoke(ctx, "/instill.model.v1alpha.ModelService/TriggerModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) TriggerModelBinaryFileUpload(ctx context.Context, opts ...grpc.CallOption) (ModelService_TriggerModelBinaryFileUploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ModelService_ServiceDesc.Streams[1], "/instill.model.v1alpha.ModelService/TriggerModelBinaryFileUpload", opts...)
	if err != nil {
		return nil, err
	}
	x := &modelServiceTriggerModelBinaryFileUploadClient{stream}
	return x, nil
}

type ModelService_TriggerModelBinaryFileUploadClient interface {
	Send(*TriggerModelBinaryFileUploadRequest) error
	CloseAndRecv() (*TriggerModelBinaryFileUploadResponse, error)
	grpc.ClientStream
}

type modelServiceTriggerModelBinaryFileUploadClient struct {
	grpc.ClientStream
}

func (x *modelServiceTriggerModelBinaryFileUploadClient) Send(m *TriggerModelBinaryFileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *modelServiceTriggerModelBinaryFileUploadClient) CloseAndRecv() (*TriggerModelBinaryFileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TriggerModelBinaryFileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ModelServiceServer is the server API for ModelService service.
// All implementations should embed UnimplementedModelServiceServer
// for forward compatibility
type ModelServiceServer interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// CreateModelByGitHub method receives a
	// CreateModelByGitHubRequest message and returns a
	// CreateModelByGitHubResponse message.
	CreateModelByGitHub(context.Context, *CreateModelByGitHubRequest) (*CreateModelByGitHubResponse, error)
	// CreateModelBinaryFileUpload method receives a
	// CreateModelBinaryFileUploadRequest message and returns a
	// CreateModelBinaryFileUploadResponse message.
	CreateModelBinaryFileUpload(ModelService_CreateModelBinaryFileUploadServer) error
	// CreateModel method receives a CreateModelRequest message and returns
	// a CreateModelResponse message.
	ListModel(context.Context, *ListModelRequest) (*ListModelResponse, error)
	// GetModel method receives a GetModelRequest message and returns
	// a GetModelResponse message.
	GetModel(context.Context, *GetModelRequest) (*GetModelResponse, error)
	// UpdateModelInstance method receives a UpdateModelInstanceRequest message and
	// returns a UpdateModelInstanceResponse message.
	UpdateModelInstance(context.Context, *UpdateModelInstanceRequest) (*UpdateModelInstanceResponse, error)
	// DeleteModel method receives a DeleteModelRequest message and returns
	// a DeleteModelResponse message.
	DeleteModel(context.Context, *DeleteModelRequest) (*DeleteModelResponse, error)
	// DeleteModelInstance method receives a DeleteModelInstanceRequest message and
	// returns a DeleteModelInstanceResponse message.
	DeleteModelInstance(context.Context, *DeleteModelInstanceRequest) (*DeleteModelInstanceResponse, error)
	// TriggerModel method receives a TriggerModelRequest message and
	// returns a TriggerModelResponse message.
	TriggerModel(context.Context, *TriggerModelRequest) (*TriggerModelResponse, error)
	// TriggerModelBinaryFileUpload method receives a
	// TriggerModelBinaryFileUploadRequest message and returns a
	// TriggerModelBinaryFileUploadResponse message.
	TriggerModelBinaryFileUpload(ModelService_TriggerModelBinaryFileUploadServer) error
}

// UnimplementedModelServiceServer should be embedded to have forward compatible implementations.
type UnimplementedModelServiceServer struct {
}

func (UnimplementedModelServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedModelServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (UnimplementedModelServiceServer) CreateModelByGitHub(context.Context, *CreateModelByGitHubRequest) (*CreateModelByGitHubResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModelByGitHub not implemented")
}
func (UnimplementedModelServiceServer) CreateModelBinaryFileUpload(ModelService_CreateModelBinaryFileUploadServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateModelBinaryFileUpload not implemented")
}
func (UnimplementedModelServiceServer) ListModel(context.Context, *ListModelRequest) (*ListModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModel not implemented")
}
func (UnimplementedModelServiceServer) GetModel(context.Context, *GetModelRequest) (*GetModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModel not implemented")
}
func (UnimplementedModelServiceServer) UpdateModelInstance(context.Context, *UpdateModelInstanceRequest) (*UpdateModelInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModelInstance not implemented")
}
func (UnimplementedModelServiceServer) DeleteModel(context.Context, *DeleteModelRequest) (*DeleteModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModel not implemented")
}
func (UnimplementedModelServiceServer) DeleteModelInstance(context.Context, *DeleteModelInstanceRequest) (*DeleteModelInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModelInstance not implemented")
}
func (UnimplementedModelServiceServer) TriggerModel(context.Context, *TriggerModelRequest) (*TriggerModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerModel not implemented")
}
func (UnimplementedModelServiceServer) TriggerModelBinaryFileUpload(ModelService_TriggerModelBinaryFileUploadServer) error {
	return status.Errorf(codes.Unimplemented, "method TriggerModelBinaryFileUpload not implemented")
}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/Readiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_CreateModelByGitHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelByGitHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).CreateModelByGitHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/CreateModelByGitHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).CreateModelByGitHub(ctx, req.(*CreateModelByGitHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_CreateModelBinaryFileUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelServiceServer).CreateModelBinaryFileUpload(&modelServiceCreateModelBinaryFileUploadServer{stream})
}

type ModelService_CreateModelBinaryFileUploadServer interface {
	SendAndClose(*CreateModelBinaryFileUploadResponse) error
	Recv() (*CreateModelBinaryFileUploadRequest, error)
	grpc.ServerStream
}

type modelServiceCreateModelBinaryFileUploadServer struct {
	grpc.ServerStream
}

func (x *modelServiceCreateModelBinaryFileUploadServer) SendAndClose(m *CreateModelBinaryFileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *modelServiceCreateModelBinaryFileUploadServer) Recv() (*CreateModelBinaryFileUploadRequest, error) {
	m := new(CreateModelBinaryFileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ModelService_ListModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).ListModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/ListModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).ListModel(ctx, req.(*ListModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_GetModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/GetModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetModel(ctx, req.(*GetModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_UpdateModelInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModelInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).UpdateModelInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/UpdateModelInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).UpdateModelInstance(ctx, req.(*UpdateModelInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_DeleteModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).DeleteModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/DeleteModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).DeleteModel(ctx, req.(*DeleteModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_DeleteModelInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).DeleteModelInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/DeleteModelInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).DeleteModelInstance(ctx, req.(*DeleteModelInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_TriggerModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).TriggerModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/instill.model.v1alpha.ModelService/TriggerModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).TriggerModel(ctx, req.(*TriggerModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_TriggerModelBinaryFileUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelServiceServer).TriggerModelBinaryFileUpload(&modelServiceTriggerModelBinaryFileUploadServer{stream})
}

type ModelService_TriggerModelBinaryFileUploadServer interface {
	SendAndClose(*TriggerModelBinaryFileUploadResponse) error
	Recv() (*TriggerModelBinaryFileUploadRequest, error)
	grpc.ServerStream
}

type modelServiceTriggerModelBinaryFileUploadServer struct {
	grpc.ServerStream
}

func (x *modelServiceTriggerModelBinaryFileUploadServer) SendAndClose(m *TriggerModelBinaryFileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *modelServiceTriggerModelBinaryFileUploadServer) Recv() (*TriggerModelBinaryFileUploadRequest, error) {
	m := new(TriggerModelBinaryFileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "instill.model.v1alpha.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _ModelService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _ModelService_Readiness_Handler,
		},
		{
			MethodName: "CreateModelByGitHub",
			Handler:    _ModelService_CreateModelByGitHub_Handler,
		},
		{
			MethodName: "ListModel",
			Handler:    _ModelService_ListModel_Handler,
		},
		{
			MethodName: "GetModel",
			Handler:    _ModelService_GetModel_Handler,
		},
		{
			MethodName: "UpdateModelInstance",
			Handler:    _ModelService_UpdateModelInstance_Handler,
		},
		{
			MethodName: "DeleteModel",
			Handler:    _ModelService_DeleteModel_Handler,
		},
		{
			MethodName: "DeleteModelInstance",
			Handler:    _ModelService_DeleteModelInstance_Handler,
		},
		{
			MethodName: "TriggerModel",
			Handler:    _ModelService_TriggerModel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateModelBinaryFileUpload",
			Handler:       _ModelService_CreateModelBinaryFileUpload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TriggerModelBinaryFileUpload",
			Handler:       _ModelService_TriggerModelBinaryFileUpload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "instill/model/v1alpha/model.proto",
}