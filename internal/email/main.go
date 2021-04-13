package email

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

type EmailClient struct {
	sendgridClient *sendgrid.Client
	Email          string
	Author         string
}

func Init(log *logrus.Entry) {
	setLogger(log)
	loadAuthConfigFromEnvs()

	logger.Info("EmailClient config sucessfylly loaded")
}

func NewClient() *EmailClient {
	return &EmailClient{
		sendgridClient: sendgrid.NewSendClient(configuration.SendgridAPIKey),
		Email:          configuration.Email,
		Author:         configuration.Author,
	}
}

func (ec *EmailClient) SendEmail(name, address, subject, content string) error {
	from := mail.NewEmail(ec.Author, ec.Email)
	// subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail(name, address)
	plainTextContent := "and easy to do anywhere, even with Go"
	// htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, content)
	resp, err := ec.sendgridClient.Send(message)

	if err != nil || (resp.StatusCode != 200 && resp.StatusCode != 202) {
		logger.WithFields(logrus.Fields{
			"status_code": resp.StatusCode,
			"body":        resp.Body,
		})
		return err
	}

	return nil
}
