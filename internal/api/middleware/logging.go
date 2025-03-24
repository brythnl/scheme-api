package middleware

import (
	"net/http"
	"time"

	"github.com/brythnl/scheme-api/internal/logger"
)

type LoggingMiddleware struct {
	logger *logger.Logger
}

func NewLoggingMiddleware(logger *logger.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

func (m *LoggingMiddleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		m.logger.Info("HTTP Request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", rw.statusCode,
			"duration", duration,
			"ip", r.RemoteAddr,
			"user_agent", r.UserAgent(),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
