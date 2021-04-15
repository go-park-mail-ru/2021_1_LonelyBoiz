package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	id, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		a.UserCase.LogInfo("GetIdFromContext error")
	}

	if id != userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Пытаетесь получить доступ к чужому аккаунту"}
		model.ResponseWithJson(w, 401, response)
		return
	}

	userInfo, err := a.UserCase.UserInfo(userId)
	if err != nil {
		a.UserCase.LogError(err)
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	userInfo.PasswordHash = nil
	model.ResponseWithJson(w, 200, userInfo)
}
