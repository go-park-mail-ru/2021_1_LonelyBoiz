package usecase

import (
	"context"
	"net/http"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"

	user_proto "server/internal/user_server/delivery/proto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChat2ProtoChat(t *testing.T) {
	chat := model.Chat{
		ChatId:              1,
		PartnerId:           1,
		PartnerName:         "Serega",
		LastMessage:         "Privet",
		LastMessageTime:     123123,
		LastMessageAuthorId: 1,
		Photos:              []string{"1", "2"},
	}

	protoChat := user_proto.Chat{
		ChatId:              1,
		PartnerId:           1,
		PartnerName:         "Serega",
		LastMessage:         "Privet",
		LastMessageTime:     123123,
		LastMessageAuthorId: 1,
		Photos:              []string{"1", "2"},
	}

	ChatUseCaseTest := ChatUsecase{}

	res := ChatUseCaseTest.Chat2ProtoChat(chat)

	assert.Equal(t, res, &protoChat)
}

func TestProtoChat2Chat(t *testing.T) {
	chat := model.Chat{
		ChatId:              1,
		PartnerId:           1,
		PartnerName:         "Serega",
		LastMessage:         "Privet",
		LastMessageTime:     123123,
		LastMessageAuthorId: 1,
		Photos:              []string{"1", "2"},
	}

	protoChat := user_proto.Chat{
		ChatId:              1,
		PartnerId:           1,
		PartnerName:         "Serega",
		LastMessage:         "Privet",
		LastMessageTime:     123123,
		LastMessageAuthorId: 1,
		Photos:              []string{"1", "2"},
	}

	ChatUseCaseTest := ChatUsecase{}

	res := ChatUseCaseTest.ProtoChat2Chat(&protoChat)

	assert.Equal(t, res, chat)
}

func TestGetIdFromContext(t *testing.T) {
	req := &http.Request{}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	ChatUseCaseTest := ChatUsecase{}

	res, ok := ChatUseCaseTest.GetIdFromContext(ctx)
	assert.Equal(t, ok, true)
	assert.Equal(t, res, 1)
}

func TestGetIdFromContext_Error(t *testing.T) {
	req := &http.Request{}

	ChatUseCaseTest := ChatUsecase{}

	_, ok := ChatUseCaseTest.GetIdFromContext(req.Context())
	assert.Equal(t, ok, false)

}
