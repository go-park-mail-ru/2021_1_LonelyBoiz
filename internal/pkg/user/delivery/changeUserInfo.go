package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	cookieId, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 403, response))
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Юзера с таким id нет"
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}

	if cookieId != userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пытаешься поменять не себя"
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 403, response))
		return
	}

	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 400, response))
		return
	}

	newUser, code, err := a.UserCase.ChangeUserInfo(newUser, userId)

	if code == 200 {
		model.Process(model.NewLogFunc("Success Change User Info", a.UserCase.LogInfo), model.NewResponseFunc(w, 200, newUser))
		return
	}

	model.Process(model.NewLogFunc(err, a.UserCase.LogInfo), model.NewResponseFunc(w, code, err))
}
