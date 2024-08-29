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
var from string

func SendMail(to []string, Name string, subject string, bodyBuf bytes.Buffer) error {

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, bodyBuf.Bytes())
	return err
}

func SendOTPMail(to []string, Name string, OTP int) error {
	// Load email template from current directory.
	t, err := template.ParseFiles("auth/services/mail/verifyTemplate.html")
	if err != nil {
		log.Println(err)
		return err
	}

	var bodyBuf bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	bodyBuf.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", " Strive OTP Verification", mimeHeaders)))

	t.Execute(&bodyBuf, struct {
		OTP int
	}{
		OTP: OTP,
	})

	err = SendMail(to, Name, "Email Verification", bodyBuf)
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
	// smtp server configuration.
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"

	// Authentication.
	auth = smtp.PlainAuth("", from, password, smtpHost)

}
