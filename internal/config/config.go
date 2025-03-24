package config

import (
	"crypto/tls"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type ServerConfig struct {
	Env            string
	Addr           string
	AllowedOrigins []string
	TLS            *tls.Config
	CertFile       string
	KeyFile        string
	IdleTimeout    time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

type DBConfig struct {
	URL string
}

type JWTConfig struct {
	JWTSecret          string
	JWTExpirationHours int
}

type AuthConfig struct {
	JWT   *JWTConfig
	OAuth *oauth2.Config
}

type Config struct {
	Server *ServerConfig
	DB     *DBConfig
	Auth   *AuthConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	jwtExpirationHours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Server: &ServerConfig{
			Env:            os.Getenv("ENV"),
			Addr:           os.Getenv("ADDR"),
			AllowedOrigins: allowedOrigins,
			TLS: &tls.Config{
				CurvePreferences: []tls.CurveID{
					tls.CurveP256,
					tls.X25519,
				},
			},
			CertFile:     os.Getenv("CERT_FILE"),
			KeyFile:      os.Getenv("KEY_FILE"),
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		DB: &DBConfig{
			URL: os.Getenv("DATABASE_URL"),
		},
		Auth: &AuthConfig{
			JWT: &JWTConfig{
				JWTSecret:          os.Getenv("JWT_SECRET"),
				JWTExpirationHours: jwtExpirationHours,
			},
			OAuth: &oauth2.Config{
				ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
				ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
				RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
				Endpoint: oauth2.Endpoint{
					AuthURL:  os.Getenv("OAUTH_AUTH_URL"),
					TokenURL: os.Getenv("OAUTH_TOKEN_URL"),
				},
				Scopes: strings.Split(os.Getenv("OAUTH_SCOPES"), ","),
			},
		},
	}

	return cfg, nil
}
