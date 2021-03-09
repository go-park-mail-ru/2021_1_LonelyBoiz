package api

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func checkAuthorization(a *App, userId int, getKey string) bool {
	mutex := sync.Mutex{}
	mutex.Lock()
	expectedCookies, ok := a.Sessions[userId]
	mutex.Unlock()
	if !ok {
		return false
	}

	for _, expectedCookie := range expectedCookies {
		if expectedCookie.Name == "token" && expectedCookie.Value == getKey {
			return true
		}
	}

	return false
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("token")
	if cookieError != nil {
		responseWithJson(w, 401, cookieError)
		return
	}

	key := cookie.Value

	args := mux.Vars(r)
	userId, err := strconv.Atoi(args["id"])
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	if !checkAuthorization(a, userId, key) {
		responseWithJson(w, 400, "Отказано в доступе")
		return
	}

	delete(a.Users, userId)

	responseWithJson(w, 200, nil)

	log.Println("deleted user", a.Users)
}
