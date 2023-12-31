// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: wg.proto

package protobuf

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
	Wireguard_AddUser_FullMethodName     = "/wireguard.Wireguard/AddUser"
	Wireguard_DeleteUser_FullMethodName  = "/wireguard.Wireguard/DeleteUser"
	Wireguard_AllowAccess_FullMethodName = "/wireguard.Wireguard/AllowAccess"
	Wireguard_DenyAccess_FullMethodName  = "/wireguard.Wireguard/DenyAccess"
)

// WireguardClient is the client API for Wireguard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WireguardClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	AllowAccess(ctx context.Context, in *UserAllowAccessRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DenyAccess(ctx context.Context, in *UserDenyAccessRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type wireguardClient struct {
	cc grpc.ClientConnInterface
}

func NewWireguardClient(cc grpc.ClientConnInterface) WireguardClient {
	return &wireguardClient{cc}
}

func (c *wireguardClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Wireguard_AddUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Wireguard_DeleteUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) AllowAccess(ctx context.Context, in *UserAllowAccessRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Wireguard_AllowAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) DenyAccess(ctx context.Context, in *UserDenyAccessRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, Wireguard_DenyAccess_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WireguardServer is the server API for Wireguard service.
// All implementations must embed UnimplementedWireguardServer
// for forward compatibility
type WireguardServer interface {
	AddUser(context.Context, *AddUserRequest) (*EmptyResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*EmptyResponse, error)
	AllowAccess(context.Context, *UserAllowAccessRequest) (*EmptyResponse, error)
	DenyAccess(context.Context, *UserDenyAccessRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedWireguardServer()
}

// UnimplementedWireguardServer must be embedded to have forward compatible implementations.
type UnimplementedWireguardServer struct {
}

func (UnimplementedWireguardServer) AddUser(context.Context, *AddUserRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedWireguardServer) DeleteUser(context.Context, *DeleteUserRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedWireguardServer) AllowAccess(context.Context, *UserAllowAccessRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowAccess not implemented")
}
func (UnimplementedWireguardServer) DenyAccess(context.Context, *UserDenyAccessRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyAccess not implemented")
}
func (UnimplementedWireguardServer) mustEmbedUnimplementedWireguardServer() {}

// UnsafeWireguardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WireguardServer will
// result in compilation errors.
type UnsafeWireguardServer interface {
	mustEmbedUnimplementedWireguardServer()
}

func RegisterWireguardServer(s grpc.ServiceRegistrar, srv WireguardServer) {
	s.RegisterService(&Wireguard_ServiceDesc, srv)
}

func _Wireguard_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wireguard_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wireguard_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_AllowAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAllowAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).AllowAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wireguard_AllowAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).AllowAccess(ctx, req.(*UserAllowAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_DenyAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDenyAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).DenyAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wireguard_DenyAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).DenyAccess(ctx, req.(*UserDenyAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wireguard_ServiceDesc is the grpc.ServiceDesc for Wireguard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wireguard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wireguard.Wireguard",
	HandlerType: (*WireguardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Wireguard_AddUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Wireguard_DeleteUser_Handler,
		},
		{
			MethodName: "AllowAccess",
			Handler:    _Wireguard_AllowAccess_Handler,
		},
		{
			MethodName: "DenyAccess",
			Handler:    _Wireguard_DenyAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wg.proto",
}
