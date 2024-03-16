package usecase

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/internal/entities"
	"github.com/maiquocthinh/go-comic/internal/user/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
)

func (uc *userUseCase) UpdateProfile(ctx context.Context, profileUpdate *models.UserProfileUpdate) (*entities.User, error) {
	if !profileUpdate.Validate() {
		return nil, common.NewBadRequestApiError(
			errors.New("UserProfileUpdate invalid"),
			"No thing to update",
		)
	}

	if err := uc.userRepo.UpdateProfile(ctx, profileUpdate); err != nil {
		return nil, err
	}
	return uc.userRepo.GetProfile(ctx, profileUpdate.ID)
}
