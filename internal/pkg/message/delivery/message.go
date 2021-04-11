package delivery

import (
	"net/http"
	mesrep "server/internal/pkg/message/repository"
	"server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
)

type MessageHandler struct {
	Db       mesrep.MessageRepository
	Sessions *session.SessionsManager
	Usecase  *usecase.MessageUsecase
}

func (m *MessageHandler) Message(w http.ResponseWriter, r *http.Request) {
	id, ok := m.Sessions.GetIdFromContext(r.Context())
	if !ok {
		m.Usecase.Logger.Error("Can't get id from context")
	}

	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: model.UserErrorInvalidData}
		model.ResponseWithJson(w, 400, response)
		m.Usecase.Logger.Info(response.Err)
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
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 400, response)
		return
	}

	go messagesWriter(&newMessage)
}

func messagesWriter(newMessage *model.Message) {
	model.MessagesChan <- newMessage
}
