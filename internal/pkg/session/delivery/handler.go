package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	auth_server "server/internal/pkg/session"
)

type AuthHandler struct {
	Usecase auth_server.SessionManagerInterface
}

func (a *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		response := model.ErrorResponse{Err: "Не удалось взять куку"}
		model.ResponseFunc(w, 400, response)
		return
	}

	err = a.Usecase.DeleteSessionByToken(cookie.Value)
	if err != nil {
		model.ResponseFunc(w, 500, nil)
		return
	}

	a.Usecase.DeleteCookie(cookie)
	http.SetCookie(w, cookie)

	model.ResponseFunc(w, 200, nil)
}
