package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/entities"
)

func (uc *userUseCase) GetProfile(ctx context.Context, userID int) (*entities.User, error) {
	return uc.userRepo.GetProfile(ctx, userID)
}
