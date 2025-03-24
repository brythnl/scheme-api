package middleware

import "net/http"

type HeadersMiddleware struct {
	securityHeaders map[string]string
}

func NewHeadersMiddleware() *HeadersMiddleware {
	securityHeaders := map[string]string{
		"Referrer-Policy":         "strict-origin-when-cross-origin",
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":         "DENY",
		"X-XSS-Protection":        "1; mode=block",
		"Content-Security-Policy": "default-src 'self'",
	}

	return &HeadersMiddleware{
		securityHeaders: securityHeaders,
	}
}

func (m *HeadersMiddleware) SetHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range m.securityHeaders {
			w.Header().Set(k, v)
		}

		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
