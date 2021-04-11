package delivery

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"server/internal/pkg/chat/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
)

type ChatHandler struct {
	Sessions *session.SessionsManager
	Usecase  *usecase.ChatUsecase
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *ChatHandler) CreateChat(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		c.Usecase.Logger.Error("Can't get id from context")
	}

	c.Usecase.Clients[id] = ws
}

func (c *ChatHandler) LikesHandler(w http.ResponseWriter, r *http.Request) {}
