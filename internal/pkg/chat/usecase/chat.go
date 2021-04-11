package usecase

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"server/internal/pkg/chat/repository"
	model "server/internal/pkg/models"
)

type ChatUsecase struct {
	Logger    *logrus.Entry
	Db        *repository.ChatRepository
	Clients   map[int]*websocket.Conn
	chatsChan chan *model.Chat
}

func (u *ChatUsecase) WebSocketResponse() {
	for {
		newMessage := <-model.MessagesChan
		partnerId, err := u.Db.GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
		if err != nil {
			u.Logger.Error(err)
			continue
		}

		client := u.Clients[partnerId]
		err = client.WriteJSON(newMessage)
		if err != nil {
			u.Logger.Error(err)
			client.Close()
			delete(u.Clients, partnerId)
		}
	}
}

func (u *ChatUsecase) chatsWriter(newChat *model.Chat) {
	u.chatsChan <- newChat
}
