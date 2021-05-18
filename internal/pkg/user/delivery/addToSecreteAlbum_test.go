package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
	"server/internal/pkg/user/usecase"
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

func TestAddToSecreteAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCase := usecase.UserUsecase{}

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	user := models.User{
		Photos: []string{"1", "2"},
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

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	res := user_proto.UserNothing{}

	protoUser := user_proto.User{
		Photos: []string{"1", "2"},
	}

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(gomock.Any()).Return(user, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(&protoUser)
	serverMock.EXPECT().AddToSecreteAlbum(ctx, useCase.User2ProtoUser(user)).Return(&res, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.AddToSecreteAlbum(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 204, response.StatusCode)
}

func TestAddToSecreteAlbum_ParseJsonToUser_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	user := models.User{
		Photos: []string{"1", "2"},
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

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(gomock.Any()).Return(user, errors.New("Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.AddToSecreteAlbum(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestAddToSecreteAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCase := usecase.UserUsecase{}

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	serverMock := serverMocks.NewMockUserServiceClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
		Server:   serverMock,
	}

	user := models.User{
		Photos: []string{"1", "2"},
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

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	res := user_proto.UserNothing{}

	protoUser := user_proto.User{
		Photos: []string{"1", "2"},
	}

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(gomock.Any()).Return(user, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(&protoUser)
	serverMock.EXPECT().AddToSecreteAlbum(ctx, useCase.User2ProtoUser(user)).Return(&res, status.Error(codes.Code(500), "Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.AddToSecreteAlbum(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
