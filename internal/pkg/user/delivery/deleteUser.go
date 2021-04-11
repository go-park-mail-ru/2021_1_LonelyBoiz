package delivery

import (
	"github.com/gorilla/mux"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	id, ok := a.Sessions.GetIdFromContext(r.Context())

	if !ok || err != nil || id != userId {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	err = a.Db.DeleteUser(userId)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		a.UserCase.Logger.Error(err.Error())
		return
	}

	a.UserCase.Logger.Info("Success Delete User")
	a.LogOut(w, r)
}
