package delivery

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	usecaseMocks "server/internal/pkg/user/usecase/mocks"
	userProto "server/internal/user_server/delivery/proto"
	serverMocks "server/internal/user_server/delivery/proto/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteChat(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/deleteChat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().DeleteChat(gomock.Any(), gomock.Any()).Return(&userProto.UserNothing{}, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteChat(rw, req)

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}

func TestDeleteChat_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/deleteChat")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}
	vars := map[string]string{
		"chatId": "1",
	}
	req = mux.SetURLVars(req, vars)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().DeleteChat(gomock.Any(), gomock.Any()).Return(&userProto.UserNothing{}, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.DeleteChat(rw, req)

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
