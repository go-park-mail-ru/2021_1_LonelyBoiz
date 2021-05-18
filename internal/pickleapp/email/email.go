package email

import (
	"net/smtp"
	"server/internal/pkg/models"
)

const (
	emailFrom = "nikita.nackaznoy@gmail.com"
	pass      = "Nn070202"
	smtpHost  = "smtp.gmail.com"
	smtpPort  = "587"
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
			e.LogError("Can't send message to " + bucket.email + " Error: " + err.Error())
			//log.Println("Can't send message to " + bucket.email + " Error: " + err.Error())
		}
	}
}
