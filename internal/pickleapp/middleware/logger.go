package middleware

//func MiddlewareLogger(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//		next.ServeHTTP(w, r)
//
//		a.Logger.WithFields(logrus.Fields{
//			"method":      r.Method,
//			"remote_addr": r.RemoteAddr,
//			"work_time":   time.Since(start),
//		}).Info(r.URL.Path)
//	})
//}
