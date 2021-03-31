package api

import (
	"net/http"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

var TestCaseValidateCookieWithId = []struct {
	userId, cookieId int
	cookie           string
	sessions         []http.Cookie
	res              bool
}{
	{userId: 0, cookieId: 0, cookie: "cookie", sessions: []http.Cookie{{Value: "cookie"}, {Value: "anotherCookie"}}, res: true},
	{userId: 1, cookieId: 1, cookie: "cookie", sessions: []http.Cookie{{Value: "anotherCookie"}}, res: false},
	{userId: 1, cookieId: 0, cookie: "cookie", sessions: []http.Cookie{{Value: "cookie"}, {Value: "anotherCookie"}}, res: false},
}

func TestValidateCookieWithId(t *testing.T) {
	var a App
	for _, v := range TestCaseValidateCookieWithId {
		a.Sessions = make(map[int][]http.Cookie)
		a.Sessions[v.cookieId] = v.sessions
		res := a.ValidateCookieWithId(v.cookie, v.userId)
		if res != v.res {
			t.Error("ValidateCookieWithId works wrong", "\nExpected:", v.res, "\nGot:", res)
		}
	}
}

var TestCaseCheckPasswordForCHanging = []struct {
	newUser User
	users   map[int]User
	res     bool
}{
	{newUser: User{Email: "mail", Id: 0, OldPassword: "password"},
		users: map[int]User{0: {Email: "mail", PasswordHash: []byte("password")}}, res: true},
	{newUser: User{Email: "mail", Id: 0, OldPassword: "password"},
		users: map[int]User{0: {Email: "mail", PasswordHash: []byte("password1")}}, res: false},
	{newUser: User{Email: "mail", Id: 0, OldPassword: "password"},
		users: map[int]User{0: {Email: "Email", PasswordHash: []byte("password")}}, res: false},
}

func TestCheckPasswordForCHanging(t *testing.T) {
	var a App
	for _, v := range TestCaseCheckPasswordForCHanging {
		firstHash := sha3.New512()
		firstHash.Write([]byte(v.users[0].PasswordHash))
		secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
		if err != nil {
			t.Error("Hash function error")
		}
		bufUser := v.users[0]
		bufUser.PasswordHash = secondHash
		v.users[0] = bufUser

		a.Users = v.users

		res := a.checkPasswordForCHanging(v.newUser)
		if res != v.res {
			t.Error("checkPasswordForCHanging works wrong", "\nExpected:", v.res, "\nGot:", res)
		}
	}
}

var TestCaseValidateSex = []struct {
	sex string
	res bool
}{
	{sex: "male", res: true},
	{sex: "female", res: true},
	{sex: "genderfluid helisexual", res: false},
}

func TestValidateSex(t *testing.T) {
	for _, v := range TestCaseValidateSex {
		res := ValidateSex(v.sex)
		if res != v.res {
			t.Error("ValidateSex works wrong", "\nExpected:", v.res, "\nGot:", res)
		}
	}
}

var TestCaseValidateDatePreference = []struct {
	preference string
	res        bool
}{
	{preference: "male", res: true},
	{preference: "female", res: true},
	{preference: "both", res: true},
	{preference: "genderfluid helisexual", res: false},
}

//func TestValidateDatePreference(t *testing.T) {
//	for _, v := range TestCaseValidateDatePreference {
//		res := ValidateDatePreferensces(v.preference)
//		if res != v.res {
//			t.Error("ValidateSex works wrong", "\nExpected:", v.res, "\nGot:", res)
//		}
//	}
//}

var TestCaseChangeUserProperties = []struct {
	newUser, oldUser User
	res              bool
}{
	{
		newUser: User{
			Id:             0,
			Name:           "Name",
			Birthday:       123,
			Description:    "description",
			City:           "city",
			Instagram:      "@inst",
			Avatar:         "avatar",
			Sex:            "male",
			DatePreference: "male"},
		oldUser: User{Id: 0},
		res:     true,
	},
	{
		newUser: User{
			Id:  0,
			Sex: "genderfluid helisexual"},
		oldUser: User{Id: 0},
		res:     false,
	},
	{
		newUser: User{
			Id:             0,
			DatePreference: "genderfluid helisexual"},
		oldUser: User{Id: 0},
		res:     false,
	},
	{
		newUser: User{
			Id:    0,
			Email: "mail@mail.ru"},
		oldUser: User{Id: 0},
		res:     true,
	},
	{
		newUser: User{
			Id:    0,
			Email: "asdf"},
		oldUser: User{Id: 0},
		res:     false,
	},
	{
		newUser: User{
			Id: 1},
		oldUser: User{Id: 0},
		res:     false,
	},
}

func TestChangeUserProperties(t *testing.T) {
	var a App
	for _, v := range TestCaseChangeUserProperties {
		a.Users = make(map[int]User)
		a.Users[0] = v.oldUser
		res := a.changeUserProperties(v.newUser)
		if (res == nil && v.res == false) || (res != nil && v.res == true) {
			t.Error("ChangeUserProperties works wrong", "\nExpected:", v.res, "\nGot:", res)
		}
	}
}

var TestCaseChangeUserPassword = []struct {
	newUser User
	users   map[int]User
	res     bool
}{
	{newUser: User{Password: "password", SecondPassword: "noPassword"},
		users: map[int]User{0: {PasswordHash: []byte("password1")}}, res: false},
	{newUser: User{Password: "1234567Qq", SecondPassword: "1234567Qq"},
		users: map[int]User{0: {PasswordHash: []byte("password1")}}, res: false},
	{newUser: User{Password: "1234567Qq", SecondPassword: "1234567q", OldPassword: "password1"},
		users: map[int]User{0: {PasswordHash: []byte("password1")}}, res: false},
	{newUser: User{Password: "1234567Qq", SecondPassword: "1234567Qq", OldPassword: "password1"},
		users: map[int]User{0: {PasswordHash: []byte("password1")}}, res: true},
}

func TestChangeUserPassword(t *testing.T) {
	var a App
	for _, v := range TestCaseChangeUserPassword {
		firstHash := sha3.New512()
		firstHash.Write([]byte(v.users[0].PasswordHash))
		secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
		if err != nil {
			t.Error("Hash function error")
		}
		bufUser := v.users[0]
		bufUser.PasswordHash = secondHash
		v.users[0] = bufUser

		a.Users = v.users

		res := a.changeUserPassword(v.newUser)
		if (res == nil && v.res == false) || (res != nil && v.res == true) {
			t.Error("changeUserPassword works wrong", "\nExpected:", v.res, "\nGot:", res)
		}
	}
}
