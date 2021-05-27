// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user_proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error)
	ChangeUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	CheckUser(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteChat(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*UserNothing, error)
	DeleteUser(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*UserNothing, error)
	GetUserById(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*User, error)
	CreateFeed(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*Feed, error)
	CreateChat(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Chat, error)
	//chat
	GetChats(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*ChatsResponse, error)
	//message
	GetMessages(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*MessagesResponse, error)
	CreateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ChangeMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	// secret album
	AddToSecreteAlbum(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserNothing, error)
	UnlockSecretAlbum(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*UserNothing, error)
	GetSecreteAlbum(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*Photos, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/session.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/session.UserService/ChangeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckUser(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/session.UserService/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteChat(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*UserNothing, error) {
	out := new(UserNothing)
	err := c.cc.Invoke(ctx, "/session.UserService/DeleteChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*UserNothing, error) {
	out := new(UserNothing)
	err := c.cc.Invoke(ctx, "/session.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/session.UserService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateFeed(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*Feed, error) {
	out := new(Feed)
	err := c.cc.Invoke(ctx, "/session.UserService/CreateFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateChat(ctx context.Context, in *Like, opts ...grpc.CallOption) (*Chat, error) {
	out := new(Chat)
	err := c.cc.Invoke(ctx, "/session.UserService/CreateChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetChats(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*ChatsResponse, error) {
	out := new(ChatsResponse)
	err := c.cc.Invoke(ctx, "/session.UserService/GetChats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMessages(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*MessagesResponse, error) {
	out := new(MessagesResponse)
	err := c.cc.Invoke(ctx, "/session.UserService/GetMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/session.UserService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/session.UserService/ChangeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddToSecreteAlbum(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserNothing, error) {
	out := new(UserNothing)
	err := c.cc.Invoke(ctx, "/session.UserService/AddToSecreteAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnlockSecretAlbum(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*UserNothing, error) {
	out := new(UserNothing)
	err := c.cc.Invoke(ctx, "/session.UserService/UnlockSecretAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSecreteAlbum(ctx context.Context, in *UserNothing, opts ...grpc.CallOption) (*Photos, error) {
	out := new(Photos)
	err := c.cc.Invoke(ctx, "/session.UserService/GetSecreteAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*UserResponse, error)
	ChangeUser(context.Context, *User) (*User, error)
	CheckUser(context.Context, *UserLogin) (*UserResponse, error)
	DeleteChat(context.Context, *UserNothing) (*UserNothing, error)
	DeleteUser(context.Context, *UserNothing) (*UserNothing, error)
	GetUserById(context.Context, *UserNothing) (*User, error)
	CreateFeed(context.Context, *UserNothing) (*Feed, error)
	CreateChat(context.Context, *Like) (*Chat, error)
	//chat
	GetChats(context.Context, *UserNothing) (*ChatsResponse, error)
	//message
	GetMessages(context.Context, *UserNothing) (*MessagesResponse, error)
	CreateMessage(context.Context, *Message) (*Message, error)
	ChangeMessage(context.Context, *Message) (*Message, error)
	// secret album
	AddToSecreteAlbum(context.Context, *User) (*UserNothing, error)
	UnlockSecretAlbum(context.Context, *UserNothing) (*UserNothing, error)
	GetSecreteAlbum(context.Context, *UserNothing) (*Photos, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ChangeUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUser not implemented")
}
func (UnimplementedUserServiceServer) CheckUser(context.Context, *UserLogin) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteChat(context.Context, *UserNothing) (*UserNothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChat not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *UserNothing) (*UserNothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserById(context.Context, *UserNothing) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) CreateFeed(context.Context, *UserNothing) (*Feed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeed not implemented")
}
func (UnimplementedUserServiceServer) CreateChat(context.Context, *Like) (*Chat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedUserServiceServer) GetChats(context.Context, *UserNothing) (*ChatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChats not implemented")
}
func (UnimplementedUserServiceServer) GetMessages(context.Context, *UserNothing) (*MessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}
func (UnimplementedUserServiceServer) CreateMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedUserServiceServer) ChangeMessage(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeMessage not implemented")
}
func (UnimplementedUserServiceServer) AddToSecreteAlbum(context.Context, *User) (*UserNothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToSecreteAlbum not implemented")
}
func (UnimplementedUserServiceServer) UnlockSecretAlbum(context.Context, *UserNothing) (*UserNothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockSecretAlbum not implemented")
}
func (UnimplementedUserServiceServer) GetSecreteAlbum(context.Context, *UserNothing) (*Photos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecreteAlbum not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/ChangeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckUser(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/DeleteChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteChat(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/CreateFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateFeed(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/CreateChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateChat(ctx, req.(*Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetChats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetChats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/GetChats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetChats(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/GetMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMessages(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/ChangeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddToSecreteAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddToSecreteAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/AddToSecreteAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddToSecreteAlbum(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnlockSecretAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnlockSecretAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/UnlockSecretAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnlockSecretAlbum(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSecreteAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNothing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSecreteAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.UserService/GetSecreteAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSecreteAlbum(ctx, req.(*UserNothing))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ChangeUser",
			Handler:    _UserService_ChangeUser_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _UserService_CheckUser_Handler,
		},
		{
			MethodName: "DeleteChat",
			Handler:    _UserService_DeleteChat_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "CreateFeed",
			Handler:    _UserService_CreateFeed_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _UserService_CreateChat_Handler,
		},
		{
			MethodName: "GetChats",
			Handler:    _UserService_GetChats_Handler,
		},
		{
			MethodName: "GetMessages",
			Handler:    _UserService_GetMessages_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _UserService_CreateMessage_Handler,
		},
		{
			MethodName: "ChangeMessage",
			Handler:    _UserService_ChangeMessage_Handler,
		},
		{
			MethodName: "AddToSecreteAlbum",
			Handler:    _UserService_AddToSecreteAlbum_Handler,
		},
		{
			MethodName: "UnlockSecretAlbum",
			Handler:    _UserService_UnlockSecretAlbum_Handler,
		},
		{
			MethodName: "GetSecreteAlbum",
			Handler:    _UserService_GetSecreteAlbum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
