package email

type EmailService interface {
	SendEmail(data EmailMessageData) error
}

type EmailMessageData struct {
	To           string
	Subject      string
	TemplateName string
	Model        interface{}
}
