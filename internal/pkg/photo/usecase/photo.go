package usecase

import (
	"server/internal/pkg/models"
	"server/internal/pkg/photo/repository"
)

type PhotoUseCaseInterface interface {
	models.LoggerInterface
}

type PhotoUseCase struct {
	Db repository.PhotoRepositoryInterface
	models.LoggerInterface
}
