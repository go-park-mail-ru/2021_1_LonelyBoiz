package delivery

import (
	"net/http"
	models "server/internal/pkg/models"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) AddToSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	user, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogError), models.ResponseFunc(w, 400, response), models.MetricFunc(400, r, response))
		return
	}

	a.UserCase.LogInfo("Передано на сервер USER")
	_, err = a.Server.AddToSecreteAlbum(r.Context(), a.UserCase.User2ProtoUser(user))
	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	models.Process(models.LoggerFunc("Success add photo to secrete album", a.UserCase.LogInfo), models.ResponseFunc(w, 204, nil), models.MetricFunc(204, r, nil))
}
