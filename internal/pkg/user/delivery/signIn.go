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
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 401, response))
		return
	}

	userResponse, err := a.Server.CheckUser(r.Context(), &userProto.UserLogin{
		Email:          newUser.Email,
		Password:       newUser.Password,
		SecondPassword: newUser.SecondPassword,
	})

	if err != nil {
		st, _ := status.FromError(err)
		models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()))
		return
	}

	user, ok := a.UserCase.ProtoUser2User(userResponse.GetUser())
	if !ok {
		models.Process(models.LoggerFunc("Proto Error", a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
	}

	cookie := a.UserCase.SetCookie(userResponse.GetToken())
	http.SetCookie(w, &cookie)
	models.Process(models.LoggerFunc("Success LogIn", a.UserCase.LogInfo), models.ResponseFunc(w, 200, user))
}
