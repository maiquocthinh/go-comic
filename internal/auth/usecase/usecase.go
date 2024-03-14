package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/internal/auth/repository"
)

type authUseCase struct {
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) *authUseCase {
	return &authUseCase{authRepo: authRepo}
}

type AuthUseCase interface {
	Register(ctx context.Context, userRegister *models.UserRegister) error
	Login(ctx context.Context, userLogin *models.UserLogin) (string, error)
}
