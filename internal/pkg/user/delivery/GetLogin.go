package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	id, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	userInfo, err := a.UserCase.GetUserInfoById(id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"

		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 401, response))
		return
	}

	model.Process(model.LoggerFunc("Get User Info for "+strconv.Itoa(id), a.UserCase.LogInfo), model.ResponseFunc(w, 200, userInfo))
}
