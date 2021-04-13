package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) GetSmiles(w http.ResponseWriter, r *http.Request) {
	_, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.Logger.Info(response.Err)
		return
	}

	//s := '\u+1F600'

}
