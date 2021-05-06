package delivery

import (
	"net/http"
	"server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"

	"github.com/gorilla/mux"
)

type MessageHandler struct {
	Usecase usecase.MessageUsecaseInterface
	Server  userProto.UserServiceClient
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
	protoMessages, err := m.Server.GetMessages(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), m.Usecase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	var nMesssages []model.Message
	for _, message := range protoMessages.GetMessages() {
		nMesssages = append(nMesssages, m.Usecase.ProtoMessage2Message(message))
	}

	model.Process(model.LoggerFunc("Success: Get Messages", m.Usecase.LogInfo), model.ResponseFunc(w, 200, nMesssages))
}

func (m *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	message, err := m.Server.CreateMessage(r.Context(), m.Usecase.Message2ProtoMessage(newMessage))
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), m.Usecase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	nMessage := m.Usecase.ProtoMessage2Message(message)
	m.Usecase.WebsocketMessage(nMessage)
	model.Process(model.LoggerFunc("Success Create Message", m.Usecase.LogInfo), model.ResponseFunc(w, 200, nMessage))
}

func (m *MessageHandler) ChangeMessage(w http.ResponseWriter, r *http.Request) {
	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, m.Usecase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	protoMessage, err := m.Server.ChangeMessage(r.Context(), m.Usecase.Message2ProtoMessage(newMessage))
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), m.Usecase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	nMessage := m.Usecase.ProtoMessage2Message(protoMessage)
	m.Usecase.WebsocketReactMessage(nMessage)
	model.Process(model.LoggerFunc("New message", m.Usecase.LogInfo), model.ResponseFunc(w, 204, nMessage))
}
