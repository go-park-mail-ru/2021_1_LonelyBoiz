package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: model.UserErrorInvalidData}
		a.UserCase.Logger.Info(model.UserErrorInvalidData)
		model.ResponseWithJson(w, 400, response)
		return
	}

	if response := a.UserCase.ValidateSignUpData(newUser); response != nil {
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(model.UserErrorInvalidData)
		return
	}

	if isSignedUp, response := a.UserCase.IsAlreadySignedUp(newUser.Email); isSignedUp {
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info("Mail is already used")
		return
	}

	err = a.UserCase.AddNewUser(&newUser)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		a.UserCase.Logger.Error(err.Error())
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		a.UserCase.Logger.Error(err.Error())
		return
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)
	a.UserCase.Logger.Info("Success SignUp")
}
