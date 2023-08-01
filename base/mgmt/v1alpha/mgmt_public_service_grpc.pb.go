// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: base/mgmt/v1alpha/mgmt_public_service.proto

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

const (
	MgmtPublicService_Liveness_FullMethodName                         = "/base.mgmt.v1alpha.MgmtPublicService/Liveness"
	MgmtPublicService_Readiness_FullMethodName                        = "/base.mgmt.v1alpha.MgmtPublicService/Readiness"
	MgmtPublicService_QueryAuthenticatedUser_FullMethodName           = "/base.mgmt.v1alpha.MgmtPublicService/QueryAuthenticatedUser"
	MgmtPublicService_PatchAuthenticatedUser_FullMethodName           = "/base.mgmt.v1alpha.MgmtPublicService/PatchAuthenticatedUser"
	MgmtPublicService_ExistUsername_FullMethodName                    = "/base.mgmt.v1alpha.MgmtPublicService/ExistUsername"
	MgmtPublicService_CreateToken_FullMethodName                      = "/base.mgmt.v1alpha.MgmtPublicService/CreateToken"
	MgmtPublicService_ListTokens_FullMethodName                       = "/base.mgmt.v1alpha.MgmtPublicService/ListTokens"
	MgmtPublicService_GetToken_FullMethodName                         = "/base.mgmt.v1alpha.MgmtPublicService/GetToken"
	MgmtPublicService_DeleteToken_FullMethodName                      = "/base.mgmt.v1alpha.MgmtPublicService/DeleteToken"
	MgmtPublicService_ListPipelineTriggerRecords_FullMethodName       = "/base.mgmt.v1alpha.MgmtPublicService/ListPipelineTriggerRecords"
	MgmtPublicService_ListPipelineTriggerTableRecords_FullMethodName  = "/base.mgmt.v1alpha.MgmtPublicService/ListPipelineTriggerTableRecords"
	MgmtPublicService_ListPipelineTriggerChartRecords_FullMethodName  = "/base.mgmt.v1alpha.MgmtPublicService/ListPipelineTriggerChartRecords"
	MgmtPublicService_ListConnectorExecuteRecords_FullMethodName      = "/base.mgmt.v1alpha.MgmtPublicService/ListConnectorExecuteRecords"
	MgmtPublicService_ListConnectorExecuteTableRecords_FullMethodName = "/base.mgmt.v1alpha.MgmtPublicService/ListConnectorExecuteTableRecords"
	MgmtPublicService_ListConnectorExecuteChartRecords_FullMethodName = "/base.mgmt.v1alpha.MgmtPublicService/ListConnectorExecuteChartRecords"
)

// MgmtPublicServiceClient is the client API for MgmtPublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MgmtPublicServiceClient interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// QueryAuthenticatedUser method receives a QueryAuthenticatedUserRequest
	// message and returns a QueryAuthenticatedUserResponse message.
	QueryAuthenticatedUser(ctx context.Context, in *QueryAuthenticatedUserRequest, opts ...grpc.CallOption) (*QueryAuthenticatedUserResponse, error)
	// PatchAuthenticatedUser method receives a PatchAuthenticatedUserRequest
	// message and returns a PatchAuthenticatedUserResponse message.
	PatchAuthenticatedUser(ctx context.Context, in *PatchAuthenticatedUserRequest, opts ...grpc.CallOption) (*PatchAuthenticatedUserResponse, error)
	// ExistUsername method receives a ExistUsernameRequest message and returns a
	// ExistUsernameResponse
	ExistUsername(ctx context.Context, in *ExistUsernameRequest, opts ...grpc.CallOption) (*ExistUsernameResponse, error)
	// CreateToken method receives a CreateTokenRequest message and returns
	// a CreateTokenResponse message.
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	// ListTokens method receives a ListTokensRequest message and returns a
	// ListTokensResponse message.
	ListTokens(ctx context.Context, in *ListTokensRequest, opts ...grpc.CallOption) (*ListTokensResponse, error)
	// GetToken method receives a GetTokenRequest message and returns a
	// GetTokenResponse message.
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	// DeleteToken method receives a DeleteTokenRequest message and returns
	// a DeleteTokenResponse message.
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error)
	// ListPipelineTriggerRecords method receives a ListPipelineTriggerRecordsRequest message and returns a
	// ListPipelineTriggerRecordsResponse message.
	ListPipelineTriggerRecords(ctx context.Context, in *ListPipelineTriggerRecordsRequest, opts ...grpc.CallOption) (*ListPipelineTriggerRecordsResponse, error)
	// ListPipelineTriggerTableRecords method receives a ListPipelineTriggerTableRecordsRequest message and returns a
	// ListPipelineTriggerTableRecordsResponse message.
	ListPipelineTriggerTableRecords(ctx context.Context, in *ListPipelineTriggerTableRecordsRequest, opts ...grpc.CallOption) (*ListPipelineTriggerTableRecordsResponse, error)
	// ListPipelineTriggerChartRecords method receives a ListPipelineTriggerChartRecordsRequest message and returns a
	// ListPipelineTriggerChartRecordsResponse message.
	ListPipelineTriggerChartRecords(ctx context.Context, in *ListPipelineTriggerChartRecordsRequest, opts ...grpc.CallOption) (*ListPipelineTriggerChartRecordsResponse, error)
	// ListConnectorExecuteRecords method receives a ListConnectorExecuteRecordsRequest message and returns a
	// ListConnectorExecuteRecordsResponse message.
	ListConnectorExecuteRecords(ctx context.Context, in *ListConnectorExecuteRecordsRequest, opts ...grpc.CallOption) (*ListConnectorExecuteRecordsResponse, error)
	// ListConnectorExecuteTableRecords method receives a ListConnectorExecuteTableRecordsRequest message and returns a
	// ListConnectorExecuteTableRecordsResponse message.
	ListConnectorExecuteTableRecords(ctx context.Context, in *ListConnectorExecuteTableRecordsRequest, opts ...grpc.CallOption) (*ListConnectorExecuteTableRecordsResponse, error)
	// ListConnectorExecuteChartRecords method receives a ListConnectorExecuteChartRecordsRequest message and returns a
	// ListConnectorExecuteChartRecordsResponse message.
	ListConnectorExecuteChartRecords(ctx context.Context, in *ListConnectorExecuteChartRecordsRequest, opts ...grpc.CallOption) (*ListConnectorExecuteChartRecordsResponse, error)
}

type mgmtPublicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMgmtPublicServiceClient(cc grpc.ClientConnInterface) MgmtPublicServiceClient {
	return &mgmtPublicServiceClient{cc}
}

func (c *mgmtPublicServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_Liveness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_Readiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) QueryAuthenticatedUser(ctx context.Context, in *QueryAuthenticatedUserRequest, opts ...grpc.CallOption) (*QueryAuthenticatedUserResponse, error) {
	out := new(QueryAuthenticatedUserResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_QueryAuthenticatedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) PatchAuthenticatedUser(ctx context.Context, in *PatchAuthenticatedUserRequest, opts ...grpc.CallOption) (*PatchAuthenticatedUserResponse, error) {
	out := new(PatchAuthenticatedUserResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_PatchAuthenticatedUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ExistUsername(ctx context.Context, in *ExistUsernameRequest, opts ...grpc.CallOption) (*ExistUsernameResponse, error) {
	out := new(ExistUsernameResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ExistUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_CreateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListTokens(ctx context.Context, in *ListTokensRequest, opts ...grpc.CallOption) (*ListTokensResponse, error) {
	out := new(ListTokensResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_GetToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*DeleteTokenResponse, error) {
	out := new(DeleteTokenResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_DeleteToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListPipelineTriggerRecords(ctx context.Context, in *ListPipelineTriggerRecordsRequest, opts ...grpc.CallOption) (*ListPipelineTriggerRecordsResponse, error) {
	out := new(ListPipelineTriggerRecordsResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListPipelineTriggerRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListPipelineTriggerTableRecords(ctx context.Context, in *ListPipelineTriggerTableRecordsRequest, opts ...grpc.CallOption) (*ListPipelineTriggerTableRecordsResponse, error) {
	out := new(ListPipelineTriggerTableRecordsResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListPipelineTriggerTableRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListPipelineTriggerChartRecords(ctx context.Context, in *ListPipelineTriggerChartRecordsRequest, opts ...grpc.CallOption) (*ListPipelineTriggerChartRecordsResponse, error) {
	out := new(ListPipelineTriggerChartRecordsResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListPipelineTriggerChartRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListConnectorExecuteRecords(ctx context.Context, in *ListConnectorExecuteRecordsRequest, opts ...grpc.CallOption) (*ListConnectorExecuteRecordsResponse, error) {
	out := new(ListConnectorExecuteRecordsResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListConnectorExecuteRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListConnectorExecuteTableRecords(ctx context.Context, in *ListConnectorExecuteTableRecordsRequest, opts ...grpc.CallOption) (*ListConnectorExecuteTableRecordsResponse, error) {
	out := new(ListConnectorExecuteTableRecordsResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListConnectorExecuteTableRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mgmtPublicServiceClient) ListConnectorExecuteChartRecords(ctx context.Context, in *ListConnectorExecuteChartRecordsRequest, opts ...grpc.CallOption) (*ListConnectorExecuteChartRecordsResponse, error) {
	out := new(ListConnectorExecuteChartRecordsResponse)
	err := c.cc.Invoke(ctx, MgmtPublicService_ListConnectorExecuteChartRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MgmtPublicServiceServer is the server API for MgmtPublicService service.
// All implementations should embed UnimplementedMgmtPublicServiceServer
// for forward compatibility
type MgmtPublicServiceServer interface {
	// Liveness method receives a LivenessRequest message and returns a
	// LivenessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// Readiness method receives a ReadinessRequest message and returns a
	// ReadinessResponse message.
	// See https://github.com/grpc/grpc/blob/master/doc/health-checking.md
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// QueryAuthenticatedUser method receives a QueryAuthenticatedUserRequest
	// message and returns a QueryAuthenticatedUserResponse message.
	QueryAuthenticatedUser(context.Context, *QueryAuthenticatedUserRequest) (*QueryAuthenticatedUserResponse, error)
	// PatchAuthenticatedUser method receives a PatchAuthenticatedUserRequest
	// message and returns a PatchAuthenticatedUserResponse message.
	PatchAuthenticatedUser(context.Context, *PatchAuthenticatedUserRequest) (*PatchAuthenticatedUserResponse, error)
	// ExistUsername method receives a ExistUsernameRequest message and returns a
	// ExistUsernameResponse
	ExistUsername(context.Context, *ExistUsernameRequest) (*ExistUsernameResponse, error)
	// CreateToken method receives a CreateTokenRequest message and returns
	// a CreateTokenResponse message.
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	// ListTokens method receives a ListTokensRequest message and returns a
	// ListTokensResponse message.
	ListTokens(context.Context, *ListTokensRequest) (*ListTokensResponse, error)
	// GetToken method receives a GetTokenRequest message and returns a
	// GetTokenResponse message.
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	// DeleteToken method receives a DeleteTokenRequest message and returns
	// a DeleteTokenResponse message.
	DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error)
	// ListPipelineTriggerRecords method receives a ListPipelineTriggerRecordsRequest message and returns a
	// ListPipelineTriggerRecordsResponse message.
	ListPipelineTriggerRecords(context.Context, *ListPipelineTriggerRecordsRequest) (*ListPipelineTriggerRecordsResponse, error)
	// ListPipelineTriggerTableRecords method receives a ListPipelineTriggerTableRecordsRequest message and returns a
	// ListPipelineTriggerTableRecordsResponse message.
	ListPipelineTriggerTableRecords(context.Context, *ListPipelineTriggerTableRecordsRequest) (*ListPipelineTriggerTableRecordsResponse, error)
	// ListPipelineTriggerChartRecords method receives a ListPipelineTriggerChartRecordsRequest message and returns a
	// ListPipelineTriggerChartRecordsResponse message.
	ListPipelineTriggerChartRecords(context.Context, *ListPipelineTriggerChartRecordsRequest) (*ListPipelineTriggerChartRecordsResponse, error)
	// ListConnectorExecuteRecords method receives a ListConnectorExecuteRecordsRequest message and returns a
	// ListConnectorExecuteRecordsResponse message.
	ListConnectorExecuteRecords(context.Context, *ListConnectorExecuteRecordsRequest) (*ListConnectorExecuteRecordsResponse, error)
	// ListConnectorExecuteTableRecords method receives a ListConnectorExecuteTableRecordsRequest message and returns a
	// ListConnectorExecuteTableRecordsResponse message.
	ListConnectorExecuteTableRecords(context.Context, *ListConnectorExecuteTableRecordsRequest) (*ListConnectorExecuteTableRecordsResponse, error)
	// ListConnectorExecuteChartRecords method receives a ListConnectorExecuteChartRecordsRequest message and returns a
	// ListConnectorExecuteChartRecordsResponse message.
	ListConnectorExecuteChartRecords(context.Context, *ListConnectorExecuteChartRecordsRequest) (*ListConnectorExecuteChartRecordsResponse, error)
}

// UnimplementedMgmtPublicServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMgmtPublicServiceServer struct {
}

func (UnimplementedMgmtPublicServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (UnimplementedMgmtPublicServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (UnimplementedMgmtPublicServiceServer) QueryAuthenticatedUser(context.Context, *QueryAuthenticatedUserRequest) (*QueryAuthenticatedUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAuthenticatedUser not implemented")
}
func (UnimplementedMgmtPublicServiceServer) PatchAuthenticatedUser(context.Context, *PatchAuthenticatedUserRequest) (*PatchAuthenticatedUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAuthenticatedUser not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ExistUsername(context.Context, *ExistUsernameRequest) (*ExistUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistUsername not implemented")
}
func (UnimplementedMgmtPublicServiceServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListTokens(context.Context, *ListTokensRequest) (*ListTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}
func (UnimplementedMgmtPublicServiceServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedMgmtPublicServiceServer) DeleteToken(context.Context, *DeleteTokenRequest) (*DeleteTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListPipelineTriggerRecords(context.Context, *ListPipelineTriggerRecordsRequest) (*ListPipelineTriggerRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineTriggerRecords not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListPipelineTriggerTableRecords(context.Context, *ListPipelineTriggerTableRecordsRequest) (*ListPipelineTriggerTableRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineTriggerTableRecords not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListPipelineTriggerChartRecords(context.Context, *ListPipelineTriggerChartRecordsRequest) (*ListPipelineTriggerChartRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineTriggerChartRecords not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListConnectorExecuteRecords(context.Context, *ListConnectorExecuteRecordsRequest) (*ListConnectorExecuteRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectorExecuteRecords not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListConnectorExecuteTableRecords(context.Context, *ListConnectorExecuteTableRecordsRequest) (*ListConnectorExecuteTableRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectorExecuteTableRecords not implemented")
}
func (UnimplementedMgmtPublicServiceServer) ListConnectorExecuteChartRecords(context.Context, *ListConnectorExecuteChartRecordsRequest) (*ListConnectorExecuteChartRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListConnectorExecuteChartRecords not implemented")
}

// UnsafeMgmtPublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MgmtPublicServiceServer will
// result in compilation errors.
type UnsafeMgmtPublicServiceServer interface {
	mustEmbedUnimplementedMgmtPublicServiceServer()
}

func RegisterMgmtPublicServiceServer(s grpc.ServiceRegistrar, srv MgmtPublicServiceServer) {
	s.RegisterService(&MgmtPublicService_ServiceDesc, srv)
}

func _MgmtPublicService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_Liveness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_Readiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_QueryAuthenticatedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthenticatedUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).QueryAuthenticatedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_QueryAuthenticatedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).QueryAuthenticatedUser(ctx, req.(*QueryAuthenticatedUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_PatchAuthenticatedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAuthenticatedUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).PatchAuthenticatedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_PatchAuthenticatedUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).PatchAuthenticatedUser(ctx, req.(*PatchAuthenticatedUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ExistUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ExistUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ExistUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ExistUsername(ctx, req.(*ExistUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_CreateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListTokens(ctx, req.(*ListTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_DeleteToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListPipelineTriggerRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineTriggerRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListPipelineTriggerRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListPipelineTriggerRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListPipelineTriggerRecords(ctx, req.(*ListPipelineTriggerRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListPipelineTriggerTableRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineTriggerTableRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListPipelineTriggerTableRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListPipelineTriggerTableRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListPipelineTriggerTableRecords(ctx, req.(*ListPipelineTriggerTableRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListPipelineTriggerChartRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineTriggerChartRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListPipelineTriggerChartRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListPipelineTriggerChartRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListPipelineTriggerChartRecords(ctx, req.(*ListPipelineTriggerChartRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListConnectorExecuteRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorExecuteRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListConnectorExecuteRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListConnectorExecuteRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListConnectorExecuteRecords(ctx, req.(*ListConnectorExecuteRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListConnectorExecuteTableRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorExecuteTableRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListConnectorExecuteTableRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListConnectorExecuteTableRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListConnectorExecuteTableRecords(ctx, req.(*ListConnectorExecuteTableRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MgmtPublicService_ListConnectorExecuteChartRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorExecuteChartRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MgmtPublicServiceServer).ListConnectorExecuteChartRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MgmtPublicService_ListConnectorExecuteChartRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MgmtPublicServiceServer).ListConnectorExecuteChartRecords(ctx, req.(*ListConnectorExecuteChartRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MgmtPublicService_ServiceDesc is the grpc.ServiceDesc for MgmtPublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MgmtPublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base.mgmt.v1alpha.MgmtPublicService",
	HandlerType: (*MgmtPublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _MgmtPublicService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _MgmtPublicService_Readiness_Handler,
		},
		{
			MethodName: "QueryAuthenticatedUser",
			Handler:    _MgmtPublicService_QueryAuthenticatedUser_Handler,
		},
		{
			MethodName: "PatchAuthenticatedUser",
			Handler:    _MgmtPublicService_PatchAuthenticatedUser_Handler,
		},
		{
			MethodName: "ExistUsername",
			Handler:    _MgmtPublicService_ExistUsername_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _MgmtPublicService_CreateToken_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _MgmtPublicService_ListTokens_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _MgmtPublicService_GetToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _MgmtPublicService_DeleteToken_Handler,
		},
		{
			MethodName: "ListPipelineTriggerRecords",
			Handler:    _MgmtPublicService_ListPipelineTriggerRecords_Handler,
		},
		{
			MethodName: "ListPipelineTriggerTableRecords",
			Handler:    _MgmtPublicService_ListPipelineTriggerTableRecords_Handler,
		},
		{
			MethodName: "ListPipelineTriggerChartRecords",
			Handler:    _MgmtPublicService_ListPipelineTriggerChartRecords_Handler,
		},
		{
			MethodName: "ListConnectorExecuteRecords",
			Handler:    _MgmtPublicService_ListConnectorExecuteRecords_Handler,
		},
		{
			MethodName: "ListConnectorExecuteTableRecords",
			Handler:    _MgmtPublicService_ListConnectorExecuteTableRecords_Handler,
		},
		{
			MethodName: "ListConnectorExecuteChartRecords",
			Handler:    _MgmtPublicService_ListConnectorExecuteChartRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/mgmt/v1alpha/mgmt_public_service.proto",
}