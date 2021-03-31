package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(ctxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	args := mux.Vars(r)
	userId, err := strconv.Atoi(args["id"])
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	if userId != id {
		responseWithJson(w, 400, "Отказано в доступе")
		return
	}

	delete(a.Users, userId)

	responseWithJson(w, 200, nil)

	log.Println("deleted user", a.Users)
}
