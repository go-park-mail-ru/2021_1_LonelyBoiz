package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/lib/pq"
)

type GoogleCaptcha struct {
	Success     bool     `json:"success"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

type User struct {
	Id               int            `json:"id"`
	Email            string         `json:"mail" valid:"email~Почта не прошла валидацию"`
	Password         string         `json:"password,omitempty" valid:"length(8|64)~Пароль не прошел валидацию"`
	SecondPassword   string         `json:"passwordRepeat,omitempty" valid:"length(8|64)"`
	PasswordHash     []byte         `json:",omitempty"`
	OldPassword      string         `json:"passwordOld,omitempty"`
	Name             string         `json:"name"`
	Birthday         int64          `json:"birthday" valid:"ageValid~Вам должно быть 18!"`
	Description      string         `json:"description"`
	City             string         `json:"city"`
	Instagram        string         `json:"instagram"`
	Sex              string         `json:"sex"`
	DatePreference   string         `json:"datePreference"`
	IsDeleted        bool           `json:"isDeleted"`
	IsActive         bool           `json:"isActive"`
	Photos           pq.StringArray `json:"photos"`
	CaptchaToken     string         `json:"captchaToken"`
	Height           int            `json:"height"`
	PartnerHeightTop int            `json:"partnerHeightTop" db:"partnerheighttop"`
	PartnerHeightBot int            `json:"partnerHeightBot" db:"partnerheightbot"`
	Weight           int            `json:"weight"`
	PartnerWeightTop int            `json:"partnerWeightTop" db:"partnerweighttop"`
	PartnerWeightBot int            `json:"partnerWeightBot" db:"partnerweightbot"`
	PartnerAgeTop    int            `json:"partnerAgeTop" db:"partneragetop"`
	PartnerAgeBot    int            `json:"partnerAgeBot" db:"partneragebot"`
	Interests        pq.Int64Array  `json:"interests"`
}

type Label struct {
	UserId int `json:"userId"`
}

type Like struct {
	UserId   int    `json:"userId"`
	Reaction string `json:"reaction"`
}

var Tarif = map[string]int{
	"1.00": 10,
	"2.00": 20,
	"3.00": 40,
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

const CtxImageId key = -2

const CharSet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

var (
	UserErrorInvalidData = "Неверный формат входных данных"
)
