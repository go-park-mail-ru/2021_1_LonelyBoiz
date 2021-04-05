package delivery

import (
	"github.com/sirupsen/logrus"
	"server/internal/pkg/session"
	"server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
)

type UserHandler struct {
	Db       repository.UserRepository
	UserCase usecase.UserUsecase
	Logger   *logrus.Entry
	Sessions *session.SessionsManager
}
