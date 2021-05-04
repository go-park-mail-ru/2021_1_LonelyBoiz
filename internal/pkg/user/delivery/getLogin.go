package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	"server/internal/pkg/models"
	user_proto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	userResponse, err := a.Server.GetUserById(r.Context(), &user_proto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	nUser, ok := a.UserCase.ProtoUser2User(userResponse)
	if !ok {
		models.Process(models.LoggerFunc("User Proto Error", a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
	}
	models.Process(models.LoggerFunc("Get User Info", a.UserCase.LogInfo), models.ResponseFunc(w, 200, nUser))
}
