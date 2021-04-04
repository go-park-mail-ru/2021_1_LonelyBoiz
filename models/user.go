package models

type User struct {
	Id             int    `json:"id"`
	Email          string `json:"mail" valid:"email~Почта не прошла валидацию"`
	Password       string `json:"password,omitempty" valid:"length(8|64)~Пароль не прошел валидацию"`
	SecondPassword string `json:"passwordRepeat,omitempty" valid:"length(8|64)"`
	PasswordHash   []byte `json:",omitempty"`
	OldPassword    string `json:"oldPassword,omitempty"`
	Name           string `json:"name"` // Введите имя
	Birthday       int64  `json:"birthday" valid:"ageValid~Вам должно быть 18!"`
	Description    string `json:"description"`
	City           string `json:"city"`
	Avatar         string `json:"avatar"`
	Instagram      string `json:"instagram"`
	Sex            string `json:"sex"`
	DatePreference string `json:"datePreference"`
	IsDeleted      bool   `json:"isDeleted"`
	IsActive       bool   `json:"isActive"`
	Photos         string `json:"photos"`
}
