// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: core/mgmt/v1beta/mgmt_private_service.proto

package mgmtv1beta

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MgmtPrivateService_ListUsersAdmin_FullMethodName                   = "/core.mgmt.v1beta.MgmtPrivateService/ListUsersAdmin"
	MgmtPrivateService_GetUserAdmin_FullMethodName                     = "/core.mgmt.v1beta.MgmtPrivateService/GetUserAdmin"
	MgmtPrivateService_LookUpUserAdmin_FullMethodName                  = "/core.mgmt.v1beta.MgmtPrivateService/LookUpUserAdmin"
	MgmtPrivateService_ListOrganizationsAdmin_FullMethodName           = "/core.mgmt.v1beta.MgmtPrivateService/ListOrganizationsAdmin"
	MgmtPrivateService_GetOrganizationAdmin_FullMethodName             = "/core.mgmt.v1beta.MgmtPrivateService/GetOrganizationAdmin"
	MgmtPrivateService_LookUpOrganizationAdmin_FullMethodName          = "/core.mgmt.v1beta.MgmtPrivateService/LookUpOrganizationAdmin"
	MgmtPrivateService_GetUserSubscriptionAdmin_FullMethodName         = "/core.mgmt.v1beta.MgmtPrivateService/GetUserSubscriptionAdmin"
	MgmtPrivateService_GetOrganizationSubscriptionAdmin_FullMethodName = "/core.mgmt.v1beta.MgmtPrivateService/GetOrganizationSubscriptionAdmin"
	MgmtPrivateService_SubtractCreditAdmin_FullMethodName              = "/core.mgmt.v1beta.MgmtPrivateService/SubtractCreditAdmin"
	MgmtPrivateService_GetRemainingCreditAdmin_FullMethodName          = "/core.mgmt.v1beta.MgmtPrivateService/GetRemainingCreditAdmin"
	MgmtPrivateService_CheckNamespaceAdmin_FullMethodName              = "/core.mgmt.v1beta.MgmtPrivateService/CheckNamespaceAdmin"
	MgmtPrivateService_CheckNamespaceByUIDAdmin_FullMethodName         = "/core.mgmt.v1beta.MgmtPrivateService/CheckNamespaceByUIDAdmin"
)

// MgmtPrivateServiceClient is the client API for MgmtPrivateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Mgmt service responds to internal access
type MgmtPrivateServiceClient interface {
	// ListUsersAdmin method receives a ListUsersAdminRequest message and returns
	// a ListUsersAdminResponse message.
	ListUsersAdmin(ctx context.Context, in *ListUsersAdminRequest, opts ...grpc.CallOption) (*ListUsersAdminResponse, error)
	// GetUserAdmin method receives a GetUserAdminRequest message and returns
	// a GetUserAdminResponse message.
	GetUserAdmin(ctx context.Context, in *GetUserAdminRequest, opts ...grpc.CallOption) (*GetUserAdminResponse, error)
	// LookUpUserAdmin method receives a LookUpUserAdminRequest message and
	// returns a LookUpUserAdminResponse
	LookUpUserAdmin(ctx context.Context, in *LookUpUserAdminRequest, opts ...grpc.CallOption) (*LookUpUserAdminResponse, error)
	// ListOrganizationsAdmin method receives a ListOrganizationsAdminRequest message and returns
	// a ListOrganizationsAdminResponse message.
	ListOrganizationsAdmin(ctx context.Context, in *ListOrganizationsAdminRequest, opts ...grpc.CallOption) (*ListOrganizationsAdminResponse, error)
	// GetOrganizationAdmin method receives a GetOrganizationAdminRequest message and returns
	// a GetOrganizationAdminResponse message.
	GetOrganizationAdmin(ctx context.Context, in *GetOrganizationAdminRequest, opts ...grpc.CallOption) (*GetOrganizationAdminResponse, error)
	// LookUpOrganizationAdmin method receives a LookUpOrganizationAdminRequest message and
	// returns a LookUpOrganizationAdminResponse
	LookUpOrganizationAdmin(ctx context.Context, in *LookUpOrganizationAdminRequest, opts ...grpc.CallOption) (*LookUpOrganizationAdminResponse, error)
	// GetUserSubscriptionAdmin returns the subscription details of a user.
	GetUserSubscriptionAdmin(ctx context.Context, in *GetUserSubscriptionAdminRequest, opts ...grpc.CallOption) (*GetUserSubscriptionAdminResponse, error)
	// GetOrganizationSubscriptionAdmin returns the subscription details of an
	// organization.
	GetOrganizationSubscriptionAdmin(ctx context.Context, in *GetOrganizationSubscriptionAdminRequest, opts ...grpc.CallOption) (*GetOrganizationSubscriptionAdminResponse, error)
	// Subtract Instill Credit from a user or organization account.
	//
	// This endpoint subtracts the specified amount of Instill Credit from an
	// account. This is intended for processes on Instill Cloud that consume
	// credit, such as the execution of pre-configured connectors.
	// Note that if the remaining credit in the account is less than the
	// requested amount, it will be subtracted anyways, leaving the account
	// credit at zero. A ResourceExhausted error will be returned in this case.
	//
	// On Instill Core, this endpoint will return an Unimplemented status.
	SubtractCreditAdmin(ctx context.Context, in *SubtractCreditAdminRequest, opts ...grpc.CallOption) (*SubtractCreditAdminResponse, error)
	// Get the remaining Instill Credit by owner UID
	//
	// This endpoint fetches the remaining unexpired credit of a user or
	// organization, referenced by UID.
	//
	// On Instill Core, this endpoint will return a 404 Not Found status.
	GetRemainingCreditAdmin(ctx context.Context, in *GetRemainingCreditAdminRequest, opts ...grpc.CallOption) (*GetRemainingCreditAdminResponse, error)
	// Check if a namespace is in use
	//
	// Returns the availability of a namespace or, alternatively, the type of
	// resource that is using it.
	CheckNamespaceAdmin(ctx context.Context, in *CheckNamespaceAdminRequest, opts ...grpc.CallOption) (*CheckNamespaceAdminResponse, error)
	// Check if a namespace is in use by UID
	//
	// Returns the availability of a namespace or, alternatively, the type of
	// resource that is using it.
	CheckNamespaceByUIDAdmin(ctx context.Context, in *CheckNamespaceByUIDAdminRequest, opts ...grpc.CallOption) (*CheckNamespaceByUIDAdminResponse, error)
}

type mgmtPrivateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMgmtPrivateServiceClient(cc grpc.ClientConnInterface) MgmtPrivateServiceClient {
	return &mgmtPrivateServiceClient{cc}
}

func (c *mgmtPrivateServiceClient) ListUsersAdmin(ctx context.Context, in *ListUsersAdminRequest, opts ...grpc.CallOption) (*ListUsersAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUsersAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_ListUsersAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) GetUserAdmin(ctx context.Context, in *GetUserAdminRequest, opts ...grpc.CallOption) (*GetUserAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_GetUserAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) LookUpUserAdmin(ctx context.Context, in *LookUpUserAdminRequest, opts ...grpc.CallOption) (*LookUpUserAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookUpUserAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_LookUpUserAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) ListOrganizationsAdmin(ctx context.Context, in *ListOrganizationsAdminRequest, opts ...grpc.CallOption) (*ListOrganizationsAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrganizationsAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_ListOrganizationsAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) GetOrganizationAdmin(ctx context.Context, in *GetOrganizationAdminRequest, opts ...grpc.CallOption) (*GetOrganizationAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrganizationAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_GetOrganizationAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) LookUpOrganizationAdmin(ctx context.Context, in *LookUpOrganizationAdminRequest, opts ...grpc.CallOption) (*LookUpOrganizationAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LookUpOrganizationAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_LookUpOrganizationAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) GetUserSubscriptionAdmin(ctx context.Context, in *GetUserSubscriptionAdminRequest, opts ...grpc.CallOption) (*GetUserSubscriptionAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserSubscriptionAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_GetUserSubscriptionAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) GetOrganizationSubscriptionAdmin(ctx context.Context, in *GetOrganizationSubscriptionAdminRequest, opts ...grpc.CallOption) (*GetOrganizationSubscriptionAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrganizationSubscriptionAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_GetOrganizationSubscriptionAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) SubtractCreditAdmin(ctx context.Context, in *SubtractCreditAdminRequest, opts ...grpc.CallOption) (*SubtractCreditAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubtractCreditAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_SubtractCreditAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) GetRemainingCreditAdmin(ctx context.Context, in *GetRemainingCreditAdminRequest, opts ...grpc.CallOption) (*GetRemainingCreditAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRemainingCreditAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_GetRemainingCreditAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) CheckNamespaceAdmin(ctx context.Context, in *CheckNamespaceAdminRequest, opts ...grpc.CallOption) (*CheckNamespaceAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckNamespaceAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_CheckNamespaceAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPrivateServiceClient) CheckNamespaceByUIDAdmin(ctx context.Context, in *CheckNamespaceByUIDAdminRequest, opts ...grpc.CallOption) (*CheckNamespaceByUIDAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckNamespaceByUIDAdminResponse)
	err := c.cc.Invoke(ctx, MgmtPrivateService_CheckNamespaceByUIDAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtPrivateServiceServer is the server API for MgmtPrivateService service.
// All implementations should embed UnimplementedMgmtPrivateServiceServer
// for forward compatibility.
//
// Mgmt service responds to internal access
type MgmtPrivateServiceServer interface {
	// ListUsersAdmin method receives a ListUsersAdminRequest message and returns
	// a ListUsersAdminResponse message.
	ListUsersAdmin(context.Context, *ListUsersAdminRequest) (*ListUsersAdminResponse, error)
	// GetUserAdmin method receives a GetUserAdminRequest message and returns
	// a GetUserAdminResponse message.
	GetUserAdmin(context.Context, *GetUserAdminRequest) (*GetUserAdminResponse, error)
	// LookUpUserAdmin method receives a LookUpUserAdminRequest message and
	// returns a LookUpUserAdminResponse
	LookUpUserAdmin(context.Context, *LookUpUserAdminRequest) (*LookUpUserAdminResponse, error)
	// ListOrganizationsAdmin method receives a ListOrganizationsAdminRequest message and returns
	// a ListOrganizationsAdminResponse message.
	ListOrganizationsAdmin(context.Context, *ListOrganizationsAdminRequest) (*ListOrganizationsAdminResponse, error)
	// GetOrganizationAdmin method receives a GetOrganizationAdminRequest message and returns
	// a GetOrganizationAdminResponse message.
	GetOrganizationAdmin(context.Context, *GetOrganizationAdminRequest) (*GetOrganizationAdminResponse, error)
	// LookUpOrganizationAdmin method receives a LookUpOrganizationAdminRequest message and
	// returns a LookUpOrganizationAdminResponse
	LookUpOrganizationAdmin(context.Context, *LookUpOrganizationAdminRequest) (*LookUpOrganizationAdminResponse, error)
	// GetUserSubscriptionAdmin returns the subscription details of a user.
	GetUserSubscriptionAdmin(context.Context, *GetUserSubscriptionAdminRequest) (*GetUserSubscriptionAdminResponse, error)
	// GetOrganizationSubscriptionAdmin returns the subscription details of an
	// organization.
	GetOrganizationSubscriptionAdmin(context.Context, *GetOrganizationSubscriptionAdminRequest) (*GetOrganizationSubscriptionAdminResponse, error)
	// Subtract Instill Credit from a user or organization account.
	//
	// This endpoint subtracts the specified amount of Instill Credit from an
	// account. This is intended for processes on Instill Cloud that consume
	// credit, such as the execution of pre-configured connectors.
	// Note that if the remaining credit in the account is less than the
	// requested amount, it will be subtracted anyways, leaving the account
	// credit at zero. A ResourceExhausted error will be returned in this case.
	//
	// On Instill Core, this endpoint will return an Unimplemented status.
	SubtractCreditAdmin(context.Context, *SubtractCreditAdminRequest) (*SubtractCreditAdminResponse, error)
	// Get the remaining Instill Credit by owner UID
	//
	// This endpoint fetches the remaining unexpired credit of a user or
	// organization, referenced by UID.
	//
	// On Instill Core, this endpoint will return a 404 Not Found status.
	GetRemainingCreditAdmin(context.Context, *GetRemainingCreditAdminRequest) (*GetRemainingCreditAdminResponse, error)
	// Check if a namespace is in use
	//
	// Returns the availability of a namespace or, alternatively, the type of
	// resource that is using it.
	CheckNamespaceAdmin(context.Context, *CheckNamespaceAdminRequest) (*CheckNamespaceAdminResponse, error)
	// Check if a namespace is in use by UID
	//
	// Returns the availability of a namespace or, alternatively, the type of
	// resource that is using it.
	CheckNamespaceByUIDAdmin(context.Context, *CheckNamespaceByUIDAdminRequest) (*CheckNamespaceByUIDAdminResponse, error)
}

// UnimplementedMgmtPrivateServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMgmtPrivateServiceServer struct{}

func (UnimplementedMgmtPrivateServiceServer) ListUsersAdmin(context.Context, *ListUsersAdminRequest) (*ListUsersAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) GetUserAdmin(context.Context, *GetUserAdminRequest) (*GetUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) LookUpUserAdmin(context.Context, *LookUpUserAdminRequest) (*LookUpUserAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpUserAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) ListOrganizationsAdmin(context.Context, *ListOrganizationsAdminRequest) (*ListOrganizationsAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizationsAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) GetOrganizationAdmin(context.Context, *GetOrganizationAdminRequest) (*GetOrganizationAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) LookUpOrganizationAdmin(context.Context, *LookUpOrganizationAdminRequest) (*LookUpOrganizationAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookUpOrganizationAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) GetUserSubscriptionAdmin(context.Context, *GetUserSubscriptionAdminRequest) (*GetUserSubscriptionAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSubscriptionAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) GetOrganizationSubscriptionAdmin(context.Context, *GetOrganizationSubscriptionAdminRequest) (*GetOrganizationSubscriptionAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrganizationSubscriptionAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) SubtractCreditAdmin(context.Context, *SubtractCreditAdminRequest) (*SubtractCreditAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubtractCreditAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) GetRemainingCreditAdmin(context.Context, *GetRemainingCreditAdminRequest) (*GetRemainingCreditAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemainingCreditAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) CheckNamespaceAdmin(context.Context, *CheckNamespaceAdminRequest) (*CheckNamespaceAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNamespaceAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) CheckNamespaceByUIDAdmin(context.Context, *CheckNamespaceByUIDAdminRequest) (*CheckNamespaceByUIDAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNamespaceByUIDAdmin not implemented")
}
func (UnimplementedMgmtPrivateServiceServer) testEmbeddedByValue() {}

// UnsafeMgmtPrivateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MgmtPrivateServiceServer will
// result in compilation errors.
type UnsafeMgmtPrivateServiceServer interface {
	mustEmbedUnimplementedMgmtPrivateServiceServer()
}

func RegisterMgmtPrivateServiceServer(s grpc.ServiceRegistrar, srv MgmtPrivateServiceServer) {
	// If the following call pancis, it indicates UnimplementedMgmtPrivateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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
		FullMethod: MgmtPrivateService_ListUsersAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).ListUsersAdmin(ctx, req.(*ListUsersAdminRequest))
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
		FullMethod: MgmtPrivateService_GetUserAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).GetUserAdmin(ctx, req.(*GetUserAdminRequest))
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
		FullMethod: MgmtPrivateService_LookUpUserAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).LookUpUserAdmin(ctx, req.(*LookUpUserAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_ListOrganizationsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationsAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).ListOrganizationsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_ListOrganizationsAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).ListOrganizationsAdmin(ctx, req.(*ListOrganizationsAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_GetOrganizationAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).GetOrganizationAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_GetOrganizationAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).GetOrganizationAdmin(ctx, req.(*GetOrganizationAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_LookUpOrganizationAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookUpOrganizationAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).LookUpOrganizationAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_LookUpOrganizationAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).LookUpOrganizationAdmin(ctx, req.(*LookUpOrganizationAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_GetUserSubscriptionAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSubscriptionAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).GetUserSubscriptionAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_GetUserSubscriptionAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).GetUserSubscriptionAdmin(ctx, req.(*GetUserSubscriptionAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_GetOrganizationSubscriptionAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrganizationSubscriptionAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).GetOrganizationSubscriptionAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_GetOrganizationSubscriptionAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).GetOrganizationSubscriptionAdmin(ctx, req.(*GetOrganizationSubscriptionAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_SubtractCreditAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractCreditAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).SubtractCreditAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_SubtractCreditAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).SubtractCreditAdmin(ctx, req.(*SubtractCreditAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_GetRemainingCreditAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRemainingCreditAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).GetRemainingCreditAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_GetRemainingCreditAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).GetRemainingCreditAdmin(ctx, req.(*GetRemainingCreditAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_CheckNamespaceAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckNamespaceAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).CheckNamespaceAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_CheckNamespaceAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).CheckNamespaceAdmin(ctx, req.(*CheckNamespaceAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPrivateService_CheckNamespaceByUIDAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckNamespaceByUIDAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPrivateServiceServer).CheckNamespaceByUIDAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPrivateService_CheckNamespaceByUIDAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPrivateServiceServer).CheckNamespaceByUIDAdmin(ctx, req.(*CheckNamespaceByUIDAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MgmtPrivateService_ServiceDesc is the grpc.ServiceDesc for MgmtPrivateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MgmtPrivateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.mgmt.v1beta.MgmtPrivateService",
	HandlerType: (*MgmtPrivateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUsersAdmin",
			Handler:    _MgmtPrivateService_ListUsersAdmin_Handler,
		},
		{
			MethodName: "GetUserAdmin",
			Handler:    _MgmtPrivateService_GetUserAdmin_Handler,
		},
		{
			MethodName: "LookUpUserAdmin",
			Handler:    _MgmtPrivateService_LookUpUserAdmin_Handler,
		},
		{
			MethodName: "ListOrganizationsAdmin",
			Handler:    _MgmtPrivateService_ListOrganizationsAdmin_Handler,
		},
		{
			MethodName: "GetOrganizationAdmin",
			Handler:    _MgmtPrivateService_GetOrganizationAdmin_Handler,
		},
		{
			MethodName: "LookUpOrganizationAdmin",
			Handler:    _MgmtPrivateService_LookUpOrganizationAdmin_Handler,
		},
		{
			MethodName: "GetUserSubscriptionAdmin",
			Handler:    _MgmtPrivateService_GetUserSubscriptionAdmin_Handler,
		},
		{
			MethodName: "GetOrganizationSubscriptionAdmin",
			Handler:    _MgmtPrivateService_GetOrganizationSubscriptionAdmin_Handler,
		},
		{
			MethodName: "SubtractCreditAdmin",
			Handler:    _MgmtPrivateService_SubtractCreditAdmin_Handler,
		},
		{
			MethodName: "GetRemainingCreditAdmin",
			Handler:    _MgmtPrivateService_GetRemainingCreditAdmin_Handler,
		},
		{
			MethodName: "CheckNamespaceAdmin",
			Handler:    _MgmtPrivateService_CheckNamespaceAdmin_Handler,
		},
		{
			MethodName: "CheckNamespaceByUIDAdmin",
			Handler:    _MgmtPrivateService_CheckNamespaceByUIDAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/mgmt/v1beta/mgmt_private_service.proto",
}
