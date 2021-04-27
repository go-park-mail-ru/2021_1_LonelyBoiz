package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/internal/pkg/models"
	sessionMocks "server/internal/pkg/session/mocks"
	mock_usecase "server/internal/pkg/user/usecase/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogIn(t *testing.T) {
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

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, nil)
	userUseCaseMock.EXPECT().SignIn(user).Return(user, 200, nil)
	userUseCaseMock.EXPECT().LogInfo("Success LogIn").Return()
	sessionManagerMock.EXPECT().SetSession(rw, user.Id).Return(nil)

	handlerTest.SignIn(rw, req)

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLogInParseToJsdnError(t *testing.T) {
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

func TestLogInSignInError(t *testing.T) {
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

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, nil)
	userUseCaseMock.EXPECT().SignIn(user).Return(user, 500, errors.New("some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.SignIn(rw, req)

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}

func TestLogInSetSessionError(t *testing.T) {
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

	userUseCaseMock.EXPECT().ParseJsonToUser(req.Body).Return(user, nil)
	userUseCaseMock.EXPECT().SignIn(user).Return(user, 200, nil)
	sessionManagerMock.EXPECT().SetSession(rw, user.Id).Return(errors.New("some error"))
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()

	handlerTest.SignIn(rw, req)

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
