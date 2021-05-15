package email

import (
	"log"
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
	AddEmailToQueue(email string)
}

type NotificationByEmail struct {
	Emails *chan string
	Body   string
	models.LoggerInterface
}

func (e *NotificationByEmail) AddEmailToQueue(email string) {
	*e.Emails <- email
}

func (e *NotificationByEmail) SendMessage() {
	for email := range *e.Emails {
		msg := "From: " + emailFrom + "\n" +
			"To: " + email + "\n" +
			"Subject: Привет от Pickle!\n\n" +
			e.Body
		err := smtp.SendMail(
			smtpHost+":"+smtpPort,
			smtp.PlainAuth("", emailFrom, pass, smtpHost),
			emailFrom,
			[]string{email},
			[]byte(msg))
		if err != nil {
			log.Println("Can't send message to " + email + " Error: " + err.Error())
		}
	}
}
