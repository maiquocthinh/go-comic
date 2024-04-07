package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/internal/auth/repository"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

type authUseCase struct {
	cfg           *config.Config
	authRepo      repository.AuthRepository
	authRedisRepo repository.AuthRedisRepository
}

func NewAuthUseCase(cfg *config.Config, authRepo repository.AuthRepository, authRedisRepo repository.AuthRedisRepository) AuthUseCase {
	return &authUseCase{
		cfg:           cfg,
		authRepo:      authRepo,
		authRedisRepo: authRedisRepo,
	}
}

type AuthUseCase interface {
	Register(ctx context.Context, userRegister *models.UserRegister) error
	Login(ctx context.Context, userLogin *models.UserLogin) (string, error)
	Logout(ctx context.Context, userClaims *utils.UserTokenClaims) error
}
