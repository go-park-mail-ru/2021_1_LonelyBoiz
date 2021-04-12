package usecase

import (
	"server/internal/pkg/chat/repository"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ChatUsecase struct {
	Clients   *map[int]*websocket.Conn
	Logger    *logrus.Entry
	Db        repository.ChatRepository
	chatsChan chan *model.Chat
}

func (u *ChatUsecase) chatsWriter(newChat *model.Chat) {
	u.chatsChan <- newChat
}
