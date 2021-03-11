package api

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
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

	a := App{Users: map[int]User{}}
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
