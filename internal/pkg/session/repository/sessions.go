package repository

import (
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/jackc/pgx/stdlib"
)

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
	err := repo.DB.Select(&id, `SELECT userId FROM sessions WHERE token=$1`, token)
	if err != nil {
		return -1, err
	}

	return id[0], nil
}

func (repo *SessionRepository) DeleteCookie(id int, token string) error {
	_, err := repo.DB.Exec(`DELETE FROM sessions WHERE userid=$1 AND token=$2`, id, token)

	return err
}
