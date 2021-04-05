package delivery

import (
	"bytes"
	json2 "encoding/json"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"server/api"
	"testing"
)

func TestValidateSignInData(t *testing.T) {
	testCases := []struct {
		name string
		in   User
		out  bool
		err  error
	}{
		{
			name: "Success case",
			in: User{
				Email:    "sdsfsf@ya.ru",
				Password: "Sk02343432",
			},
			out: true,
			err: nil,
		},
		{
			name: "Invalid email",
			in: User{
				Email:    "",
				Password: "Sk02343432",
			},
			out: false,
			err: errorDescriptionResponse{
				Description: map[string]string{
					"mail": "Введите почту",
				},
				Err: "Неверный формат входных данных",
			},
		},
		{
			name: "Invalid email",
			in: User{
				Email:    "asds@ya.ru",
				Password: "",
			},
			out: false,
			err: errorDescriptionResponse{
				Description: map[string]string{
					"password": "Введите пароль",
				},
				Err: "Неверный формат входных данных",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			out, _ := validateSignInData(testCase.in)
			assert.Equal(t, testCase.out, out)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		name string
		in   User
		out  bool
	}{
		{
			name: "Success case",
			in: User{
				Email:    "test@ya.ru",
				Password: "Sk02343432",
			},
			out: true,
		},
		{
			name: "User is not registered",
			in: User{
				Email:    "dfgdfg@ya.ru",
				Password: "Sk02343432",
			},
			out: false,
		},
		{
			name: "Incorrect password",
			in: User{
				Email:    "test@ya.ru",
				Password: "Sk023434sdfsd32",
			},
			out: false,
		},
	}

	a := api.App{Users: map[int]User{}}
	newUser := User{
		Email:    "test@ya.ru",
		Password: "Sk02343432",
	}

	pass := sha3.New512()
	pass.Write([]byte(newUser.Password))
	newUser.PasswordHash = pass.Sum(nil)

	_ = a.addNewUser(&newUser)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, a.checkPassword(&testCase.in))
		})
	}
}

func TestApp_SignInInvalidData(t *testing.T) {
	testCases := []struct {
		name        string
		in          User
		outCode     int
		outResponse errorDescriptionResponse
	}{
		{
			name: "Incorrect password",
			in: User{
				Email:    "test@ya.ru",
				Password: "Kps123456",
			},
			outCode:     401,
			outResponse: errorDescriptionResponse{Err: "Неверный логин или пароль"},
		},
		{
			name: "Missing field mail",
			in: User{
				Password: "Kps123456",
			},
			outCode: 400,
			outResponse: errorDescriptionResponse{
				Description: map[string]string{
					"mail": "Введите почту",
				},
				Err: "Неверный формат входных данных",
			},
		},
	}

	a := api.App{Users: map[int]User{}}
	newUser := User{
		Email:    "test@ya.ru",
		Password: "Kp123456",
	}
	pass := sha3.New512()
	pass.Write([]byte(newUser.Password))
	newUser.PasswordHash = pass.Sum(nil)
	_ = a.addNewUser(&newUser)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			json, err := json2.Marshal(testCase.in)
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
			a.SignIn(rw, req)
			response := rw.Result()

			assert.Equal(t, testCase.outCode, response.StatusCode)
			postBody, err := ioutil.ReadAll(response.Body)
			if err != nil {
				t.Error(err)
			}

			var resultBody errorDescriptionResponse
			err = json2.Unmarshal(postBody, &resultBody)

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, testCase.outResponse, resultBody)
		})
	}
}

func TestApp_SignInCorrectUser(t *testing.T) {
	a := api.App{Users: map[int]User{}, Sessions: map[int][]http.Cookie{}}
	newUser := User{
		Email:          "test@ya.ru",
		Password:       "Kp123456",
		DatePreference: "male",
	}
	pass := sha3.New512()
	pass.Write([]byte(newUser.Password))
	newUser.PasswordHash = pass.Sum(nil)
	err := a.addNewUser(&newUser)
	if err != nil {
		t.Error(nil)
	}

	UserTest := User{
		Email:    "test@ya.ru",
		Password: "Kp123456",
	}

	json, err := json2.Marshal(UserTest)
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
	a.SignIn(rw, req)
	response := rw.Result()

	assert.Equal(t, 200, response.StatusCode)

	postBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	var resultUser User

	err = json2.Unmarshal(postBody, &resultUser)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "male", resultUser.DatePreference)
	assert.Equal(t, UserTest.Email, resultUser.Email)
}

func TestApp_SignInInvalidData2(t *testing.T) {
	a := api.App{Users: map[int]User{}, Sessions: map[int][]http.Cookie{}}
	murl, er := url.Parse("http://localhost:8000/users")
	if er != nil {
		t.Error(er)
	}

	req := &http.Request{
		Method: "POST",
		URL:    murl,
		Body:   ioutil.NopCloser(bytes.NewBuffer(nil)),
	}

	rw := httptest.NewRecorder()
	a.SignIn(rw, req)
	response := rw.Result()

	assert.Equal(t, 400, response.StatusCode)

	postBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	var resultresp errorDescriptionResponse

	err = json2.Unmarshal(postBody, &resultresp)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Неверный формат входных данных", resultresp.Err)
}
