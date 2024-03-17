package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/internal/user/repository"
	"github.com/maiquocthinh/go-comic/pkg/uploadprovider"
)

type userUseCase struct {
	userRepo       repository.UserRepository
	uploadProvider uploadprovider.UploadProvider
}

func NewUserUseCase(userRepo repository.UserRepository, uploadProvider uploadprovider.UploadProvider) *userUseCase {
	return &userUseCase{
		userRepo:       userRepo,
		uploadProvider: uploadProvider,
	}
}

type UserUseCase interface {
	GetProfile(ctx context.Context, userID int) (*entities.User, error)
	UpdateProfile(ctx context.Context, profileUpdate *models.UserProfileUpdate) (*entities.User, error)
	UpdateAvatar(ctx context.Context, userAvatarUpdate *models.UserAvatarUpdate) error
}
