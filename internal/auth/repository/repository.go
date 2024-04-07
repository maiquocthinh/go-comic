package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepo{db: db}
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entities.User, error)
	CreateUser(ctx context.Context, userRegister *models.UserRegister) error
}
