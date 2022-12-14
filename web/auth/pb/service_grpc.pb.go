// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: service.proto

package pb

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

// ProjectServerClient is the client API for ProjectServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServerClient interface {
	RegisterService(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	LoginService(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetUsersToFollow(ctx context.Context, in *GetUsersToFollowRequest, opts ...grpc.CallOption) (*GetUsersToFollowResponse, error)
	GetUsersToUnfollow(ctx context.Context, in *GetUsersToUnFollowRequest, opts ...grpc.CallOption) (*GetUsersToUnFollowResponse, error)
	StartFollowing(ctx context.Context, in *StartFollowingRequest, opts ...grpc.CallOption) (*StartFollowingResponse, error)
	StopFollowing(ctx context.Context, in *StopFollowingRequest, opts ...grpc.CallOption) (*StopFollowingResponse, error)
	Tweet(ctx context.Context, in *TweetRequest, opts ...grpc.CallOption) (*TweetResponse, error)
	FeedService(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
}

type projectServerClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServerClient(cc grpc.ClientConnInterface) ProjectServerClient {
	return &projectServerClient{cc}
}

func (c *projectServerClient) RegisterService(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/RegisterService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) LoginService(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/LoginService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) GetUsersToFollow(ctx context.Context, in *GetUsersToFollowRequest, opts ...grpc.CallOption) (*GetUsersToFollowResponse, error) {
	out := new(GetUsersToFollowResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/GetUsersToFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) GetUsersToUnfollow(ctx context.Context, in *GetUsersToUnFollowRequest, opts ...grpc.CallOption) (*GetUsersToUnFollowResponse, error) {
	out := new(GetUsersToUnFollowResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/GetUsersToUnfollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) StartFollowing(ctx context.Context, in *StartFollowingRequest, opts ...grpc.CallOption) (*StartFollowingResponse, error) {
	out := new(StartFollowingResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/StartFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) StopFollowing(ctx context.Context, in *StopFollowingRequest, opts ...grpc.CallOption) (*StopFollowingResponse, error) {
	out := new(StopFollowingResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/StopFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) Tweet(ctx context.Context, in *TweetRequest, opts ...grpc.CallOption) (*TweetResponse, error) {
	out := new(TweetResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/Tweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServerClient) FeedService(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/proto.ProjectServer/FeedService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServerServer is the server API for ProjectServer service.
// All implementations must embed UnimplementedProjectServerServer
// for forward compatibility
type ProjectServerServer interface {
	RegisterService(context.Context, *RegisterRequest) (*RegisterResponse, error)
	LoginService(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetUsersToFollow(context.Context, *GetUsersToFollowRequest) (*GetUsersToFollowResponse, error)
	GetUsersToUnfollow(context.Context, *GetUsersToUnFollowRequest) (*GetUsersToUnFollowResponse, error)
	StartFollowing(context.Context, *StartFollowingRequest) (*StartFollowingResponse, error)
	StopFollowing(context.Context, *StopFollowingRequest) (*StopFollowingResponse, error)
	Tweet(context.Context, *TweetRequest) (*TweetResponse, error)
	FeedService(context.Context, *FeedRequest) (*FeedResponse, error)
	mustEmbedUnimplementedProjectServerServer()
}

// UnimplementedProjectServerServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServerServer struct {
}

func (UnimplementedProjectServerServer) RegisterService(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterService not implemented")
}
func (UnimplementedProjectServerServer) LoginService(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginService not implemented")
}
func (UnimplementedProjectServerServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedProjectServerServer) GetUsersToFollow(context.Context, *GetUsersToFollowRequest) (*GetUsersToFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersToFollow not implemented")
}
func (UnimplementedProjectServerServer) GetUsersToUnfollow(context.Context, *GetUsersToUnFollowRequest) (*GetUsersToUnFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersToUnfollow not implemented")
}
func (UnimplementedProjectServerServer) StartFollowing(context.Context, *StartFollowingRequest) (*StartFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFollowing not implemented")
}
func (UnimplementedProjectServerServer) StopFollowing(context.Context, *StopFollowingRequest) (*StopFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopFollowing not implemented")
}
func (UnimplementedProjectServerServer) Tweet(context.Context, *TweetRequest) (*TweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tweet not implemented")
}
func (UnimplementedProjectServerServer) FeedService(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedService not implemented")
}
func (UnimplementedProjectServerServer) mustEmbedUnimplementedProjectServerServer() {}

// UnsafeProjectServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServerServer will
// result in compilation errors.
type UnsafeProjectServerServer interface {
	mustEmbedUnimplementedProjectServerServer()
}

func RegisterProjectServerServer(s grpc.ServiceRegistrar, srv ProjectServerServer) {
	s.RegisterService(&ProjectServer_ServiceDesc, srv)
}

func _ProjectServer_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).RegisterService(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_LoginService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).LoginService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/LoginService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).LoginService(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_GetUsersToFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersToFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).GetUsersToFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/GetUsersToFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).GetUsersToFollow(ctx, req.(*GetUsersToFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_GetUsersToUnfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersToUnFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).GetUsersToUnfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/GetUsersToUnfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).GetUsersToUnfollow(ctx, req.(*GetUsersToUnFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_StartFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).StartFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/StartFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).StartFollowing(ctx, req.(*StartFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_StopFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).StopFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/StopFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).StopFollowing(ctx, req.(*StopFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_Tweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).Tweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/Tweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).Tweet(ctx, req.(*TweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectServer_FeedService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServerServer).FeedService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProjectServer/FeedService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServerServer).FeedService(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectServer_ServiceDesc is the grpc.ServiceDesc for ProjectServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProjectServer",
	HandlerType: (*ProjectServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _ProjectServer_RegisterService_Handler,
		},
		{
			MethodName: "LoginService",
			Handler:    _ProjectServer_LoginService_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ProjectServer_Logout_Handler,
		},
		{
			MethodName: "GetUsersToFollow",
			Handler:    _ProjectServer_GetUsersToFollow_Handler,
		},
		{
			MethodName: "GetUsersToUnfollow",
			Handler:    _ProjectServer_GetUsersToUnfollow_Handler,
		},
		{
			MethodName: "StartFollowing",
			Handler:    _ProjectServer_StartFollowing_Handler,
		},
		{
			MethodName: "StopFollowing",
			Handler:    _ProjectServer_StopFollowing_Handler,
		},
		{
			MethodName: "Tweet",
			Handler:    _ProjectServer_Tweet_Handler,
		},
		{
			MethodName: "FeedService",
			Handler:    _ProjectServer_FeedService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
