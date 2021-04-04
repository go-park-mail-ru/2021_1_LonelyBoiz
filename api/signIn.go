package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"log"
	"net/http"
	model "server/models"
)

func validateSignInData(newUser model.User) (bool, error) {
	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}

	_, err := govalidator.ValidateStruct(newUser)
	if err != nil {
		response.Description["mail"] = govalidator.ErrorByField(err, "email")
		response.Description["password"] = govalidator.ErrorByField(err, "password")
		return false, response
	}

	return true, nil
}

func (a *App) checkPassword(newUser *model.User) bool {
	user, err := a.Db.SignIn(newUser.Email)
	if err != nil || user.IsDeleted == true {
		return false
	}

	pass := sha3.New512()
	pass.Write([]byte(newUser.Password))
	err = bcrypt.CompareHashAndPassword(user.PasswordHash, pass.Sum(nil))
	if err != nil {
		return false
	}

	*newUser = user

	return true
}

func (a *App) SignIn(w http.ResponseWriter, r *http.Request) {
	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	isValid, response := validateSignInData(newUser)
	if !isValid {
		responseWithJson(w, 400, response)
		return
	}

	if isCorrect := a.checkPassword(&newUser); !isCorrect {
		response := errorResponse{Err: "Неверный логин или пароль"}
		responseWithJson(w, 401, response)
		return
	}

	err = a.setSession(w, newUser.Id)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		responseWithJson(w, 500, response)
		return
	}

	newUser.PasswordHash = nil
	responseWithJson(w, 200, newUser)

	log.Println("successful login", newUser)
}
