package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func (a *App) validateCookieForChanging(cookie string, id int) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, v := range a.Sessions[id] {
		if v.Value == cookie {
			return true
		}
	}

	return false
}

func (a *App) checkPasswordForCHanging(newUser User) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	oldUser := a.Users[newUser.Id]
	mutex.Unlock()

	if oldUser.Email == newUser.Email {
		pass := sha3.New512()
		pass.Write([]byte(newUser.OldPassword))
		err := bcrypt.CompareHashAndPassword(oldUser.PasswordHash, pass.Sum(nil))
		if err != nil {
			return false
		}

		return true
	}

	return false
}

func validateSex(sex string) bool {
	if sex != "male" && sex != "female" {
		return false
	}

	return true
}

func validateDatePreferensces(pref string) bool {
	if pref != "male" && pref != "female" && pref != "both" {
		return false
	}

	return true
}

func (a *App) changeUserProperties(newUser User) error {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	bufUser := a.Users[newUser.Id]

	if newUser.Email != "" {
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

	response := errorResponse{Description: map[string]string{}, Err: "Не удалось поменять данные"}
	if newUser.Sex != "" {
		if !validateSex(newUser.Sex) {
			response.Description["sex"] = "Неверно введен пол"
			return response
		}
		bufUser.Sex = newUser.Sex
	}

	if newUser.DatePreference != "" {
		if !validateDatePreferensces(newUser.DatePreference) {
			response.Description["datePreferences"] = "Неверно введены предпочтения"
			return response
		}
		bufUser.DatePreference = newUser.DatePreference
	}

	a.Users[newUser.Id] = bufUser

	return nil
}

func (a *App) ChangeUserPassword(newUser User) error {
	err := validatePass(newUser.Password)
	if err != nil {
		return err
	}

	response := errorResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
	if newUser.SecondPassword != newUser.Password {
		response.Description["password"] = "Пароли не совпадают"
		return response
	}

	if !a.checkPasswordForCHanging(newUser) {
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["password"] = "Неверный пароль"
		return response
	}

	newPassHash, err := hashPassword(newUser.Password)
	if err != nil {
		response.Description["password"] = "Не удалось поменять пароль"
		return response
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	bufUser, ok := a.Users[newUser.Id]
	if !ok {
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пользователя с таким id не существует"
		return response
	}
	bufUser.PasswordHash = newPassHash
	a.Users[newUser.Id] = bufUser

	return nil
}

func (a *App) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	if !a.validateCookieForChanging(token.Value, userId) {
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 401, response)
		return
	}

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}
	newUser.Id = userId

	response := a.changeUserProperties(newUser)
	if response != nil {
		responseWithJson(w, 400, response)
		return
	}

	response = a.ChangeUserPassword(newUser)
	if response != nil {
		responseWithJson(w, 400, response)
		return
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	userInfo := a.Users[newUser.Id]
	mutex.Unlock()

	userInfo.PasswordHash = nil
	responseWithJson(w, 200, userInfo)

	fmt.Println("successful change")
	fmt.Println(userInfo)
}

/*
curl -b 'token=hCHjWRWy09NPrA8f0c3PedGTUPPbrK7stm2MVNrm' \
	 --header "Content-Type: application/json" \
  	 --request PATCH \
  	 --data '{"mail":"xyz","pass":"1234567Qq","passRepeat":"1234567Qq","oldPass":"1234567Qq","name":"Kolyan","sex":"female"}' \
  	 http://localhost:8003/users/0
*/
