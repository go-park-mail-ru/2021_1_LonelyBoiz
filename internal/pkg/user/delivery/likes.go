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
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 403, response))
		return
	}

	var like model.Like
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&like)
	defer r.Body.Close()
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	newChat, code, err := a.UserCase.CreateChat(userId, like)
	switch code {
	case 204:
		a.UserCase.LogInfo("Success like")
		w.WriteHeader(code)
		return
	case 402:
		model.Process(model.LoggerFunc("Create Feed", a.UserCase.LogInfo), model.ResponseFunc(w, code, err))
		return
	case 500:
		model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, code, nil))
		return
	}

	model.Process(model.LoggerFunc("Successful like", a.UserCase.LogInfo), model.ResponseFunc(w, code, newChat))

	a.UserCase.WebsocketChat(&newChat)
}
