package api

import (
	"context"
	"net/http"
	"sync"
)

type key int

const ctxUserId key = -1

func (a *App) MiddlewareValidateCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			response := errorResponse{Err: "Вы не авторизованы"}
			responseWithJson(w, 401, response)
			return
		}

		//здесь будет поход в базу
		var mutex = &sync.Mutex{}
		mutex.Lock()
		defer mutex.Unlock()
		for id, userSessions := range a.Sessions {
			for _, v := range userSessions {
				if v.Value == token.Value {
					ctx := r.Context()
					ctx = context.WithValue(ctx,
						ctxUserId,
						id,
					)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
		}

		response := errorResponse{Err: "Вы не авторизованы"}
		responseWithJson(w, 401, response)
		return
	})
}
