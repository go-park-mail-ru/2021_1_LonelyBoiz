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

		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 403, response))
		return
	}

	query := r.URL.Query()
	limit, ok := query["count"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан count"}

		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogInfo), model.NewResponseFunc(w, 400, response))
		return
	}

	limitInt, err := strconv.Atoi(limit[0])

	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		model.Process(model.NewLogFunc(response.Err, a.UserCase.LogError), model.NewResponseFunc(w, 400, response))
		return
	}

	feed, code, err := a.UserCase.CreateFeed(id, limitInt)
	switch code {
	case 200:
		model.Process(model.NewLogFunc("Create Feed", a.UserCase.LogInfo), model.NewResponseFunc(w, code, feed))
	case 500:
		model.Process(model.NewLogFunc(err, a.UserCase.LogError), model.NewResponseFunc(w, code, err))
	default:
		model.Process(model.NewLogFunc(err, a.UserCase.LogInfo), model.NewResponseFunc(w, code, err))
	}

	return
}
