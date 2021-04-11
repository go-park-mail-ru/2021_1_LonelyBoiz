package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)

	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: model.UserErrorInvalidData}
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Error(model.UserErrorInvalidData)
		return
	}

	isValid, response := a.UserCase.ValidateSignInData(newUser)
	if !isValid {
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(response.Error())
		return
	}

	if isCorrect := a.UserCase.CheckPassword(&newUser); !isCorrect {
		response := model.ErrorResponse{Err: "Неверный логин или пароль"}
		model.ResponseWithJson(w, 401, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		a.UserCase.Logger.Error(err.Error())
		return
	}

	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)

	a.UserCase.Logger.Info("Success LogIn")
}
