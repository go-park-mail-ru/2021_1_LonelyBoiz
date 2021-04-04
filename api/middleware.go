package api

import (
	"context"
	"net/http"
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
		id, err := a.Db.GetCookie(token.Value)
		if err != nil {
			response := errorResponse{Err: err.Error()}
			responseWithJson(w, 401, response)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx,
			ctxUserId,
			id,
		)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}
