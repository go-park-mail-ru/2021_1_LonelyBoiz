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
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 403, response))
		return
	}

	var like model.Like
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&like)
	defer r.Body.Close()
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}

		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 400, response))
		return
	}

	newChat, code, err := a.UserCase.CreateChat(userId, like)
	switch code {
	case 200:
		model.Process(model.NewLogFunc("Create Feed", a.UserCase.LogError), model.NewResponseFunc(w, code, newChat))
	case 204:
		w.WriteHeader(204)
		model.Process(model.NewLogFunc("Return 204 header", a.UserCase.LogInfo), model.NewResponseFunc(w, code, nil))
		return
	case 500:
		model.Process(model.NewLogFunc(err, a.UserCase.LogError), model.NewResponseFunc(w, code, err))
		return
	default:
		model.Process(model.NewLogFunc(err, a.UserCase.LogInfo), model.NewResponseFunc(w, code, err))
		return
	}

	go chatsWriter(&newChat)
}

func chatsWriter(newChat *model.Chat) {
	model.ChatsChan <- newChat
}

func (a *UserHandler) WebSocketChatResponse() {
	for {
		newChat := <-model.ChatsChan
		a.UserCase.WebsocketChat(newChat)
	}
}
