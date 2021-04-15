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
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	a.UserCase.Logger.Debug("Новое подключение по вэбсокету с id = ", id)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		model.ResponseWithJson(w, 500, nil)
		a.UserCase.Logger.Error(err)
		return
	}

	(*a.UserCase.Clients)[id] = ws
	a.UserCase.Logger.Debug("Текущие подключиения к вэбсокету", (*a.UserCase.Clients))
}
