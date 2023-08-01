// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: base/usage/v1alpha/usage_service.proto

package usagev1alpha

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
	UsageService_Liveness_FullMethodName          = "/base.usage.v1alpha.UsageService/Liveness"
	UsageService_Readiness_FullMethodName         = "/base.usage.v1alpha.UsageService/Readiness"
	UsageService_CreateSession_FullMethodName     = "/base.usage.v1alpha.UsageService/CreateSession"
	UsageService_SendSessionReport_FullMethodName = "/base.usage.v1alpha.UsageService/SendSessionReport"
)

// UsageServiceClient is the client API for UsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsageServiceClient interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// CreateSession method receives a CreateSessionRequest message and returns
	// a CreateSessionResponse message.
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	// SendSessionReport method receives a SendSessionReportRequest message and
	// returns a SendSessionReportResponse message.
	SendSessionReport(ctx context.Context, in *SendSessionReportRequest, opts ...grpc.CallOption) (*SendSessionReportResponse, error)
}

type usageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsageServiceClient(cc grpc.ClientConnInterface) UsageServiceClient {
	return &usageServiceClient{cc}
}

func (c *usageServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, UsageService_Liveness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, UsageService_Readiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, UsageService_CreateSession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usageServiceClient) SendSessionReport(ctx context.Context, in *SendSessionReportRequest, opts ...grpc.CallOption) (*SendSessionReportResponse, error) {
	out := new(SendSessionReportResponse)
	err := c.cc.Invoke(ctx, UsageService_SendSessionReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsageServiceServer is the server API for UsageService service.
// All implementations should embed UnimplementedUsageServiceServer
// for forward compatibility
type UsageServiceServer interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// CreateSession method receives a CreateSessionRequest message and returns
	// a CreateSessionResponse message.
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	// SendSessionReport method receives a SendSessionReportRequest message and
	// returns a SendSessionReportResponse message.
	SendSessionReport(context.Context, *SendSessionReportRequest) (*SendSessionReportResponse, error)
}

// UnimplementedUsageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUsageServiceServer struct {
}

func (UnimplementedUsageServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedUsageServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (UnimplementedUsageServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedUsageServiceServer) SendSessionReport(context.Context, *SendSessionReportRequest) (*SendSessionReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSessionReport not implemented")
}

// UnsafeUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsageServiceServer will
// result in compilation errors.
type UnsafeUsageServiceServer interface {
	mustEmbedUnimplementedUsageServiceServer()
}

func RegisterUsageServiceServer(s grpc.ServiceRegistrar, srv UsageServiceServer) {
	s.RegisterService(&UsageService_ServiceDesc, srv)
}

func _UsageService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsageService_Liveness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsageService_Readiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsageService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsageService_SendSessionReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSessionReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageServiceServer).SendSessionReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsageService_SendSessionReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageServiceServer).SendSessionReport(ctx, req.(*SendSessionReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsageService_ServiceDesc is the grpc.ServiceDesc for UsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base.usage.v1alpha.UsageService",
	HandlerType: (*UsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _UsageService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _UsageService_Readiness_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _UsageService_CreateSession_Handler,
		},
		{
			MethodName: "SendSessionReport",
			Handler:    _UsageService_SendSessionReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/usage/v1alpha/usage_service.proto",
}