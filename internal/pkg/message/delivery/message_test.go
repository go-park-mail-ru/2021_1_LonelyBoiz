package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	mockUsecase "server/internal/pkg/message/usecase/mocks"
	"server/internal/pkg/models"
	sessionMocks "server/internal/pkg/session/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	chatId := 1
	limit := 20
	offset := 20

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ManageMessage(userId, chatId, limit, offset).Return([]models.Message{}, 200, nil)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetMessagesGetIdFromCOntextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestGetMessagesVarsError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}

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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetMessagesQueryCountError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	q.Add("NotCount", "20")
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
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetMessagesCountAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetMessagesOffsetError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetMessagesOffsetAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestGetMessagesManageMessageError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	chatId := 1
	limit := 20
	offset := 20

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ManageMessage(userId, chatId, limit, offset).Return([]models.Message{}, 500, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestGetMessagesManageMessageNonValidInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
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
	chatId := 1
	limit := 20
	offset := 20

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ManageMessage(userId, chatId, limit, offset).Return([]models.Message{}, 400, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetMessages(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestSendMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ParseJsonToMessage(req.Body).Return(message, nil)
	messageUseCaseMock.EXPECT().CreateMessage(message, message.ChatId, userId).Return(message, 200, nil)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestSendMessageAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "notInt",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestSendMessageGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, false)
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 401, response.StatusCode)
}

func TestSendMessageParseJsonToMessageError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ParseJsonToMessage(req.Body).Return(message, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestSendMessageCreateMessageError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ParseJsonToMessage(req.Body).Return(message, nil)
	messageUseCaseMock.EXPECT().CreateMessage(message, message.ChatId, userId).Return(message, 500, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestSendMessageNonValidInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ParseJsonToMessage(req.Body).Return(message, nil)
	messageUseCaseMock.EXPECT().CreateMessage(message, message.ChatId, userId).Return(message, 400, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SendMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestChangeMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"messageId": "1",
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ParseJsonToMessage(req.Body).Return(message, nil)
	messageUseCaseMock.EXPECT().ChangeMessage(userId, message.MessageId, message).Return(message, 204, nil)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}

func TestChangeMessageGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"messageId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, false)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestChangeMessageAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"notMessageId": "1",
	}
	req = mux.SetURLVars(req, vars)

	userId := 1

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userId,
	)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestChangeMessageParseToJsonError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	messageUseCaseMock := mockUsecase.NewMockMessageUsecaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := MessageHandler{
		Usecase:  messageUseCaseMock,
		Sessions: sessionManagerMock,
	}

	murl, er := url.Parse("chat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "PATCH",
		URL:    murl,
	}
	vars := map[string]string{
		"messageId": "1",
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

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(userId, true)
	messageUseCaseMock.EXPECT().ParseJsonToMessage(req.Body).Return(message, errors.New("Some error"))
	messageUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.ChangeMessage(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}
