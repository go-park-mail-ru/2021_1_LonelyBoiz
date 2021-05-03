package middleware

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func SetContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, ok := vars["id"]
		ctx := r.Context()
		if ok {
			ctx = metadata.AppendToOutgoingContext(r.Context(), "id", vars["id"])
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
