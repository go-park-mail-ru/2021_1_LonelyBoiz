package delivery

import (
	"io/ioutil"
	"log"
	"net/http"
	model "server/internal/pkg/models"
)

func (a *UserHandler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		log.Println("error: get id from context")
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	image := string(bodyBytes)

	/*
		//парсит картинку
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
			response.Description["image"] = "Неверный формат фото"
			model.ResponseWithJson(w, 400, response)
			return
		}

		//вытаскивает картинку
		file, _, err := r.FormFile("image")
		if err != nil {
			response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
			response.Description["image"] = "Неверный формат фото"
			model.ResponseWithJson(w, 400, response)
			return
		}*/

	photoId, err := a.Db.AddPhoto(id, image)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}

	/*err = repository.SavePhoto(photoId, file)
	if err != nil {
		a.UserCase.Logger.Logger.Error(err)
		model.ResponseWithJson(w, 500, nil)
		return
	}*/

	model.ResponseWithJson(w, 200, photoId)
	return
}
