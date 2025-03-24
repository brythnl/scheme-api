package repository

import (
	"context"
	"database/sql"

	"github.com/brythnl/scheme-api/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}
