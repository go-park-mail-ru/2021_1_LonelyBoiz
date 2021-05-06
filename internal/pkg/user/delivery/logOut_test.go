package delivery

/*import (
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
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogOut(t *testing.T) {
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

	userUseCaseMock.EXPECT().LogInfo("Success LogOut").Return()

	handlerTest.LogOut(rw, req)

	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)
}

func TestLogOutCookieError(t *testing.T) {
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

	userUseCaseMock.EXPECT().LogError("Не удалось взять куку").Return()

	handlerTest.LogOut(rw, req)

	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)
}

func TestLogOutDeleteSessionError(t *testing.T) {
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

	sessionManagerMock.EXPECT().DeleteSession(gomock.Any()).Return(errors.New("Some error"))
	userUseCaseMock.EXPECT().LogError("Some error").Return()

	handlerTest.LogOut(rw, req)

	response := rw.Result()

	assert.Equal(t, 500, response.StatusCode)
}
*/
