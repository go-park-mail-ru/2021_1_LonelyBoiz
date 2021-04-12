package delivery

import (
	chatrep "server/internal/pkg/chat/repository"
	"server/internal/pkg/chat/usecase"
	"server/internal/pkg/session"
)

type ChatHandler struct {
	Db       chatrep.ChatRepository
	Sessions *session.SessionsManager
	Usecase  *usecase.ChatUsecase
}

/*
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

	(*c.Usecase).Clients[id] = ws
}

func (c *ChatHandler) LikesHandler(w http.ResponseWriter, r *http.Request) {
	cookieId, ok := c.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		c.Usecase.Logger.Info(response.Err)
		return
	}

	var like model.Like
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&like)
	if err != nil {
		c.Usecase.Logger.Info(err)
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	if like.Reaction != "like" || like.Reaction != "skip" {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["like"] = "неправильный формат реацкции ожидается skip или like"
		model.ResponseWithJson(w, 400, response)
		return
	}

	rowsAffected, err := c.Usecase.Db.Rating(id, like.UserId, like.Reaction)
	if err != nil {
		//залогировать ошибку
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		model.ResponseWithJson(w, 500, response)
		return
	}
	if rowsAffected == -1 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userID"] = "Пытаешься поставить лайк человеку не со своей ленты"
		mpdel.ResponseWithJson(w, 403, response)
		return
	}
	reciprocity, err := DB.CheckReciprocity(like.UserId, id)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось проверить взаимность"}
		responseWithJson(w, 500, response)
		return
	}
	if reciprocity == false {
		w.WriteHeader(204)
		return
	}
	var newChat Chat
	newChat.ChatId, err = DB.CreateChat(id, like.UserId)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось создать чат"}
		responseWithJson(w, 500, response)
		return
	}
	newChat, err = DB.GetNewChat(newChat.ChatId, like.UserId)
	if err != nil {
		//залогировать ошибку
		response := errorDescriptionResponse{Description: map[string]string{}, Err: "Не удалось получить новый чат из базы"}
		responseWithJson(w, 500, response)
		return
	}
	go chatsWriter(&newChat)
	responseWithJson(w, 200, newChat)
}
*/
