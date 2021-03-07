package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func validateSignInData(newUser User) (bool, errorResponse) {
	response := errorResponse{map[string]string{}, "Неверный формат входных данных"}
	switch {
	case newUser.Email == "":
		response.Description["mail"] = "Введите почту"
	case newUser.Password == "":
		response.Description["password"] = "Введите пароль"
	}

	if len(response.Description) > 0 {
		return false, response
	}
	return true, response
}

func (a *App) checkPassword(newUser *User) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, v := range a.Users {
		if v.Email == newUser.Email {
			pass := sha3.New512()
			pass.Write([]byte(newUser.Password))
			err := bcrypt.CompareHashAndPassword(v.PasswordHash, pass.Sum(nil))
			if err != nil {
				return false
			}

			*newUser = v
			return true
		}
	}

	return false
}

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	var newUser User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		response := errorResponse{map[string]string{}, "Неверный запрос"}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid, response := validateSignInData(newUser)
	if !isValid {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	var isCorrect bool
	isCorrect = a.checkPassword(&newUser)

	if !isCorrect {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Авторизация не прошла"}
		response.Description["password"] = "Неверный пароль"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(200)
	expiration := time.Now().Add(24 * time.Hour)
	key := KeyGen()
	cookie := http.Cookie{Name: "token", Value: key, Expires: expiration}
	http.SetCookie(w, &cookie)
	newUser.PasswordHash = nil
	json.NewEncoder(w).Encode(newUser)

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	a.Sessions[newUser.Id] = append(a.Sessions[newUser.Id], cookie)

	fmt.Println("------------", key, "------------")
	fmt.Println("successful login\n", newUser)
}

/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"mail":"xyz","pass":"xyz"}' \
  http://localhost:8000/login
*/
