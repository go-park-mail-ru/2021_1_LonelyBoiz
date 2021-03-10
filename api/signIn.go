package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func validateSignInData(newUser User) (bool, error) {
	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
	switch {
	case newUser.Email == "":
		response.Description["mail"] = "Введите почту"
	case newUser.Password == "":
		response.Description["password"] = "Введите пароль"
	}

	if len(response.Description) > 0 {
		return false, response
	}

	return true, nil
}

func (a *App) checkPassword(newUser *User) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	users := a.Users
	mutex.Unlock()
	for _, v := range users {
		if v.Email == newUser.Email {
			pass := sha3.New512()
			pass.Write([]byte(newUser.Password))
			err := bcrypt.CompareHashAndPassword(v.PasswordHash, pass.Sum(nil))
			if err != nil {
				return false
			}

			mutex.Lock()
			*newUser = v
			mutex.Unlock()

			return true
		}
	}

	return false
}

func (a *App) validateCookieAndId(cookie string) (bool, int) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for userId, userSessions := range a.Sessions {
		for _, v := range userSessions {
			if v.Value == cookie {

				return true, userId
			}
		}
	}

	return false, -1
}

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err == nil {
		if ok, userId := a.validateCookieAndId(token.Value); ok {
			var mutex = &sync.Mutex{}
			mutex.Lock()
			userInfo, ok := a.Users[userId]
			mutex.Unlock()
			if !ok {
				response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе, кука устарела"}
				response.Description["id"] = "Пользователя с таким id нет"
				responseWithJson(w, 400, response)
				return
			}

			userInfo.PasswordHash = nil
			responseWithJson(w, 200, userInfo)
			return
		}
	}

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	isValid, response := validateSignInData(newUser)
	if !isValid {
		responseWithJson(w, 400, response)
		return
	}

	if isCorrect := a.checkPassword(&newUser); !isCorrect {
		response := errorResponse{Err: "Неверный логин или пароль"}
		responseWithJson(w, 401, response)
		return
	}

	a.setSession(w, newUser.Id)

	newUser.PasswordHash = nil
	responseWithJson(w, 200, newUser)

	log.Println("successful login", newUser, a.Sessions[newUser.Id])
}
