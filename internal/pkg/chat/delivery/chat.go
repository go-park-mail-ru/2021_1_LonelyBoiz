package delivery

import (
	"net/http"
	chatrep "server/internal/pkg/chat/repository"
	"server/internal/pkg/chat/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
	"strconv"
)

type ChatHandler struct {
	Db       chatrep.ChatRepository
	Sessions *session.SessionsManager
	Usecase  *usecase.ChatUsecase
}

func (c *ChatHandler) GetChats(w http.ResponseWriter, r *http.Request) {
	userId, ok := c.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		c.Usecase.Logger.Error(response.Err)
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

	chats, err := c.Db.GetChats(userId, limitInt, offsetInt)
	if err != nil {
		c.Usecase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	model.ResponseWithJson(w, 200, chats)
}
