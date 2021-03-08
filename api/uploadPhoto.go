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
	defer mutex.Unlock()
	user, ok := a.Users[userId]
	if !ok {
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
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
		responseWithJson(w, 400, err)
		return
	}
	fmt.Println(photoAddr)

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

