package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/internal/user/models"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

type UserRepository interface {
	GetProfile(ctx context.Context, userID int) (*entities.User, error)
	UpdateProfile(ctx context.Context, profileUpdate *models.UserProfileUpdate) error
	UpdateAvatar(ctx context.Context, userID int, avatar string) error
	UpdatePassword(ctx context.Context, userID int, hashedPassword string) error
}
