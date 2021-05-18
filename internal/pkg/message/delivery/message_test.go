package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	mockUsecase "server/internal/pkg/message/usecase/mocks"
	"server/internal/pkg/models"
	serverMocks "server/internal/user_server/delivery/proto/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
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

	serverMock.EXPECT().GetMessages(ctx, gomock.Any()).Return(nil, nil)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetMessagesError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
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

	serverMock.EXPECT().GetMessages(ctx, gomock.Any()).Return(nil, errors.New("some error"))
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 2, response.StatusCode)
}

func TestSendMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  2,
		Reaction:  1,
		ChatId:    1,
	}

	messageUseCaseMock.EXPECT().ParseJsonToMessage(gomock.Any()).Return(message, nil)
	messageUseCaseMock.EXPECT().Message2ProtoMessage(message).Return(nil)
	serverMock.EXPECT().CreateMessage(ctx, nil).Return(nil, nil)
	messageUseCaseMock.EXPECT().ProtoMessage2Message(nil).Return(message)
	messageUseCaseMock.EXPECT().WebsocketMessage(message).Return()
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestSendMessage_ParseJsonToMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  2,
		Reaction:  1,
		ChatId:    1,
	}

	messageUseCaseMock.EXPECT().ParseJsonToMessage(gomock.Any()).Return(message, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestSendMessage_CreateMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  2,
		Reaction:  1,
		ChatId:    1,
	}

	messageUseCaseMock.EXPECT().ParseJsonToMessage(gomock.Any()).Return(message, nil)
	messageUseCaseMock.EXPECT().Message2ProtoMessage(message).Return(nil)
	serverMock.EXPECT().CreateMessage(ctx, nil).Return(nil, status.Error(codes.Code(500), "Some error"))
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestChangeMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  2,
		Reaction:  1,
		ChatId:    1,
	}

	messageUseCaseMock.EXPECT().ParseJsonToMessage(gomock.Any()).Return(message, nil)
	messageUseCaseMock.EXPECT().Message2ProtoMessage(message).Return(nil)
	serverMock.EXPECT().ChangeMessage(ctx, nil).Return(nil, nil)
	messageUseCaseMock.EXPECT().ProtoMessage2Message(nil).Return(message)
	messageUseCaseMock.EXPECT().WebsocketReactMessage(message).Return()
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}

func TestChangeMessage_ParseJsonToMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  2,
		Reaction:  1,
		ChatId:    1,
	}

	messageUseCaseMock.EXPECT().ParseJsonToMessage(gomock.Any()).Return(message, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestChangeMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := MessageHandler{
		Usecase: messageUseCaseMock,
		Server:  serverMock,
	}

	murl, er := url.Parse("/chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  2,
		Reaction:  1,
		ChatId:    1,
	}

	messageUseCaseMock.EXPECT().ParseJsonToMessage(gomock.Any()).Return(message, nil)
	messageUseCaseMock.EXPECT().Message2ProtoMessage(message).Return(nil)
	serverMock.EXPECT().ChangeMessage(ctx, nil).Return(nil, status.Error(codes.Code(500), "Some error"))
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
