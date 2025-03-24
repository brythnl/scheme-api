package api

import (
	"net/http"

	"github.com/brythnl/scheme-api/internal/api/handler"
	"github.com/brythnl/scheme-api/internal/api/middleware"
)

func (s *Server) addRoutes() http.Handler {
	mux := http.NewServeMux()

	authMiddleware := middleware.NewAuthMiddleware(s.config.Auth.JWT, s.logger)

	// Auth routes (public)
	mux.Handle("POST /api/v1/auth/login", handler.Login(s.authService))
	mux.Handle("POST /api/v1/auth/register", handler.Register(s.authService))

	// Global middleware
	corsMiddleware := middleware.NewCorsMiddleware(s.config.Server.AllowedOrigins)
	loggingMiddleware := middleware.NewLoggingMiddleware(s.logger)
	headersMiddleware := middleware.NewHeadersMiddleware()

	return applyMiddleware(
		mux,
		corsMiddleware.HandleCors,
		loggingMiddleware.LogRequest,
		headersMiddleware.SetHeaders,
	)
}

func applyMiddleware(
	h http.Handler,
	middleware ...func(http.Handler) http.Handler,
) http.Handler {
	for _, m := range middleware {
		h = m(h)
	}
	return h
}
