package usecase

import (
	"server/internal/pkg/chat/repository"
	"server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"github.com/lib/pq"
	"golang.org/x/net/context"
)

type ChatUsecaseInterface interface {
	models.LoggerInterface
	GetChat(userId, limitInt, offsetInt int) ([]models.Chat, error)
	GetIdFromContext(ctx context.Context) (int, bool)
	Photos2ProtoPhotos(userPhotos pq.StringArray) (photos []string)
	ProtoPhotos2Photos(userPhotos []string) (photos pq.StringArray)
	ProtoChat2Chat(chat *userProto.Chat) models.Chat
	Chat2ProtoChat(chat models.Chat) *userProto.Chat
}

type ChatUsecase struct {
	models.LoggerInterface
	Db repository.ChatRepositoryInterface
}

func (u *ChatUsecase) GetChat(userId, limitInt, offsetInt int) ([]models.Chat, error) {
	return u.Db.GetChats(userId, limitInt, offsetInt)
}

func (u *ChatUsecase) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(models.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}

func (u *ChatUsecase) Photos2ProtoPhotos(userPhotos pq.StringArray) (photos []string) {
	for _, photo := range userPhotos {
		photos = append(photos, photo)
	}
	return photos
}

func (u *ChatUsecase) ProtoPhotos2Photos(userPhotos []string) (photos pq.StringArray) {
	for _, photo := range userPhotos {
		photos = append(photos, photo)
	}
	return photos
}

func (u *ChatUsecase) ProtoChat2Chat(chat *userProto.Chat) models.Chat {
	return models.Chat{
		ChatId:              int(chat.GetChatId()),
		PartnerId:           int(chat.GetPartnerId()),
		PartnerName:         chat.GetPartnerName(),
		LastMessage:         chat.GetLastMessage(),
		LastMessageTime:     chat.GetLastMessageTime(),
		LastMessageAuthorId: int(chat.GetLastMessageAuthorId()),
		Photos:              u.ProtoPhotos2Photos(chat.GetPhotos()),
	}
}

func (u *ChatUsecase) Chat2ProtoChat(chat models.Chat) *userProto.Chat {
	return &userProto.Chat{
		ChatId:              int32(chat.ChatId),
		PartnerId:           int32(chat.PartnerId),
		PartnerName:         chat.PartnerName,
		LastMessage:         chat.LastMessage,
		LastMessageTime:     chat.LastMessageTime,
		LastMessageAuthorId: int32(chat.LastMessageAuthorId),
		Photos:              u.Photos2ProtoPhotos(chat.Photos),
	}
}
