// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/userUcase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	io "io"
	reflect "reflect"
	models "server/internal/pkg/models"

	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
)

// MockUserUseCaseInterface is a mock of UserUseCaseInterface interface.
type MockUserUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseInterfaceMockRecorder
}

// MockUserUseCaseInterfaceMockRecorder is the mock recorder for MockUserUseCaseInterface.
type MockUserUseCaseInterfaceMockRecorder struct {
	mock *MockUserUseCaseInterface
}

// NewMockUserUseCaseInterface creates a new mock instance.
func NewMockUserUseCaseInterface(ctrl *gomock.Controller) *MockUserUseCaseInterface {
	mock := &MockUserUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCaseInterface) EXPECT() *MockUserUseCaseInterfaceMockRecorder {
	return m.recorder
}

// AddNewUser mocks base method.
func (m *MockUserUseCaseInterface) AddNewUser(newUser *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewUser indicates an expected call of AddNewUser.
func (mr *MockUserUseCaseInterfaceMockRecorder) AddNewUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewUser", reflect.TypeOf((*MockUserUseCaseInterface)(nil).AddNewUser), newUser)
}

// ChangeUserInfo mocks base method.
func (m *MockUserUseCaseInterface) ChangeUserInfo(newUser models.User, id int) (models.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserInfo", newUser, id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ChangeUserInfo indicates an expected call of ChangeUserInfo.
func (mr *MockUserUseCaseInterfaceMockRecorder) ChangeUserInfo(newUser, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserInfo", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ChangeUserInfo), newUser, id)
}

// ChangeUserPassword mocks base method.
func (m *MockUserUseCaseInterface) ChangeUserPassword(newUser *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockUserUseCaseInterfaceMockRecorder) ChangeUserPassword(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ChangeUserPassword), newUser)
}

// ChangeUserProperties mocks base method.
func (m *MockUserUseCaseInterface) ChangeUserProperties(newUser *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserProperties", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserProperties indicates an expected call of ChangeUserProperties.
func (mr *MockUserUseCaseInterfaceMockRecorder) ChangeUserProperties(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserProperties", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ChangeUserProperties), newUser)
}

// CheckCaptch mocks base method.
func (m *MockUserUseCaseInterface) CheckCaptch(token string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCaptch", token)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCaptch indicates an expected call of CheckCaptch.
func (mr *MockUserUseCaseInterfaceMockRecorder) CheckCaptch(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCaptch", reflect.TypeOf((*MockUserUseCaseInterface)(nil).CheckCaptch), token)
}

// CheckPasswordWithEmail mocks base method.
func (m *MockUserUseCaseInterface) CheckPasswordWithEmail(passToCheck, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordWithEmail", passToCheck, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPasswordWithEmail indicates an expected call of CheckPasswordWithEmail.
func (mr *MockUserUseCaseInterfaceMockRecorder) CheckPasswordWithEmail(passToCheck, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordWithEmail", reflect.TypeOf((*MockUserUseCaseInterface)(nil).CheckPasswordWithEmail), passToCheck, email)
}

// CheckPasswordWithId mocks base method.
func (m *MockUserUseCaseInterface) CheckPasswordWithId(passToCheck string, id int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPasswordWithId", passToCheck, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPasswordWithId indicates an expected call of CheckPasswordWithId.
func (mr *MockUserUseCaseInterfaceMockRecorder) CheckPasswordWithId(passToCheck, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPasswordWithId", reflect.TypeOf((*MockUserUseCaseInterface)(nil).CheckPasswordWithId), passToCheck, id)
}

// CreateChat mocks base method.
func (m *MockUserUseCaseInterface) CreateChat(id int, like models.Like) (models.Chat, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChat", id, like)
	ret0, _ := ret[0].(models.Chat)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateChat indicates an expected call of CreateChat.
func (mr *MockUserUseCaseInterfaceMockRecorder) CreateChat(id, like interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChat", reflect.TypeOf((*MockUserUseCaseInterface)(nil).CreateChat), id, like)
}

// CreateFeed mocks base method.
func (m *MockUserUseCaseInterface) CreateFeed(id, limitInt int) ([]int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFeed", id, limitInt)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateFeed indicates an expected call of CreateFeed.
func (mr *MockUserUseCaseInterfaceMockRecorder) CreateFeed(id, limitInt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFeed", reflect.TypeOf((*MockUserUseCaseInterface)(nil).CreateFeed), id, limitInt)
}

// CreateNewUser mocks base method.
func (m *MockUserUseCaseInterface) CreateNewUser(newUser models.User) (models.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", newUser)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateNewUser indicates an expected call of CreateNewUser.
func (mr *MockUserUseCaseInterfaceMockRecorder) CreateNewUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockUserUseCaseInterface)(nil).CreateNewUser), newUser)
}

// DeleteUser mocks base method.
func (m *MockUserUseCaseInterface) DeleteUser(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserUseCaseInterfaceMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserUseCaseInterface)(nil).DeleteUser), id)
}

// GetUserInfoById mocks base method.
func (m *MockUserUseCaseInterface) GetUserInfoById(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoById", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoById indicates an expected call of GetUserInfoById.
func (mr *MockUserUseCaseInterfaceMockRecorder) GetUserInfoById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoById", reflect.TypeOf((*MockUserUseCaseInterface)(nil).GetUserInfoById), id)
}

// HashPassword mocks base method.
func (m *MockUserUseCaseInterface) HashPassword(pass string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", pass)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockUserUseCaseInterfaceMockRecorder) HashPassword(pass interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockUserUseCaseInterface)(nil).HashPassword), pass)
}

// IsAlreadySignedUp mocks base method.
func (m *MockUserUseCaseInterface) IsAlreadySignedUp(newEmail string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlreadySignedUp", newEmail)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAlreadySignedUp indicates an expected call of IsAlreadySignedUp.
func (mr *MockUserUseCaseInterfaceMockRecorder) IsAlreadySignedUp(newEmail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlreadySignedUp", reflect.TypeOf((*MockUserUseCaseInterface)(nil).IsAlreadySignedUp), newEmail)
}

// LogError mocks base method.
func (m *MockUserUseCaseInterface) LogError(data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogError", data)
}

// LogError indicates an expected call of LogError.
func (mr *MockUserUseCaseInterfaceMockRecorder) LogError(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockUserUseCaseInterface)(nil).LogError), data)
}

// LogInfo mocks base method.
func (m *MockUserUseCaseInterface) LogInfo(data interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogInfo", data)
}

// LogInfo indicates an expected call of LogInfo.
func (mr *MockUserUseCaseInterfaceMockRecorder) LogInfo(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogInfo", reflect.TypeOf((*MockUserUseCaseInterface)(nil).LogInfo), data)
}

// ParseJsonToUser mocks base method.
func (m *MockUserUseCaseInterface) ParseJsonToUser(body io.ReadCloser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseJsonToUser", body)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseJsonToUser indicates an expected call of ParseJsonToUser.
func (mr *MockUserUseCaseInterfaceMockRecorder) ParseJsonToUser(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseJsonToUser", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ParseJsonToUser), body)
}

// SetChat mocks base method.
func (m *MockUserUseCaseInterface) SetChat(ws *websocket.Conn, id int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetChat", ws, id)
}

// SetChat indicates an expected call of SetChat.
func (mr *MockUserUseCaseInterfaceMockRecorder) SetChat(ws, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChat", reflect.TypeOf((*MockUserUseCaseInterface)(nil).SetChat), ws, id)
}

// SignIn mocks base method.
func (m *MockUserUseCaseInterface) SignIn(user models.User) (models.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", user)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserUseCaseInterfaceMockRecorder) SignIn(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUserUseCaseInterface)(nil).SignIn), user)
}

// ValidateDatePreferences mocks base method.
func (m *MockUserUseCaseInterface) ValidateDatePreferences(pref string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDatePreferences", pref)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateDatePreferences indicates an expected call of ValidateDatePreferences.
func (mr *MockUserUseCaseInterfaceMockRecorder) ValidateDatePreferences(pref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDatePreferences", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ValidateDatePreferences), pref)
}

// ValidatePassword mocks base method.
func (m *MockUserUseCaseInterface) ValidatePassword(password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePassword", password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidatePassword indicates an expected call of ValidatePassword.
func (mr *MockUserUseCaseInterfaceMockRecorder) ValidatePassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePassword", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ValidatePassword), password)
}

// ValidateSex mocks base method.
func (m *MockUserUseCaseInterface) ValidateSex(sex string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSex", sex)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateSex indicates an expected call of ValidateSex.
func (mr *MockUserUseCaseInterfaceMockRecorder) ValidateSex(sex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSex", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ValidateSex), sex)
}

// ValidateSignInData mocks base method.
func (m *MockUserUseCaseInterface) ValidateSignInData(newUser models.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSignInData", newUser)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateSignInData indicates an expected call of ValidateSignInData.
func (mr *MockUserUseCaseInterfaceMockRecorder) ValidateSignInData(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSignInData", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ValidateSignInData), newUser)
}

// ValidateSignUpData mocks base method.
func (m *MockUserUseCaseInterface) ValidateSignUpData(newUser models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSignUpData", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateSignUpData indicates an expected call of ValidateSignUpData.
func (mr *MockUserUseCaseInterfaceMockRecorder) ValidateSignUpData(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSignUpData", reflect.TypeOf((*MockUserUseCaseInterface)(nil).ValidateSignUpData), newUser)
}

// WebsocketChat mocks base method.
func (m *MockUserUseCaseInterface) WebsocketChat(newChat *models.Chat) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WebsocketChat", newChat)
}

// WebsocketChat indicates an expected call of WebsocketChat.
func (mr *MockUserUseCaseInterfaceMockRecorder) WebsocketChat(newChat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WebsocketChat", reflect.TypeOf((*MockUserUseCaseInterface)(nil).WebsocketChat), newChat)
}

// IsActive mocks base method.
func (m *MockUserUseCaseInterface) IsActive(newUser *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsActive", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsActive indicates an expected call of IsActive.
func (mr *MockUserUseCaseInterfaceMockRecorder) IsActive(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockUserUseCaseInterface)(nil).IsActive), newUser)
}
