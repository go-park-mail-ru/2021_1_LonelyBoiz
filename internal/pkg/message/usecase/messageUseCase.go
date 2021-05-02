package usecase

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	mesrep "server/internal/pkg/message/repository"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
)

type MessageUsecaseInterface interface {
	ParseJsonToMessage(body io.ReadCloser) (model.Message, error)
	ManageMessage(userId, chatId, limitInt, offsetInt int) ([]model.Message, int, error)
	model.LoggerInterface
	CreateMessage(newMessage model.Message, chatId, id int) (model.Message, int, error)
	ChangeMessage(userId, messageId int, newMessage model.Message) (model.Message, int, error)
	WebsocketMessage(message model.Message)
	GetIdFromContext(ctx context.Context) (int, bool)
}

type MessageUsecase struct {
	Clients *map[int]*websocket.Conn
	Db      messageRepository.MessageRepositoryInterface

	model.LoggerInterface

	Sanitizer *bluemonday.Policy
}

func (m *MessageUsecase) WebsocketMessage(newMessage model.Message) {
	partnerId, err := m.Db.GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
	if err != nil {
		m.LogError(err)
		return
	}
	if partnerId == -1 {
		m.LogInfo("Нет такого чата")
		return
	}

	var response model.WebsocketResponse

	response.ResponseType = "message"
	response.Object = newMessage

	client, ok := (*m.Clients)[partnerId]
	if !ok {
		m.LogInfo("Пользователь не подключен")
		return
	}

	err = client.WriteJSON(response)
	if err != nil {
		m.LogError("Не удалось отправить сообщение")
		return
	}

	m.LogInfo("Сообщение отправлено")
}

func (m *MessageUsecase) WebsocketReactMessage(newMessage model.Message) {
	fmt.Println(newMessage)
	var response model.WebsocketResponse

	response.ResponseType = "editMessage"
	response.Object = newMessage

	client, ok := (*m.Clients)[newMessage.AuthorId]
	if !ok {
		m.LogInfo("Пользователь не подключен")
		return
	}

	err := client.WriteJSON(response)
	if err != nil {
		m.LogError("Не удалось отправить сообщение")
		return
	}

	m.LogInfo("Сообщение отправлено")
}

func (m *MessageUsecase) ChangeMessage(userId, messageId int, newMessage model.Message) (model.Message, int, error) {
	authorId, err := m.Db.CheckMessageForReacting(userId, messageId)
	if err != nil {
		m.LogError(err)
		return model.Message{}, 500, nil
	}

	if authorId == -1 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userId"] = "Пытаешь поменять сообщение не из своего чата"
		m.LogInfo(response)
		return model.Message{}, 403, response
	}

	if authorId == userId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userId"] = "Пытаешь поставить реакцию на свое сообщение"
		m.LogInfo(response)
		return model.Message{}, 403, response
	}

	err = m.Db.ChangeMessageReaction(messageId, newMessage.Reaction)
	if err != nil {
		m.LogError(err)
		return model.Message{}, 500, err
	}

	newMessage, err = m.Db.GetMessage(messageId)
	if err != nil {
		m.LogError(err)
		return model.Message{}, 500, err
	}

	return newMessage, 204, nil
}

func (m *MessageUsecase) CreateMessage(newMessage model.Message, chatId, id int) (model.Message, int, error) {
	newMessage.ChatId = chatId

	if id != newMessage.AuthorId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userId"] = "Пытаешься отправить сообщение не от своего имени"
		return model.Message{}, 403, response
	}

	if len(newMessage.Text) > 250 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Ошибка валидации"}
		response.Description["text"] = "Слишком длинный текст"
		return model.Message{}, 400, response
	}

	ok, err := m.Db.CheckChat(id, newMessage.ChatId)
	if err != nil {
		return model.Message{}, 500, err
	}

	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["chatId"] = "Пытаешься писать не в свой чат"
		return model.Message{}, 403, response
	}

	newMessage, err = m.Db.AddMessage(newMessage.AuthorId, newMessage.ChatId, newMessage.Text)
	if err != nil {
		return model.Message{}, 500, err
	}

	return newMessage, 200, nil
}

func (m *MessageUsecase) ManageMessage(userId, chatId, limitInt, offsetInt int) ([]model.Message, int, error) {
	ok, err := m.Db.CheckChat(userId, chatId)
	if err != nil {
		m.LogError(err)
		return nil, 500, nil
	}
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["chatId"] = "Пытаешься получить не свой чат"
		m.LogInfo(response)
		return nil, 403, response
	}

	messages, err := m.Db.GetMessages(chatId, limitInt, offsetInt)
	if err != nil {
		m.LogError(err)
		return nil, 500, err
	}

	if len(messages) == 0 {
		messages = make([]model.Message, 0)
	}

	return messages, 200, nil
}

func (m MessageUsecase) ParseJsonToMessage(body io.ReadCloser) (model.Message, error) {
	var message model.Message
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&message)
	defer body.Close()

	message.Text = m.Sanitizer.Sanitize(message.Text)

	return message, err
}

func (m MessageUsecase) messagesWriter(newMessage *model.Message) {
	m.messagesChan <- newMessage
}

func (u *MessageUsecase) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
