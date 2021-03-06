package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func validateSignInData(newUser User) (bool, errorResponse) {
	response := errorResponse{map[string]string{}, "Неверный формат входных данных"}
	switch {
	case newUser.Email == "":
		response.Description["Email"] = "Введите почту"
	case newUser.Password == "":
		response.Description["Password"] = "Введите пароль"
	}

	if len(response.Description) > 0 {
		return false, response
	}
	return true, response
}

func (a *App) checkPassword(newEmail string, password string) (bool, int) {
	for _, v := range a.Users {
		if v.Email == newEmail {
			pass := sha3.New512()
			pass.Write([]byte(password))
			err := bcrypt.CompareHashAndPassword(v.PasswordHash, pass.Sum(nil))
			if err != nil {
				return false, -1
			}

			return true, v.Id
		}
	}

	return false, -1
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
	isCorrect, newUser.Id = a.checkPassword(newUser.Email, newUser.Password)

	if !isCorrect {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Авторизация не прошла"}
		response.Description["Password"] = "Неверный пароль"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(200)
	expiration := time.Now().Add(24 * time.Hour)
	key := KeyGen()
	a.Sessions[newUser.Id] = append(a.Sessions[newUser.Id], session{key: key, expirationDate: expiration})
	cookie := http.Cookie{Name: "token", Value: key, Expires: expiration}
	http.SetCookie(w, &cookie)
	fmt.Println("successful login\n", newUser)
}

/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"mail":"xyz","pass":"xyz"}' \
  http://localhost:8000/login
*/
