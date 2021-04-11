package usecase

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	mesrep "server/internal/pkg/message/repository"
	model "server/internal/pkg/models"
)

type MessageUsecase struct {
	Db     mesrep.MessageRepository
	Logger *logrus.Entry
}

func (m MessageUsecase) ParseJsonToMessage(body io.ReadCloser) (model.Message, error) {
	var message model.Message
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&message)
	defer body.Close()
	return message, err
}
