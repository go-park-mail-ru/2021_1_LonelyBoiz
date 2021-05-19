package session

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
	DB     repository.SessionRepositoryInterface
	Logger model.LoggerInterface
}

func (session *SessionsManager) CheckSession(tokens []string) (int, bool) {
	session.Logger.LogInfo("Check Session")
	for _, token := range tokens {
		id, _ := session.DB.GetCookie(token)

		if id != -1 {
			session.Logger.LogInfo("Find Session for id" + string(rune(id)))
			return id, true
		}
	}

	session.Logger.LogInfo("Session Not Found")
	return -1, false
}

func (session *SessionsManager) keyGen() string {
	session.Logger.LogInfo("Generate token")
	b := make([]byte, 40)
	for i := range b {
		b[i] = model.CharSet[rand.Intn(len(model.CharSet))]
	}

	return string(b)
}

func (session *SessionsManager) SetSession(id int) (string, error) {
	key := session.keyGen()

	err := session.DB.AddCookie(id, key)
	if err != nil {
		session.Logger.LogError(err)
		return "", err
	}

	session.Logger.LogInfo("Add Cookie to DB")
	return key, nil
}

func (session *SessionsManager) DeleteSessionById(id int) error {
	session.Logger.LogInfo("Delete")
	if err := session.DB.DeleteCookie(id, ""); err != nil {
		session.Logger.LogInfo("Delete Cookie error: " + err.Error())
		return err
	}

	session.Logger.LogInfo("Delete Session for " + string(rune(id)))
	return nil
}

func (session *SessionsManager) DeleteSessionByToken(token string) error {
	if err := session.DB.DeleteCookie(-1, token); err != nil {
		session.Logger.LogInfo("Delete Cookie error: " + err.Error())
		return err
	}

	session.Logger.LogInfo("Delete Session by token")
	return nil
}

func (session *SessionsManager) DeleteCookie(cookie *http.Cookie) {
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	cookie.HttpOnly = true
	//cookie.Domain = "lepick.online:8000"
	cookie.Domain = "localhost:8000"
}

func (session *SessionsManager) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
