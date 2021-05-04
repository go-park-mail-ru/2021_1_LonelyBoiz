package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	user_proto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	user, err := a.Server.GetUserById(r.Context(), &user_proto.UserNothing{Dummy: true})
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, 401, response))
		return
	}

	model.Process(model.LoggerFunc("Get User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, user))
}
