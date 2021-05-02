package delivery

import (
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	model "server/internal/pkg/models"
	"server/internal/pkg/photo/usecase"
	"strconv"
	"strings"
)

type PhotoDeliveryInterface interface {
	UploadPhoto(w http.ResponseWriter, r *http.Request)
	DownloadPhoto(w http.ResponseWriter, r *http.Request)
}

type PhotoHandler struct {
	Usecase usecase.PhotoUseCase
}

func (a *PhotoHandler) SetPhotoHandlers(subRouter *mux.Router) {
	// загрузить новую фотку на сервер
	subRouter.HandleFunc("/images", a.UploadPhoto).Methods("POST")
	// выгрузить фотку с сервера
	subRouter.HandleFunc("/images/{id:[0-9]+}", a.DownloadPhoto).Methods("GET")
	// удалить фотку
	//subRouter.HandleFunc("/images/{id:[0-9]+}", a.DeletePhoto).Methods("DELETE")
}

func (a *PhotoHandler) DownloadPhoto(w http.ResponseWriter, r *http.Request) {
	_, ok := a.Usecase.GetIdFromContext(r.Context())

	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.Process(model.LoggerFunc(response.Err, a.Usecase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	vars := mux.Vars(r)
	photoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Фото с таким id нет"
		model.ResponseWithJson(w, 400, response)

		model.Process(model.LoggerFunc(response.Err, a.Usecase.LogInfo), model.ResponseFunc(w, 403, response))
		return
	}

	res, err := a.Usecase.Db.GetPhoto(photoId)
	if err != nil {
		a.Usecase.LogError(err)
		model.ResponseWithJson(w, 500, err)
		return
	}

	res = res[strings.IndexByte(res, ',')+1:]
	res = res[0 : len(res)-1]
	bytes, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Decode error")
	}
	w.Header().Set("Content-Type", "image/*")
	_, err = w.Write(bytes)
	if err != nil {
		fmt.Println("Bytes error")
	}

	//model.ResponseWithJson(w, 200, res)

	/*img, err := repository.GetPhoto(photoId)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, err)
	}
	_, err = io.Copy(w, img)
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, err)
	}

	fileInfo, err := img.Stat()
	if err != nil {
		a.UserCase.Logger.Error(err)
		model.ResponseWithJson(w, 500, err)
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(int(fileInfo.Size())))
	w.WriteHeader(200)*/

}

func (a *PhotoHandler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
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

	photoId, err := a.Usecase.Db.AddPhoto(id, image)
	if err != nil {
		a.Usecase.LogError(err)
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
