package delivery

import (
	"net/http"
	model "server/internal/pkg/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (a *UserHandler) WsHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}

		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 403, response))
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		model.Process(model.NewLogFunc(err.Error(), a.UserCase.LogError), model.NewResponseFunc(w, 500, nil))
		return
	}

	a.UserCase.SetChat(ws, id)
	a.UserCase.LogError("Set Websocket connection")
}
