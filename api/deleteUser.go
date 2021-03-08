package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
)

func checkAuthorization(a *App, userId int, getKey string) bool {
	if expectedCookies, ok := a.Sessions[userId]; ok {
		for _, expectedCookie := range expectedCookies {
			if expectedCookie.Name == "token" {
				if expectedCookie.Value != getKey {
					return true //TODO:: КОД??
				}
				break
			}
		}
	}
	return false
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("token")

	if cookieError == http.ErrNoCookie || cookie == nil {
		return //TODO:: КОД??
	}

	getKey := cookie.Value

	args := mux.Vars(r)
	userId, err := strconv.Atoi(args["id"])

	if err != nil {
		return //TODO:: КОД??
	}

	mutex := sync.Mutex{}

	mutex.Lock()
	{
		if checkAuthorization(a, userId, getKey) {
			return
		}

		delete(a.Users, userId)
	}
	mutex.Unlock()
}
