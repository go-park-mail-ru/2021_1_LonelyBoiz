package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) BlockSecreteAlbum(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	_, err := a.Server.BlockSecretAlbum(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()), model.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	model.Process(model.LoggerFunc("Successful lock secret album", a.UserCase.LogInfo), model.ResponseFunc(w, 204, nil), model.MetricFunc(204, r, nil))
}
