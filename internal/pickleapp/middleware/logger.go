package middleware

import (
	"math/rand"
	"net/http"
	chatUsecase "server/internal/pkg/chat/usecase"
	imageUsecase "server/internal/pkg/image/usecase"
	messageUsecase "server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
	"server/internal/pkg/session"
	"server/internal/pkg/user/usecase"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Logger  *models.Logger
	User    *usecase.UserUsecase
	Image   *imageUsecase.ImageUsecase
	Chat    *chatUsecase.ChatUsecase
	Session *session.SessionsManager
	Message *messageUsecase.MessageUsecase
}

func (logger *LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.LogError(err)
			}
		}()

		start := time.Now()
		reqId := rand.Int63()

		logger.Logger.Logger = logger.Logger.Logger.WithFields(logrus.Fields{
			"requestId":   reqId,
			"method":      r.Method,
			"url":         r.URL.Path,
			"remote_addr": r.RemoteAddr,
			"user_agent":  r.UserAgent(),
			"work_time":   time.Since(start),
			"time":        time.Now(),
		})

		logger.Logger.LogInfo("Entry Point - Logger Middleware")
		logger.User.LoggerInterface = logger.Logger
		logger.Image.LoggerInterface = logger.Logger
		logger.Session.Logger = logger.Logger
		logger.Chat.LoggerInterface = logger.Logger
		logger.Message.LoggerInterface = logger.Logger

		next.ServeHTTP(w, r)
	})
}
