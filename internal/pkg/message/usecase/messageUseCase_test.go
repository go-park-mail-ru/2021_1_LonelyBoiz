package usecase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"server/internal/pkg/message/repository/mocks"
	"server/internal/pkg/models"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	message := models.Message{
		Text:     "some text",
		AuthorId: 1,
	}

	dbMock.EXPECT().CheckChat(id, chatId).Return(true, nil)
	dbMock.EXPECT().AddMessage(message.AuthorId, chatId, message.Text).Return(message, nil)

	_, code, err := UserUsecaseTest.CreateMessage(message, chatId, id)

	assert.Equal(t, 200, code)
	assert.Equal(t, nil, err)
}

func TestCreateMessageIdNotEqual(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	message := models.Message{
		Text:     "some text",
		AuthorId: 2,
	}

	_, code, err := UserUsecaseTest.CreateMessage(message, chatId, id)

	assert.Equal(t, 403, code)
	assert.NotEqual(t, nil, err)
}

func TestCreateMessageTextLenError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	asdf := make([]byte, 251)

	message := models.Message{
		Text:     string(asdf),
		AuthorId: 1,
	}

	_, code, err := UserUsecaseTest.CreateMessage(message, chatId, id)

	assert.Equal(t, 400, code)
	assert.NotEqual(t, nil, err)
}

func TestCreateMessageCheckChatError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	message := models.Message{
		Text:     "some text",
		AuthorId: 1,
	}

	dbMock.EXPECT().CheckChat(id, chatId).Return(true, errors.New("some error"))

	_, code, err := UserUsecaseTest.CreateMessage(message, chatId, id)

	assert.Equal(t, 500, code)
	assert.NotEqual(t, nil, err)
}

func TestCreateMessageCheckChatNotOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	message := models.Message{
		Text:     "some text",
		AuthorId: 1,
	}

	dbMock.EXPECT().CheckChat(id, chatId).Return(false, nil)

	_, code, err := UserUsecaseTest.CreateMessage(message, chatId, id)

	assert.Equal(t, 403, code)
	assert.NotEqual(t, nil, err)
}

func TestCreateMessageAddMessageError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	message := models.Message{
		Text:     "some text",
		AuthorId: 1,
	}

	dbMock.EXPECT().CheckChat(id, chatId).Return(true, nil)
	dbMock.EXPECT().AddMessage(message.AuthorId, chatId, message.Text).Return(message, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateMessage(message, chatId, id)

	assert.Equal(t, 500, code)
	assert.NotEqual(t, nil, err)
}

func TestManageMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	limit := 10
	offset := 10

	dbMock.EXPECT().CheckChat(id, chatId).Return(true, nil)
	dbMock.EXPECT().GetMessages(chatId, limit, offset).Return([]models.Message{}, nil)

	_, code, err := UserUsecaseTest.ManageMessage(id, chatId, limit, offset)

	assert.Equal(t, 200, code)
	assert.Equal(t, nil, err)
}

func TestManageMessageCheckChatError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	limit := 10
	offset := 10

	dbMock.EXPECT().CheckChat(id, chatId).Return(true, errors.New("Some error"))

	_, code, err := UserUsecaseTest.ManageMessage(id, chatId, limit, offset)

	assert.Equal(t, 500, code)
	assert.Equal(t, nil, err)
}

func TestManageMessageCheckChatNotOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	limit := 10
	offset := 10

	dbMock.EXPECT().CheckChat(id, chatId).Return(false, nil)

	_, code, err := UserUsecaseTest.ManageMessage(id, chatId, limit, offset)

	assert.Equal(t, 403, code)
	assert.NotEqual(t, nil, err)
}

func TestManageMessageGetMessagesError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	chatId := 1
	id := 1

	limit := 10
	offset := 10

	message := models.Message{
		Text:     "some text",
		AuthorId: 1,
	}

	dbMock.EXPECT().CheckChat(id, chatId).Return(true, nil)
	dbMock.EXPECT().GetMessages(chatId, limit, offset).Return([]models.Message{message}, errors.New("Some error"))

	_, code, err := UserUsecaseTest.ManageMessage(id, chatId, limit, offset)

	assert.Equal(t, 500, code)
	assert.NotEqual(t, nil, err)
}

func TestChangeMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	id := 1

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  1,
		Reaction:  1,
	}

	dbMock.EXPECT().CheckMessageForReacting(id, message.MessageId).Return(2, nil)
	dbMock.EXPECT().ChangeMessageReaction(message.MessageId, message.Reaction).Return(nil)

	_, code, err := UserUsecaseTest.ChangeMessage(id, message.MessageId, message)

	assert.Equal(t, 204, code)
	assert.Equal(t, nil, err)
}

func TestChangeMessageCheckMessageForReactingError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	id := 1

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  1,
		Reaction:  1,
	}

	dbMock.EXPECT().CheckMessageForReacting(id, message.MessageId).Return(2, errors.New("Some string"))

	_, code, err := UserUsecaseTest.ChangeMessage(id, message.MessageId, message)

	assert.Equal(t, 500, code)
	assert.Equal(t, nil, err)
}

func TestChangeMessageNoAuthor(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	id := 1

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  1,
		Reaction:  1,
	}

	dbMock.EXPECT().CheckMessageForReacting(id, message.MessageId).Return(-1, nil)

	_, code, err := UserUsecaseTest.ChangeMessage(id, message.MessageId, message)

	assert.Equal(t, 403, code)
	assert.NotEqual(t, nil, err)
}

func TestChangeMessageSelfREactingError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	id := 1

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  1,
		Reaction:  1,
	}

	dbMock.EXPECT().CheckMessageForReacting(id, message.MessageId).Return(1, nil)

	_, code, err := UserUsecaseTest.ChangeMessage(id, message.MessageId, message)

	assert.Equal(t, 403, code)
	assert.NotEqual(t, nil, err)
}

func TestChangeMessageReqctionError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	id := 1

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  1,
		Reaction:  1,
	}

	dbMock.EXPECT().CheckMessageForReacting(id, message.MessageId).Return(2, nil)
	dbMock.EXPECT().ChangeMessageReaction(message.MessageId, message.Reaction).Return(errors.New("Some error"))

	_, code, err := UserUsecaseTest.ChangeMessage(id, message.MessageId, message)

	assert.Equal(t, 500, code)
	assert.NotEqual(t, nil, err)
}

func TestParseJsonToMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockMessageRepositoryInterface(mockCtrl)

	UserUsecaseTest := MessageUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	message := models.Message{
		MessageId: 1,
		Text:      "some text",
		AuthorId:  1,
		Reaction:  1,
	}

	buf, err := json.Marshal(message)
	assert.Equal(t, nil, err)

	r := ioutil.NopCloser(strings.NewReader(string(buf)))

	_, err = UserUsecaseTest.ParseJsonToMessage(r)
	assert.Equal(t, nil, err)
}
