package email

import (
	"log"
	"net/smtp"
	"os"
	"server/internal/pkg/models"
)

type NotificationInterface interface {
	SendMessage()
	AddEmailLetterToQueue(email string, body string)
}

type NotificationByEmail struct {
	Buckets *chan EmailBucket
	models.LoggerInterface
}

type EmailBucket struct {
	email string
	body  string
}

func (e *NotificationByEmail) AddEmailLetterToQueue(email string, body string) {
	*e.Buckets <- EmailBucket{
		email: email,
		body:  body,
	}
}

func (e *NotificationByEmail) SendMessage() {
	emailFrom := os.Getenv("EMAILFROM")
	pass := os.Getenv("PASS")
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := os.Getenv("SMTPPORT")

	for bucket := range *e.Buckets {
		msg := "From: " + emailFrom + "\n" +
			"To: " + bucket.email + "\n" +
			"Subject: Привет от Pickle!\n\n" +
			bucket.body

		err := smtp.SendMail(
			smtpHost+":"+smtpPort,
			smtp.PlainAuth("", emailFrom, pass, smtpHost),
			emailFrom,
			[]string{bucket.email},
			[]byte(msg))

		if err != nil {
			log.Println("Can't send message to " + bucket.email + " Error: " + err.Error())
		}
	}
}
