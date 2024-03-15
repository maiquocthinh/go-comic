package usecase

import (
	"context"
	"github.com/maiquocthinh/go-comic/pkg/utils"
)

func (uc *authUseCase) Logout(ctx context.Context, userClaims *utils.UserTokenClaims) error {
	return uc.authRedisRepo.AddTokenToBlackList(ctx, userClaims.ID, &userClaims.ExpiresAt.Time)
}
