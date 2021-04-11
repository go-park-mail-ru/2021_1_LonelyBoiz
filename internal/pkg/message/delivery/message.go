package delivery

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	mesrep "server/internal/pkg/message/repository"
	"server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
	"server/internal/pkg/user/repository"
)

type MessageHandler struct {
	Db       mesrep.MessageRepository
	Sessions *session.SessionsManager
	Logger   *logrus.Entry
	usecase  *usecase.MessageUsecase
}

func (m *MessageHandler) Message(w http.ResponseWriter, r *http.Request) {
	id, ok := m.Sessions.GetIdFromContext(r.Context())
	if !ok {
		m.Logger.Info("Can't get id from context")
	}

	newMessage, err := m.usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: model.UserErrorInvalidData}
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	if id != newMessage.AuthorId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 401, response)
		return
	}

	if len(newMessage.Text) > 250 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Ошибка валидации"}
		response.Description["text"] = "Слишком длинный текст"
		model.ResponseWithJson(w, 400, response)
		return
	}

	newMessage, err = m.Db.AddMessage(newMessage.AuthorId, newMessage.ChatId, newMessage.Text)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err}
		model.ResponseWithJson(w, 400, response)
		return
	}

	go messagesWriter(&newMessage)
}

func messagesWriter(newMessage *model.Message) {
	messagesChan <- newMessage
}
