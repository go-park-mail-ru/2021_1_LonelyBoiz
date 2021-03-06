package api

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type errorResponse struct {
	Description map[string]string
	Err         string
}

func validateSignUpData(newUser User) (bool, errorResponse) {
	response := errorResponse{map[string]string{}, "Не удалось зарегестрироваться"}
	switch {
	case newUser.Email == "":
		response.Description["Email"] = "Введите почту"
	case newUser.Name == "":
		response.Description["Name"] = "Введите имя"
	case newUser.Password != newUser.SecondPassword:
		response.Description["Password"] = "Неверный повторный пароль"
	case math.Floor(time.Now().Sub(newUser.Birthday).Hours()/24/365) < 18:
		response.Description["Birthday"] = "Вам должно быть 18"
	}

	if len(response.Description) > 0 {
		return false, response
	}
	return true, response
}

func (a *App) isAlreadySignedUp(newEmail string) bool {
	for _, v := range a.Users {
		if v.Email == newEmail {
			return true
		}
	}

	return false
}

func hashPassword(pass string) ([]byte, error) {
	firstHash := sha3.New512()
	firstHash.Write([]byte(pass))
	secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
	return secondHash, err
}

func (a *App) addNewUser(newUser User) (string, error) {
	var err error
	newUser.Id = a.UserIds
	a.UserIds++
	newUser.PasswordHash, err = hashPassword(newUser.Password)
	if err != nil {
		return "", err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""
	a.Users = append(a.Users, newUser)

	key := KeyGen()
	a.Sessions = append(a.Sessions, Session{newUser.Id, key, time.Now()})

	return key, nil
}

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		response := errorResponse{map[string]string{}, "Неверный запрос"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid, response := validateSignUpData(newUser)
	if !isValid {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	if a.isAlreadySignedUp(newUser.Email) {
		response := errorResponse{map[string]string{}, "Не удалось зарегестрироваться"}
		response.Description["Email"] = "Почта уже зарегестрирована"
		json.NewEncoder(w).Encode(response)
		return
	}

	key, err := a.addNewUser(newUser)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: key, Expires: expiration}
	http.SetCookie(w, &cookie)
}

/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"Email":"xyz","pass_1":"xyz","pass_2":"xyz","Birthday":"2021-03-06 20:27:42.502497477 +0300 MSK m=+16.623219324"}' \
  http://localhost:8000/users
*/
