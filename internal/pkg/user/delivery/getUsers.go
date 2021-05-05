package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	"server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	feed, err := a.Server.CreateFeed(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	var users []int
	for _, user := range feed.GetUsers() {
		users = append(users, int(user.GetId()))
	}

	models.Process(models.LoggerFunc("Create Feed", a.UserCase.LogInfo), models.ResponseFunc(w, 200, users))
	return
}
