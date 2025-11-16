package helper

import (
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
	from := "juniansyahanwar@gmail.com"
	password := "mhlw utii qoam bkur"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dailer := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	if err := dailer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

func SendOTPByEmail(to, message string) error {
	from := "juniansyahanwar@gmail.com"
	password := "mhlw utii qoam bkur"

	msg := "Subject: Reset Password\r\n\r\n"
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	return smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
}