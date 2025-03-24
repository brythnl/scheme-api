package service

import (
	"context"

	"github.com/brythnl/scheme-api/internal/config"
	"github.com/brythnl/scheme-api/internal/model"
	"github.com/brythnl/scheme-api/internal/repository"
	"github.com/brythnl/scheme-api/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, loginData *model.UserLogin) (*model.TokenResponse, error)
	Register(ctx context.Context, userData *model.UserCreate) (*model.User, error)
}

type authService struct {
	userRepository repository.UserRepository
	config         *config.AuthConfig
}

func NewAuthService(repository repository.UserRepository, config *config.AuthConfig) AuthService {
	return &authService{
		userRepository: repository,
		config:         config,
	}
}

func (s *authService) Login(
	ctx context.Context,
	loginData *model.UserLogin,
) (*model.TokenResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, loginData.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID, user.Username, s.config.JWT)
	if err != nil {
		return nil, err
	}

	return &model.TokenResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   s.config.JWT.JWTExpirationHours * 3600,
	}, nil
}

func (s *authService) Register(
	ctx context.Context,
	userData *model.UserCreate,
) (*model.User, error) {
	_, err := s.userRepository.GetByEmail(ctx, userData.Email)
	if err == nil {
		// TODO: handle and centralize error
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userData.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: string(hashedPassword),
	}

	return s.userRepository.Create(ctx, user)
}
