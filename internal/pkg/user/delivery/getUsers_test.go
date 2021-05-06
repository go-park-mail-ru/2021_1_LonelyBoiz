package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
	mock_usecase "server/internal/pkg/user/usecase/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	q := req.URL.Query()
	q.Add("count", "20")
	req.URL.RawQuery = q.Encode()

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().CreateFeed(user.Id, 20).Return(nil, 200, nil)
	userUseCaseMock.EXPECT().LogInfo("Create Feed").Return()

	handlerTest.GetUsers(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetUsersGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	q := req.URL.Query()
	q.Add("count", "20")
	req.URL.RawQuery = q.Encode()

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetUsers(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestGetUsersNoQuery(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetUsers(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetUsersAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	q := req.URL.Query()
	q.Add("count", "not number")
	req.URL.RawQuery = q.Encode()

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetUsers(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetUsersCreateFeedError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	q := req.URL.Query()
	q.Add("count", "20")
	req.URL.RawQuery = q.Encode()

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().CreateFeed(user.Id, 20).Return(nil, 500, errors.New("Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetUsers(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
