package delivery

import (
	"net/http"
	"server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	a.UserCase.LogInfo("Передано на сервер USER")
	feed, err := a.Server.CreateFeed(r.Context(), &userProto.UserNothing{})
	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	var users []int
	for _, user := range feed.GetUsers() {
		users = append(users, int(user.GetId()))
	}
	if len(users) == 0 {
		users = make([]int, 0)
	}

	models.Process(models.LoggerFunc("Create Feed", a.UserCase.LogInfo), models.ResponseFunc(w, 200, users), models.MetricFunc(200, r, nil))
}
