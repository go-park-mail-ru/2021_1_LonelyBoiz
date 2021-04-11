package middleware

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"server/internal/pkg/user/delivery"
	"time"
)

type LoggerMiddleware struct {
	Logger *logrus.Logger
	User   *delivery.UserHandler
}

func (logger *LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rand.Seed(time.Now().UnixNano())
		reqId := rand.Int63()

		logger.User.UserCase.Logger = logger.Logger.WithFields(logrus.Fields{
			"requestId":   reqId,
			"method":      r.Method,
			"url":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start),
		})

		logger.User.Sessions.Logger = logger.User.UserCase.Logger

		logger.User.UserCase.Logger.Info("Entry")

		next.ServeHTTP(w, r)
	})
}
