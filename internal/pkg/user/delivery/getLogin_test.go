package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/internal/pkg/models"

	mock_usecase "server/internal/pkg/user/usecase/mocks"
	"testing"

	sessionMocks "server/internal/auth_server/delivery/session/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().GetUserInfoById(user.Id).Return(user, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	handlerTest.GetLogin(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestGetLoginGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	handlerTest.GetLogin(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestGetLoginGetUserInfoByIdError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	murl, er := url.Parse("auth")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "GET",
		URL:    murl,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().GetUserInfoById(user.Id).Return(user, errors.New("some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()
	handlerTest.GetLogin(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 401, response.StatusCode)
}
