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
	_, ok := a.Sessions.GetIdFromContext(r.Context()) //id
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.LogInfo(response.Err)
		return
	}

	_, err := upgrader.Upgrade(w, r, nil) //ws
	if err != nil {
		model.ResponseWithJson(w, 500, nil)
		a.UserCase.LogError(err)
		return
	}

	//(*a.UserCase.Clients)[id] = ws TODO:: поменять срочно!
}
