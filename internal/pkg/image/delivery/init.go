package delivery

import (
	"io"
	"io/ioutil"
	"net/http"
	"server/internal/pkg/image/usecase"
	"server/internal/pkg/models"
	"server/internal/pkg/session"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ImageDeliveryInterface interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
	DeleteImage(w http.ResponseWriter, r *http.Request)
}

type ImageHandler struct {
	Usecase  usecase.ImageUsecaseInterface
	Sessions *session.SessionsManager
}

func (a *ImageHandler) SetHandlers(subRouter *mux.Router) {
	// добавить фотку
	subRouter.HandleFunc("/images", a.UploadImage).Methods("POST")
	// удалить фотку
	subRouter.HandleFunc("/images/{uuid}", a.DeleteImage).Methods("DELETE")
}

func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userId, ok := h.Sessions.GetIdFromContext(r.Context())
	if !ok {
		responseBody := models.ErrorResponse{Err: models.SessionErrorDenAccess}
		h.Usecase.LogInfo(responseBody)
		models.ResponseWithJson(w, http.StatusForbidden, responseBody)
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		responseBody := models.ErrorResponse{Err: "Не удалось прочитать файл"}
		models.ResponseWithJson(w, http.StatusBadRequest, responseBody)
		return
	}

	model, err := h.Usecase.AddImage(userId, body)
	status := setStatusCode(err, http.StatusBadRequest)
	if err != nil {
		responseBody := models.ErrorResponse{Err: err.Error()}
		h.Usecase.LogInfo(responseBody)
		models.ResponseWithJson(w, status, responseBody)
		return
	}

	models.ResponseWithJson(w, http.StatusOK, model)
}

func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userId, ok := h.Sessions.GetIdFromContext(r.Context())
	if !ok {
		responseBody := models.ErrorResponse{Err: models.SessionErrorDenAccess}
		models.ResponseWithJson(w, http.StatusForbidden, responseBody)
		return
	}

	vars := mux.Vars(r)
	sUuid := vars["uuid"]
	uid, err := uuid.Parse(sUuid)
	if err != nil {
		responseBody := models.ErrorResponse{Err: "invalid uuid"}
		models.ResponseWithJson(w, http.StatusBadRequest, responseBody)
		return
	}

	err = h.Usecase.DeleteImage(userId, uid)
	status := setStatusCode(err, http.StatusBadRequest)
	if err != nil {
		responseBody := models.ErrorResponse{Err: err.Error()}
		models.ResponseWithJson(w, status, responseBody)
		return
	}

	models.ResponseWithJson(w, http.StatusNoContent, nil)
}

func setStatusCode(err error, initStatus int) int {
	status := initStatus
	switch err {
	case usecase.ErrUsecaseImageTooThick:
		status = http.StatusBadRequest
	case usecase.ErrUsecaseFailedToUpload:
		status = http.StatusUnprocessableEntity
	case usecase.ErrUsecaseImageNotBelongToUser:
		status = http.StatusBadRequest
	case usecase.ErrUsecaseImageNotFound:
		status = http.StatusBadRequest
	case usecase.ErrUsecaseUserHaveNoImages:
		status = http.StatusUnprocessableEntity
	case usecase.ErrUsecaseFailedToDelete:
		status = http.StatusUnprocessableEntity
	case usecase.ErrUsecaseFatal:
		status = http.StatusInternalServerError
	}

	return status
}
