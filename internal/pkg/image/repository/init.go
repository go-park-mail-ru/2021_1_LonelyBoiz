package repository

import (
	"bytes"
	"database/sql"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"server/internal/pkg/models"
)

var (
	ErrRepositoryImageNotFound = errors.New("image not found")
	ErrRepositoryConnection    = errors.New("something wrong with db connection")
	ErrQueryFailure            = errors.New("database query failed")
)

type DbRepositoryInterface interface {
	AddImage(userId int, uuid uuid.UUID) (models.Image, error)
	GetImages(userId int) ([]models.Image, error)
	RemoveImage(uuid uuid.UUID) error
}

type PostgresRepository struct {
	Db *sqlx.DB
}

type StorageRepositoryInterface interface {
	AddImage(uuid uuid.UUID, image []byte) error
	DeleteImage(uuid uuid.UUID) error
}

type AwsImageRepository struct {
	Bucket   string
	Svc      *s3.S3
	Uploader *s3manager.Uploader
}

func (r *PostgresRepository) AddImage(userId int, uuid uuid.UUID) (models.Image, error) {
	_, err := r.Db.Exec("INSERT INTO photos (photoUuid, userId) VALUES ($1, $2)", uuid, userId)
	if err == sql.ErrConnDone {
		return models.Image{}, ErrRepositoryConnection
	} else if err != nil {
		return models.Image{}, ErrQueryFailure
	}

	return models.Image{Uuid: uuid}, nil
}

func (r *PostgresRepository) GetImages(userId int) ([]models.Image, error) {
	images := make([]models.Image, 0)
	err := r.Db.Select(&images, "SELECT photoUuid FROM photos WHERE userId = $1", userId)
	if err == sql.ErrConnDone {
		return nil, ErrRepositoryConnection
	} else if err != nil {
		return nil, ErrQueryFailure
	}

	return images, nil
}

func (r *PostgresRepository) RemoveImage(uuid uuid.UUID) error {
	_, err := r.Db.Exec("DELETE FROM photos WHERE photoUuid = $1", uuid)
	if err == sql.ErrConnDone {
		return ErrRepositoryConnection
	} else if err != nil {
		return ErrQueryFailure
	}

	return nil
}

func (a *AwsImageRepository) AddImage(uuid uuid.UUID, image []byte) error {
	_, err := a.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.Bucket),
		Key:    aws.String(uuid.String()),
		Body:   aws.ReadSeekCloser(bytes.NewReader(image)),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *AwsImageRepository) DeleteImage(uuid uuid.UUID) error {
	_, err := a.Svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(a.Bucket),
		Key:    aws.String(uuid.String()),
	})
	if aerr, ok := err.(awserr.Error); ok {
		switch aerr.Code() {
		case s3.ErrCodeNoSuchKey:
			return ErrRepositoryImageNotFound
		default:
			return err
		}
	}

	err = a.Svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(a.Bucket),
		Key:    aws.String(uuid.String()),
	})
	if err != nil {
		return err
	}

	return nil
}
