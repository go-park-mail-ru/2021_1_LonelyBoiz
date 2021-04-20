package delivery

import (
	"net/http"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.NewLogFunc(response.Err, a.UserCase.LogInfo), models.NewResponseFunc(w, 400, response))
		return
	}

	ok, err := a.UserCase.CheckCaptch(newUser.CaptchaToken)
	if err != nil {
		models.Process(models.NewLogFunc(err.Error(), a.UserCase.LogError), models.NewResponseFunc(w, 500, nil))
		return
	}

	if !ok {
		response := model.ErrorResponse{Err: "Не удалось пройти капчу"}
		models.Process(models.NewLogFunc(response.Err, a.UserCase.LogInfo), models.NewResponseFunc(w, 400, response))
		return
	}

	newUser, code, err := a.UserCase.CreateNewUser(newUser)
	switch code {
	case 200:
	case 500:
		models.Process(models.NewLogFunc(err, a.UserCase.LogError), models.NewResponseFunc(w, code, err))
		return
	default:
		models.Process(models.NewLogFunc(err, a.UserCase.LogInfo), models.NewResponseFunc(w, code, err))
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		models.Process(models.NewLogFunc(err, a.UserCase.LogError), models.NewResponseFunc(w, 500, nil))
		return
	}

	models.Process(models.NewLogFunc("Success SignUp", a.UserCase.LogInfo), models.NewResponseFunc(w, 200, newUser))
}
