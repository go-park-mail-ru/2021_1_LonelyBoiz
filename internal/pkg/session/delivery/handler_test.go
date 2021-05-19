package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/internal/pkg/models"
	sessionMocks "server/internal/pkg/session/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogOut(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := AuthHandler{
		Usecase: sessionManagerMock,
	}

	murl, er := url.Parse("/login")
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

	sessionManagerMock.EXPECT().DeleteSessionByToken(cookie.Value).Return(nil)
	sessionManagerMock.EXPECT().DeleteCookie(cookie).Return()

	handlerTest.LogOut(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLogOut_DeleteSessionByToken_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := AuthHandler{
		Usecase: sessionManagerMock,
	}

	murl, er := url.Parse("/login")
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

	sessionManagerMock.EXPECT().DeleteSessionByToken(cookie.Value).Return(errors.New("Some error"))

	handlerTest.LogOut(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLogOut_NoCookie(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	sessionManagerMock := sessionMocks.NewMockSessionManagerInterface(mockCtrl)

	handlerTest := AuthHandler{
		Usecase: sessionManagerMock,
	}

	murl, er := url.Parse("/login")
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

	handlerTest.LogOut(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}
