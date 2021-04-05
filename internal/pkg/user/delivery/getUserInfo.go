package delivery

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	if id != userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	userInfo, err := a.Db.GetUser(id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 500, response)
		return
	}

	userInfo.PasswordHash = nil
	model.ResponseWithJson(w, 200, userInfo)

	log.Println("successful get user", userInfo)
}
