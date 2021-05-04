package delivery

import (
	"google.golang.org/grpc/status"
	"net/http"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 400, response))
		return
	}

	userResponse, err := a.Server.CreateUser(r.Context(), a.UserCase.User2ProtoUser(newUser))
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			models.Process(models.LoggerFunc(st.Message(), a.UserCase.LogError), models.ResponseFunc(w, int(st.Code()), st.Message()))
			return
		}
	}

	//if err != nil || userResponse.GetCode() != 200 {
	//	models.ResponseFunc(w, int(userResponse.GetCode()), err.Error())
	//	return
	//}

	token, err := a.Sessions.Create(r.Context(), &session_proto2.SessionId{Id: userResponse.GetUser().GetId()})
	if err != nil {
		models.Process(models.LoggerFunc(err, a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}

	cookie := a.UserCase.SetCookie(token.GetToken())
	http.SetCookie(w, &cookie)

	models.Process(models.LoggerFunc("Success SignUp", a.UserCase.LogInfo), models.ResponseFunc(w, 200, userResponse.GetUser()))
}
