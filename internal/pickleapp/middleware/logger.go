package middleware

import (
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"server/internal/pkg/models"
	usecase2 "server/internal/pkg/photo/usecase"
	"server/internal/pkg/session"
	"server/internal/pkg/user/usecase"
	"time"
)

type LoggerMiddleware struct {
	Logger  *models.Logger
	User    *usecase.UserUsecase
	Photo   *usecase2.PhotoUseCase
	Session *session.SessionsManager
}

func (logger *LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.LogError(err)
			}
		}()

		start := time.Now()

		rand.Seed(time.Now().UnixNano())
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
		logger.Photo.LoggerInterface = logger.Logger
		logger.Session.Logger = logger.Logger

		next.ServeHTTP(w, r)
	})
}
