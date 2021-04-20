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
		models.Process(models.NewLogFunc(response.Err, a.UserCase.LogInfo), models.NewResponseFunc(w, 401, response))
		return
	}

	newUser, code, err := a.UserCase.SignIn(newUser)
	if code != 200 {
		models.Process(models.NewLogFunc(err.Error(), a.UserCase.LogError), models.NewResponseFunc(w, code, err))
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		models.Process(models.NewLogFunc(err.Error(), a.UserCase.LogError), models.NewResponseFunc(w, 500, nil))
		return
	}

	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]int, 0)
	}

	newUser.PasswordHash = nil

	models.Process(models.NewLogFunc("Success LogIn", a.UserCase.LogInfo), models.NewResponseFunc(w, 200, newUser))
}
