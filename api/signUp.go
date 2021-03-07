package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func validatePass(pass string) errorResponse {
	response := errorResponse{map[string]string{}, "Неверный формат входных данных"}
	switch {
	case len(pass) < 8:
		response.Description["password"] = "Пароль должен содержать 8 символов"
	case pass == strings.ToLower(pass) || pass == strings.ToUpper(pass):
		response.Description["password"] = "Пароль должен состоять из символов разного регистра"
	case !strings.ContainsAny(pass, "1234567890"):
		response.Description["password"] = "Пароль должен содержать цифру"
	}

	return response
}

func validateSignUpData(newUser User) (bool, errorResponse) {
	response := validatePass(newUser.Password)
	if len(response.Description) != 0 {
		return false, response
	}

	switch {
	case newUser.Email == "":
		response.Description["mail"] = "Введите почту"
	case newUser.Name == "":
		response.Description["name"] = "Введите имя"
	case newUser.Password != newUser.SecondPassword:
		response.Description["password"] = "Пароли не совпадают"
	case math.Floor(time.Now().Sub(newUser.Birthday).Hours()/24/365) < 18:
		response.Description["Birthday"] = "Вам должно быть 18"
	}

	if len(response.Description) != 0 {
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

func (a *App) addNewUser(newUser User) error {
	var err error
	newUser.Id = a.UserIds
	a.UserIds++
	newUser.PasswordHash, err = hashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""
	a.Users = append(a.Users, newUser)

	return nil
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
		response.Description["mail"] = "Почта уже зарегестрирована"
		json.NewEncoder(w).Encode(response)
		return
	}

	err = a.addNewUser(newUser)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	key := KeyGen()
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: key, Expires: expiration}
	a.Sessions[newUser.Id] = append(a.Sessions[newUser.Id], cookie)
	cookie.HttpOnly = true
	http.SetCookie(w, &cookie)

	fmt.Println("successful resistration\n", a.Users)
}

/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"mail":"xyz","pass":"xyz","passRepeat":"xyz","name":"vasya"|тут должна быть дата рождения но я в душе не чаю как в курле ее ввести|}' \
  http://localhost:8000/users
*/
