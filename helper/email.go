package helper

import "gopkg.in/gomail.v2"

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