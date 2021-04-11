package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	if response := a.UserCase.ValidateSignUpData(newUser); response != nil {
		model.ResponseWithJson(w, 400, response)
		return
	}

	if isSignedUp, response := a.UserCase.IsAlreadySignedUp(newUser.Email); isSignedUp {
		model.ResponseWithJson(w, 400, response)
		return
	}

	err = a.UserCase.AddNewUser(&newUser)
	if err != nil {
		model.ResponseWithJson(w, 500, err)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		model.ResponseWithJson(w, 500, err)
		return
	}

	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)

	log.Println("successful registration\n", newUser)
}
