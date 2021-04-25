package delivery

import (
	"net/http"
	"server/internal/pkg/models"
)

func (a *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.LogError(err.Error())
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 401, response))
		return
	}

	newUser, code, err := a.UserCase.SignIn(newUser)
	if code != 200 {
		models.ResponseFunc(w, code, err)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}

	models.Process(models.LoggerFunc("Success LogIn", a.UserCase.LogInfo), models.ResponseFunc(w, 200, newUser))
}
