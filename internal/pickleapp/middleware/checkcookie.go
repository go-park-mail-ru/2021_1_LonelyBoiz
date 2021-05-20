package middleware

import (
	"context"
	"log"
	"net/http"
	session_proto "server/internal/auth_server/delivery/session"
	model "server/internal/pkg/models"
	"strconv"

	"google.golang.org/grpc/metadata"
)

type ValidateCookieMiddleware struct {
	Session session_proto.AuthCheckerClient
}

func (m *ValidateCookieMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			log.Println(err)
			response := model.ErrorResponse{Err: "Вы не авторизованы"}
			model.ResponseWithJson(w, 401, response)
			return
		}

		idProto, err := m.Session.Check(r.Context(), &session_proto.SessionToken{Token: token.Value})
		if err != nil {
			response := model.ErrorResponse{Err: err.Error()}
			model.ResponseWithJson(w, 401, response)
			return
		}

		id := int(idProto.GetId())

		ctx := r.Context()
		ctx = context.WithValue(ctx,
			model.CtxUserId,
			id,
		)

		ctx = metadata.AppendToOutgoingContext(ctx, "cookieId", strconv.Itoa(id))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
