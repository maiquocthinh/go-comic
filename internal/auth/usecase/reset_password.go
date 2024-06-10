package usecase

import (
	"context"
	"encoding/json"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/pkg/common"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"time"
)

func (uc authUseCase) ResetPassword(ctx context.Context, userResetPassword *models.UserResetPassword) error {
	// get hashed code from redis
	hashedCode, err := uc.authRedisRepo.GetHashedCodeResetPassword(ctx, userResetPassword.Email)
	if err != nil {
		return common.NewBadRequestApiError(err, "Verification code is incorrect")
	}

	if err := utils.ComparePassword(hashedCode, userResetPassword.Code); err != nil {
		return common.NewBadRequestApiError(err, "Verification code is incorrect")
	}

	hashedPassword, err := utils.HashPassword(userResetPassword.Password)
	if err != nil {
		return err
	}

	if err := uc.authRepo.ResetPassword(ctx, userResetPassword.Email, hashedPassword); err != nil {
		return err
	}

	// delete hashed code from redis
	uc.authRedisRepo.RemoveHashedCodeResetPassword(ctx, userResetPassword.Email)

	return nil
}

func (uc authUseCase) SendCodeResetPassword(ctx context.Context, userSendCodeResetPassword *models.UserSendCodeResetPassword) error {
	user, err := uc.authRepo.GetUserByEmail(ctx, userSendCodeResetPassword.Email)
	if err != nil {
		return err
	}

	code := utils.GenerateOTPCode(6)
	hashedCode, err := utils.HashPassword(code)
	if err != nil {
		return err
	}

	// save code to redis
	{
		expirationTime := time.Duration(uc.cfg.Server.ResetPasswordCodeExpiration) * time.Second
		expiresAt := time.Now().Add(expirationTime)
		err = uc.authRedisRepo.AddHashedCodeResetPassword(
			ctx,
			user.Email,
			hashedCode,
			&expiresAt,
		)
		if err != nil {
			return err
		}
	}

	// send code via email
	{
		if user.FirstName == nil {
			user.FirstName = new(string)
		}

		jsonData, err := json.Marshal(&models.UserResetPasswordPubSub{
			Email:     user.Email,
			Firstname: *user.FirstName,
			Username:  user.Username,
			Code:      code,
			ExpiredIn: utils.SecondsToMinutes(uc.cfg.Server.ResetPasswordCodeExpiration),
		})

		err = uc.pubsub.Publish(ctx, common.TopicSendCodeResetPassword, string(jsonData))
		if err != nil {
			return err
		}
	}

	return nil
}
