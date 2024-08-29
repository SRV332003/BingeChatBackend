package mail

import (
	"HangAroundBackend/config"

	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

var auth smtp.Auth
var smtpHost string
var smtpPort string
var frontendUrl string
var from string

func sendMail(to []string, bodyBuf bytes.Buffer) error {

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, bodyBuf.Bytes())
	return err
}

func SendVerificationMail(to []string, Name string, token string) error {
	// Load email template from current directory.
	t, err := template.ParseFiles("./services/mail/template.html")
	if err != nil {
		log.Println(err)
		return err
	}

	url := frontendUrl + "/verifyUser?token=" + token

	var bodyBuf bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	bodyBuf.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", "Email Verification for HangAround", mimeHeaders)))

	t.Execute(&bodyBuf, struct {
		LINK string
		NAME string
	}{
		LINK: url,
		NAME: Name,
	})

	err = sendMail(to, bodyBuf)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent successfully")
	}

	return err
}

func init() {
	from = config.GetEnv("MAIL_SERVICE_EMAIL")
	password := config.GetEnv("MAIL_SERVICE_PASSWORD")
	frontendUrl = config.GetEnv("CORS_ORIGIN")
	// smtp server configuration.
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"

	// Authentication.
	auth = smtp.PlainAuth("", from, password, smtpHost)

}
