package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	mockUsecase "server/internal/pkg/chat/usecase/mocks"
	"server/internal/pkg/models"
	sessionMocks "server/internal/pkg/session/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetChats(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	//sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("offset", "20")
	req.URL.RawQuery = q.Encode()

	userId := 1
	//limit := 20
	//offset := 20

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	//sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)

	//chatUseCaseMock.EXPECT().GetChat(userId, limit, offset).Return([]models.Chat{}, nil)
	chatUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetChatsGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("offset", "20")
	req.URL.RawQuery = q.Encode()

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, false)
	chatUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestGetChatsCountError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("notCount", "20")
	q.Add("offset", "20")
	req.URL.RawQuery = q.Encode()

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	chatUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetChatsLimitAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("count", "notInt")
	q.Add("offset", "20")
	req.URL.RawQuery = q.Encode()

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	chatUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetChatsOffset(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("notOffset", "20")
	req.URL.RawQuery = q.Encode()

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	chatUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetChatsOffsetAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("offset", "notInt")
	req.URL.RawQuery = q.Encode()

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	chatUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetChatsGetChatError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	q := req.URL.Query()
	q.Add("count", "20")
	q.Add("offset", "20")
	req.URL.RawQuery = q.Encode()

	userId := 1
	limit := 20
	offset := 20

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	chatUseCaseMock.EXPECT().GetChat(userId, limit, offset).Return([]models.Chat{}, errors.New("Some error"))
	chatUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
