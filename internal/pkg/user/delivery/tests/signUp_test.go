package tests

import (
	"github.com/stretchr/testify/assert"
	"server/api"
	"testing"
)

func TestApp_ValidatePass(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  error
	}{
		{
			name: "Correct",
			in:   "Ak12345678",
			out:  nil,
		},
		{
			name: "Invalid pass < 8",
			in:   "1",
			out: errorDescriptionResponse{
				Description: map[string]string{
					"password": "Пароль должен содержать 8 символов",
				},
				Err: "Неверный формат входных данных",
			},
		},
		{
			name: "Invalid case",
			in:   "aaaaaaaaaaa",
			out: errorDescriptionResponse{
				Description: map[string]string{
					"password": "Пароль должен состоять из символов разного регистра",
				},
				Err: "Неверный формат входных данных",
			},
		},
		{
			name: "Missing nums",
			in:   "Abcdfefrrf",
			out: errorDescriptionResponse{
				Description: map[string]string{
					"password": "Пароль должен содержать цифру",
				},
				Err: "Неверный формат входных данных",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, validatePass(testCase.in))
		})
	}
}

func TestApp_ValidateEmail(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "Correct",
			in:   "yfrfpyjq@yandex.ru",
			out:  true,
		},
		{
			name: "Email < 3",
			in:   "a",
			out:  false,
		},
		{
			name: "Email > 255",
			in:   "ertyuioplertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjertyuioplkjhgfdcvbnmkjkjhgfdcvbnmkj",
			out:  false,
		},
		{
			name: "Email don't match regexp",
			in:   "sfdsfsdf@",
			out:  false,
		},
		{
			name: "Email don't match regexp",
			in:   "sfdsfsdf.sdfsdfsdfd",
			out:  false,
		},
		{
			name: "Email don't match regexp",
			in:   "sfdsfsdf sdfsdfsdfd",
			out:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.out != validateEmail(testCase.in) {
				t.Error(testCase.name)
			}
		})
	}
}

func TestApp_ValidateSignUpData(t *testing.T) {
	testCases := []struct {
		name string
		in   User
		out  error
	}{
		{
			name: "correct",
			in: User{
				Name:           "Nick",
				Email:          "yfrfpyjq@yandex.ru",
				Password:       "Sk08820342",
				SecondPassword: "Sk08820342",
				Birthday:       123,
			},
			out: nil,
		},
		{
			name: "incorrect name",
			in: User{
				Name:           "",
				Email:          "yfrfpyjq@yandex.ru",
				Password:       "Sk08820342",
				SecondPassword: "Sk08820342",
				Birthday:       123,
			},
			out: errorDescriptionResponse{
				Description: map[string]string{
					"name": "Введите имя",
				},
				Err: "Не удалось зарегестрироваться",
			},
		},
		{
			name: "password1 != password2",
			in: User{
				Name:           "Nick",
				Email:          "yfrfpyjq@yandex.ru",
				Password:       "Sk088203422",
				SecondPassword: "Sk08820342",
				Birthday:       123,
			},
			out: errorDescriptionResponse{
				Description: map[string]string{
					"password": "Пароли не совпадают",
				},
				Err: "Не удалось зарегестрироваться",
			},
		},
		{
			name: "password1 != password2",
			in: User{
				Name:           "Nick",
				Email:          "yfrfpyjq@yandex.ru",
				Password:       "Sk088203422",
				SecondPassword: "Sk08820342",
				Birthday:       123,
			},
			out: errorDescriptionResponse{
				Description: map[string]string{
					"password": "Пароли не совпадают",
				},
				Err: "Не удалось зарегестрироваться",
			},
		},
		{
			name: "Age < 18",
			in: User{
				Name:           "Nick",
				Email:          "yfrfpyjq@yandex.ru",
				Password:       "Sk08820342",
				SecondPassword: "Sk08820342",
				Birthday:       99999999999999999,
			},

			out: errorDescriptionResponse{
				Description: map[string]string{
					"Birthday": "Вам должно быть 18",
				},
				Err: "Не удалось зарегестрироваться",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, validateSignUpData(testCase.in))
		})
	}
}

func TestApp_IsAlreadySignedUpSuccess(t *testing.T) {
	app := api.App{Users: map[int]User{
		0: {
			Email: "test@t.ru",
		},
	}}

	if st, err := app.isAlreadySignedUp("test@t.ru"); !(st == true && err != nil) {
		t.Error("Mail is already add!")
	}
}

func TestApp_IsAlreadySignedUpFailed(t *testing.T) {
	app := api.App{Users: map[int]User{
		0: {
			Email: "test@t.ru",
		},
	}}

	if st, err := app.isAlreadySignedUp("tesst@t.ru"); !(st == false && err == nil) {
		t.Error("This was new mail!")
	}
}

func TestApp_AddNewUserSuccess(t *testing.T) {
	app := api.App{Users: map[int]User{}}
	newUser := User{
		Email: "us@y.ru",
	}

	err := app.addNewUser(&newUser)
	if err != nil {
		t.Error(err)
	}

	if app.Users[0].Email != newUser.Email {
		t.Error("Added User != User")
	}
}
