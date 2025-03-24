package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/brythnl/scheme-api/internal/api"
	"github.com/brythnl/scheme-api/internal/config"
	"github.com/brythnl/scheme-api/internal/database"
	"github.com/brythnl/scheme-api/internal/logger"
	"github.com/brythnl/scheme-api/internal/repository"
	"github.com/brythnl/scheme-api/internal/service"
)

func main() {
	ctx := context.Background()
	if err := Run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Set up logger
	logger := logger.NewLogger(w)

	// Load configuration
	config, err := config.Load()
	if err != nil {
		logger.Error(err.Error())
		return fmt.Errorf("error loading config: %w", err)
	}

	// Connect database
	db, err := database.Connect(config.DB.URL)
	if err != nil {
		logger.Error(err.Error())
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer db.Close()

	// Set up repositories
	userRepository := repository.NewUserRepository(db)

	// Set up services
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository, config.Auth)

	// Set up server
	s := api.NewServer(
		config,
		logger,
		authService,
		userService,
	)

	// Start server
	s.Start()

	// Wait for shutdown signal
	<-ctx.Done()

	// Shutdown server gracefully
	if err := s.Stop(); err != nil {
		return err
	}

	return nil
}
