package delivery

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"
)

func (a *UserHandler) AddToSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	ownerId, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	defer r.Body.Close()
	if err != nil {
		a.UserCase.LogError(err)
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	code, err := a.UserCase.AddToSecreteAlbum(ownerId, user.Photos)
	if code == 500 {
		model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, code, nil))
	}

	model.ResponseFunc(w, 204, nil)
	a.UserCase.LogInfo("Success add photo to secrete album")
}
