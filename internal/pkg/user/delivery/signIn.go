package delivery

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	model "server/internal/pkg/models"
)

func ParseJsonToUser(body io.ReadCloser) (model.User, error) {
	var newUser model.User
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&newUser)
	defer body.Close()
	return newUser, err
}

func (a *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	newUser, err := ParseJsonToUser(r.Body)

	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	isValid, response := a.UserCase.ValidateSignInData(newUser)
	if !isValid {
		model.ResponseWithJson(w, 400, response)
		return
	}

	if isCorrect := a.UserCase.CheckPassword(&newUser); !isCorrect {
		response := model.ErrorResponse{Err: "Неверный логин или пароль"}
		model.ResponseWithJson(w, 401, response)
		return
	}

	err = a.Sessions.SetSession(w, newUser.Id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		return
	}

	newUser.PasswordHash = nil
	model.ResponseWithJson(w, 200, newUser)

	log.Println("successful login", newUser)
}
