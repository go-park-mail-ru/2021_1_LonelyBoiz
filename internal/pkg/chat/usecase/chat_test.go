package usecase

import (
	"context"
	"net/http"
	"reflect"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"

	chat_rep "server/internal/pkg/chat/repository"
	user_proto "server/internal/user_server/delivery/proto"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
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

func TestGetChats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	userId := 1
	limit := 10
	offset := 10

	chat := models.Chat{
		ChatId:              1,
		PartnerId:           1,
		PartnerName:         "Name",
		LastMessage:         "last message",
		LastMessageTime:     123,
		LastMessageAuthorId: 1,
		Photos:              []string{"1", "2"},
	}

	rows := sqlmock.NewRows([]string{
		"chatid",
		"partnerid",
		"partnername",
		"lastmessage",
		"lastmessagetime",
		"lastmessageauthorid",
		"photos",
	}).AddRow(
		chat.ChatId,
		chat.PartnerId,
		chat.PartnerName,
		chat.LastMessage,
		chat.LastMessageTime,
		chat.LastMessageAuthorId,
		chat.Photos,
	)

	mock.
		ExpectQuery("SELECT chats.id AS chatId,").
		WithArgs(userId, limit, offset).
		WillReturnRows(rows)

	repo := chat_rep.ChatRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetChats(userId, limit, offset)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res[0], chat) {
		t.Errorf("results not match, want %v, have %v", chat, res)
		return
	}
}
