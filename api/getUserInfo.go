package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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

	userInfo, err := a.Db.GetUser(id)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		responseWithJson(w, 500, response)
		return
	}

	userInfo.PasswordHash = nil
	responseWithJson(w, 200, userInfo)

	log.Println("successful get user", userInfo)
}
