package repository

type PhotoRepositoryInterface interface {
	AddPhoto(userId int, image string) (int, error)
	GetPhoto(photoId int) (string, error)
	GetPhotos(userId int) ([]int, error)
	CheckPhoto(photoId int, userId int) (bool, error)
}
