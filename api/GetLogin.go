package api

import (
	"fmt"
	"log"
	"net/http"
)

func (a *App) GetLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(ctxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}
	fmt.Println("id from context =", id)

	userInfo, err := a.Db.GetUser(id)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		responseWithJson(w, 401, response)
		return
	}

	userInfo.PasswordHash = nil
	responseWithJson(w, 200, userInfo)
	return
}
