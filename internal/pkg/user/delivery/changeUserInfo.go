package delivery

import (
	"net/http"
	"reflect"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	cookieId, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Юзера с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	if cookieId != userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пытаешься поменять не себя"
		model.ResponseWithJson(w, 403, response)
		return
	}

	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	a.UserCase.Logger.Info(newUser, "после чтения")

	newUser.Id = userId

	if newUser.Password != "" {
		err := a.UserCase.ChangeUserPassword(&newUser)
		if err != nil {
			model.ResponseWithJson(w, 400, err)
			return
		}
		newUser.Password = ""
		newUser.OldPassword = ""
		newUser.SecondPassword = ""
	}

	newUser.Id = userId
	err = a.UserCase.ChangeUserProperties(&newUser)
	if err != nil {
		if reflect.TypeOf(err) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
			a.UserCase.Logger.Logger.Error(err)
			model.ResponseWithJson(w, 500, nil)
			return
		}
		model.ResponseWithJson(w, 400, err)
		return
	}

	a.UserCase.Logger.Info(newUser, "после обновления")

	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)

	a.UserCase.Logger.Info("Success Change User Info")
}
