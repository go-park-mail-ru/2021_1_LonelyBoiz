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
	if code == 500 {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogError), models.ResponseFunc(w, code, nil))
		return
	}
	if code != 200 {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogInfo), models.ResponseFunc(w, code, err))
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}

	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]string, 0)
	}

	newUser.PasswordHash = nil
	models.ResponseWithJson(w, 200, newUser)

	a.UserCase.LogInfo("Success LogIn")
}
