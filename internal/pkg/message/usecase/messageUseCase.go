package usecase

import (
	"encoding/json"
	"io"
	mesrep "server/internal/pkg/message/repository"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type MessageUsecase struct {
	Clients      *map[int]*websocket.Conn
	Db           mesrep.MessageRepository
	Logger       *logrus.Entry
	messagesChan chan *model.Message
}

func (m MessageUsecase) ParseJsonToMessage(body io.ReadCloser) (model.Message, error) {
	var message model.Message
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&message)
	defer body.Close()
	return message, err
}

func (m MessageUsecase) messagesWriter(newMessage *model.Message) {
	m.messagesChan <- newMessage
}
