package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type inputJson struct {
	Addr string `json:"addr"`
}

func (a *App) uploadPhoto(userId int, photoId string) (User, error) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	user, ok := a.Users[userId]
	defer mutex.Unlock()
	if !ok {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пользователя с таким id не найдено"
		return user, response
	}

	user.AvatarAddr = append(user.AvatarAddr, photoId)
	a.Users[userId] = user
	return user, nil
}

func (a *App) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	var photoAddr inputJson
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&photoAddr)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		responseWithJson(w, 400, response)
		return
	}
	fmt.Println(photoAddr)

	token, err := r.Cookie("token")
	if err != nil {
		response := errorResponse{Err: "Вы не авторизованы"}
		responseWithJson(w, 401, response)
		return
	}

	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Пользователя с таким id нет"
		responseWithJson(w, 400, response)
		return
	}

	if !a.validateCookieForChanging(token.Value, userId) {
		response := errorResponse{Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 401, response)
		return
	}

	user, response := a.uploadPhoto(userId, photoAddr.Addr)
	if response != nil {
		responseWithJson(w, 400, response)
		return
	}

	user.PasswordHash = nil
	responseWithJson(w, 200, user)

	fmt.Println("photo uploaded", user)
}

/*
	curl -b 'token=4PmZNbRkhQdftce2BiBZ8cyL3huAZGUbShdgY5FN' \
	--header "Content-Type: application/json" \
  --request POST \
  --data '{"addr":"chetiotr"}' \
  http://localhost:8003/users/0/photos
*/
