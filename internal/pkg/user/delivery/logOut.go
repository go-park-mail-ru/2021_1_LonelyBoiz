package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось взять куку"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	err = a.Sessions.DeleteSession(cookie)
	http.SetCookie(w, cookie)
	if err != nil {
		model.Process(model.LoggerFunc(err.Error(), a.UserCase.LogError), model.ResponseFunc(w, 500, nil))
		return
	}

	model.Process(model.LoggerFunc("Success LogOut", a.UserCase.LogInfo), model.ResponseFunc(w, 200, nil))
}
