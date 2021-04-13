package delivery

import (
	"encoding/json"
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) LikesHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	var like model.Like
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&like)
	defer r.Body.Close()
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		a.UserCase.Logger.Error(err)
		return
	}

	if like.Reaction != "like" && like.Reaction != "skip" {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["like"] = "неправильный формат реацкции, ожидается skip или like"
		model.ResponseWithJson(w, 400, response)
		return
	}

	rowsAffected, err := a.Db.Rating(userId, like.UserId, like.Reaction)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if rowsAffected != 1 {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["userID"] = "Пытаешься поставить лайк человеку не со своей ленты"
		model.ResponseWithJson(w, 403, response)
		return
	}

	reciprocity, err := a.Db.CheckReciprocity(like.UserId, userId)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if reciprocity == false || like.Reaction == "skip" {
		w.WriteHeader(204)
		return
	}

	var newChat model.Chat
	newChat.ChatId, err = a.Db.CreateChat(userId, like.UserId)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	newChat.PartnerId = like.UserId
	newChat.Photos, err = a.Db.GetPhotos(newChat.PartnerId)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	go chatsWriter(&newChat)

	model.ResponseWithJson(w, 200, newChat)
}

func chatsWriter(newChat *model.Chat) {
	model.ChatsChan <- newChat
}

func (a *UserHandler) WebSocketChatResponse() {
	for {
		newChat := <-model.ChatsChan

		client, ok := (*a.UserCase.Clients)[newChat.PartnerId]
		if !ok {
			a.UserCase.Logger.Info("Пользователь с id = ", newChat.PartnerId, " не в сети")
			continue
		}

		newChatToSend, err := a.Db.GetNewChatById(newChat.ChatId, newChat.PartnerId)
		if err != nil {
			a.UserCase.Logger.Error("Не удалось составить чат", err)
			continue
		}

		if len(newChatToSend.Photos) == 0 {
			newChatToSend.Photos = make([]int, 0)
		}

		response := model.WebsocketReesponse{ResponseType: "chat", Object: newChatToSend}

		err = client.WriteJSON(response)
		if err != nil {
			a.UserCase.Logger.Error("Не удалось отправить сообщение")
			client.Close()
			delete(*a.UserCase.Clients, newChat.PartnerId)
			continue
		}
	}
}
