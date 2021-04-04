package api

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func deleteCookie(key string, sessions *map[int][]http.Cookie) {
	mutex := &sync.Mutex{}

	mutex.Lock()
	{
		for userId, userSessions := range *sessions {
			for sessionId, session := range userSessions {
				if session.Name == "token" && session.Value == key {
					userSessions = append(userSessions[:sessionId], userSessions[sessionId+1:]...)
					(*sessions)[userId] = userSessions
				}
			}
		}
	}
	mutex.Unlock()
}

func (a *App) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		responseWithJson(w, 400, err)
		return
	}

	key := cookie.Value

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)

	err = a.Db.DeleteCookie(0, key)
	if err != nil {
		response := errorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		responseWithJson(w, 500, response)
		return
	}

	responseWithJson(w, 200, nil)

	log.Println("logout")
}
