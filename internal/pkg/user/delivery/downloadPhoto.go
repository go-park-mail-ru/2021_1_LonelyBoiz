package delivery

import (
	"encoding/base64"
	"fmt"
	"net/http"
	model "server/internal/pkg/models"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (a *UserHandler) DownloadPhoto(w http.ResponseWriter, r *http.Request) {
	_, ok := a.Sessions.GetIdFromContext(r.Context())
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		model.ResponseWithJson(w, 403, response)
		a.UserCase.LogInfo(response.Err)
		return
	}

	vars := mux.Vars(r)
	photoId, err := strconv.Atoi(vars["id"])
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Фото с таким id нет"
		model.ResponseWithJson(w, 400, response)
		return
	}

	res, err := a.UserCase.GetPhoto(photoId)
	if err != nil {
		a.UserCase.LogError(err)
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
