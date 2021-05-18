package delivery

import (
	"context"
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
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("login")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}

	cookie := &http.Cookie{
		Name:  "token",
		Value: "123",
	}

	req.Header = make(http.Header)
	req.AddCookie(cookie)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().DeleteUser(ctx, &user_proto.UserNothing{}).Return(&user_proto.UserNothing{}, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().DeleteSession(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteUser_DB_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("login")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}

	cookie := &http.Cookie{
		Name:  "token",
		Value: "123",
	}

	req.Header = make(http.Header)
	req.AddCookie(cookie)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().DeleteUser(ctx, &user_proto.UserNothing{}).Return(&user_proto.UserNothing{}, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestDeleteUser_Cookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("login")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().DeleteUser(ctx, &user_proto.UserNothing{}).Return(&user_proto.UserNothing{}, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}
