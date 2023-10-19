// Message API

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: message_api/v3/mls.proto

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MlsApi_PublishToGroup_FullMethodName       = "/xmtp.message_api.v3.MlsApi/PublishToGroup"
	MlsApi_RegisterInstallation_FullMethodName = "/xmtp.message_api.v3.MlsApi/RegisterInstallation"
	MlsApi_UploadKeyPackages_FullMethodName    = "/xmtp.message_api.v3.MlsApi/UploadKeyPackages"
	MlsApi_ConsumeKeyPackages_FullMethodName   = "/xmtp.message_api.v3.MlsApi/ConsumeKeyPackages"
	MlsApi_RevokeInstallation_FullMethodName   = "/xmtp.message_api.v3.MlsApi/RevokeInstallation"
	MlsApi_GetIdentityUpdates_FullMethodName   = "/xmtp.message_api.v3.MlsApi/GetIdentityUpdates"
)

// MlsApiClient is the client API for MlsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MlsApiClient interface {
	// Publish a MLS payload, that would be validated before being stored to the
	// network
	PublishToGroup(ctx context.Context, in *PublishToGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Register a new installation, which would be validated before storage
	RegisterInstallation(ctx context.Context, in *RegisterInstallationRequest, opts ...grpc.CallOption) (*RegisterInstallationResponse, error)
	// Upload one or more Key Packages, which would be validated before storage
	UploadKeyPackages(ctx context.Context, in *UploadKeyPackagesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get one or more Key Packages by installation_id
	ConsumeKeyPackages(ctx context.Context, in *ConsumeKeyPackagesRequest, opts ...grpc.CallOption) (*ConsumeKeyPackagesResponse, error)
	// Would delete all key packages associated with the installation and mark
	// the installation as having been revoked
	RevokeInstallation(ctx context.Context, in *RevokeInstallationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Used to check for changes related to members of a group.
	// Would return an array of any new installations associated with the wallet
	// address, and any revocations that have happened.
	GetIdentityUpdates(ctx context.Context, in *GetIdentityUpdatesRequest, opts ...grpc.CallOption) (*GetIdentityUpdatesResponse, error)
}

type mlsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMlsApiClient(cc grpc.ClientConnInterface) MlsApiClient {
	return &mlsApiClient{cc}
}

func (c *mlsApiClient) PublishToGroup(ctx context.Context, in *PublishToGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MlsApi_PublishToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mlsApiClient) RegisterInstallation(ctx context.Context, in *RegisterInstallationRequest, opts ...grpc.CallOption) (*RegisterInstallationResponse, error) {
	out := new(RegisterInstallationResponse)
	err := c.cc.Invoke(ctx, MlsApi_RegisterInstallation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mlsApiClient) UploadKeyPackages(ctx context.Context, in *UploadKeyPackagesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MlsApi_UploadKeyPackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mlsApiClient) ConsumeKeyPackages(ctx context.Context, in *ConsumeKeyPackagesRequest, opts ...grpc.CallOption) (*ConsumeKeyPackagesResponse, error) {
	out := new(ConsumeKeyPackagesResponse)
	err := c.cc.Invoke(ctx, MlsApi_ConsumeKeyPackages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mlsApiClient) RevokeInstallation(ctx context.Context, in *RevokeInstallationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MlsApi_RevokeInstallation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mlsApiClient) GetIdentityUpdates(ctx context.Context, in *GetIdentityUpdatesRequest, opts ...grpc.CallOption) (*GetIdentityUpdatesResponse, error) {
	out := new(GetIdentityUpdatesResponse)
	err := c.cc.Invoke(ctx, MlsApi_GetIdentityUpdates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MlsApiServer is the server API for MlsApi service.
// All implementations must embed UnimplementedMlsApiServer
// for forward compatibility
type MlsApiServer interface {
	// Publish a MLS payload, that would be validated before being stored to the
	// network
	PublishToGroup(context.Context, *PublishToGroupRequest) (*emptypb.Empty, error)
	// Register a new installation, which would be validated before storage
	RegisterInstallation(context.Context, *RegisterInstallationRequest) (*RegisterInstallationResponse, error)
	// Upload one or more Key Packages, which would be validated before storage
	UploadKeyPackages(context.Context, *UploadKeyPackagesRequest) (*emptypb.Empty, error)
	// Get one or more Key Packages by installation_id
	ConsumeKeyPackages(context.Context, *ConsumeKeyPackagesRequest) (*ConsumeKeyPackagesResponse, error)
	// Would delete all key packages associated with the installation and mark
	// the installation as having been revoked
	RevokeInstallation(context.Context, *RevokeInstallationRequest) (*emptypb.Empty, error)
	// Used to check for changes related to members of a group.
	// Would return an array of any new installations associated with the wallet
	// address, and any revocations that have happened.
	GetIdentityUpdates(context.Context, *GetIdentityUpdatesRequest) (*GetIdentityUpdatesResponse, error)
	mustEmbedUnimplementedMlsApiServer()
}

// UnimplementedMlsApiServer must be embedded to have forward compatible implementations.
type UnimplementedMlsApiServer struct {
}

func (UnimplementedMlsApiServer) PublishToGroup(context.Context, *PublishToGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishToGroup not implemented")
}
func (UnimplementedMlsApiServer) RegisterInstallation(context.Context, *RegisterInstallationRequest) (*RegisterInstallationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInstallation not implemented")
}
func (UnimplementedMlsApiServer) UploadKeyPackages(context.Context, *UploadKeyPackagesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadKeyPackages not implemented")
}
func (UnimplementedMlsApiServer) ConsumeKeyPackages(context.Context, *ConsumeKeyPackagesRequest) (*ConsumeKeyPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsumeKeyPackages not implemented")
}
func (UnimplementedMlsApiServer) RevokeInstallation(context.Context, *RevokeInstallationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeInstallation not implemented")
}
func (UnimplementedMlsApiServer) GetIdentityUpdates(context.Context, *GetIdentityUpdatesRequest) (*GetIdentityUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentityUpdates not implemented")
}
func (UnimplementedMlsApiServer) mustEmbedUnimplementedMlsApiServer() {}

// UnsafeMlsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MlsApiServer will
// result in compilation errors.
type UnsafeMlsApiServer interface {
	mustEmbedUnimplementedMlsApiServer()
}

func RegisterMlsApiServer(s grpc.ServiceRegistrar, srv MlsApiServer) {
	s.RegisterService(&MlsApi_ServiceDesc, srv)
}

func _MlsApi_PublishToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlsApiServer).PublishToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MlsApi_PublishToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlsApiServer).PublishToGroup(ctx, req.(*PublishToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MlsApi_RegisterInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInstallationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlsApiServer).RegisterInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MlsApi_RegisterInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlsApiServer).RegisterInstallation(ctx, req.(*RegisterInstallationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MlsApi_UploadKeyPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadKeyPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlsApiServer).UploadKeyPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MlsApi_UploadKeyPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlsApiServer).UploadKeyPackages(ctx, req.(*UploadKeyPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MlsApi_ConsumeKeyPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeKeyPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlsApiServer).ConsumeKeyPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MlsApi_ConsumeKeyPackages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlsApiServer).ConsumeKeyPackages(ctx, req.(*ConsumeKeyPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MlsApi_RevokeInstallation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeInstallationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlsApiServer).RevokeInstallation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MlsApi_RevokeInstallation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlsApiServer).RevokeInstallation(ctx, req.(*RevokeInstallationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MlsApi_GetIdentityUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIdentityUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MlsApiServer).GetIdentityUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MlsApi_GetIdentityUpdates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MlsApiServer).GetIdentityUpdates(ctx, req.(*GetIdentityUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MlsApi_ServiceDesc is the grpc.ServiceDesc for MlsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MlsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xmtp.message_api.v3.MlsApi",
	HandlerType: (*MlsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishToGroup",
			Handler:    _MlsApi_PublishToGroup_Handler,
		},
		{
			MethodName: "RegisterInstallation",
			Handler:    _MlsApi_RegisterInstallation_Handler,
		},
		{
			MethodName: "UploadKeyPackages",
			Handler:    _MlsApi_UploadKeyPackages_Handler,
		},
		{
			MethodName: "ConsumeKeyPackages",
			Handler:    _MlsApi_ConsumeKeyPackages_Handler,
		},
		{
			MethodName: "RevokeInstallation",
			Handler:    _MlsApi_RevokeInstallation_Handler,
		},
		{
			MethodName: "GetIdentityUpdates",
			Handler:    _MlsApi_GetIdentityUpdates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message_api/v3/mls.proto",
}
