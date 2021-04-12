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
	limit, ok := query["limit"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан limit"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	limitInt, err := strconv.Atoi(limit[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат limit"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	offset, ok := query["offset"]
	if !ok {
		response := model.ErrorResponse{Err: "Не указан offset"}
		model.ResponseWithJson(w, 400, response)
		return
	}
	offsetInt, err := strconv.Atoi(offset[0])
	if err != nil {
		response := model.ErrorResponse{Err: "Неверный формат offset"}
		model.ResponseWithJson(w, 400, response)
		return
	}

	feed, err := a.Db.GetFeed(id, limitInt, offsetInt)
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
		feed, err = a.Db.GetFeed(id, limitInt, offsetInt)
		if err != nil {
			a.UserCase.Logger.Info(err)
			model.ResponseWithJson(w, 500, nil)
			return
		}
	}
	if len(feed) == 0 {
		response := model.ErrorResponse{Err: "Лента закончилась"}
		model.ResponseWithJson(w, 503, response)
		return
	}

	model.ResponseWithJson(w, 200, feed)
	return
}
