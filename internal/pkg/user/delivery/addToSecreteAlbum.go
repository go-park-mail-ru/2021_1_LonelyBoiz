package delivery

import (
	"net/http"
	model "server/internal/pkg/models"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) AddToSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	user, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response), model.MetricFunc(400, r, response))
		return
	}

	a.UserCase.LogInfo("Передано на сервер USER")
	_, err = a.Server.AddToSecreteAlbum(r.Context(), a.UserCase.User2ProtoUser(user))
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()), model.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	model.Process(model.LoggerFunc("Success add photo to secrete album", a.UserCase.LogInfo), model.ResponseFunc(w, 204, nil), model.MetricFunc(204, r, nil))
}
