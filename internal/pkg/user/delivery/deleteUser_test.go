package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/internal/pkg/models"
	sessionMocks "server/internal/pkg/session/mocks"
	mock_usecase "server/internal/pkg/user/usecase/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

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
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	cookie := http.Cookie{
		Name:     "token",
		Value:    "key",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost:3000",
	}

	req.Header = make(http.Header)

	req.AddCookie(&cookie)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, true)
	userUseCaseMock.EXPECT().DeleteUser(user.Id).Return(nil)
	sessionManagerMock.EXPECT().DeleteSession(gomock.Any()).Return(nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteUserGetIdFromContextError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

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
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	cookie := http.Cookie{
		Name:     "token",
		Value:    "key",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost:3000",
	}

	req.Header = make(http.Header)

	req.AddCookie(&cookie)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, false)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestDeleteUserAtoiError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

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
	vars := map[string]string{
		"id": "a",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	cookie := http.Cookie{
		Name:     "token",
		Value:    "key",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost:3000",
	}

	req.Header = make(http.Header)

	req.AddCookie(&cookie)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, true)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestDeleteUserIdNotEqual(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

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
	vars := map[string]string{
		"id": "2",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	cookie := http.Cookie{
		Name:     "token",
		Value:    "key",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost:3000",
	}

	req.Header = make(http.Header)

	req.AddCookie(&cookie)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, true)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 403, response.StatusCode)
}

func TestDeleteUserDeleteUserError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

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
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		user.Id,
	)

	cookie := http.Cookie{
		Name:     "token",
		Value:    "key",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost:3000",
	}

	req.Header = make(http.Header)

	req.AddCookie(&cookie)

	rw := httptest.NewRecorder()

	sessionManagerMock.EXPECT().GetIdFromContext(ctx).Return(user.Id, true)
	userUseCaseMock.EXPECT().DeleteUser(user.Id).Return(errors.New("Some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.DeleteUser(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
