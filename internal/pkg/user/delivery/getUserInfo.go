package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	user_proto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	user, err := a.Server.GetUserById(r.Context(), &user_proto.UserNothing{Dummy: true})
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	ret := a.UserCase.ProtoUser2User(user)

	if len(ret.Photos) == 0 {
		ret.Photos = make([]string, 0)
	}
	model.Process(model.LoggerFunc("Get User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, ret))
}
