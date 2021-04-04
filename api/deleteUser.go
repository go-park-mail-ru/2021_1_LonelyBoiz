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

	err = a.Db.DeleteUser(userId)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		responseWithJson(w, 500, response)
		return
	}

	log.Println("deleted user")
	a.LogOut(w, r)

	//responseWithJson(w, 200, nil)

}
