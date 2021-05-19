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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/login")
	if er != nil {
		t.Error(er)
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	protoUser := user_proto.User{
		Photos: []string{"1", "2"},
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().GetUserById(ctx, gomock.Any()).Return(&protoUser, nil)
	userUseCaseMock.EXPECT().ProtoUser2User(&protoUser).Return(user)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetLogin(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetLogin_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/login")
	if er != nil {
		t.Error(er)
	}

	user := models.User{
		Id:     1,
		Photos: []string{"1", "2"},
	}

	protoUser := user_proto.User{
		Photos: []string{"1", "2"},
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().GetUserById(ctx, gomock.Any()).Return(&protoUser, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.GetLogin(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
