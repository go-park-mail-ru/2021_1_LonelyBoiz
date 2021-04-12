package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	isValid, response := a.UserCase.ValidateSignInData(newUser)
	if !isValid {
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(response.Error())
		return
	}

	isCorrect, err := a.UserCase.CheckPassword(&newUser)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if !isCorrect {
		response := model.ErrorResponse{Err: "Неверный логин или пароль"}
		model.ResponseWithJson(w, 401, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	newUser, err = a.Db.GetUser(newUser.Id)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)

	a.UserCase.Logger.Info("Success LogIn")
}
