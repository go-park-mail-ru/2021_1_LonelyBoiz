package api

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func (a *App) validateCookie(cookie string) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, userSessions := range a.Sessions {
		for _, v := range userSessions {
			if v.Value == cookie {

				return true
			}
		}
	}

	return false
}

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		response := errorResponse{Err: "Вы не авторизованы"}
		responseWithJson(w, 401, response)
		return
	}

	if !a.validateCookie(token.Value) {
		w.WriteHeader(401)
		response := errorResponse{Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 401, response)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользоватея с таким id нет"
		responseWithJson(w, 400, response)
		return
	}

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

	log.Println("successful get user", userInfo)
}
