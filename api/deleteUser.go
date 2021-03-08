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
					return false
				}
				break
			}
		}
	}
	return true
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	cookie, cookieError := r.Cookie("token")

	if cookieError == http.ErrNoCookie || cookie == nil {
		return
	}

	key := cookie.Value

	args := mux.Vars(r)
	userId, err := strconv.Atoi(args["id"])
	if err != nil {
		return
	}

	mutex := sync.Mutex{}

	mutex.Lock()
	{
		if !checkAuthorization(a, userId, key) {
			return
		}

		delete(a.Users, userId)
	}
	mutex.Unlock()
}
