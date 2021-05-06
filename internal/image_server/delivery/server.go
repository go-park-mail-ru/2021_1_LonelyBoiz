package delivery

import (
	"context"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	imageProto "server/internal/image_server/delivery/proto"
	imageUsecase "server/internal/pkg/image/usecase"
	"server/internal/pkg/models"
)

type ImageServer struct {
	imageProto.UnimplementedImageServiceServer
	Usecase imageUsecase.ImageUsecaseInterface
}

func setStatusCode(err error, initStatus int) int {
	s := initStatus
	switch err {
	case imageUsecase.ErrUsecaseImageTooThick:
		s = http.StatusBadRequest
	case imageUsecase.ErrUsecaseFailedToUpload:
		s = http.StatusUnprocessableEntity
	case imageUsecase.ErrUsecaseImageNotBelongToUser:
		s = http.StatusBadRequest
	case imageUsecase.ErrUsecaseImageNotFound:
		s = http.StatusBadRequest
	case imageUsecase.ErrUsecaseUserHaveNoImages:
		s = http.StatusUnprocessableEntity
	case imageUsecase.ErrUsecaseFailedToDelete:
		s = http.StatusUnprocessableEntity
	case imageUsecase.ErrUsecaseFatal:
		s = http.StatusInternalServerError
	}

	return s
}

func (i ImageServer) UploadImage(ctx context.Context, request *imageProto.ImageRequest) (*imageProto.ImageResponse, error) {
	userId, ok := i.Usecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		responseBody := models.ErrorResponse{Err: models.SessionErrorDenAccess}
		i.Usecase.LogInfo(responseBody)
		return &imageProto.ImageResponse{}, status.Error(http.StatusForbidden, responseBody.Error())
	}

	model, err := i.Usecase.AddImage(userId, request.Image)
	statusCode := setStatusCode(err, http.StatusBadRequest)
	if err != nil {
		responseBody := models.ErrorResponse{Err: err.Error()}
		i.Usecase.LogInfo(responseBody)
		return &imageProto.ImageResponse{}, status.Error(codes.Code(statusCode), responseBody.Error())
	}

	return &imageProto.ImageResponse{Image: model.Uuid.String()}, nil
}

func (i ImageServer) DeleteImage(ctx context.Context, nothing *imageProto.ImageNothing) (*imageProto.ImageNothing, error) {
	userId, ok := i.Usecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		responseBody := models.ErrorResponse{Err: models.SessionErrorDenAccess}
		i.Usecase.LogInfo(responseBody)
		return &imageProto.ImageNothing{}, status.Error(http.StatusForbidden, responseBody.Error())
	}

	uid, ok := i.Usecase.GetUUID(ctx)
	if !ok {
		responseBody := models.ErrorResponse{Err: "invalid uuid"}
		i.Usecase.LogInfo(responseBody)
		return &imageProto.ImageNothing{}, status.Error(http.StatusBadRequest, responseBody.Error())
	}

	err := i.Usecase.DeleteImage(userId, uid)
	statusCode := setStatusCode(err, http.StatusBadRequest)
	if err != nil {
		responseBody := models.ErrorResponse{Err: err.Error()}
		i.Usecase.LogInfo(responseBody)
		return &imageProto.ImageNothing{}, status.Error(codes.Code(statusCode), responseBody.Error())
	}

	return &imageProto.ImageNothing{}, nil
}
