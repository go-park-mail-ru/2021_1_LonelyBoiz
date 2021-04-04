package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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

func (a *App) MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		a.Logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start),
		}).Info(r.URL.Path)
	})
}
