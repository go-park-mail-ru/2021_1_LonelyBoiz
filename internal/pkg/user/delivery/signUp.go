package delivery

import (
	"net/http"
	"reflect"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	if response := a.UserCase.ValidateSignUpData(newUser); response != nil {
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(model.UserErrorInvalidData)
		return
	}

	isSignedUp, response := a.UserCase.IsAlreadySignedUp(newUser.Email)
	if response != nil && reflect.TypeOf(response) != reflect.TypeOf(model.ErrorDescriptionResponse{}) {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if isSignedUp {
		a.UserCase.Logger.Info("Already Signed Up")
		model.ResponseWithJson(w, 400, response)
		return
	}

	err = a.UserCase.AddNewUser(&newUser)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	if len(newUser.Photos) == 0 {
		newUser.Photos = make([]int, 0)
	}

	model.ResponseWithJson(w, 200, newUser)
	a.UserCase.Logger.Info("Success SignUp")
}
