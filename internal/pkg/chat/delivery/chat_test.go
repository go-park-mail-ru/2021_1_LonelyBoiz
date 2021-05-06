package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	mockUsecase "server/internal/pkg/chat/usecase/mocks"
	"server/internal/pkg/models"
	serverMocks "server/internal/user_server/delivery/proto/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetChats(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
		Server:  serverMock,
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

	chatUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().GetChats(ctx, gomock.Any()).Return(nil, nil)

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetChatsError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	chatUseCaseMock := mockUsecase.NewMockChatUsecaseInterface(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := ChatHandler{
		Usecase: chatUseCaseMock,
		Server:  serverMock,
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

	chatUseCaseMock.EXPECT().LogError(gomock.Any()).Return()
	serverMock.EXPECT().GetChats(ctx, gomock.Any()).Return(nil, errors.New("Some error"))

	handlerTest.GetChats(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 2, response.StatusCode)
}
