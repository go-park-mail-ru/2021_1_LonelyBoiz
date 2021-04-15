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
	for {
		newChat := <-model.ChatsChan
		a.UserCase.ChatResponse(newChat)
	}
}
