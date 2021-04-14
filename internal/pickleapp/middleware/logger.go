package middleware

import (
	"math/rand"
	"net/http"
	delivery2 "server/internal/pkg/chat/delivery"
	delivery3 "server/internal/pkg/message/delivery"
	"server/internal/pkg/user/delivery"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Logger  *logrus.Logger
	User    *delivery.UserHandler
	Chat    *delivery2.ChatHandler
	Message *delivery3.MessageHandler
}

func (logger *LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error(err)
			}
		}()

		start := time.Now()

		rand.Seed(time.Now().UnixNano())
		reqId := rand.Int63()

		logger.User.UserCase.Logger = logger.Logger.WithFields(logrus.Fields{
			"requestId":   reqId,
			"method":      r.Method,
			"url":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
			"work_time":   time.Since(start),
			"time":        time.Now(),
		})

		logger.User.Sessions.Logger = logger.User.UserCase.Logger
		logger.Chat.Usecase.Logger = logger.User.UserCase.Logger
		logger.Message.Usecase.Logger = logger.User.UserCase.Logger
		logger.User.UserCase.Logger.Info("Entry")

		next.ServeHTTP(w, r)

	})
}
