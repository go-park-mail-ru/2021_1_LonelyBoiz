package api

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id нет"
		responseWithJson(w, 400, response)
		return
	}

	ctx := r.Context()
	id, ok := ctx.Value(ctxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	if id != userId {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id нет"
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
