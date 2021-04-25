package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookieId, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Юзера с таким id нет"
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	if cookieId != userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пытаешься удалить не себя"
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	err = a.UserCase.DeleteUser(userId)
	if err != nil {
		model.ResponseWithJson(w, 500, nil)
		model.Process(model.LoggerFunc(err.Error(), a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
		return
	}

	a.LogOut(w, r)
}
