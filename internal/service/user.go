package service

import (
	"context"

	"github.com/brythnl/scheme-api/internal/model"
	"github.com/brythnl/scheme-api/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repository.Create(ctx, user)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repository.GetByEmail(ctx, email)
}
