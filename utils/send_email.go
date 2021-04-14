package utils

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"net/smtp"
)

const (
	SmtpURL  = "smtp.gmail.com"
	SmtpPort = ":587"
)

func SendEmail(email, pass, to, content, subject, typ string) error {
	return smtp.SendMail(SmtpURL+SmtpPort, smtp.PlainAuth("", email, pass, SmtpURL),
		email, []string{to}, []byte(PrepareEmailContent(email, to, typ, subject, content)))

}

func PrepareEmailContent(from, to, contentType, subject, msg string) string {
	header := make(map[string]string)
	header["From"] = from

	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", contentType)
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	message := ""

	for key, value := range header {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	var encodedMessage bytes.Buffer

	finalMessage := quotedprintable.NewWriter(&encodedMessage)
	finalMessage.Write([]byte(msg))
	finalMessage.Close()

	message += "\r\n" + encodedMessage.String()

	return message
}
