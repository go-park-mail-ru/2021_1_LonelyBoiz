package models

type Chat struct {
	ChatId              int    `json:"chatId"`
	PartnerId           int    `json:"partnerId"`
	PartnerName         string `json:"partnerName"`
	LastMessage         string `json:"lastMessage"`
	LastMessageTime     int64  `json:"lastMessageTime"`
	LastMessageAuthorId int    `json:"lastMessageAuthor"`
	Avatar              string `json:"pathToAvatar"`
}

type Like struct {
	UserId   int    `json:"userId"`
	Reaction string `json:"reaction"`
}
