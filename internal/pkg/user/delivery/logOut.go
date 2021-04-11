package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		model.ResponseWithJson(w, 400, err)
		return
	}

	err = a.Sessions.DeleteSession(cookie)
	http.SetCookie(w, cookie)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		model.ResponseWithJson(w, 500, response)
		a.UserCase.Logger.Error(err.Error())
		return
	}

	model.ResponseWithJson(w, 200, nil)
	a.UserCase.Logger.Info("Success LogOut")
}
