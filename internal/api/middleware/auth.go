package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/brythnl/scheme-api/internal/config"
	"github.com/brythnl/scheme-api/internal/logger"
	"github.com/brythnl/scheme-api/pkg/auth"
)

type AuthMiddleware struct {
	jwtConfig *config.JWTConfig
	logger    *logger.Logger
}

func NewAuthMiddleware(jwtConfig *config.JWTConfig, logger *logger.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		jwtConfig: jwtConfig,
		logger:    logger,
	}
}

func (m *AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
			http.Error(
				w,
				"Authorization header format must be 'Bearer {token}'",
				http.StatusUnauthorized,
			)
			return
		}

		claims, err := auth.VerifyToken(headerParts[1], m.jwtConfig)
		if err != nil {
			m.logger.Warn("Invalid token", "error", err)
			http.Error(w, "Token invalid or expired", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
