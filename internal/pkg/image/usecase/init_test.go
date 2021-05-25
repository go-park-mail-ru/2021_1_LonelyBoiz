package usecase

import (
	"errors"
	"server/internal/pkg/image/repository"
	repMocks "server/internal/pkg/image/repository/mocks"
	"server/internal/pkg/models"

	"testing"

	"server/internal/pkg/utils/metrics"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAddImage(t *testing.T) {
	metrics.New()
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	image := []byte{1, 2}
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageStorageRep.EXPECT().AddImage(gomock.Any(), image).Return(nil)
	imageRep.EXPECT().AddImage(userId, gomock.Any()).Return(res, nil)

	ret, err := useCaseTest.AddImage(userId, image)

	assert.Equal(t, res, ret)
	assert.Equal(t, err, nil)
}

func TestAddImage_Storage_AddImage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage:    imageStorageRep,
		Db:              imageRep,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	userId := 1
	image := []byte{1, 2}
	res := errors.New("some error")

	imageStorageRep.EXPECT().AddImage(gomock.Any(), image).Return(res)

	_, err := useCaseTest.AddImage(userId, image)

	assert.Equal(t, ErrUsecaseFailedToUpload, err)
}

func TestAddImage_DB_AddImage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage:    imageStorageRep,
		Db:              imageRep,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	userId := 1
	image := []byte{1, 2}
	res := errors.New("some error")

	imageStorageRep.EXPECT().AddImage(gomock.Any(), image).Return(nil)
	imageRep.EXPECT().AddImage(userId, gomock.Any()).Return(models.Image{}, res)

	_, err := useCaseTest.AddImage(userId, image)

	assert.Equal(t, ErrUsecaseFailedToUpload, err)
}

func TestAddImage_DB_AddImage_Error2(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage:    imageStorageRep,
		Db:              imageRep,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	userId := 1
	image := []byte{1, 2}

	imageStorageRep.EXPECT().AddImage(gomock.Any(), image).Return(nil)
	imageRep.EXPECT().AddImage(userId, gomock.Any()).Return(models.Image{}, repository.ErrRepositoryConnection)

	_, err := useCaseTest.AddImage(userId, image)

	assert.Equal(t, ErrUsecaseFatal, err)
}

func TestDeleteImage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(gomock.Any()).Return([]models.Image{res}, nil)
	imageRep.EXPECT().RemoveImage(gomock.Any()).Return(nil)
	imageStorageRep.EXPECT().DeleteImage(gomock.Any()).Return(nil)

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, nil)
}

func TestDeleteImage_GetImages_NoConError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, repository.ErrRepositoryConnection)

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseFatal)
}

func TestDeleteImage_GetImages_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, errors.New("Some error"))

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseUserHaveNoImages)
}

func TestDeleteImage_ImageNotBelong(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{}, nil)

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseImageNotBelongToUser)
}

func TestDeleteImage_RemoveImage_ImaeNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, nil)
	imageRep.EXPECT().RemoveImage(uuid).Return(repository.ErrQueryFailure)

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseImageNotFound)
}

func TestDeleteImage_RemoveImage_NoConnection(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, nil)
	imageRep.EXPECT().RemoveImage(uuid).Return(repository.ErrRepositoryConnection)

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseFatal)
}

func TestDeleteImage_RemoveImage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, nil)
	imageRep.EXPECT().RemoveImage(uuid).Return(errors.New("Some error"))

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseFailedToDelete)
}

func TestDeleteImage_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, nil)
	imageRep.EXPECT().RemoveImage(uuid).Return(nil)
	imageStorageRep.EXPECT().DeleteImage(uuid).Return(repository.ErrRepositoryImageNotFound)

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseImageNotFound)
}

func TestDeleteImage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	imageRep := repMocks.NewMockDbRepositoryInterface(mockCtrl)
	imageStorageRep := repMocks.NewMockStorageRepositoryInterface(mockCtrl)

	useCaseTest := ImageUsecase{
		ImageStorage: imageStorageRep,
		Db:           imageRep,
	}

	userId := 1
	uuid := uuid.New()
	res := models.Image{Uuid: uuid}

	imageRep.EXPECT().GetImages(userId).Return([]models.Image{res}, nil)
	imageRep.EXPECT().RemoveImage(uuid).Return(nil)
	imageStorageRep.EXPECT().DeleteImage(uuid).Return(errors.New("Some error"))

	err := useCaseTest.DeleteImage(userId, uuid)

	assert.Equal(t, err, ErrUsecaseFailedToDelete)
}
