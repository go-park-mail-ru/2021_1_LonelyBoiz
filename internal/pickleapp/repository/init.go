package repository

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

func getPostgres() *sql.DB {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("cant parse config" + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("can`t ping db" + err.Error())
	}

	db.SetMaxOpenConns(10)

	return db
}

func Init() *sqlx.DB {
	return sqlx.NewDb(getPostgres(), "psx")
}
