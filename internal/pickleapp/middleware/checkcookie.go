package middleware

import (
	"context"
	"net/http"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
)

type ValidateCookieMiddleware struct {
	Session *session.SessionsManager
}

func (m *ValidateCookieMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			response := model.ErrorResponse{Err: "Вы не авторизованы"}
			model.ResponseWithJson(w, 401, response)
			return
		}

		id, err := m.Session.DB.GetCookie(token.Value)
		if err != nil {
			response := model.ErrorResponse{Err: err.Error()}
			model.ResponseWithJson(w, 401, response)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx,
			model.CtxUserId,
			id,
		)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}
