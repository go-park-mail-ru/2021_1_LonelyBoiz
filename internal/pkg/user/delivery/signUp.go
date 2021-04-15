package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.LogError(err)
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	code, response := a.UserCase.SignUp(&newUser)
	if code != 200 {
		model.ResponseWithJson(w, code, response)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		a.UserCase.LogError(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	model.ResponseWithJson(w, 200, newUser)
	a.UserCase.LogInfo("Success SignUp")
}
