package delivery

import (
	"net/http"
	"server/internal/pkg/chat/usecase"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"github.com/gorilla/mux"
	"google.golang.org/grpc/status"
)

type ChatHandlerInterface interface {
	GetChats(w http.ResponseWriter, r *http.Request)
	SetChatHandlers(subRouter *mux.Router)
}

type ChatHandler struct {
	Usecase usecase.ChatUsecaseInterface
	Server  userProto.UserServiceClient
}

func (c *ChatHandler) GetChats(w http.ResponseWriter, r *http.Request) {
	chats, err := c.Server.GetChats(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), c.Usecase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()), model.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}

	var nChats []model.Chat
	for _, chat := range chats.GetChats() {
		nChats = append(nChats, c.Usecase.ProtoChat2Chat(chat))
	}

	model.Process(model.LoggerFunc("Success Get Chat", c.Usecase.LogInfo), model.ResponseFunc(w, 200, nChats), model.MetricFunc(200, r, nil))
}

func (c *ChatHandler) SetChatHandlers(subRouter *mux.Router) {
	// получить чаты юзера
	subRouter.HandleFunc("/users/{userId:[0-9]+}/chats", c.GetChats).Methods("GET")
}
