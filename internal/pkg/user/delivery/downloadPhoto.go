package delivery

import (
	"io"
	"log"
	"net/http"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/repository"
	"strconv"

	"github.com/gorilla/mux"
)

func (a *UserHandler) DownloadPhoto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		//вернуть ошибку залогировать
		log.Println("error: get id from context")
	}

	vars := mux.Vars(r)
	photoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Фото с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	ok, err = a.Db.CheckPhoto(photoId, userId)
	if err != nil {
		model.ResponseWithJson(w, 500, err)
	}
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["image"] = "Пытаешься получить не свое фото"
		model.ResponseWithJson(w, 403, response)
	}

	img, err := repository.SendPhoto(photoId)
	if err != nil {
		model.ResponseWithJson(w, 500, err)
	}
	_, err = io.Copy(w, img)
	if err != nil {
		model.ResponseWithJson(w, 500, err)
	}

	fileInfo, err := img.Stat()
	if err != nil {
		model.ResponseWithJson(w, 500, err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(int(fileInfo.Size())))
	w.WriteHeader(200)
}
