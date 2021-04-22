package models

type Chat struct {
	ChatId              int    `json:"chatId" db:"ChatId"`
	PartnerId           int    `json:"partnerId" db:"partnerId"`
	PartnerName         string `json:"partnerName" db:"partnerName"`
	LastMessage         string `json:"lastMessage,omitempty" db:"lastMessage"`
	LastMessageTime     int64  `json:"lastMessageTime,omitempty" db:"lastMessageTime"`
	LastMessageAuthorId int    `json:"lastMessageAuthor,omitempty" db:"lastMessageAuthorId"`
	Photos              []int  `json:"photos" db:"photos"`
}

var ChatsChan = make(chan *Chat)
