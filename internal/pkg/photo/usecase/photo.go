package usecase

import (
	"golang.org/x/net/context"
	"server/internal/pkg/models"
	"server/internal/pkg/photo/repository"
)

type PhotoUseCaseInterface interface {
	models.LoggerInterface
	GetIdFromContext(ctx context.Context) (int, bool)
}

type PhotoUseCase struct {
	Db repository.PhotoRepositoryInterface
	models.LoggerInterface
}

func (u *PhotoUseCase) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(models.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
