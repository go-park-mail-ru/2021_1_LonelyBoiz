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
			model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
			return
		}
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	cookie, _ := r.Cookie("token")

	a.UserCase.DeleteSession(cookie)
	http.SetCookie(w, cookie)

	a.UserCase.LogInfo(cookie.Expires)
	model.ResponseFunc(w, 200, nil)
}
