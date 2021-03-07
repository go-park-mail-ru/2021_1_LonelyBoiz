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
	defer mutex.Unlock()
	user, ok := a.Users[userId]
	if !ok {
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
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
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["photo_id"] = "Фото с таким id не найдено"
		return user, response
	}

	user.AvatarAddr = remove(user.AvatarAddr, num)
	a.Users[userId] = user

	return user, nil
}

func (a *App) DeletePhoto(w http.ResponseWriter, r *http.Request) {
	var photoAddr inputJson
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&photoAddr)
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

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

	user, response := a.deletePhoto(userId, photoAddr.addr)
	if response != nil {
		responseWithJson(w, 400, response)
		return
	}

	user.PasswordHash = nil
	responseWithJson(w, 200, user)

	fmt.Println("photo deleted")
}

/*
curl -b 'token=M3S7xRA8wv0EuDXTk1b0y7OPwQJ5ciqUBi28qiLE' \
	 --header "Content-Type: application/json" \
  	 --request PATCH \
  	 --data '{"mail":"xyz","pass":"xyz","passRepeat":"xyz","oldPass":"xyz1","name":"Kolyan","sex":"female"}' \
  	 http://localhost:8003/users/0/photos/0
*/
