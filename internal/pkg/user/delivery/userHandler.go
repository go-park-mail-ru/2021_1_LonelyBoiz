package delivery

import (
	"server/internal/pkg/session"
	"server/internal/pkg/user/usecase"
)

type UserHandler struct {
	UserCase usecase.UserUsecaseInterface
	Sessions *session.SessionsManager
}
