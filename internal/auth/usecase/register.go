package usecase

import (
	"context"
	"errors"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

func (uc *authUseCase) Register(ctx context.Context, userRegister *models.UserRegister) error {
	userExists, _ := uc.authRepo.GetUserByEmail(ctx, userRegister.Email)
	if userExists != nil {
		return common.NewConflictApiError(
			errors.New("This Email already in our sysytem"),
			"This Email already in our sysytem",
		)
	}
	userExists, _ = uc.authRepo.GetUserByUsername(ctx, userRegister.Username)
	if userExists != nil {
		return common.NewConflictApiError(
			errors.New("This Username already taken."),
			"This Username already taken.",
		)
	}

	// hash password
	hashedPassword, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return err
	}

	userRegister.HashedPassword = hashedPassword

	return uc.authRepo.CreateUser(ctx, userRegister)
}
