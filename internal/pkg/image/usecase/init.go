package usecase

import (
	"errors"
	"server/internal/pkg/image/repository"
	"server/internal/pkg/models"

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
}

type ImageUsecase struct {
	Db           repository.DbRepositoryInterface
	ImageStorage repository.StorageRepositoryInterface
}

func (u *ImageUsecase) AddImage(userId int, image []byte) (models.Image, error) {
	if b2mb(len(image)) > MAXIMAGESIZE {
		return models.Image{}, ErrUsecaseImageTooThick
	}

	newUuid := uuid.New()

	err := u.ImageStorage.AddImage(newUuid, image)
	if err != nil {
		return models.Image{}, ErrUsecaseFailedToUpload
	}

	model, err := u.Db.AddImage(userId, newUuid)
	if err == repository.ErrRepositoryConnection {
		return models.Image{}, ErrUsecaseFatal
	} else if err != nil {
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
