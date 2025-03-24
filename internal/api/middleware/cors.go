package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

type CorsMiddleware struct {
	allowedOrigins []string
}

func NewCorsMiddleware(allowedOrigins []string) *CorsMiddleware {
	return &CorsMiddleware{
		allowedOrigins: allowedOrigins,
	}
}

func (m *CorsMiddleware) HandleCors(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   m.allowedOrigins,
		AllowCredentials: true,
		Debug:            true,
	}).Handler(next)
}
