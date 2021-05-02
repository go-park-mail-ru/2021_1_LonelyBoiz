package delivery

import (
	"net/http"
	"server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

type MessageHandler struct {
	Usecase usecase.MessageUsecaseInterface
}

func (m *MessageHandler) SetMessageHandlers(subRouter *mux.Router) {
	// получить сообщения из чата
	subRouter.HandleFunc("/chats/{chatId:[0-9]+}/messages", m.GetMessages).Methods("GET")
	// отправка нового сообщения
	subRouter.HandleFunc("/chats/{chatId:[0-9]+}/messages", m.SendMessage).Methods("POST")
	// реакция
	subRouter.HandleFunc("/messages/{messageId:[0-9]+}", m.ChangeMessage).Methods("PATCH")
}

func (m *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	userId, ok := m.Usecase.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogError), model.ResponseFunc(w, 403, response))
		return
	}

	vars := mux.Vars(r)
	chatId, err := strconv.Atoi(vars["chatId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	query := r.URL.Query()
	limit, ok := query["count"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан count"}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}
	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}
	offset, ok := query["offset"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан offset"}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}
	offsetInt, err := strconv.Atoi(offset[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат offset"}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	messages, code, err := m.Usecase.ManageMessage(userId, chatId, limitInt, offsetInt)
	switch code {
	case 200:
		model.Process(model.LoggerFunc("Success: Get Messages", m.Usecase.LogInfo), model.ResponseFunc(w, code, messages))
	case 500:
		model.Process(model.LoggerFunc(err, m.Usecase.LogError), model.ResponseFunc(w, code, nil))
	default:
		model.Process(model.LoggerFunc(err, m.Usecase.LogInfo), model.ResponseFunc(w, code, err))
	}
}

func (m *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatId, err := strconv.Atoi(vars["chatId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["chatId"] = "Чата с таким id нет"
		model.Process(model.LoggerFunc(err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	id, ok := m.Usecase.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogError), model.ResponseFunc(w, 401, response))
		return
	}

	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	newMessage, code, err := m.Usecase.CreateMessage(newMessage, chatId, id)
	switch code {
	case 200:
	case 500:
		model.Process(model.LoggerFunc(err, m.Usecase.LogError), model.ResponseFunc(w, 500, nil))
		return
	default:
		model.Process(model.LoggerFunc(err, m.Usecase.LogInfo), model.ResponseFunc(w, code, err))
		return
	}

	model.Process(model.LoggerFunc("Success Create Message", m.Usecase.LogInfo), model.ResponseFunc(w, 200, newMessage))

	// отправить сообщение по вэбсокету
	m.Usecase.WebsocketMessage(newMessage)
}

func (m *MessageHandler) ChangeMessage(w http.ResponseWriter, r *http.Request) {
	userId, ok := m.Usecase.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	vars := mux.Vars(r)
	messageId, err := strconv.Atoi(vars["messageId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	newMessage, code, err := m.Usecase.ChangeMessage(userId, messageId, newMessage)
	if code != 204 {
		model.ResponseFunc(w, code, err)
		return
	}

	model.Process(model.LoggerFunc("New message", m.Usecase.LogInfo), model.ResponseFunc(w, 204, err))

	// отправить сообщение в вэбсокет
	m.Usecase.WebsocketReactMessage(newMessage)
}
