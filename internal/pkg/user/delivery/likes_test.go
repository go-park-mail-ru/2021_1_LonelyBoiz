package delivery

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
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

func TestLikesHandler_ReadBody_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	errorJson := "some string"

	murl, er := url.Parse("/likes")
	if er != nil {
		t.Error(er)
	}

	json, err := json.Marshal(errorJson)
	if err != nil {
		t.Error(err)
	}

	req := &http.Request{
		Method: "GET",
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

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestLikesHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/likes")
	if er != nil {
		t.Error(er)
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	json, err := json.Marshal(like)
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

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().CreateChat(ctx, gomock.Any()).Return(&user_proto.Chat{}, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().ProtoPhotos2Photos(gomock.Any()).Return([]string{})
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().WebsocketChat(gomock.Any())

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLikesHandler_CreateCHat_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/likes")
	if er != nil {
		t.Error(er)
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	json, err := json.Marshal(like)
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

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().CreateChat(ctx, gomock.Any()).Return(&user_proto.Chat{}, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestLikesHandler_NotMutual(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/likes")
	if er != nil {
		t.Error(er)
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	json, err := json.Marshal(like)
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

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().CreateChat(ctx, gomock.Any()).Return(&user_proto.Chat{}, status.Error(codes.Code(204), "Success like"))
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}
