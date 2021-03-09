package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (a *App) deletePhoto(userId int, photoId string) (User, error) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	user, ok := a.Users[userId]
	mutex.Unlock()
	if !ok {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пользователя с таким id не найдено"
		return user, response
	}

	num := -1
	for i, v := range user.AvatarAddr {
		if v == photoId {
			num = i
			break
		}

	}
	if num == -1 {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["photo_id"] = "Фото с таким id не найдено"
		return user, response
	}

	user.AvatarAddr = remove(user.AvatarAddr, num)

	mutex.Lock()
	a.Users[userId] = user
	mutex.Unlock()

	return user, nil
}

func (a *App) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		responseWithJson(w, 401, err)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	if !a.validateCookieForChanging(token.Value, userId) {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 401, response)
		return
	}

	var photoAddr inputJson
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&photoAddr)
	defer r.Body.Close()
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	user, response := a.deletePhoto(userId, photoAddr.Addr)
	if response != nil {
		responseWithJson(w, 400, response)
		return
	}

	user.PasswordHash = nil
	responseWithJson(w, 200, user)

	fmt.Println("photo deleted")
}

/*
curl -b 'token=4PmZNbRkhQdftce2BiBZ8cyL3huAZGUbShdgY5FN' \
	 --header "Content-Type: application/json" \
  	 --request DELETE \
  	 --data '{"addr":"chet123iotr"}' \
  	 http://localhost:8003/users/0/photos
*/
