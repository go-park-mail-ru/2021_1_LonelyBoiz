package delivery

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
	usecaseMocks "server/internal/pkg/user/usecase/mocks"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPayment(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	data := url.Values{}
	data.Set("label", "{\"userid\":1}")
	data.Set("withdraw_amount", "3.00")
	req, _ := http.NewRequest(http.MethodPost, "urlStr", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().UpdatePayment(1, 40).Return(nil)

	handlerTest.Payment(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestPayment_Empty_Label(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	data := url.Values{}
	data.Set("label", "")
	data.Set("withdraw_amount", "3.00")
	req, _ := http.NewRequest(http.MethodPost, "urlStr", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.Payment(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestPayment_Empty_Withdraw(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	data := url.Values{}
	data.Set("label", "{\"userid\":1}")
	data.Set("withdraw_amount", "")
	req, _ := http.NewRequest(http.MethodPost, "urlStr", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.Payment(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestPayment_Wrong_Tarif(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	data := url.Values{}
	data.Set("label", "{\"userid\":1}")
	data.Set("withdraw_amount", "3.10")
	req, _ := http.NewRequest(http.MethodPost, "urlStr", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.Payment(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestPayment_Unmarshal_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	data := url.Values{}
	data.Set("label", ":1}")
	data.Set("withdraw_amount", "3.00")
	req, _ := http.NewRequest(http.MethodPost, "urlStr", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.Payment(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestPayment_UpdatePayment_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := usecaseMocks.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	handlerTest := UserHandler{
		UserCase: userUseCaseMock,
		Sessions: sessionManagerMock,
	}

	data := url.Values{}
	data.Set("label", "{\"userid\":1}")
	data.Set("withdraw_amount", "3.00")
	req, _ := http.NewRequest(http.MethodPost, "urlStr", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()
	userUseCaseMock.EXPECT().UpdatePayment(1, 40).Return(errors.New("Some error"))

	handlerTest.Payment(rw, req.WithContext(ctx))

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}
