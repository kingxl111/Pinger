package handlers

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5/middleware"
)

func NewLogger(log *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		log.Info("logger middleware enabled")

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			entry := log.WithFields(logrus.Fields{
				"method":      r.Method,
				"path":        r.URL.Path,
				"remote_addr": r.RemoteAddr,
				"user_agent":  r.UserAgent(),
				"request_id":  middleware.GetReqID(r.Context()),
			})

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.WithFields(logrus.Fields{
					"status":   ww.Status(),
					"bytes":    ww.BytesWritten(),
					"duration": time.Since(t1).String(),
				}).Info("request completed")
			}()

			next.ServeHTTP(ww, r)
		})
	}
}
