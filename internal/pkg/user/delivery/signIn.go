package delivery

import (
	"net/http"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pkg/models"

	"github.com/google/uuid"
)

func (a *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	newUser, err := a.UserCase.ParseJsonToUser(r.Body)
	if err != nil {
		a.UserCase.LogError(err.Error())
		response := models.ErrorResponse{Err: "Не удалось прочитать тело запроса"}
		models.Process(models.LoggerFunc(response.Err, a.UserCase.LogInfo), models.ResponseFunc(w, 401, response))
		return
	}

	newUser, code, err := a.UserCase.SignIn(newUser)
	if code == 500 {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogError), models.ResponseFunc(w, code, nil))
		return
	}
	if code != 200 {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogInfo), models.ResponseFunc(w, code, err))
		return
	}
	println(newUser.Id)
	token, err := a.Sessions.Create(r.Context(), &session_proto2.SessionId{Id: int32(newUser.Id)})
	println(token)
	if err != nil {
		models.Process(models.LoggerFunc(err.Error(), a.UserCase.LogError), models.ResponseFunc(w, 500, nil))
		return
	}
	//if len(newUser.Photos) == 0 {
	//		newUser.Photos = make([]uuid.UUID, 0)
	//	}
	cookie := a.UserCase.SetCookie(token.GetToken())
	http.SetCookie(w, &cookie)
	models.Process(models.LoggerFunc("Success LogIn", a.UserCase.LogInfo), models.ResponseFunc(w, 200, newUser))
}
