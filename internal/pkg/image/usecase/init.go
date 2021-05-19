package usecase

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"server/internal/pkg/image/repository"
	"server/internal/pkg/models"
	"strconv"

	"github.com/google/uuid"
)

const (
	MAXIMAGESIZE = 10
	MB           = 1000000
)

var (
	ErrUsecaseFatal                = errors.New("ошибка сервера")
	ErrUsecaseImageTooThick        = errors.New("размер изображения слишком большой")
	ErrUsecaseImageNotBelongToUser = errors.New("изображение не принадлежит пользователю")
	ErrUsecaseUserHaveNoImages     = errors.New("у пользователя нет фотографий")
	ErrUsecaseImageNotFound        = errors.New("изображнение не найдено")
	ErrUsecaseFailedToDelete       = errors.New("не удалось удалить изображение")
	ErrUsecaseFailedToUpload       = errors.New("не удалось загрузить изображение")
)

type ImageUsecaseInterface interface {
	AddImage(userId int, image []byte) (models.Image, error)
	DeleteImage(userId int, imageUuid uuid.UUID) error
	GetParamFromContext(ctx context.Context, param string) (int, bool)
	GetUUID(ctx context.Context) (uuid.UUID, bool)
	GetIdFromContext(ctx context.Context) (int, bool)
	GetUUIDFromContext(ctx context.Context) (uuid.UUID, bool)

	models.LoggerInterface
}

type ImageUsecase struct {
	Db           repository.DbRepositoryInterface
	ImageStorage repository.StorageRepositoryInterface
	models.LoggerInterface
}

func (u *ImageUsecase) GetUUID(ctx context.Context) (uuid.UUID, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.New(), false
	}

	sUUID := md.Get("urlUUID")
	if len(sUUID) == 0 {
		return uuid.New(), false
	}

	uid, err := uuid.Parse(sUUID[0])
	if err != nil {
		return uuid.New(), false
	}

	return uid, true
}

func (u *ImageUsecase) GetParamFromContext(ctx context.Context, param string) (int, bool) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return -1, false
	}

	dataByParam := data.Get(param)
	if len(dataByParam) == 0 {
		return -1, false
	}

	value, err := strconv.Atoi(dataByParam[0])
	if err != nil {
		return -1, false
	}

	return value, true
}

func (u *ImageUsecase) AddImage(userId int, image []byte) (models.Image, error) {
	if b2mb(len(image)) > MAXIMAGESIZE {
		return models.Image{}, ErrUsecaseImageTooThick
	}

	newUuid := uuid.New()

	err := u.ImageStorage.AddImage(newUuid, image)
	if err != nil {
		u.LogInfo("1")
		u.LogError(err)
		return models.Image{}, ErrUsecaseFailedToUpload
	}

	model, err := u.Db.AddImage(userId, newUuid)
	if err == repository.ErrRepositoryConnection {
		u.LogInfo("2")
		u.LogError(err)
		return models.Image{}, ErrUsecaseFatal
	} else if err != nil {
		u.LogInfo("3")
		u.LogError(err)
		return models.Image{}, ErrUsecaseFailedToUpload
	}

	return model, nil
}

func (u *ImageUsecase) DeleteImage(userId int, imageUuid uuid.UUID) error {
	err := userImagesContains(u.Db, userId, imageUuid)
	if err != nil {
		return err
	}

	err = u.Db.RemoveImage(imageUuid)
	if err == repository.ErrQueryFailure {
		return ErrUsecaseImageNotFound
	} else if err == repository.ErrRepositoryConnection {
		return ErrUsecaseFatal
	} else if err != nil {
		return ErrUsecaseFailedToDelete
	}

	err = u.ImageStorage.DeleteImage(imageUuid)
	if err == repository.ErrRepositoryImageNotFound {
		return ErrUsecaseImageNotFound
	} else if err != nil {
		return ErrUsecaseFailedToDelete
	}

	return nil
}

func userImagesContains(db repository.DbRepositoryInterface, userId int, uuid uuid.UUID) error {
	images, err := db.GetImages(userId)
	if err == repository.ErrRepositoryConnection {
		return ErrUsecaseFatal
	} else if err != nil {
		return ErrUsecaseUserHaveNoImages
	}

	for _, image := range images {
		if image.Uuid == uuid {
			return nil
		}
	}

	return ErrUsecaseImageNotBelongToUser
}

func b2mb(bytes int) int {
	return bytes / MB
}

func (u *ImageUsecase) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(models.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}

func (u *ImageUsecase) GetUUIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	uuidString, ok := ctx.Value(models.CtxImageId).(string)
	if !ok {
		return uuid.New(), false
	}

	id := uuid.MustParse(uuidString)

	return id, true
}
