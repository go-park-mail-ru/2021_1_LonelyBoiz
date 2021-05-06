// Code generated by MockGen. DO NOT EDIT.
// Source: ./users.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	models "server/internal/pkg/models"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepositoryInterface is a mock of UserRepositoryInterface interface.
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryInterfaceMockRecorder
}

// MockUserRepositoryInterfaceMockRecorder is the mock recorder for MockUserRepositoryInterface.
type MockUserRepositoryInterfaceMockRecorder struct {
	mock *MockUserRepositoryInterface
}

// NewMockUserRepositoryInterface creates a new mock instance.
func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryInterface) EXPECT() *MockUserRepositoryInterfaceMockRecorder {
	return m.recorder
}

// AddToSecreteAlbum mocks base method.
func (m *MockUserRepositoryInterface) AddToSecreteAlbum(ownerId int, photos []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToSecreteAlbum", ownerId, photos)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToSecreteAlbum indicates an expected call of AddToSecreteAlbum.
func (mr *MockUserRepositoryInterfaceMockRecorder) AddToSecreteAlbum(ownerId, photos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToSecreteAlbum", reflect.TypeOf((*MockUserRepositoryInterface)(nil).AddToSecreteAlbum), ownerId, photos)
}

// AddUser mocks base method.
func (m *MockUserRepositoryInterface) AddUser(newUser models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", newUser)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) AddUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).AddUser), newUser)
}

// ChangePassword mocks base method.
func (m *MockUserRepositoryInterface) ChangePassword(userId int, hash []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", userId, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserRepositoryInterfaceMockRecorder) ChangePassword(userId, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserRepositoryInterface)(nil).ChangePassword), userId, hash)
}

// ChangeUser mocks base method.
func (m *MockUserRepositoryInterface) ChangeUser(newUser models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUser indicates an expected call of ChangeUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) ChangeUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).ChangeUser), newUser)
}

// CheckMail mocks base method.
func (m *MockUserRepositoryInterface) CheckMail(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckMail", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckMail indicates an expected call of CheckMail.
func (mr *MockUserRepositoryInterfaceMockRecorder) CheckMail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckMail", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CheckMail), email)
}

// CheckPermission mocks base method.
func (m *MockUserRepositoryInterface) CheckPermission(ownerId, getterId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPermission", ownerId, getterId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPermission indicates an expected call of CheckPermission.
func (mr *MockUserRepositoryInterfaceMockRecorder) CheckPermission(ownerId, getterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPermission", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CheckPermission), ownerId, getterId)
}

// CheckReciprocity mocks base method.
func (m *MockUserRepositoryInterface) CheckReciprocity(userId1, userId2 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReciprocity", userId1, userId2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckReciprocity indicates an expected call of CheckReciprocity.
func (mr *MockUserRepositoryInterfaceMockRecorder) CheckReciprocity(userId1, userId2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReciprocity", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CheckReciprocity), userId1, userId2)
}

// ClearFeed mocks base method.
func (m *MockUserRepositoryInterface) ClearFeed(userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearFeed", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearFeed indicates an expected call of ClearFeed.
func (mr *MockUserRepositoryInterfaceMockRecorder) ClearFeed(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearFeed", reflect.TypeOf((*MockUserRepositoryInterface)(nil).ClearFeed), userId)
}

// CreateChat mocks base method.
func (m *MockUserRepositoryInterface) CreateChat(userId1, userId2 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", userId1, userId2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateChat(userId1, userId2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateChat), userId1, userId2)
}

// CreateFeed mocks base method.
func (m *MockUserRepositoryInterface) CreateFeed(userId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeed", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateFeed indicates an expected call of CreateFeed.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateFeed(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeed", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateFeed), userId)
}

// CreateSecretAlbum mocks base method.
func (m *MockUserRepositoryInterface) CreateSecretAlbum(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecretAlbum", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSecretAlbum indicates an expected call of CreateSecretAlbum.
func (mr *MockUserRepositoryInterfaceMockRecorder) CreateSecretAlbum(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecretAlbum", reflect.TypeOf((*MockUserRepositoryInterface)(nil).CreateSecretAlbum), id)
}

// DeleteUser mocks base method.
func (m *MockUserRepositoryInterface) DeleteUser(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).DeleteUser), id)
}

// GetChatById mocks base method.
func (m *MockUserRepositoryInterface) GetChatById(chatId, userId int) (models.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChatById", chatId, userId)
	ret0, _ := ret[0].(models.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChatById indicates an expected call of GetChatById.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetChatById(chatId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChatById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetChatById), chatId, userId)
}

// GetFeed mocks base method.
func (m *MockUserRepositoryInterface) GetFeed(userId, limit int) ([]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeed", userId, limit)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeed indicates an expected call of GetFeed.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetFeed(userId, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetFeed), userId, limit)
}

// GetNewChatById mocks base method.
func (m *MockUserRepositoryInterface) GetNewChatById(chatId, userId int) (models.Chat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewChatById", chatId, userId)
	ret0, _ := ret[0].(models.Chat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewChatById indicates an expected call of GetNewChatById.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetNewChatById(chatId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewChatById", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetNewChatById), chatId, userId)
}

// GetPassWithEmail mocks base method.
func (m *MockUserRepositoryInterface) GetPassWithEmail(email string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassWithEmail", email)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassWithEmail indicates an expected call of GetPassWithEmail.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetPassWithEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassWithEmail", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetPassWithEmail), email)
}

// GetPassWithId mocks base method.
func (m *MockUserRepositoryInterface) GetPassWithId(id int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassWithId", id)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassWithId indicates an expected call of GetPassWithId.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetPassWithId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassWithId", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetPassWithId), id)
}

// GetPhotos mocks base method.
func (m *MockUserRepositoryInterface) GetPhotos(userId int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhotos", userId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhotos indicates an expected call of GetPhotos.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetPhotos(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhotos", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetPhotos), userId)
}

// GetSecretePhotos mocks base method.
func (m *MockUserRepositoryInterface) GetSecretePhotos(ownerId int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretePhotos", ownerId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretePhotos indicates an expected call of GetSecretePhotos.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetSecretePhotos(ownerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretePhotos", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetSecretePhotos), ownerId)
}

// GetUser mocks base method.
func (m *MockUserRepositoryInterface) GetUser(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryInterfaceMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepositoryInterface)(nil).GetUser), id)
}

// Rating mocks base method.
func (m *MockUserRepositoryInterface) Rating(userIdFrom, userIdTo int, reaction string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rating", userIdFrom, userIdTo, reaction)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rating indicates an expected call of Rating.
func (mr *MockUserRepositoryInterfaceMockRecorder) Rating(userIdFrom, userIdTo, reaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rating", reflect.TypeOf((*MockUserRepositoryInterface)(nil).Rating), userIdFrom, userIdTo, reaction)
}

// ReduceScrolls mocks base method.
func (m *MockUserRepositoryInterface) ReduceScrolls(userId int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReduceScrolls", userId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReduceScrolls indicates an expected call of ReduceScrolls.
func (mr *MockUserRepositoryInterfaceMockRecorder) ReduceScrolls(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceScrolls", reflect.TypeOf((*MockUserRepositoryInterface)(nil).ReduceScrolls), userId)
}

// SignIn mocks base method.
func (m *MockUserRepositoryInterface) SignIn(email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserRepositoryInterfaceMockRecorder) SignIn(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserRepositoryInterface)(nil).SignIn), email)
}

// UnblockSecreteAlbum mocks base method.
func (m *MockUserRepositoryInterface) UnblockSecreteAlbum(ownerId, getterId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnblockSecreteAlbum", ownerId, getterId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnblockSecreteAlbum indicates an expected call of UnblockSecreteAlbum.
func (mr *MockUserRepositoryInterfaceMockRecorder) UnblockSecreteAlbum(ownerId, getterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnblockSecreteAlbum", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UnblockSecreteAlbum), ownerId, getterId)
}

// UpdatePayment mocks base method.
func (m *MockUserRepositoryInterface) UpdatePayment(userId, amount int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePayment", userId, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePayment indicates an expected call of UpdatePayment.
func (mr *MockUserRepositoryInterfaceMockRecorder) UpdatePayment(userId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePayment", reflect.TypeOf((*MockUserRepositoryInterface)(nil).UpdatePayment), userId, amount)
}