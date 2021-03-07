package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func (a *App) validateCookieForChanging(cookie string, id int) bool {
	for _, v := range a.Sessions[id] {
		fmt.Println(v.Value, cookie)
		if v.Value == cookie {
			return true
		}
	}

	return false
}

func (a *App) checkPasswordForCHanging(newUser User) bool {
	oldUser := a.Users[newUser.Id]
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

func (a *App) changeUserProperties(newUser User) errorResponse {
	response := errorResponse{map[string]string{}, "Не удалось поменять данные"}

	switch {
	case newUser.Email != "":
		a.Users[newUser.Id].Email = newUser.Email
	case newUser.Name != "":
		a.Users[newUser.Id].Name = newUser.Name
	case newUser.Birthday != time.Time{}:
		a.Users[newUser.Id].Birthday = newUser.Birthday
	case newUser.Description != "":
		a.Users[newUser.Id].Description = newUser.Description
	case newUser.City != "":
		a.Users[newUser.Id].City = newUser.City
	case newUser.AvatarAddr != "":
		a.Users[newUser.Id].AvatarAddr = newUser.AvatarAddr
	case newUser.Instagram != "":
		a.Users[newUser.Id].Instagram = newUser.Instagram
	case newUser.Sex != "":
		if !validateSex(newUser.Sex) {
			response.Description["sex"] = "Неверно введен пол"
			return response
		}
		a.Users[newUser.Id].Sex = newUser.Sex
	case newUser.DatePreference != "":
		if !validateDatePreferensces(newUser.DatePreference) {
			response.Description["datePreferences"] = "Неверно введены предпочтения"
			return response
		}
		a.Users[newUser.Id].DatePreference = newUser.DatePreference
	}

	return response
}

func (a *App) ChangeUserPassword(newUser User) errorResponse {
	response := validatePass(newUser.Password)
	if len(response.Description) != 0 {
		return response
	}

	if newUser.SecondPassword != newUser.Password {
		response.Description["password"] = "Пароли не совпадают"
		return response
	}

	if !a.checkPasswordForCHanging(newUser) {
		response := errorResponse{map[string]string{}, "Отказано в доступе"}
		response.Description["password"] = "Неверный пароль"
		return response
	}

	newPassHash, err := hashPassword(newUser.Password)
	if err != nil {
		response.Description["password"] = "Не удалось поменять пароль"
		return response
	}
	a.Users[newUser.Id].PasswordHash = newPassHash
	newUser.Password = ""
	newUser.SecondPassword = ""
	newUser.OldPassword = ""

	return errorResponse{}
}

func (a *App) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Не залогинен 1"}
		json.NewEncoder(w).Encode(response)
		return
	}

	userId, err := strconv.Atoi(strings.SplitAfter(r.URL.String(), "/")[2])
	if err != nil {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Неверный id пользователя"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if !a.validateCookieForChanging(token.Value, userId) {
		w.WriteHeader(401)
		response := errorResponse{map[string]string{}, "Кука устарела"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	if err != nil {
		w.WriteHeader(400)
		response := errorResponse{map[string]string{}, "Неверный запрос"}
		json.NewEncoder(w).Encode(response)
		return
	}
	newUser.Id = userId

	response := a.changeUserProperties(newUser)
	if len(response.Description) != 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = a.ChangeUserPassword(newUser)
	if len(response.Description) != 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(newUser)
	fmt.Println("successful change")
}

/*
curl -b 'token=abcdef' \
	 --header "Content-Type: application/json" \
  	 --request PATCH \
  	 --data '{"mail":"xyz","pass":"xyz","passRepeat":"xyz","oldPass":"xyz1","name":"vasya"}' \
  	 http://localhost:8002/users
*/
