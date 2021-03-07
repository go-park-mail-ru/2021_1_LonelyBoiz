package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

type errorResponse struct {
	Description map[string]string `json:"description"`
	Err         string            `json:"error"`
}

func (e errorResponse) Error() string {
	ret, _ := json.Marshal(e)

	return string(ret)
}

func responseWithJson(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

func validatePass(pass string) error {
	response := errorResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
	switch {
	case len(pass) < 8:
		response.Description["password"] = "Пароль должен содержать 8 символов"
		return response
	case pass == strings.ToLower(pass) || pass == strings.ToUpper(pass):
		response.Description["password"] = "Пароль должен состоять из символов разного регистра"
		return response
	case !strings.ContainsAny(pass, "1234567890"):
		response.Description["password"] = "Пароль должен содержать цифру"
		return response
	}

	return nil
}

func validateSignUpData(newUser User) error {
	err := validatePass(newUser.Password)
	if err != nil {
		return err
	}

	response := errorResponse{Description: map[string]string{}, Err: "Не удалось зарегестрироваться"}

	switch {
	case newUser.Email == "":
		response.Description["mail"] = "Введите почту"
	case newUser.Name == "":
		response.Description["name"] = "Введите имя"
	case newUser.Password != newUser.SecondPassword:
		response.Description["password"] = "Пароли не совпадают"
	case math.Floor(time.Now().Sub(time.Unix(newUser.Birthday, 0)).Hours()) < 18*24*365:
		response.Description["Birthday"] = "Вам должно быть 18"
	}

	if len(response.Description) != 0 {
		return response
	}

	return nil
}

func (a *App) isAlreadySignedUp(newEmail string) (bool, error) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, v := range a.Users {
		if v.Email == newEmail {
			response := errorResponse{Description: map[string]string{}, Err: "Не удалось зарегестрироваться"}
			response.Description["mail"] = "Почта уже зарегестрирована"
			return true, response
		}
	}

	return false, nil
}

func hashPassword(pass string) ([]byte, error) {
	firstHash := sha3.New512()
	firstHash.Write([]byte(pass))
	secondHash, err := bcrypt.GenerateFromPassword(firstHash.Sum(nil), 14)
	return secondHash, err
}

func (a *App) addNewUser(newUser User) error {
	var err error
	newUser.PasswordHash, err = hashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = ""
	newUser.SecondPassword = ""

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	newUser.Id = a.UserIds
	a.UserIds++
	a.Users[newUser.Id] = newUser

	return nil
}

func (a *App) setSession(w http.ResponseWriter, id int) {
	key := KeyGen()
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{Name: "token", Value: key, Expires: expiration}
	cookie.HttpOnly = true
	http.SetCookie(w, &cookie)

	var mutex = &sync.Mutex{}
	mutex.Lock()
	a.Sessions[id] = append(a.Sessions[id], cookie)
	mutex.Unlock()

}

func (a *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	if response := validateSignUpData(newUser); response != nil {
		responseWithJson(w, 400, response)
		return
	}

	if isSignedUp, response := a.isAlreadySignedUp(newUser.Email); isSignedUp {
		responseWithJson(w, 400, response)
		return
	}

	err = a.addNewUser(newUser)
	if err != nil {
		responseWithJson(w, 500, err)
		return
	}

	a.setSession(w, newUser.Id)

	responseWithJson(w, 200, newUser)

	fmt.Println("successful resistration\n", a.Users)
}

/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"mail":"xyz","pass":"1234567Qq","passRepeat":"1234567Qq","name":"vasya","birthday":1016048654}' \
  http://localhost:8003/users
*/
