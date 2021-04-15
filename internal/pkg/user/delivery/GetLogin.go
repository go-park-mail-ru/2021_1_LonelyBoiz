package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	id, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.LogInfo(response.Err)
		return
	}

	userInfo, err := a.UserCase.UserInfo(id)
	if err != nil {
		a.UserCase.LogError(err)
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 401, response)
		return
	}

	userInfo.PasswordHash = nil
	model.ResponseWithJson(w, 200, userInfo)
}
