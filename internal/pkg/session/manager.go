package session

import (
	"context"
	"math/rand"
	"net/http"
	model "server/internal/pkg/models"
	"server/internal/pkg/session/repository"
	"time"

	"github.com/sirupsen/logrus"
)

type SessionsManager struct {
	DB     repository.SessionRepository
	Logger *logrus.Entry
}

func (session *SessionsManager) KeyGen() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = model.CharSet[rand.Intn(len(model.CharSet))]
	}

	return string(b)
}

func (session *SessionsManager) SetSession(w http.ResponseWriter, id int) error {
	key := session.KeyGen()
	expiration := time.Now().Add(24 * time.Hour)

	logrus.Println(id)

	//cookie := http.Cookie{
	//	Name:     "token",
	//	Value:    key,
	//	Expires:  expiration,
	//	SameSite: http.SameSiteNoneMode,
	//	Secure:   true,
	//	//Domain:   "p1ckle.herokuapp.com", // TODO:: поменять перед пушем
	//	HttpOnly: true,
	//}

	cookie := http.Cookie{
		Name:     "token",
		Value:    key,
		Expires:  expiration,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		Domain:   "localhost:8000",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	logrus.Println(cookie)

	err := session.DB.AddCookie(id, key)
	if err != nil {
		return err
	}

	session.Logger.Info("Success Set Cookie")
	return nil
}

func (session *SessionsManager) DeleteSession(cookie *http.Cookie) error {
	key := cookie.Value
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	if err := session.DB.DeleteCookie(0, key); err != nil {
		session.Logger.Info("Delete Cookie : " + err.Error())
		return err
	}

	session.Logger.Info("Success Delete Cookie")
	return nil
}

func (session *SessionsManager) GetIdFromContext(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(model.CtxUserId).(int)
	if !ok {
		return 0, false
	}
	return id, true
}
