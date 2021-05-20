package delivery

import (
	"io"
	"io/ioutil"
	"net/http"
	"server/internal/pkg/image/usecase"
	"server/internal/pkg/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ImageDeliveryInterface interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
	DeleteImage(w http.ResponseWriter, r *http.Request)
}

type ImageHandler struct {
	Usecase usecase.ImageUsecaseInterface
}

func (h *ImageHandler) SetHandlers(subRouter *mux.Router) {
	// добавить фотку
	subRouter.HandleFunc("/images", h.UploadImage).Methods("POST")
	// удалить фотку
	subRouter.HandleFunc("/images/{uuid}", h.DeleteImage).Methods("DELETE")
}

func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil || len(body) == 0 {
		responseBody := models.ErrorResponse{Err: "Не удалось прочитать файл"}
		models.Process(models.LoggerFunc(responseBody.Err, h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusBadRequest, responseBody), models.MetricFunc(http.StatusBadRequest, r, err))
		return
	}

	userId, ok := h.Usecase.GetIdFromContext(r.Context())
	if !ok {
		responseBody := models.ErrorResponse{Err: models.SessionErrorDenAccess}
		models.Process(models.LoggerFunc(responseBody.Err, h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusForbidden, responseBody), models.MetricFunc(http.StatusForbidden, r, err))
		return
	}

	model, err := h.Usecase.AddImage(userId, body)
	status := setStatusCode(err, http.StatusBadRequest)
	if err != nil {
		responseBody := models.ErrorResponse{Err: err.Error()}
		models.Process(models.LoggerFunc(responseBody.Err, h.Usecase.LogInfo), models.ResponseFunc(w, status, responseBody), models.MetricFunc(status, r, err))
		return
	}

	models.Process(models.LoggerFunc("Success Upload Image", h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusOK, model), models.MetricFunc(status, r, err))

}

func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	uid, ok := vars["uuid"]
	if !ok {
		responseBody := models.ErrorResponse{Err: "invalid uuid"}
		models.Process(models.LoggerFunc(responseBody, h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusBadRequest, responseBody), models.MetricFunc(http.StatusBadRequest, r, responseBody))
		return
	}

	userId, ok := h.Usecase.GetIdFromContext(r.Context())
	if !ok {
		responseBody := models.ErrorResponse{Err: models.SessionErrorDenAccess}
		models.Process(models.LoggerFunc(responseBody, h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusForbidden, responseBody), models.MetricFunc(http.StatusForbidden, r, responseBody))
		return
	}

	err := h.Usecase.DeleteImage(userId, uuid.MustParse(uid))
	status := setStatusCode(err, http.StatusBadRequest)
	if err != nil {
		responseBody := models.ErrorResponse{Err: err.Error()}
		models.Process(models.LoggerFunc(responseBody, h.Usecase.LogInfo), models.ResponseFunc(w, status, responseBody), models.MetricFunc(status, r, err))
		return
	}

	models.Process(models.LoggerFunc("Success Delete Image", h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusNoContent, nil), models.MetricFunc(http.StatusNoContent, r, nil))
}

func setStatusCode(err error, initStatus int) int {
	s := initStatus
	switch err {
	case usecase.ErrUsecaseImageTooThick:
		s = http.StatusBadRequest
	case usecase.ErrUsecaseFailedToUpload:
		s = http.StatusUnprocessableEntity
	case usecase.ErrUsecaseImageNotBelongToUser:
		s = http.StatusBadRequest
	case usecase.ErrUsecaseImageNotFound:
		s = http.StatusBadRequest
	case usecase.ErrUsecaseUserHaveNoImages:
		s = http.StatusUnprocessableEntity
	case usecase.ErrUsecaseFailedToDelete:
		s = http.StatusUnprocessableEntity
	case usecase.ErrUsecaseFatal:
		s = http.StatusInternalServerError
	}
	return s
}
