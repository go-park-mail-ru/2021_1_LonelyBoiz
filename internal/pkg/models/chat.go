package models

import "github.com/lib/pq"

type Chat struct {
	ChatId              int            `json:"chatId" db:"chatid"`
	PartnerId           int            `json:"partnerId" db:"partnerid"`
	PartnerName         string         `json:"partnerName" db:"partnername"`
	LastMessage         string         `json:"lastMessage,omitempty" db:"lastmessage"`
	LastMessageTime     int64          `json:"lastMessageTime,omitempty" db:"lastmessagetime"`
	LastMessageAuthorId int            `json:"lastMessageAuthor,omitempty" db:"lastmessageauthorid"`
	Photos              pq.StringArray `json:"photos" db:"photos"`
	IsOpened            bool           `json:"isOpened" db:"isopened"`
}

var ChatsChan = make(chan *Chat)
