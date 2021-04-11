package delivery

import (
	"github.com/gorilla/mux"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])

	id, ok := a.Sessions.GetIdFromContext(r.Context())

	if !ok || err != nil || id != userId {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	userInfo, err := a.Db.GetUser(id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 500, response)
		a.UserCase.Logger.Error(err.Error())
		return
	}

	userInfo.PasswordHash = nil
	model.ResponseWithJson(w, 200, userInfo)

	a.UserCase.Logger.Info("Success")
}
