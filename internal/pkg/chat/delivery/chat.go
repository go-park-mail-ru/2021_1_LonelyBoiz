package delivery

import (
	"github.com/gorilla/mux"
	"net/http"
	"server/internal/pkg/chat/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
	"strconv"
)

type ChatHandlerInterface interface {
	GetChats(w http.ResponseWriter, r *http.Request)
	SetChatHandlers(subRouter *mux.Router)
}

type ChatHandler struct {
	Sessions session.SessionManagerInterface
	Usecase  usecase.ChatUsecaseInterface
}

func (c *ChatHandler) GetChats(w http.ResponseWriter, r *http.Request) {
	userId, ok := c.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.NewLogFunc(response.Err, c.Usecase.LogError), model.NewResponseFunc(w, 403, response))
		return
	}

	query := r.URL.Query()
	limit, ok := query["count"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан count"}
		model.Process(model.NewLogFunc(response.Err, c.Usecase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}
	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		model.Process(model.NewLogFunc(response.Err, c.Usecase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}
	offset, ok := query["offset"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан offset"}
		model.Process(model.NewLogFunc(response.Err, c.Usecase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}
	offsetInt, err := strconv.Atoi(offset[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат offset"}
		model.Process(model.NewLogFunc(response.Err, c.Usecase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}

	chats, err := c.Usecase.GetChat(userId, limitInt, offsetInt)
	if err != nil {
		model.Process(model.NewLogFunc(err, c.Usecase.LogError), model.NewResponseFunc(w, 500, nil))
		return
	}

	model.Process(model.NewLogFunc("Success Get Chat", c.Usecase.LogInfo), model.NewResponseFunc(w, 200, chats))
}

func (c *ChatHandler) SetChatHandlers(subRouter *mux.Router) {
	// получить чаты юзера
	subRouter.HandleFunc("/users/{userId:[0-9]+}/chats", c.GetChats).Methods("GET")
}
