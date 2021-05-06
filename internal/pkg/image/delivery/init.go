package delivery

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"net/http"
	imageProto "server/internal/image_server/delivery/proto"
	"server/internal/pkg/image/usecase"
	"server/internal/pkg/models"
)

type ImageDeliveryInterface interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
	DeleteImage(w http.ResponseWriter, r *http.Request)
}

type ImageHandler struct {
	Usecase usecase.ImageUsecaseInterface
	Server  imageProto.ImageServiceClient
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
	if err != nil {
		responseBody := models.ErrorResponse{Err: "Не удалось прочитать файл"}
		models.ResponseWithJson(w, http.StatusBadRequest, responseBody)
		return
	}

	h.Usecase.LogInfo("Передано на сервер IMAGE")
	model, err := h.Server.UploadImage(r.Context(), &imageProto.ImageRequest{Image: body})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			models.Process(models.LoggerFunc(st.Message(), h.Usecase.LogError), models.ResponseFunc(w, int(st.Code()), models.ParseGrpcError(st.Message())))
			return
		}
	}

	models.Process(models.LoggerFunc("Success Upload Image", h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusOK, models.Image{Uuid: uuid.MustParse(model.GetImage())}))
}

func (h *ImageHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	_, err := h.Server.DeleteImage(h.Usecase.SetUUID(r.Context(), mux.Vars(r)), &imageProto.ImageNothing{})
	h.Usecase.LogInfo("Передано на сервер IMAGE")
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() != 200 {
			models.Process(models.LoggerFunc(st.Message(), h.Usecase.LogError), models.ResponseFunc(w, int(st.Code()), models.ParseGrpcError(st.Message())))
			return
		}
	}

	models.Process(models.LoggerFunc("Success Delete Image", h.Usecase.LogInfo), models.ResponseFunc(w, http.StatusNoContent, nil))
}
