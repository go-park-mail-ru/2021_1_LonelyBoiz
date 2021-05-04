package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) GetSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	getterId, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	vars := mux.Vars(r)
	ownerId, err := strconv.Atoi(vars["userId"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	photos, code, err := a.UserCase.GetSecreteAlbum(ownerId, getterId)
	if code == 500 {
		model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, code, nil))
	}

	model.ResponseFunc(w, code, photos)
	a.UserCase.LogInfo("Success unblock secrete album")
}