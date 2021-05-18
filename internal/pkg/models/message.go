package models

type Message struct {
	MessageId    int    `json:"messageId,omitempty"`
	AuthorId     int    `json:"authorId"`
	ChatId       int    `json:"chatId"`
	Text         string `json:"text"`
	Reaction     int    `json:"reactionId,omitempty"`
	Time         int64  `json:"date,omitempty"`
	MessageOrder int    `json:"messageOrder,omitempty"`
}

type WebsocketResponse struct {
	ResponseType string      `json:"type"`
	Object       interface{} `json:"obj"`
}

type EditedMessage struct {
	Reaction  int `json:"reactionId"`
	MessageId int `json:"messageId"`
	ChatId    int `json:"chatId"`
}

var MessagesChan = make(chan *Message)
