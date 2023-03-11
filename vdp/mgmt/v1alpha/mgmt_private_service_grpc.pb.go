// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mgmtv1alpha

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

// MgmtPrivateServiceClient is the client API for MgmtPrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MgmtPrivateServiceClient interface {
	// ListUsersAdmin method receives a ListUsersAdminRequest message and returns a
	// ListUsersAdminResponse message.
	ListUsersAdmin(ctx context.Context, in *ListUsersAdminRequest, opts ...grpc.CallOption) (*ListUsersAdminResponse, error)
	// CreateUserAdmin receives a CreateUserAdminRequest message and returns a
	// a GetUserAdminResponse
	CreateUserAdmin(ctx context.Context, in *CreateUserAdminRequest, opts ...grpc.CallOption) (*CreateUserAdminResponse, error)
	// GetUserAdmin method receives a GetUserAdminRequest message and returns
	// a GetUserAdminResponse message.
	GetUserAdmin(ctx context.Context, in *GetUserAdminRequest, opts ...grpc.CallOption) (*GetUserAdminResponse, error)
	// UpdateUserAdmin method receives a UpdateUserAdminRequest message and returns
	// a UpdateUserAdminResponse
	UpdateUserAdmin(ctx context.Context, in *UpdateUserAdminRequest, opts ...grpc.CallOption) (*UpdateUserAdminResponse, error)
	// DeleteUserAdmin method receives a DeleteUserAdminRequest message and returns a
	// DeleteUserAdminResponse
	DeleteUserAdmin(ctx context.Context, in *DeleteUserAdminRequest, opts ...grpc.CallOption) (*DeleteUserAdminResponse, error)
	// LookUpUserAdmin method receives a LookUpUserAdminRequest message and returns a
	// LookUpUserAdminResponse
	LookUpUserAdmin(ctx context.Context, in *LookUpUserAdminRequest, opts ...grpc.CallOption) (*LookUpUserAdminResponse, error)
}

type mgmtPrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMgmtPrivateServiceClient(cc grpc.ClientConnInterface) MgmtPrivateServiceClient {
	return &mgmtPrivateServiceClient{cc}
}

func (c *mgmtPrivateServiceClient) ListUsersAdmin(ctx context.Context, in *ListUsersAdminRequest, opts ...grpc.CallOption) (*ListUsersAdminResponse, error) {
	out := new(ListUsersAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.mgmt.v1alpha.MgmtPrivateService/ListUsersAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) CreateUserAdmin(ctx context.Context, in *CreateUserAdminRequest, opts ...grpc.CallOption) (*CreateUserAdminResponse, error) {
	out := new(CreateUserAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.mgmt.v1alpha.MgmtPrivateService/CreateUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) GetUserAdmin(ctx context.Context, in *GetUserAdminRequest, opts ...grpc.CallOption) (*GetUserAdminResponse, error) {
	out := new(GetUserAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.mgmt.v1alpha.MgmtPrivateService/GetUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) UpdateUserAdmin(ctx context.Context, in *UpdateUserAdminRequest, opts ...grpc.CallOption) (*UpdateUserAdminResponse, error) {
	out := new(UpdateUserAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.mgmt.v1alpha.MgmtPrivateService/UpdateUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) DeleteUserAdmin(ctx context.Context, in *DeleteUserAdminRequest, opts ...grpc.CallOption) (*DeleteUserAdminResponse, error) {
	out := new(DeleteUserAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.mgmt.v1alpha.MgmtPrivateService/DeleteUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) LookUpUserAdmin(ctx context.Context, in *LookUpUserAdminRequest, opts ...grpc.CallOption) (*LookUpUserAdminResponse, error) {
	out := new(LookUpUserAdminResponse)
	err := c.cc.Invoke(ctx, "/vdp.mgmt.v1alpha.MgmtPrivateService/LookUpUserAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtPrivateServiceServer is the server API for MgmtPrivateService service.
// All implementations should embed UnimplementedMgmtPrivateServiceServer
// for forward compatibility
type MgmtPrivateServiceServer interface {
	// ListUsersAdmin method receives a ListUsersAdminRequest message and returns a
	// ListUsersAdminResponse message.
	ListUsersAdmin(context.Context, *ListUsersAdminRequest) (*ListUsersAdminResponse, error)
	// CreateUserAdmin receives a CreateUserAdminRequest message and returns a
	// a GetUserAdminResponse
	CreateUserAdmin(context.Context, *CreateUserAdminRequest) (*CreateUserAdminResponse, error)
	// GetUserAdmin method receives a GetUserAdminRequest message and returns
	// a GetUserAdminResponse message.
	GetUserAdmin(context.Context, *GetUserAdminRequest) (*GetUserAdminResponse, error)
	// UpdateUserAdmin method receives a UpdateUserAdminRequest message and returns
	// a UpdateUserAdminResponse
	UpdateUserAdmin(context.Context, *UpdateUserAdminRequest) (*UpdateUserAdminResponse, error)
	// DeleteUserAdmin method receives a DeleteUserAdminRequest message and returns a
	// DeleteUserAdminResponse
	DeleteUserAdmin(context.Context, *DeleteUserAdminRequest) (*DeleteUserAdminResponse, error)
	// LookUpUserAdmin method receives a LookUpUserAdminRequest message and returns a
	// LookUpUserAdminResponse
	LookUpUserAdmin(context.Context, *LookUpUserAdminRequest) (*LookUpUserAdminResponse, error)
}

// UnimplementedMgmtPrivateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMgmtPrivateServiceServer struct {
}

func (UnimplementedMgmtPrivateServiceServer) ListUsersAdmin(context.Context, *ListUsersAdminRequest) (*ListUsersAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) CreateUserAdmin(context.Context, *CreateUserAdminRequest) (*CreateUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) GetUserAdmin(context.Context, *GetUserAdminRequest) (*GetUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) UpdateUserAdmin(context.Context, *UpdateUserAdminRequest) (*UpdateUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) DeleteUserAdmin(context.Context, *DeleteUserAdminRequest) (*DeleteUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) LookUpUserAdmin(context.Context, *LookUpUserAdminRequest) (*LookUpUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpUserAdmin not implemented")
}

// UnsafeMgmtPrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MgmtPrivateServiceServer will
// result in compilation errors.
type UnsafeMgmtPrivateServiceServer interface {
	mustEmbedUnimplementedMgmtPrivateServiceServer()
}

func RegisterMgmtPrivateServiceServer(s grpc.ServiceRegistrar, srv MgmtPrivateServiceServer) {
	s.RegisterService(&MgmtPrivateService_ServiceDesc, srv)
}

func _MgmtPrivateService_ListUsersAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).ListUsersAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.mgmt.v1alpha.MgmtPrivateService/ListUsersAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).ListUsersAdmin(ctx, req.(*ListUsersAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_CreateUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).CreateUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.mgmt.v1alpha.MgmtPrivateService/CreateUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).CreateUserAdmin(ctx, req.(*CreateUserAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_GetUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).GetUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.mgmt.v1alpha.MgmtPrivateService/GetUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).GetUserAdmin(ctx, req.(*GetUserAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_UpdateUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).UpdateUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.mgmt.v1alpha.MgmtPrivateService/UpdateUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).UpdateUserAdmin(ctx, req.(*UpdateUserAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_DeleteUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).DeleteUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.mgmt.v1alpha.MgmtPrivateService/DeleteUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).DeleteUserAdmin(ctx, req.(*DeleteUserAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_LookUpUserAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpUserAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).LookUpUserAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vdp.mgmt.v1alpha.MgmtPrivateService/LookUpUserAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).LookUpUserAdmin(ctx, req.(*LookUpUserAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MgmtPrivateService_ServiceDesc is the grpc.ServiceDesc for MgmtPrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MgmtPrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vdp.mgmt.v1alpha.MgmtPrivateService",
	HandlerType: (*MgmtPrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsersAdmin",
			Handler:    _MgmtPrivateService_ListUsersAdmin_Handler,
		},
		{
			MethodName: "CreateUserAdmin",
			Handler:    _MgmtPrivateService_CreateUserAdmin_Handler,
		},
		{
			MethodName: "GetUserAdmin",
			Handler:    _MgmtPrivateService_GetUserAdmin_Handler,
		},
		{
			MethodName: "UpdateUserAdmin",
			Handler:    _MgmtPrivateService_UpdateUserAdmin_Handler,
		},
		{
			MethodName: "DeleteUserAdmin",
			Handler:    _MgmtPrivateService_DeleteUserAdmin_Handler,
		},
		{
			MethodName: "LookUpUserAdmin",
			Handler:    _MgmtPrivateService_LookUpUserAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vdp/mgmt/v1alpha/mgmt_private_service.proto",
}