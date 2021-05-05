package delivery

import (
	"encoding/json"
	"net/http"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) AddToSecreteAlbum(w http.ResponseWriter, r *http.Request) {
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

	a.UserCase.LogInfo("Передано на сервер USER")
	_, err = a.Server.AddToSecreteAlbum(r.Context(), a.UserCase.User2ProtoUser(user))
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	models.Process(models.LoggerFunc("Success add photo to secrete album", a.UserCase.LogInfo), models.ResponseFunc(w, 204, nil))
	return
}
