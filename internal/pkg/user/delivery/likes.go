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
	if reciprocity == false {
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
	go chatsWriter(&newChat)

	model.ResponseWithJson(w, 200, newChat)
}

func chatsWriter(newChat *model.Chat) {
	model.ChatsChan <- newChat
}

/*func (c *UserHandler) WebSocketChatResponse() {
	for {
		newChat := <-model.ChatsChan
		partnerId, err := m.Db.GetPartnerId(newMessage.ChatId, newMessage.AuthorId)
		if err != nil {
			m.Usecase.Logger.Error("Пользователь с id = ", newMessage.AuthorId, " не найден")
			continue
		}

		response := model.WebsocketReesponse{ResponseType: "message", Object: newMessage}
		client, ok := (*m.Usecase.Clients)[partnerId]
		if !ok {
			m.Usecase.Logger.Info("Пользователь с id = ", partnerId, " не в сети")
			client.Close()
			delete(*m.Usecase.Clients, partnerId)
			continue
		}

		err = client.WriteJSON(response)
		if err != nil {
			m.Usecase.Logger.Error("Не удалось отправить сообщение")
			client.Close()
			delete(*m.Usecase.Clients, partnerId)
			continue
		}
	}
}

/*for {
		newChat := <-chatsChan
		partnerId := newChat.PartnerId
		response := Resp{ResponseType: "chat", Object: newChat}
		client, ok := clients[partnerId]
		if !ok {
			client.Close()
			delete(clients, partnerId)
			continue
		}
		err := client.WriteJSON(response)
		if err != nil {
			//тут тоже залогировать
			client.Close()
			delete(clients, partnerId)
		}
	}
}*/
