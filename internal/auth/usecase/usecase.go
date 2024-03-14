package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/internal/auth/repository"
)

type authUseCase struct {
	cfg      *config.Config
	authRepo repository.AuthRepository
}

func NewAuthUseCase(cfg *config.Config, authRepo repository.AuthRepository) *authUseCase {
	return &authUseCase{
		cfg:      cfg,
		authRepo: authRepo,
	}
}

type AuthUseCase interface {
	Register(ctx context.Context, userRegister *models.UserRegister) error
	Login(ctx context.Context, userLogin *models.UserLogin) (string, error)
}
