package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	_, err := a.Server.DeleteChat(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()), model.MetricFunc(int(st.Code()), r, st.Err()))
			return
		}
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	model.Process(model.LoggerFunc("Delete Chat", a.UserCase.LogInfo), model.ResponseFunc(w, 204, nil), model.MetricFunc(204, r, nil))
}
