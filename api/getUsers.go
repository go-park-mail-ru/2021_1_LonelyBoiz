package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func (a *App) getUsers(newUser User) []User {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	users := a.Users
	mutex.Unlock()

	var usersRet []User

	for _, v := range users {
		if v.Id == newUser.Id {
			continue
		}

		if (v.DatePreference == "both" || v.DatePreference == newUser.Sex) &&
			(newUser.DatePreference == "both" || newUser.DatePreference == v.Sex) {
			v.PasswordHash = nil
			usersRet = append(usersRet, v)
		}

		if len(usersRet) == 5 {
			break
		}
	}

	return usersRet
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		response := errorResponse{Err: "Вы не авторизованы"}
		responseWithJson(w, 401, response)
		return
	}

	var newUser User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newUser)
	defer r.Body.Close()
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}

	if !a.ValidateCookieWithId(token.Value, newUser.Id) {
		response := errorResponse{Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 401, response)
		return
	}

	if !ValidateSex(newUser.Sex) {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["sex"] = "Неверно введен пол"
		responseWithJson(w, 400, response)
		return
	}

	if !ValidateDatePreferensces(newUser.DatePreference) {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["datePreferences"] = "Неверно введены предпочтения"
		responseWithJson(w, 400, response)
		return
	}

	response := a.getUsers(newUser)

	responseWithJson(w, 200, response)

	log.Println("successful get user", response)
}
