package email

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/maiquocthinh/go-comic/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"html/template"
	"path/filepath"
	"strings"
)

type GmailService struct {
	refreshToken string
	accessToken  string
	oauthConfig  *oauth2.Config
	templatePath string
}

func NewGmailService(cfg *config.Gmail) *GmailService {
	return &GmailService{
		refreshToken: cfg.RefreshToken,
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.ClientId,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     google.Endpoint,
			Scopes:       []string{gmail.GmailSendScope},
		},
		templatePath: "./pkg/email/template/",
	}
}

func NewGmailServiceWithTemplatePath(cfg *config.Gmail, templatePath string) *GmailService {
	emailService := NewGmailService(cfg)
	emailService.templatePath = templatePath
	return emailService
}

func (gs *GmailService) SendEmail(data EmailMessageData) error {
	// Parse the email template
	htmlContent, err := gs.parseTemplate(data.TemplateName, data.Model)
	if err != nil {
		return err
	}

	// Check if the access token is valid before sending the email
	if !gs.isAccessTokenValid() {
		err := gs.refreshAccessToken()
		if err != nil {
			return fmt.Errorf("Unable to refresh access token: %v", err)
		}
	}

	// Create an HTTP client with the current access token
	client := gs.oauthConfig.Client(
		context.Background(),
		&oauth2.Token{AccessToken: gs.accessToken},
	)

	// Create a Gmail service
	srv, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("Unable to create Gmail client: %v", err)
	}

	// Create the email to be sent
	emailContent := fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
			"%s",
		data.To, data.Subject, htmlContent)

	var message gmail.Message
	message.Raw = base64.URLEncoding.EncodeToString([]byte(emailContent))
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)

	// Send the email
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return fmt.Errorf("Unable to send email: %v", err)
	}

	return nil
}

func (gs *GmailService) isAccessTokenValid() bool {
	tokenSource := gs.oauthConfig.TokenSource(context.Background(), &oauth2.Token{AccessToken: gs.accessToken})

	token, err := tokenSource.Token()
	if err != nil {
		return false
	}

	return token.Valid()
}

func (gs *GmailService) refreshAccessToken() error {
	tokenSource := gs.oauthConfig.TokenSource(
		context.Background(),
		&oauth2.Token{RefreshToken: gs.refreshToken},
	)

	token, err := tokenSource.Token()
	if err != nil {
		return fmt.Errorf("Unable to refresh access token: %v", err)
	}

	gs.accessToken = token.AccessToken
	return nil
}

func (gs GmailService) parseTemplate(templateName string, model interface{}) (string, error) {
	filePath, _ := filepath.Abs(filepath.Join(gs.templatePath, templateName))

	tpl, err := template.ParseFiles(filePath)
	if err != nil {
		return "", fmt.Errorf("Unable to parse email template: %v", err)
	}

	bfr := new(bytes.Buffer)
	if err = tpl.Execute(bfr, model); err != nil {
		return "", fmt.Errorf("Unable to execute email template: %v", err)
	}

	return bfr.String(), nil
}
