package usecase

import (
	"server/internal/pkg/chat/repository"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
)

type ChatUsecaseInterface interface {
	model.LoggerInterface
	ChatsWriter(newChat *model.Chat)
	GetChat(userId, limitInt, offsetInt int) ([]model.Chat, error)
}

type ChatUsecase struct {
	Clients *map[int]*websocket.Conn
	model.LoggerInterface
	Db        repository.ChatRepositoryInterface
	chatsChan chan *model.Chat
}

func (u *ChatUsecase) GetChat(userId, limitInt, offsetInt int) ([]model.Chat, error) {
	return u.Db.GetChats(userId, limitInt, offsetInt)
}

func (u *ChatUsecase) ChatsWriter(newChat *model.Chat) {
	u.chatsChan <- newChat
}
