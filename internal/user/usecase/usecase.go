package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/internal/user/repository"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *userUseCase {
	return &userUseCase{userRepo: userRepo}
}

type UserUseCase interface {
	GetProfile(ctx context.Context, userID int) (*entities.User, error)
	UpdateProfile(ctx context.Context, profileUpdate *models.UserProfileUpdate) (*entities.User, error)
}
