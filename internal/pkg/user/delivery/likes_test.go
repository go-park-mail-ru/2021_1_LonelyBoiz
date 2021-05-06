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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLikesHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	newChat := models.Chat{
		ChatId:    1,
		PartnerId: 2,
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	json, err := json.Marshal(like)
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

	userUseCaseMock.EXPECT().CreateChat(1, like).Return(newChat, 200, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().WebsocketChat(gomock.Any())

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLikesHandlerGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	json, err := json.Marshal(like)
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

	assert.Equal(t, 403, response.StatusCode)
}

func TestLikesHandlerReadBodyError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	errorJson := "some string"

	murl, er := url.Parse("auth")
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

func TestLikesHandlerCreateNewChatError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	newChat := models.Chat{
		ChatId:    1,
		PartnerId: 2,
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	json, err := json.Marshal(like)
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

	userUseCaseMock.EXPECT().CreateChat(1, like).Return(newChat, 500, errors.New("Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestLikesHandlerEmptyResponse(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	like := models.Like{
		UserId:   2,
		Reaction: "like",
	}

	newChat := models.Chat{
		ChatId:    1,
		PartnerId: 2,
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	json, err := json.Marshal(like)
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

	userUseCaseMock.EXPECT().CreateChat(1, like).Return(newChat, 204, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.LikesHandler(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}
