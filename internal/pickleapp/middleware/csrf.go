package middleware

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	model "server/internal/pkg/models"
)

func keyGen() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = model.CharSet[rand.Intn(len(model.CharSet))]
	}

	return string(b)
}

func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.RequestURI == "/login" || r.RequestURI == "/users") && r.Method == "POST" {
			key := keyGen()
			expiration := time.Now().Add(24 * time.Hour)
			cookie := http.Cookie{
				Name:     "csrf-token",
				Value:    key,
				Expires:  expiration,
				SameSite: http.SameSiteNoneMode,
				Domain:   "p1ckle.herokuapp.com",
				//Domain: "localhost:8000",
				Secure:   true,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			w.Header().Set("X-CSRF-Token", key)
			w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
		} else {
			fmt.Println("enter CSRF middleware")
			csrfTokenHeader := r.Header.Get("X-CSRF-Token")
			csrfTokenCookie, err := r.Cookie("csrf-token")
			if err != nil {
				fmt.Println(err)
				response := model.ErrorResponse{Err: "Вы не авторизованы"}
				model.ResponseWithJson(w, 401, response)
				return
			}
			if csrfTokenHeader != csrfTokenCookie.Value {
				fmt.Println("csrf не равны")
				response := model.ErrorResponse{Err: "csrf токены не совпадают"}
				model.ResponseWithJson(w, 403, response)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
