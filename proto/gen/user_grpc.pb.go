// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user.proto

package gen

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
	UserService_CloseUserAccount_FullMethodName = "/ffauzann.cai.user.UserService/CloseUserAccount"
	UserService_UpdatePasscode_FullMethodName   = "/ffauzann.cai.user.UserService/UpdatePasscode"
	UserService_IsUserExist_FullMethodName      = "/ffauzann.cai.user.UserService/IsUserExist"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CloseUserAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CloseUserAccountResponse, error)
	UpdatePasscode(ctx context.Context, in *UpdatePasscodeRequest, opts ...grpc.CallOption) (*UpdatePasscodeResponse, error)
	IsUserExist(ctx context.Context, in *IsUserExistRequest, opts ...grpc.CallOption) (*IsUserExistResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CloseUserAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CloseUserAccountResponse, error) {
	out := new(CloseUserAccountResponse)
	err := c.cc.Invoke(ctx, UserService_CloseUserAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePasscode(ctx context.Context, in *UpdatePasscodeRequest, opts ...grpc.CallOption) (*UpdatePasscodeResponse, error) {
	out := new(UpdatePasscodeResponse)
	err := c.cc.Invoke(ctx, UserService_UpdatePasscode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) IsUserExist(ctx context.Context, in *IsUserExistRequest, opts ...grpc.CallOption) (*IsUserExistResponse, error) {
	out := new(IsUserExistResponse)
	err := c.cc.Invoke(ctx, UserService_IsUserExist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CloseUserAccount(context.Context, *emptypb.Empty) (*CloseUserAccountResponse, error)
	UpdatePasscode(context.Context, *UpdatePasscodeRequest) (*UpdatePasscodeResponse, error)
	IsUserExist(context.Context, *IsUserExistRequest) (*IsUserExistResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CloseUserAccount(context.Context, *emptypb.Empty) (*CloseUserAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseUserAccount not implemented")
}
func (UnimplementedUserServiceServer) UpdatePasscode(context.Context, *UpdatePasscodeRequest) (*UpdatePasscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePasscode not implemented")
}
func (UnimplementedUserServiceServer) IsUserExist(context.Context, *IsUserExistRequest) (*IsUserExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserExist not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CloseUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CloseUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CloseUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CloseUserAccount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePasscode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePasscode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdatePasscode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePasscode(ctx, req.(*UpdatePasscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_IsUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsUserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).IsUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_IsUserExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).IsUserExist(ctx, req.(*IsUserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ffauzann.cai.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CloseUserAccount",
			Handler:    _UserService_CloseUserAccount_Handler,
		},
		{
			MethodName: "UpdatePasscode",
			Handler:    _UserService_UpdatePasscode_Handler,
		},
		{
			MethodName: "IsUserExist",
			Handler:    _UserService_IsUserExist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
