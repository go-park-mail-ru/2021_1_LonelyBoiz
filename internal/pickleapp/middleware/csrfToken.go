package middleware

import (
	"github.com/gorilla/csrf"
	"log"
	"net/http"
)

func CSRFMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			token := csrf.Token(r)
			w.Header().Set("X-CSRF-Token", token)
			log.Print("hi token = " + token)
		}
		next.ServeHTTP(w, r)
	})
}
