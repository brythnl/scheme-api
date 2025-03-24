package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/brythnl/scheme-api/internal/config"
	"github.com/brythnl/scheme-api/internal/logger"
	"github.com/brythnl/scheme-api/internal/service"
)

type Server struct {
	httpServer *http.Server
	config     *config.Config
	logger     *logger.Logger

	authService    service.AuthService
	userService    service.UserService
}

func NewServer(
	config *config.Config,
	logger *logger.Logger,

	authService service.AuthService,
	userService service.UserService,
) *Server {
	return &Server{
		config:         config,
		logger:         logger,
		authService:    authService,
		userService:    userService,
	}
}

func (s *Server) Start() {
	router := s.addRoutes()
	s.httpServer = &http.Server{
		Addr:         s.config.Server.Addr,
		Handler:      router,
		ErrorLog:     slog.NewLogLogger(s.logger.Handler(), slog.LevelError),
		TLSConfig:    s.config.Server.TLS,
		IdleTimeout:  s.config.Server.IdleTimeout,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
	}

	go func() {
		s.logger.Info("server starting and listening", "addr", s.httpServer.Addr)

		if err := s.httpServer.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem"); err != nil &&
			err != http.ErrServerClosed {
			s.logger.Error(err.Error())
			fmt.Fprintf(os.Stderr, "error listening and serving on %s: %v", s.httpServer.Addr, err)
		}
	}()
}

func (s *Server) Stop() error {
	if s.httpServer == nil {
		return nil
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.httpServer.Shutdown(shutdownCtx)
}
