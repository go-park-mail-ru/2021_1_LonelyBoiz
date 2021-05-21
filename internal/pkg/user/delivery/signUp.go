package delivery

import (
	"google.golang.org/grpc/status"

	"net/http"
	"server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 400, response), models.MetricFunc(401, r, response))
		return
	}

	a.UserCase.LogInfo("Передано на сервер USER")
	userResponse, err := a.Server.CreateUser(r.Context(), a.UserCase.User2ProtoUser(newUser))
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()), models.MetricFunc(int(st.Code()), r, st.Err()))
			return
		}
	}
	a.UserCase.LogInfo("Получен результат из сервера USER")
	println(userResponse.GetToken())

	cookie := a.UserCase.SetCookie(userResponse.GetToken())
	http.SetCookie(w, &cookie)

	models.Process(models.LoggerFunc("Success SignUp", a.UserCase.LogInfo), models.ResponseFunc(w, 200, a.UserCase.ProtoUser2User(userResponse.GetUser())), models.MetricFunc(200, r, nil))
}
