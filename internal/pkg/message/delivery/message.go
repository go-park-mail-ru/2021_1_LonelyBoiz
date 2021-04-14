package delivery

import (
	"net/http"
	mesrep "server/internal/pkg/message/repository"
	"server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
	"strconv"

	"github.com/gorilla/mux"
)

func messagesWriter(newMessage *model.Message) {
	model.MessagesChan <- newMessage
}

type MessageHandler struct {
	Db       mesrep.MessageRepository
	Sessions *session.SessionsManager
	Usecase  *usecase.MessageUsecase
}

func (m *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	userId, ok := m.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		m.Usecase.Logger.Error(response.Err)
		return
	}

	vars := mux.Vars(r)
	chatId, err := strconv.Atoi(vars["chatId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	query := r.URL.Query()
	limit, ok := query["count"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан count"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	offset, ok := query["offset"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан offset"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	offsetInt, err := strconv.Atoi(offset[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат offset"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	ok, err = m.Db.CheckChat(userId, chatId)
	if err != nil {
		m.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["chatId"] = "Пытаешься получить не свой чат"
		model.ResponseWithJson(w, 403, response)
		return
	}

	messages, err := m.Db.GetMessages(chatId, limitInt, offsetInt)
	if err != nil {
		m.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	if len(messages) == 0 {
		messages = make([]model.Message, 0)
	}
	model.ResponseWithJson(w, 200, messages)
}

func (m *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatId, err := strconv.Atoi(vars["chatId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["chatId"] = "Чата с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	id, ok := m.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 401, response)
		m.Usecase.Logger.Error(response.Err)
		return
	}

	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		m.Usecase.Logger.Error(err)
		return
	}

	newMessage.ChatId = chatId

	if id != newMessage.AuthorId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userId"] = "Пытаешься отправить сообщение не от своего имени"
		model.ResponseWithJson(w, 403, response)
		return
	}

	if len(newMessage.Text) > 250 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Ошибка валидации"}
		response.Description["text"] = "Слишком длинный текст"
		model.ResponseWithJson(w, 400, response)
		return
	}

	ok, err = m.Db.CheckChat(id, newMessage.ChatId)
	if err != nil {
		m.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["chatId"] = "Пытаешься писать не в свой чат"
		model.ResponseWithJson(w, 403, response)
		return
	}

	newMessage, err = m.Db.AddMessage(newMessage.AuthorId, newMessage.ChatId, newMessage.Text)
	if err != nil {
		m.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	go messagesWriter(&newMessage)

	model.ResponseWithJson(w, 200, newMessage)
}

func (m *MessageHandler) ChangeMessage(w http.ResponseWriter, r *http.Request) {
	userId, ok := m.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		m.Usecase.Logger.Error(response.Err)
		return
	}

	vars := mux.Vars(r)
	messageId, err := strconv.Atoi(vars["messageId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	authorId, err := m.Db.CheckMessageForReacting(userId, messageId)
	if err != nil {
		m.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if authorId == -1 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userId"] = "Пытаешь поменять сообщение не из своего чата"
		model.ResponseWithJson(w, 403, response)
		return
	}
	if authorId == userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userId"] = "Пытаешь поставить реакцию на свое сообщение"
		model.ResponseWithJson(w, 403, response)
		return
	}

	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		m.Usecase.Logger.Error(err)
		return
	}

	err = m.Db.ChangeMessageReaction(messageId, newMessage.Reaction)
	if err != nil {
		m.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	model.ResponseWithJson(w, 204, nil)
}

func (m *MessageHandler) WebSocketMessageResponse() {
	for {
		newMessage := <-model.MessagesChan
		partnerId, err := m.Db.GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
		if err != nil {
			m.Usecase.Logger.Error("Пользователь с id = ", newMessage.AuthorId, " не найден")
			continue
		}

		response := model.WebsocketReesponse{ResponseType: "message", Object: newMessage}
		client, ok := (*m.Usecase.Clients)[partnerId]
		if !ok {
			m.Usecase.Logger.Info("Пользователь с id = ", partnerId, " не в сети")
			client.Close()
			delete(*m.Usecase.Clients, partnerId)
			continue
		}

		err = client.WriteJSON(response)
		if err != nil {
			m.Usecase.Logger.Error("Не удалось отправить сообщение")
			client.Close()
			delete(*m.Usecase.Clients, partnerId)
			continue
		}
	}
}
