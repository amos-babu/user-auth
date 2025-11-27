package repository

import (
	"context"
	"database/sql"

	"github.com/amos-babu/user-auth/types"
)

type UserRepository interface {
	Register(ctx context.Context, data types.RegisterRequest) error
	Login(ctx context.Context, data types.LoginRequest) error
}

type Repository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Register(ctx context.Context, data types.RegisterRequest) error {

	return nil
}

func (r *Repository) Login(ctx context.Context, data types.LoginRequest) error {

	return nil
}
