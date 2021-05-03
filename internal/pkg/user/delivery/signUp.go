package delivery

import (
	"github.com/google/uuid"

	"net/http"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pkg/models"
)

func (a *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 400, response))
		return
	}

	userResponse, err := a.Server.CreateUser(r.Context(), a.UserCase.User2ProtoUser(newUser))

	if err != nil || userResponse.GetCode() != 200 {
		models.ResponseFunc(w, int(userResponse.GetCode()), err.Error())
		return
	}

	token, err := a.Sessions.Create(r.Context(), &session_proto2.SessionId{Id: userResponse.GetUser().GetId()})
	if err != nil {
		models.Process(models.LoggerFunc(err, a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}
	//	newUser.Password = ""
	//	newUser.SecondPassword = ""
	//	newUser.PasswordHash = nil
	//	if len(newUser.Photos) == 0 {
	//		newUser.Photos = make([]uuid.UUID, 0)
	//	}
	cookie := a.UserCase.SetCookie(token.GetToken())
	http.SetCookie(w, &cookie)

	models.Process(models.LoggerFunc("Success SignUp", a.UserCase.LogInfo), models.ResponseFunc(w, 200, userResponse.GetUser()))
}
