package delivery

import (
	"github.com/google/uuid"

	"net/http"
	"server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 400, response))
		return
	}

	newUser, code, responseError := a.UserCase.CreateNewUser(newUser)
	if code == 500 {
		models.Process(models.LoggerFunc(err, a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}
	if responseError != nil {
		models.Process(models.LoggerFunc(err, a.UserCase.LogInfo), models.ResponseFunc(w, code, responseError))
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		models.Process(models.LoggerFunc(err, a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]uuid.UUID, 0)
	}
	models.ResponseWithJson(w, 200, newUser)
	a.UserCase.LogInfo("Success SignUp")
}
