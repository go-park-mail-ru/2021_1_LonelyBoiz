package middleware

import (
	"math/rand"
	"net/http"
	chatUsecase "server/internal/pkg/chat/usecase"
	imageUsecase "server/internal/pkg/image/usecase"
	messageUsecase "server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
	"server/internal/pkg/user/usecase"
	"server/internal/pkg/utils/metrics"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"google.golang.org/grpc/metadata"

	"github.com/sirupsen/logrus"
)

type LoggerMiddleware struct {
	Logger *models.Logger
	User   *usecase.UserUsecase
	Image  *imageUsecase.ImageUsecase
	Chat   *chatUsecase.ChatUsecase
	//Session *auth_server2.SessionsManager
	Message *messageUsecase.MessageUsecase
}

func (logger *LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := prometheus.NewTimer(metrics.Duration.WithLabelValues(strings.Split(r.URL.Path, "/")[1], r.Method))

		defer func() {
			tm.ObserveDuration()
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
		if logger.User != nil {
			logger.User.LoggerInterface = logger.Logger
		}

		if logger.Image != nil {
			logger.Image.LoggerInterface = logger.Logger
		}

		if logger.Chat != nil {
			logger.Chat.LoggerInterface = logger.Logger
		}

		if logger.Message != nil {
			logger.Message.LoggerInterface = logger.Logger
		}

		ctx := r.Context()
		ctx = metadata.AppendToOutgoingContext(ctx, "requestId", strconv.FormatInt(reqId, 10))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
