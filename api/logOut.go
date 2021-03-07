package api

import (
	"net/http"
	"time"
)

func convertStringToRune40(str string) (runes [40]rune) {
	for i := 0; i < 40; i++ {
		runes[i] = rune(str[i])
	}
	return
}

func deleteCookie(key string, sessions *[]Session) {
	for i, session := range *sessions {
		if session.Key == convertStringToRune40(key) {
			*sessions = append((*sessions)[:i], (*sessions)[i+1:]...)
			return
		}
	}
}

func (a *App) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")

	if err == http.ErrNoCookie || cookie == nil {
		http.Redirect(w, r, "/", http.StatusFound) //TODO:: куда делать редирект?
		return
	}

	key := cookie.Value
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)

	deleteCookie(key, &a.Sessions)

	http.Redirect(w, r, "/", http.StatusFound) //TODO:: куда делать редирект?
}
