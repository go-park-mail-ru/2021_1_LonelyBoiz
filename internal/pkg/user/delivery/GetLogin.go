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
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 403, response))
		return
	}

	userInfo, err := a.UserCase.GetUserInfoById(id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"

		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 401, response))
		return
	}

	model.Process(model.NewLogFunc("Get User Info for "+strconv.Itoa(id), a.UserCase.LogInfo), model.NewResponseFunc(w, 200, userInfo))
}
