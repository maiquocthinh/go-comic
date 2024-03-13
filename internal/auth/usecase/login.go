package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

func (uc *authUseCase) Login(ctx context.Context, userLogin *models.UserLogin) (string, error) {
	user, err := uc.authRepo.GetUserByEmail(ctx, userLogin.Email)
	if err != nil {
		return "", common.NewUnauthorizedApiError(err, "Email or Password wrong!")
	}

	if err := utils.ComparePassword(user.HashPassword, userLogin.Password); err != nil {
		return "", common.NewUnauthorizedApiError(err, "Email or Password wrong!")
	}

	return "jwt of user", nil
}
