package delivery

import (
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	id, ok := a.Sessions.GetIdFromContext(r.Context())

	if !ok || err != nil || id != userId {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: model.UserErrorInvalidData}
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	var response error
	if newUser.Password != "" {
		response = a.UserCase.ChangeUserPassword(&newUser)
		if response != nil {
			model.ResponseWithJson(w, 400, response)
			a.UserCase.Logger.Info(response.Error())
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
		a.UserCase.Logger.Info(response.Error())
		return
	}

	model.ResponseWithJson(w, 200, newUser)

	a.UserCase.Logger.Info("Success Change User Info")
}
