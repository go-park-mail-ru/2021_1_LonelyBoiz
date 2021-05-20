package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	"server/internal/pkg/models"
	userProto "server/internal/user_server/delivery/proto"
)

func (a *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.LogError(err.Error())
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 401, response), models.MetricFunc(401, r, response))
		return
	}

	a.UserCase.LogInfo("Передано на сервер USER")
	userResponse, err := a.Server.CheckUser(r.Context(), &userProto.UserLogin{
		Email:          newUser.Email,
		Password:       newUser.Password,
		SecondPassword: newUser.SecondPassword,
	})

	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
		return
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")

	cookie := a.UserCase.SetCookie(userResponse.GetToken())
	http.SetCookie(w, &cookie)
	models.Process(models.LoggerFunc("Success LogIn", a.UserCase.LogInfo), models.ResponseFunc(w, 200, a.UserCase.ProtoUser2User(userResponse.GetUser())), models.MetricFunc(200, r, nil))
}
