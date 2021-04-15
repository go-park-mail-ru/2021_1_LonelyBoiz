package middleware

import (
	"math/rand"
	"net/http"
	delivery2 "server/internal/pkg/chat/delivery"
	delivery3 "server/internal/pkg/message/delivery"
	"server/internal/pkg/session"
	"server/internal/pkg/user/usecase"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Logger  *logrus.Logger
	Chat    *delivery2.ChatHandler
	Message *delivery3.MessageHandler
	Session *session.SessionsManager
	User    *usecase.UserUsecase
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

		logger.User.Logger = logger.Logger.WithFields(logrus.Fields{
			"requestId":   reqId,
			"method":      r.Method,
			"url":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
			"work_time":   time.Since(start),
			"time":        time.Now(),
		})

		logger.Session.Logger = logger.User.Logger
		logger.Chat.Usecase.Logger = logger.User.Logger
		logger.Message.Usecase.Logger = logger.User.Logger
		logger.User.Logger.Info("Entry")
		next.ServeHTTP(w, r)

	})
}
