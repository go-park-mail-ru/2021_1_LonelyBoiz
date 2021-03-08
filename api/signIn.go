package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func validateSignInData(newUser User) (bool, error) {
	response := errorResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
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
		responseWithJson(w, 400, err)
		return
	}

	isValid, response := validateSignInData(newUser)
	if !isValid {
		responseWithJson(w, 400, response)
		return
	}

	if isCorrect := a.checkPassword(&newUser); !isCorrect {
		w.WriteHeader(401)
		response := errorResponse{Description: map[string]string{}, Err: "Не удалось авторизоваться"}
		response.Description["password"] = "Неверный пароль"
		responseWithJson(w, 401, response)
		return
	}

	newUser.PasswordHash = nil

	a.setSession(w, newUser.Id)

	responseWithJson(w, 200, newUser)

	fmt.Println("------------", a.Sessions[newUser.Id], "------------")
	fmt.Println("successful login\n", newUser)
}

/*
curl -H "Origin: http://localhost:3000" --verbose \
 --header "Content-Type: application/json" \
  --request POST \
  --data '{"mail":"2xyz","pass":"1234567Qq"}' \
  http://localhost:8000/login
*/
