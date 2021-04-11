package delivery

import (
	"server/internal/pkg/session"
	"server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
)

type UserHandler struct {
	Db       repository.UserRepository
	UserCase usecase.UserUsecase
	Sessions *session.SessionsManager
}
