package delivery

import (
	"fmt"
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
	a.UserCase.LogError("Попытка подключиться по вэбсокету")
	id, ok := a.UserCase.GetIdFromContext(r.Context())
	if !ok {
		fmt.Println("Ne ok")
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Ne up")
		model.Process(model.LoggerFunc(err.Error(), a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
		return
	}

	a.UserCase.SetChat(ws, id)
	a.UserCase.LogInfo(fmt.Sprintf("Новое подключение по вэбсокету id =%d", id))
}
