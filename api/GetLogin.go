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

	userInfo, ok := a.Users[id]

	if !ok {
		responseWithJson(w, 401, nil)
		return
	}

	userInfo.PasswordHash = nil
	responseWithJson(w, 200, userInfo)
	return
}
