package api

import (
	"net/http"
	"sync"
)

func (a *App) validateCookieAndId(cookie string) (bool, int) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	for userId, userSessions := range a.Sessions {
		for _, v := range userSessions {
			if v.Value == cookie {

				return true, userId
			}
		}
	}

	return false, -1
}

func (a *App) GetLogin(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		responseWithJson(w, 401, nil)
		return
	}

	if ok, userId := a.validateCookieAndId(token.Value); ok {
		var mutex = &sync.Mutex{}
		mutex.Lock()
		userInfo, ok := a.Users[userId]
		mutex.Unlock()
		if !ok {
			responseWithJson(w, 401, nil)
			return
		}

		userInfo.PasswordHash = nil
		responseWithJson(w, 200, userInfo)
		return
	}
}
