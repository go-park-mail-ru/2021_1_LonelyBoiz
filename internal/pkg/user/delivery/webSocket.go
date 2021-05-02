package delivery

import (
	"github.com/gorilla/websocket"
	"net/http"
	model "server/internal/pkg/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *UserHandler) WsHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := a.UserCase.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}

		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		model.Process(model.LoggerFunc(err.Error(), a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
		return
	}

	a.UserCase.SetChat(ws, id)
	a.UserCase.LogError("Set Websocket connection")
}
