package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	Logger *logrus.Entry
}

func (logger *LoggerMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start),
		}).Info(r.URL.Path)
	})
}
