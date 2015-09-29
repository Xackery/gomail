package gomail

import (
	"errors"
	"net/smtp"
)

func Send(email string, password string, to []string, subject string, content string) (err error) {
	if len(to) < 1 {
		err = errors.New("No recipient specified")
		return
	}
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		hostname,
	)

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		content)

	err = smtp.SendMail(
		hostname+":587",
		auth,
		email,
		to,
		msg,
	)
	return
}
