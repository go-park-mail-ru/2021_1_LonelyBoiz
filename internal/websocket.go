package internal

import (
	"github.com/gorilla/websocket"
	"net/http"
	_ "net/http"
	model "server/internal/pkg/models"
)

var (
	clients      = make(map[int]*websocket.Conn)
	messagesChan = make(chan *model.Message)
	chatsChan    = make(chan *model.Chat)
	upgrader     = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func webSocketResponse() {
	for {
		newMessage := <-messagesChan
		partnerId, err := GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
		if err != nil {
			//надо залогировать ошибку
			continue
		}

		client := clients[partnerId]
		err = client.WriteJSON(newMessage)
		if err != nil {
			//тут тоже залогировать
			client.Close()
			delete(clients, partnerId)
		}
	}
}
