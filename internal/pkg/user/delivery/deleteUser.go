package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"
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

	cookie, ok := r.Cookie("token")
	if ok != nil {
		model.ResponseFunc(w, 401, nil)
	}

	a.UserCase.DeleteSession(cookie)
	http.SetCookie(w, cookie)

	a.UserCase.LogInfo(cookie.Expires)
	model.ResponseFunc(w, 200, nil)
}
