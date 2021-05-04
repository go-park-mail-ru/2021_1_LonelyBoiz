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
			ctx = metadata.AppendToOutgoingContext(r.Context(), "urlId", vars["id"])
		}

		query := r.URL.Query()
		limit, ok := query["count"]
		if ok && len(limit) != 0 {
			ctx = metadata.AppendToOutgoingContext(r.Context(), "urlCount", limit[0])
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
