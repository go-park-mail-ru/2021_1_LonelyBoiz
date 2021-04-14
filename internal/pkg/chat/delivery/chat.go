package delivery

import (
	"net/http"
	chatrep "server/internal/pkg/chat/repository"
	"server/internal/pkg/chat/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/session"
	"strconv"

	"github.com/gorilla/mux"
)

type ChatHandler struct {
	Db       chatrep.ChatRepository
	Sessions *session.SessionsManager
	Usecase  *usecase.ChatUsecase
}

func (c *ChatHandler) GetChats(w http.ResponseWriter, r *http.Request) {
	cookieId, ok := c.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		c.Usecase.Logger.Error(response.Err)
		return
	}

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["chatId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	if userId != cookieId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["messageId"] = "Пытаешься взять не свои чаты"
		model.ResponseWithJson(w, 403, response)
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
