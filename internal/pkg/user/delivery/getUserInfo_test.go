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

func TestGetUserInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/secretAlbum")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	protoUser := user_proto.User{
		Id:    1,
		Email: "email",
	}
	user := models.User{
		Id:    1,
		Email: "email",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().GetUserById(ctx, &user_proto.UserNothing{Dummy: true}).Return(&protoUser, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().ProtoUser2User(&protoUser).Return(user)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.GetUserInfo(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetUserInfo_GetUserById_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	murl, er := url.Parse("/secretAlbum")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
	}
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	protoUser := user_proto.User{
		Id:    1,
		Email: "email",
	}
	user := models.User{
		Id:    1,
		Email: "email",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	serverMock.EXPECT().GetUserById(ctx, &user_proto.UserNothing{Dummy: true}).Return(&protoUser, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any())

	handlerTest.GetUserInfo(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
