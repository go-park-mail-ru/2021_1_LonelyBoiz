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
	usecaseMocks "server/internal/pkg/user/usecase/mocks"
	user_proto "server/internal/user_server/delivery/proto"
	serverMocks "server/internal/user_server/delivery/proto/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSignIn(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("likes")
	if er != nil {
		t.Error(er)
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

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().CheckUser(ctx, gomock.Any()).Return(&user_proto.UserResponse{}, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().SetCookie(gomock.Any()).Return(http.Cookie{})
	userUseCaseMock.EXPECT().ProtoUser2User(gomock.Any()).Return(user)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SignIn(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLogIn_ParseToJson_Error(t *testing.T) {
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
	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, errors.New("some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SignIn(rw, req)

	response := rw.Result()

	assert.Equal(t, 401, response.StatusCode)
}

func TestSignIn_CheckUser_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("likes")
	if er != nil {
		t.Error(er)
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

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().CheckUser(ctx, gomock.Any()).Return(nil, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.SignIn(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
