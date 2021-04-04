package api

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
	"log"
	"net/http"
	model "server/models"
	"strconv"
)

func ValidateSex(sex string) bool {
	if sex != "male" && sex != "female" {
		return false
	}

	return true
}

func ValidateDatePreferensces(pref string) bool {
	if pref != "male" && pref != "female" && pref != "both" {
		return false
	}

	return true
}

func (a *App) checkPasswordForCHanging(newUser model.User) bool {
	oldUserPass, err := a.Db.GetPass(newUser.Id)
	if err != nil {
		return false
	}

	pass := sha3.New512()
	pass.Write([]byte(newUser.OldPassword))
	err = bcrypt.CompareHashAndPassword(oldUserPass, pass.Sum(nil))
	if err != nil {
		return false
	}

	return true
}

func (a *App) changeUserProperties(newUser *model.User) error {
	bufUser, err := a.Db.GetUser(newUser.Id)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		return response
	}

	if newUser.Email != "" {
		if !govalidator.IsEmail(newUser.Email) {
			response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
			response.Description["mail"] = "Почта не прошла валидацию"
			return response
		}
		bufUser.Email = newUser.Email
	}

	if newUser.Name != "" {
		bufUser.Name = newUser.Name
	}

	if newUser.Birthday != 0 {
		bufUser.Birthday = newUser.Birthday
	}

	if newUser.Description != "" {
		bufUser.Description = newUser.Description
	}

	if newUser.City != "" {
		bufUser.City = newUser.City
	}

	if newUser.Instagram != "" {
		bufUser.Instagram = newUser.Instagram
	}

	if newUser.Avatar != "" {
		bufUser.Avatar = newUser.Avatar
	}

	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось поменять данные"}
	if newUser.Sex != "" {
		if !ValidateSex(newUser.Sex) {
			response.Description["sex"] = "Неверно введен пол"
			return response
		}
		bufUser.Sex = newUser.Sex
	}

	if newUser.DatePreference != "" {
		if !ValidateDatePreferensces(newUser.DatePreference) {
			response.Description["datePreferences"] = "Неверно введены предпочтения"
			return response
		}
		bufUser.DatePreference = newUser.DatePreference
	}

	if len(newUser.PasswordHash) != 0 {
		bufUser.PasswordHash = newUser.PasswordHash
	}

	err = a.Db.ChangeUser(bufUser)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		return response
	}

	*newUser = bufUser

	return nil
}

func (a *App) changeUserPassword(newUser *model.User) error {
	response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}

	if !validatePass(newUser.Password) {
		response.Description["password"] = "Введите пароль"
		return response
	}

	if newUser.SecondPassword != newUser.Password {
		response.Description["password"] = "Пароли не совпадают"
		return response
	}

	if !a.checkPasswordForCHanging(*newUser) {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["password"] = "Неверный пароль"
		return response
	}

	hash, err := hashPassword(newUser.Password)
	if err != nil {
		response.Description["password"] = "Не удалось поменять пароль"
		return response
	}
	newUser.PasswordHash = hash

	//bufUser, err := a.Db.GetUser(newUser.Id)
	//if err != nil{
	//	response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
	//	response.Description["id"] = "Пользователя с таким id нет"
	//	return response
	//}
	//
	//err = a.Db.ChangeUser(bufUser)
	//if err != nil {
	//	response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
	//	return response
	//}
	return nil
}

func validatePass(password string) bool {
	if len(password) >= 8 && len(password) <= 64 {
		return true
	}
	return false
}

func (a *App) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id не ceotcndetn"
		responseWithJson(w, 400, response)
		return
	}

	ctx := r.Context()
	id, ok := ctx.Value(ctxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	if id != userId {
		response := errorResponse{Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 403, response)
		return
	}

	var newUser model.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	defer r.Body.Close()
	newUser.Id = id
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	var response error
	if newUser.Password != "" {
		response = a.changeUserPassword(&newUser)
		if response != nil {
			responseWithJson(w, 400, response)
			return
		}
		newUser.Password = ""
		newUser.OldPassword = ""
		newUser.SecondPassword = ""
	}

	newUser.Id = userId
	response = a.changeUserProperties(&newUser)
	if response != nil {
		responseWithJson(w, 400, response)
		return
	}

	newUser.PasswordHash = nil
	responseWithJson(w, 200, newUser)

	log.Println("successful change", newUser)
}
