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
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, 401, response), model.MetricFunc(401, r, err))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	ret := a.UserCase.ProtoUser2User(user)

	if len(ret.Photos) == 0 {
		ret.Photos = make([]string, 0)
	}

	a.UserCase.LogInfo("Получен результат из сервера USER")
	model.Process(model.LoggerFunc("Get User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, ret), model.MetricFunc(200, r, nil))
}
