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
		model.ResponseWithJson(w, 403, response)
		return
	}

	query := r.URL.Query()
	limit, ok := query["count"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан count"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	feed, err := a.Db.GetFeed(id, limitInt)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}
	if len(feed) < limitInt {
		err = a.Db.CreateFeed(id)
		if err != nil {
			a.UserCase.Logger.Logger.Error(err)
			model.ResponseWithJson(w, 500, nil)
			return
		}
		feed, err = a.Db.GetFeed(id, limitInt)
		if err != nil {
			a.UserCase.Logger.Info(err)
			model.ResponseWithJson(w, 500, nil)
			return
		}
	}

	model.ResponseWithJson(w, 200, feed)
	return
}
