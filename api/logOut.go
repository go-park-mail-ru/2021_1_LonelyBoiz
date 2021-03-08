package api

import (
	"net/http"
	"sync"
	"time"
)

func deleteCookie(key string, sessions *map[int][]http.Cookie) {
	mutex := &sync.Mutex{}

	mutex.Lock()
	{
		for id, sessionCookies := range *sessions {
			for _, sessionCookie := range sessionCookies {
				if sessionCookie.Name == "token" && sessionCookie.Value == key {
					delete(*sessions, id)
				}
			}
		}
	}
	mutex.Unlock()
}

func (a *App) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err == http.ErrNoCookie || cookie == nil {
		return
	}

	key := cookie.Value

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
	deleteCookie(key, &a.Sessions)
}
