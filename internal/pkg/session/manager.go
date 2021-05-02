package auth_server

import (
	"context"
	"math/rand"
	"net/http"
	model "server/internal/pkg/models"
	"server/internal/pkg/session/repository"
	"time"
)

type SessionManagerInterface interface {
	SetSession(id int) (string, error)
	DeleteSessionById(id int) error
	DeleteSessionByToken(key string) error
	DeleteCookie(cookie *http.Cookie)
	GetIdFromContext(ctx context.Context) (int, bool)
	CheckSession(tokens []string) (int, bool)
}

type SessionsManager struct {
	DB repository.SessionRepositoryInterface
}

func (session *SessionsManager) CheckSession(tokens []string) (int, bool) {
	for _, token := range tokens {
		id, _ := session.DB.GetCookie(token)

		if id != -1 {
			return id, true
		}
	}

	return -1, false
}

func (session *SessionsManager) keyGen() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = model.CharSet[rand.Intn(len(model.CharSet))]
	}

	return string(b)
}

func (session *SessionsManager) createCookie(key string) []string {
	expiration := time.Now().Add(24 * time.Hour)

	cookie := []string{
		"Name=token;",
		"Value=" + key + ";",
		"Expires=" + expiration.String() + ";",
		"SameSite=http.SameSiteLaxMode;",
		"Domain=localhost:3000",
	}

	return cookie
}

func (session *SessionsManager) SetSession(id int) (string, error) {
	key := session.keyGen()

	err := session.DB.AddCookie(id, key)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (session *SessionsManager) DeleteSessionById(id int) error {
	if err := session.DB.DeleteCookie(id, ""); err != nil {
		//session.Logger.Info("Delete Cookie : " + err.Error())
		return err
	}
	return nil
}

func (session *SessionsManager) DeleteSessionByToken(token string) error {
	if err := session.DB.DeleteCookie(-1, token); err != nil {
		//session.Logger.Info("Delete Cookie : " + err.Error())
		return err
	}
	return nil
}

func (session *SessionsManager) DeleteCookie(cookie *http.Cookie) {
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.Domain = "localhost:3000"
}

func (session *SessionsManager) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
