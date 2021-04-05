package delivery

import (
	"fmt"
	"log"
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) GetLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}
	fmt.Println("id from context =", id)

	userInfo, err := a.Db.GetUser(id)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		response.Description["id"] = "Пользователя с таким id нет"
		model.ResponseWithJson(w, 401, response)
		return
	}

	userInfo.PasswordHash = nil
	model.ResponseWithJson(w, 200, userInfo)
	return
}
