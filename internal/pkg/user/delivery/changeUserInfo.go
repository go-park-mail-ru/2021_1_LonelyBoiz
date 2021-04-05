package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id не ceotcndetn"
		model.ResponseWithJson(w, 400, response)
		return
	}

	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	if id != userId {
		response := model.ErrorResponse{Err: "Отказано в доступе, кука устарела"}
		model.ResponseWithJson(w, 403, response)
		return
	}

	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	defer r.Body.Close()
	newUser.Id = id
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	var response error
	if newUser.Password != "" {
		response = a.UserCase.ChangeUserPassword(&newUser)
		if response != nil {
			model.ResponseWithJson(w, 400, response)
			return
		}
		newUser.Password = ""
		newUser.OldPassword = ""
		newUser.SecondPassword = ""
	}

	newUser.Id = userId
	response = a.UserCase.ChangeUserProperties(&newUser)
	if response != nil {
		model.ResponseWithJson(w, 400, response)
		return
	}

	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)

	log.Println("successful change", newUser)
}
