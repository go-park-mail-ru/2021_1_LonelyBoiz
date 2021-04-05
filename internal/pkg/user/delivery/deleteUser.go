package delivery

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	args := mux.Vars(r)
	userId, err := strconv.Atoi(args["id"])
	if err != nil {
		model.ResponseWithJson(w, 400, err)
		return
	}

	if userId != id {
		model.ResponseWithJson(w, 400, "Отказано в доступе")
		return
	}

	err = a.Db.DeleteUser(userId)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		return
	}

	log.Println("deleted user")
	a.LogOut(w, r)
}
