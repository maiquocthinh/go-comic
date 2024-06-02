package email

import (
	"github.com/maiquocthinh/go-comic/config"
	"github.com/maiquocthinh/go-comic/pkg/email/template"
	"testing"
)
import "github.com/maiquocthinh/go-comic/pkg/email/model"

func TestSendEmail(t *testing.T) {
	cfg := &config.Gmail{
		ClientId:     "Your ClientId",
		ClientSecret: "Your ClientSecret",
		RefreshToken: "Your RefreshToken",
	}

	emailService := NewGmailServiceWithTemplatePath(cfg, "./template/")

	err := emailService.SendEmail(EmailMessageData{
		To:           "user@gmail.com",
		Subject:      "Reset your password",
		TemplateName: template.TemplateResetPassword,
		Model: model.ResetPassword{
			Firstname: "John",
			Username:  "john.doe",
			Code:      "123456",
		},
	})

	if err != nil {
		t.Errorf(err.Error())
	}
}
