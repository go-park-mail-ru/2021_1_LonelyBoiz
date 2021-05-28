// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_grpc.pb.go

// Package mock_user_proto is a generated GoMock package.
package mock_user_proto

import (
	context "context"
	reflect "reflect"
	user_proto "server/internal/user_server/delivery/proto"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// AddToSecreteAlbum mocks base method.
func (m *MockUserServiceClient) AddToSecreteAlbum(ctx context.Context, in *user_proto.User, opts ...grpc.CallOption) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddToSecreteAlbum", varargs...)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToSecreteAlbum indicates an expected call of AddToSecreteAlbum.
func (mr *MockUserServiceClientMockRecorder) AddToSecreteAlbum(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToSecreteAlbum", reflect.TypeOf((*MockUserServiceClient)(nil).AddToSecreteAlbum), varargs...)
}

// BlockSecretAlbum mocks base method.
func (m *MockUserServiceClient) BlockSecretAlbum(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BlockSecretAlbum", varargs...)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockSecretAlbum indicates an expected call of BlockSecretAlbum.
func (mr *MockUserServiceClientMockRecorder) BlockSecretAlbum(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSecretAlbum", reflect.TypeOf((*MockUserServiceClient)(nil).BlockSecretAlbum), varargs...)
}

// ChangeMessage mocks base method.
func (m *MockUserServiceClient) ChangeMessage(ctx context.Context, in *user_proto.Message, opts ...grpc.CallOption) (*user_proto.Message, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeMessage", varargs...)
	ret0, _ := ret[0].(*user_proto.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMessage indicates an expected call of ChangeMessage.
func (mr *MockUserServiceClientMockRecorder) ChangeMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessage", reflect.TypeOf((*MockUserServiceClient)(nil).ChangeMessage), varargs...)
}

// ChangeUser mocks base method.
func (m *MockUserServiceClient) ChangeUser(ctx context.Context, in *user_proto.User, opts ...grpc.CallOption) (*user_proto.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ChangeUser", varargs...)
	ret0, _ := ret[0].(*user_proto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeUser indicates an expected call of ChangeUser.
func (mr *MockUserServiceClientMockRecorder) ChangeUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUser", reflect.TypeOf((*MockUserServiceClient)(nil).ChangeUser), varargs...)
}

// CheckUser mocks base method.
func (m *MockUserServiceClient) CheckUser(ctx context.Context, in *user_proto.UserLogin, opts ...grpc.CallOption) (*user_proto.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUser", varargs...)
	ret0, _ := ret[0].(*user_proto.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockUserServiceClientMockRecorder) CheckUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockUserServiceClient)(nil).CheckUser), varargs...)
}

// CreateChat mocks base method.
func (m *MockUserServiceClient) CreateChat(ctx context.Context, in *user_proto.Like, opts ...grpc.CallOption) (*user_proto.Chat, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateChat", varargs...)
	ret0, _ := ret[0].(*user_proto.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockUserServiceClientMockRecorder) CreateChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockUserServiceClient)(nil).CreateChat), varargs...)
}

// CreateFeed mocks base method.
func (m *MockUserServiceClient) CreateFeed(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.Feed, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateFeed", varargs...)
	ret0, _ := ret[0].(*user_proto.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeed indicates an expected call of CreateFeed.
func (mr *MockUserServiceClientMockRecorder) CreateFeed(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeed", reflect.TypeOf((*MockUserServiceClient)(nil).CreateFeed), varargs...)
}

// CreateMessage mocks base method.
func (m *MockUserServiceClient) CreateMessage(ctx context.Context, in *user_proto.Message, opts ...grpc.CallOption) (*user_proto.Message, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMessage", varargs...)
	ret0, _ := ret[0].(*user_proto.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockUserServiceClientMockRecorder) CreateMessage(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockUserServiceClient)(nil).CreateMessage), varargs...)
}

// CreateUser mocks base method.
func (m *MockUserServiceClient) CreateUser(ctx context.Context, in *user_proto.User, opts ...grpc.CallOption) (*user_proto.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateUser", varargs...)
	ret0, _ := ret[0].(*user_proto.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceClientMockRecorder) CreateUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceClient)(nil).CreateUser), varargs...)
}

// DeleteChat mocks base method.
func (m *MockUserServiceClient) DeleteChat(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteChat", varargs...)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChat indicates an expected call of DeleteChat.
func (mr *MockUserServiceClientMockRecorder) DeleteChat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChat", reflect.TypeOf((*MockUserServiceClient)(nil).DeleteChat), varargs...)
}

// DeleteUser mocks base method.
func (m *MockUserServiceClient) DeleteUser(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteUser", varargs...)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceClientMockRecorder) DeleteUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServiceClient)(nil).DeleteUser), varargs...)
}

// GetChats mocks base method.
func (m *MockUserServiceClient) GetChats(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.ChatsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetChats", varargs...)
	ret0, _ := ret[0].(*user_proto.ChatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChats indicates an expected call of GetChats.
func (mr *MockUserServiceClientMockRecorder) GetChats(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChats", reflect.TypeOf((*MockUserServiceClient)(nil).GetChats), varargs...)
}

// GetMessages mocks base method.
func (m *MockUserServiceClient) GetMessages(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.MessagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMessages", varargs...)
	ret0, _ := ret[0].(*user_proto.MessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockUserServiceClientMockRecorder) GetMessages(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockUserServiceClient)(nil).GetMessages), varargs...)
}

// GetSecreteAlbum mocks base method.
func (m *MockUserServiceClient) GetSecreteAlbum(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.Photos, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecreteAlbum", varargs...)
	ret0, _ := ret[0].(*user_proto.Photos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecreteAlbum indicates an expected call of GetSecreteAlbum.
func (mr *MockUserServiceClientMockRecorder) GetSecreteAlbum(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecreteAlbum", reflect.TypeOf((*MockUserServiceClient)(nil).GetSecreteAlbum), varargs...)
}

// GetUserById mocks base method.
func (m *MockUserServiceClient) GetUserById(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserById", varargs...)
	ret0, _ := ret[0].(*user_proto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceClientMockRecorder) GetUserById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserServiceClient)(nil).GetUserById), varargs...)
}

// UnlockSecretAlbum mocks base method.
func (m *MockUserServiceClient) UnlockSecretAlbum(ctx context.Context, in *user_proto.UserNothing, opts ...grpc.CallOption) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnlockSecretAlbum", varargs...)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlockSecretAlbum indicates an expected call of UnlockSecretAlbum.
func (mr *MockUserServiceClientMockRecorder) UnlockSecretAlbum(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockSecretAlbum", reflect.TypeOf((*MockUserServiceClient)(nil).UnlockSecretAlbum), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// AddToSecreteAlbum mocks base method.
func (m *MockUserServiceServer) AddToSecreteAlbum(arg0 context.Context, arg1 *user_proto.User) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToSecreteAlbum", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToSecreteAlbum indicates an expected call of AddToSecreteAlbum.
func (mr *MockUserServiceServerMockRecorder) AddToSecreteAlbum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToSecreteAlbum", reflect.TypeOf((*MockUserServiceServer)(nil).AddToSecreteAlbum), arg0, arg1)
}

// BlockSecretAlbum mocks base method.
func (m *MockUserServiceServer) BlockSecretAlbum(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockSecretAlbum", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockSecretAlbum indicates an expected call of BlockSecretAlbum.
func (mr *MockUserServiceServerMockRecorder) BlockSecretAlbum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockSecretAlbum", reflect.TypeOf((*MockUserServiceServer)(nil).BlockSecretAlbum), arg0, arg1)
}

// ChangeMessage mocks base method.
func (m *MockUserServiceServer) ChangeMessage(arg0 context.Context, arg1 *user_proto.Message) (*user_proto.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeMessage", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeMessage indicates an expected call of ChangeMessage.
func (mr *MockUserServiceServerMockRecorder) ChangeMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeMessage", reflect.TypeOf((*MockUserServiceServer)(nil).ChangeMessage), arg0, arg1)
}

// ChangeUser mocks base method.
func (m *MockUserServiceServer) ChangeUser(arg0 context.Context, arg1 *user_proto.User) (*user_proto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUser", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangeUser indicates an expected call of ChangeUser.
func (mr *MockUserServiceServerMockRecorder) ChangeUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUser", reflect.TypeOf((*MockUserServiceServer)(nil).ChangeUser), arg0, arg1)
}

// CheckUser mocks base method.
func (m *MockUserServiceServer) CheckUser(arg0 context.Context, arg1 *user_proto.UserLogin) (*user_proto.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockUserServiceServerMockRecorder) CheckUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockUserServiceServer)(nil).CheckUser), arg0, arg1)
}

// CreateChat mocks base method.
func (m *MockUserServiceServer) CreateChat(arg0 context.Context, arg1 *user_proto.Like) (*user_proto.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockUserServiceServerMockRecorder) CreateChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockUserServiceServer)(nil).CreateChat), arg0, arg1)
}

// CreateFeed mocks base method.
func (m *MockUserServiceServer) CreateFeed(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeed", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFeed indicates an expected call of CreateFeed.
func (mr *MockUserServiceServerMockRecorder) CreateFeed(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeed", reflect.TypeOf((*MockUserServiceServer)(nil).CreateFeed), arg0, arg1)
}

// CreateMessage mocks base method.
func (m *MockUserServiceServer) CreateMessage(arg0 context.Context, arg1 *user_proto.Message) (*user_proto.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockUserServiceServerMockRecorder) CreateMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockUserServiceServer)(nil).CreateMessage), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockUserServiceServer) CreateUser(arg0 context.Context, arg1 *user_proto.User) (*user_proto.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceServerMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceServer)(nil).CreateUser), arg0, arg1)
}

// DeleteChat mocks base method.
func (m *MockUserServiceServer) DeleteChat(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChat", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteChat indicates an expected call of DeleteChat.
func (mr *MockUserServiceServerMockRecorder) DeleteChat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChat", reflect.TypeOf((*MockUserServiceServer)(nil).DeleteChat), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockUserServiceServer) DeleteUser(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceServerMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServiceServer)(nil).DeleteUser), arg0, arg1)
}

// GetChats mocks base method.
func (m *MockUserServiceServer) GetChats(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.ChatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChats", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.ChatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChats indicates an expected call of GetChats.
func (mr *MockUserServiceServerMockRecorder) GetChats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChats", reflect.TypeOf((*MockUserServiceServer)(nil).GetChats), arg0, arg1)
}

// GetMessages mocks base method.
func (m *MockUserServiceServer) GetMessages(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.MessagesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.MessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockUserServiceServerMockRecorder) GetMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockUserServiceServer)(nil).GetMessages), arg0, arg1)
}

// GetSecreteAlbum mocks base method.
func (m *MockUserServiceServer) GetSecreteAlbum(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.Photos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecreteAlbum", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.Photos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecreteAlbum indicates an expected call of GetSecreteAlbum.
func (mr *MockUserServiceServerMockRecorder) GetSecreteAlbum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecreteAlbum", reflect.TypeOf((*MockUserServiceServer)(nil).GetSecreteAlbum), arg0, arg1)
}

// GetUserById mocks base method.
func (m *MockUserServiceServer) GetUserById(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserServiceServerMockRecorder) GetUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUserServiceServer)(nil).GetUserById), arg0, arg1)
}

// UnlockSecretAlbum mocks base method.
func (m *MockUserServiceServer) UnlockSecretAlbum(arg0 context.Context, arg1 *user_proto.UserNothing) (*user_proto.UserNothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockSecretAlbum", arg0, arg1)
	ret0, _ := ret[0].(*user_proto.UserNothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnlockSecretAlbum indicates an expected call of UnlockSecretAlbum.
func (mr *MockUserServiceServerMockRecorder) UnlockSecretAlbum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockSecretAlbum", reflect.TypeOf((*MockUserServiceServer)(nil).UnlockSecretAlbum), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}
