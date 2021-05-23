package delivery

import (
	"net/http"
	"server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
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
		models.Process(models.LoggerFunc(st.Message(), m.Usecase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}

	var nMesssages []models.Message
	for _, message := range protoMessages.GetMessages() {
		nMesssages = append(nMesssages, m.Usecase.ProtoMessage2Message(message))
	}
	if len(nMesssages) == 0 {
		nMesssages = make([]models.Message, 0)
	}

	models.Process(models.LoggerFunc("Success: Get Messages", m.Usecase.LogInfo), models.ResponseFunc(w, 200, nMesssages), models.MetricFunc(200, r, nil))
}

func (m *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(err, m.Usecase.LogInfo), models.ResponseFunc(w, 400, response), models.MetricFunc(400, r, response))
		return
	}

	message, err := m.Server.CreateMessage(r.Context(), m.Usecase.Message2ProtoMessage(newMessage))
	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), m.Usecase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}

	nMessage := m.Usecase.ProtoMessage2Message(message)
	m.Usecase.WebsocketMessage(nMessage)
	models.Process(models.LoggerFunc("Success Create Message", m.Usecase.LogInfo), models.ResponseFunc(w, 200, nMessage), models.MetricFunc(200, r, nil))
}

func (m *MessageHandler) ChangeMessage(w http.ResponseWriter, r *http.Request) {
	newMessage, err := m.Usecase.ParseJsonToMessage(r.Body)
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, m.Usecase.LogInfo), models.ResponseFunc(w, 400, response), models.MetricFunc(400, r, response))
		return
	}

	protoMessage, err := m.Server.ChangeMessage(r.Context(), m.Usecase.Message2ProtoMessage(newMessage))
	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), m.Usecase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}

	nMessage := m.Usecase.ProtoMessage2Message(protoMessage)
	m.Usecase.WebsocketReactMessage(nMessage)
	models.Process(models.LoggerFunc("New message", m.Usecase.LogInfo), models.ResponseFunc(w, 204, nMessage), models.MetricFunc(204, r, nil))
}
