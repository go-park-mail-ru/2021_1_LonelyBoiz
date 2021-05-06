package delivery

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
	mock_usecase "server/internal/pkg/user/usecase/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUpParseJsonToUserError(t *testing.T) {
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

	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, errors.New("Some error"))
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SignUp(rw, req)

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}
