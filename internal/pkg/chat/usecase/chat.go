package usecase

import (
	"server/internal/pkg/chat/repository"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
)

type ChatUsecaseInterface interface {
	model.LoggerInterface
	chatsWriter(newChat *model.Chat)
}

type ChatUsecase struct {
	Clients *map[int]*websocket.Conn
	model.LoggerInterface
	Db        repository.ChatRepositoryInterface
	chatsChan chan *model.Chat
}

func (u *ChatUsecase) chatsWriter(newChat *model.Chat) {
	u.chatsChan <- newChat
}
