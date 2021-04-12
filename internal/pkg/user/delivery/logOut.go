package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		response := model.ErrorResponse{Err: "Не удалось взять куку"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	err = a.Sessions.DeleteSession(cookie)
	http.SetCookie(w, cookie)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	model.ResponseWithJson(w, 200, nil)
	a.UserCase.Logger.Info("Success LogOut")
}
