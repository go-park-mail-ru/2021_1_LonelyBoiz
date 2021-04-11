package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	Id             int    `json:"id"`
	Email          string `json:"mail" valid:"email~Почта не прошла валидацию"`
	Password       string `json:"password,omitempty" valid:"length(8|64)~Пароль не прошел валидацию"`
	SecondPassword string `json:"passwordRepeat,omitempty" valid:"length(8|64)"`
	PasswordHash   []byte `json:",omitempty"`
	OldPassword    string `json:"passwordOld,omitempty"`
	Name           string `json:"name"` // Введите имя
	Birthday       int64  `json:"birthday" valid:"ageValid~Вам должно быть 18!"`
	Description    string `json:"description"`
	City           string `json:"city"`
	Instagram      string `json:"instagram"`
	Sex            string `json:"sex"`
	DatePreference string `json:"datePreference"`
	IsDeleted      bool   `json:"isDeleted"`
	IsActive       bool   `json:"isActive"`
	Photos         []int  `json:"photos"`
}

func init() {
	govalidator.CustomTypeTagMap.Set(
		"ageValid",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			birthday, ok := i.(int64)
			if !ok {
				return false
			}

			tm := time.Unix(birthday, 0)
			diff := time.Now().Sub(tm)

			if diff/24/365 < 18 {
				return false
			}
			return true
		}),
	)
}

type key int

const CtxUserId key = -1

const CharSet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

type ErrorDescriptionResponse struct {
	Description map[string]string `json:"description"`
	Err         string            `json:"error"`
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func (e ErrorDescriptionResponse) Error() string {
	ret, _ := json.Marshal(e)

	return string(ret)
}

func ResponseWithJson(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

var (
	UserErrorInvalidData = "Неверный формат входных данных"
)
