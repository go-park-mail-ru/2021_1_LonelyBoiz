package delivery

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc/status"
	"net/http"
	"server/internal/pkg/chat/usecase"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"
	"strconv"
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
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	//TODO:: конвертация в обычный чат из прото
	model.Process(model.LoggerFunc("Success Get Chat", c.Usecase.LogInfo), model.ResponseFunc(w, 200, chats))
}

func (c *ChatHandler) SetChatHandlers(subRouter *mux.Router) {
	// получить чаты юзера
	subRouter.HandleFunc("/users/{userId:[0-9]+}/chats", c.GetChats).Methods("GET")
}
