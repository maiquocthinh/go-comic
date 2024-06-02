package subscriber

import (
	"context"
	"encoding/json"
	"github.com/maiquocthinh/go-comic/internal/auth/models"
	"github.com/maiquocthinh/go-comic/pkg/email"
	"github.com/maiquocthinh/go-comic/pkg/email/model"
	"github.com/maiquocthinh/go-comic/pkg/email/template"
)

func SendCodeResetPasswordViaEmail(emailSvr email.EmailService) *consumerJob {
	return &consumerJob{
		Title: "Send code reset password via email",
		Hld: func(ctx context.Context, msg string) error {
			var userResetPassword models.UserResetPasswordPubSub
			if err := json.Unmarshal([]byte(msg), &userResetPassword); err != nil {
				return err
			}

			err := emailSvr.SendEmail(email.EmailMessageData{
				To:           userResetPassword.Email,
				Subject:      "Reset your password",
				TemplateName: template.TemplateResetPassword,
				Model: model.ResetPassword{
					Firstname: *userResetPassword.Firstname,
					Username:  userResetPassword.Username,
					Code:      userResetPassword.Code,
				},
			})

			return err
		},
	}
}
