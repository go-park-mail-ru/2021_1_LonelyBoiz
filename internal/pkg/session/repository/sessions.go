package repository

import (
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type SessionRepositoryInterface interface {
	AddCookie(id int, token string) error
	GetCookie(token string) (int, error)
	DeleteCookie(id int, token string) error
}

type SessionRepository struct {
	DB *sqlx.DB
}

func (repo *SessionRepository) AddCookie(id int, token string) error {
	_, err := repo.DB.Exec(
		`INSERT INTO sessions (userid, token, expiration) VALUES ($1, $2, $3)`,
		id, token, time.Now().Unix())

	return err
}

func (repo *SessionRepository) GetCookie(token string) (int, error) {
	var id []int
	err := repo.DB.Select(&id, `SELECT userId FROM sessions WHERE token = $1`, token)
	if err != nil {
		return -1, err
	}
	if len(id) == 0 {
		return -1, nil
	}

	return id[0], nil
}

func (repo *SessionRepository) DeleteCookie(id int, token string) error {
	_, err := repo.DB.Exec(`DELETE FROM sessions WHERE userid=$1 AND token=$2`, id, token)

	return err
}
