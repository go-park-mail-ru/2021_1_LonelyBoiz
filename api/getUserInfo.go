package api

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func (a *App) validateCookie(cookie string) bool {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	for _, userSessions := range a.Sessions {
		for _, v := range userSessions {
			if v.Value == cookie {

				return true
			}
		}
	}

	return false
}

func (a *App) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	if !a.validateCookie(token.Value) {
		w.WriteHeader(401)
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе, кука устарела"}
		responseWithJson(w, 401, response)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithJson(w, 500, err)
		return
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	userInfo, ok := a.Users[userId]
	if !ok {
		response := errorResponse{Description: map[string]string{}, Err: "Отказано в доступе, кука устарела"}
		response.Description["id"] = "Пользователя с таким id не существует"
		responseWithJson(w, 400, response)
		return
	}

	userInfo.PasswordHash = nil
	responseWithJson(w, 200, userInfo)

	fmt.Println("successful get user")
}

/*
curl -b 'token=UbalFd4mtHQkvXsMsKKEvidTj1YmuUN2FWtEiVDp' http://localhost:8001/users/1
*/
