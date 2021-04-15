package models

type Chat struct {
	ChatId              int    `json:"chatId"`
	PartnerId           int    `json:"partnerId"`
	PartnerName         string `json:"partnerName"`
	LastMessage         string `json:"lastMessage,omitempty"`
	LastMessageTime     int64  `json:"lastMessageTime,omitempty"`
	LastMessageAuthorId int    `json:"lastMessageAuthor,omitempty"`
	Photos              []int  `json:"photos"`
}

var ChatsChan = make(chan *Chat)
