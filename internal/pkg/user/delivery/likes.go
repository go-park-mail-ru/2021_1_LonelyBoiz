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
		a.UserCase.LogInfo(response.Err)
		return
	}

	var like model.Like
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&like)
	defer r.Body.Close()
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.ResponseWithJson(w, 400, response)
		a.UserCase.LogError(err)
		return
	}

	chat, code, err := a.UserCase.SetLike(like, userId)

	switch code {
	case 200:
		go chatsWriter(&chat)
		model.ResponseWithJson(w, 200, chat)
	case 204:
		w.WriteHeader(204)
	default:
		model.ResponseWithJson(w, code, err)
	}
}

func chatsWriter(newChat *model.Chat) {
	model.ChatsChan <- newChat
}

func (a *UserHandler) WebSocketChatResponse() {
	//for {
	//	newChat := <-model.ChatsChan
	//
	//	client, ok := (*a.UserCase.Clients)[newChat.PartnerId]
	//	if !ok {
	//		a.UserCase.Logger.Info("Пользователь с id = ", newChat.PartnerId, " не в сети")
	//		continue
	//	}
	//
	//	newChatToSend, err := a.UserCase.Db.GetNewChatById(newChat.ChatId, newChat.PartnerId)
	//
	//	if err != nil {
	//		a.UserCase.Logger.Error("Не удалось составить чат", err)
	//		continue
	//	}
	//
	//	if len(newChatToSend.Photos) == 0 {
	//		newChatToSend.Photos = make([]int, 0)
	//	}
	//
	//	response := model.WebsocketReesponse{ResponseType: "chat", Object: newChatToSend}
	//
	//	err = client.WriteJSON(response)
	//	if err != nil {
	//		a.UserCase.Logger.Error("Не удалось отправить сообщение")
	//		client.Close()
	//		delete(*a.UserCase.Clients, newChat.PartnerId)
	//		continue
	//	}
	//}
}
