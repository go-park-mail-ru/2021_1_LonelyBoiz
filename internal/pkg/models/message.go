package models

type Message struct {
	MessageId    int    `json:"messageId,omitempty"`
	AuthorId     int    `json:"authorId"`
	ChatId       int    `json:"chatId"`
	Text         string `json:"text"`
	Reaction     int    `json:"reactioId,omitempty"`
	Time         int64  `json:"date,omitempty"`
	MessageOrder int    `json:"messageOrder,omitempty"`
}
