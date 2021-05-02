package delivery

import (
	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
	"net/http"
	model "server/internal/pkg/models"
	user_proto "server/internal/user_server/delivery/proto"
	"strconv"
)

func (a *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}
	//
	//userInfo, err := a.UserCase.GetUserInfoById(userId)
	//if err != nil {
	//	response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
	//	response.Description["id"] = "Пользователя с таким id нет"
	//
	//	model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 401, response))
	//	return
	//}
	//
	//model.Process(model.LoggerFunc("Get User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, userInfo))

	ctx := metadata.AppendToOutgoingContext(r.Context(), "id", vars["id"])

	user, err := a.Server.GetUserById(ctx, &user_proto.UserNothing{Dummy: true})
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неправильные входные данные"}
		response.Description["id"] = "Пользователя с таким id нет"
		model.Process(model.LoggerFunc(err, a.UserCase.LogError), model.ResponseFunc(w, 401, response))
		return
	}

	model.Process(model.LoggerFunc("Get User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, user))
}
