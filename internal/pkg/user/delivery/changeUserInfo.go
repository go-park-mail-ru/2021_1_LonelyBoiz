package delivery

import (
	"net/http"
	model "server/internal/pkg/models"

	"google.golang.org/grpc/status"
)

func (a *UserHandler) ChangeUserInfo(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	a.UserCase.LogInfo("Передано на сервер USER")
	user, err := a.Server.ChangeUser(r.Context(), a.UserCase.User2ProtoUser(newUser))
	if err != nil {
		st, _ := status.FromError(err)
		model.Process(model.LoggerFunc(st.Message(), a.UserCase.LogError), model.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")
	model.Process(model.LoggerFunc("Success Change User Info", a.UserCase.LogInfo), model.ResponseFunc(w, 200, a.UserCase.ProtoUser2User(user)))
}
