package models

import (
	"github.com/google/uuid"
)

type Chat struct {
	ChatId              int         `json:"chatId" db:"chatid"`
	PartnerId           int         `json:"partnerId" db:"partnerid"`
	PartnerName         string      `json:"partnerName" db:"partnername"`
	LastMessage         string      `json:"lastMessage,omitempty" db:"lastmessage"`
	LastMessageTime     int64       `json:"lastMessageTime,omitempty" db:"lastmessagetime"`
	LastMessageAuthorId int         `json:"lastMessageAuthor,omitempty" db:"lastmessageauthorid"`
	Photos              []uuid.UUID `json:"photos" db:"photos"`
}

var ChatsChan = make(chan *Chat)
