package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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

func TestLogInParseToJsdnError(t *testing.T) {
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

	json, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(json)),
	}
	rw := httptest.NewRecorder()

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, errors.New("some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

	handlerTest.SignIn(rw, req)

	response := rw.Result()

	assert.Equal(t, 401, response.StatusCode)
}
