package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось взять куку"}
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 400, response))
		return
	}

	err = a.Sessions.DeleteSession(cookie)
	http.SetCookie(w, cookie)
	if err != nil {
		model.Process(model.NewLogFunc(err.Error(), a.UserCase.LogError), model.NewResponseFunc(w, 500, nil))
		return
	}

	model.Process(model.NewLogFunc("Success LogOut", a.UserCase.LogInfo), model.NewResponseFunc(w, 200, nil))
}
