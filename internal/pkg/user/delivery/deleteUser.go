package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	_, err := a.Server.DeleteUser(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()), model.MetricFunc(int(st.Code()), r, st.Err()))
			return
		}
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	cookie, _ := r.Cookie("token")

	a.UserCase.DeleteSession(cookie)
	http.SetCookie(w, cookie)

	model.Process(model.LoggerFunc("Delete User", a.UserCase.LogInfo), model.ResponseFunc(w, 200, nil), model.MetricFunc(200, r, nil))
}
