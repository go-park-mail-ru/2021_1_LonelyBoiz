package delivery

import (
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
)

func (a *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	id, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}

		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	query := r.URL.Query()
	limit, ok := query["count"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан count"}

		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogInfo), model.ResponseFunc(w, 400, response))
		return
	}

	limitInt, err := strconv.Atoi(limit[0])

	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		model.Process(model.LoggerFunc(response.Err, a.UserCase.LogError), model.ResponseFunc(w, 400, response))
		return
	}

	feed, code, err := a.UserCase.CreateFeed(id, limitInt)
	if code != 200 {
		model.ResponseFunc(w, code, err)
		return
	}

	model.Process(model.LoggerFunc("Create Feed", a.UserCase.LogInfo), model.ResponseFunc(w, 200, feed))
	return
}
