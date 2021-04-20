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
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}

	userInfo, err := a.UserCase.GetUserInfoById(userId)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"

		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 401, response))
		return
	}

	model.Process(model.NewLogFunc("Get User Info", a.UserCase.LogInfo), model.NewResponseFunc(w, 200, userInfo))
}
