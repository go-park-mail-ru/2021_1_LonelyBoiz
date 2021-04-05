package session

import (
	"math/rand"
	"net/http"
	model "server/internal/pkg/models"
	"server/internal/pkg/session/repository"
	"time"
)

type SessionsManager struct {
	DB repository.SessionRepository
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

	cookie := http.Cookie{
		Name:     "token",
		Value:    key,
		Expires:  expiration,
		SameSite: http.SameSiteNoneMode,
		//Secure:   true,
		Domain:   "localhost:3000",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	err := session.DB.AddCookie(id, key)
	if err != nil {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: err.Error()}
		return response
	}

	return nil
}

func (session *SessionsManager) DeleteSession(cookie *http.Cookie) error {
	key := cookie.Value
	cookie.Expires = time.Now().AddDate(0, 0, -1)
	if err := session.DB.DeleteCookie(0, key); err != nil {
		return err
	}

	return nil
}
