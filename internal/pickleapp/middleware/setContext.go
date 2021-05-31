package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
)

func SetContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := r.Context()
		_, ok := vars["id"]
		if ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "urlId", vars["id"])
		}

		_, ok = vars["userId"]
		if ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "userId", vars["userId"])
		}

		_, ok = vars["messageId"]
		if ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "urlMessageId", vars["messageId"])
		}

		_, ok = vars["chatId"]
		if ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "urlChatId", vars["chatId"])
		}

		query := r.URL.Query()
		limit, ok := query["count"]
		if ok && len(limit) != 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, "urlCount", limit[0])
		}

		offset, ok := query["offset"]
		if ok && len(offset) != 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, "urlOffset", offset[0])
		}

		_, ok = vars["ownerId"]
		if ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "ownerId", vars["ownerId"])
		}

		_, ok = vars["getterId"]
		if ok {
			ctx = metadata.AppendToOutgoingContext(ctx, "getterId", vars["getterId"])
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
