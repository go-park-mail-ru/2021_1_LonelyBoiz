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
	SetSession(w http.ResponseWriter, id int) error
	DeleteSession(cookie *http.Cookie) error
	GetIdFromContext(ctx context.Context) (int, bool)
}

type SessionsManager struct {
	DB     repository.SessionRepositoryInterface
	Logger model.LoggerInterface
}

func (session *SessionsManager) keyGen() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = model.CharSet[rand.Intn(len(model.CharSet))]
	}

	return string(b)
}

func (session *SessionsManager) SetSession(w http.ResponseWriter, id int) error {
	key := session.keyGen()
	expiration := time.Now().Add(24 * time.Hour)

	cookie := http.Cookie{
		Name:     "token",
		Value:    key,
		Expires:  expiration,
		SameSite: http.SameSiteLaxMode,
		Domain:   "localhost:3000",
		//Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)

	err := session.DB.AddCookie(id, key)
	if err != nil {
		return err
	}

	session.Logger.LogInfo("Success Set Cookie")
	return nil
}

func (session *SessionsManager) DeleteSession(cookie *http.Cookie) error {
	key := cookie.Value
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.Domain = "p1ckle.herokuapp.com"
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.Domain = "p1ckle.herokuapp.com"
	if err := session.DB.DeleteCookie(0, key); err != nil {
		session.Logger.Info("Delete Cookie : " + err.Error())
		return err
	}

	session.Logger.LogInfo("Success Delete Cookie")
	return nil
}

func (session *SessionsManager) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
