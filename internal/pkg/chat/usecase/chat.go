package usecase

import (
	"golang.org/x/net/context"
	"server/internal/pkg/chat/repository"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
)

type ChatUsecaseInterface interface {
	model.LoggerInterface
	chatsWriter(newChat *model.Chat)
	GetChat(userId, limitInt, offsetInt int) ([]model.Chat, error)
	GetIdFromContext(ctx context.Context) (int, bool)
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

func (u *ChatUsecase) chatsWriter(newChat *model.Chat) {
	u.chatsChan <- newChat
}

func (u *ChatUsecase) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
