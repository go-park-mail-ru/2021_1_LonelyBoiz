package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
)

func getPostgres() *sql.DB {
	dsn := "user=postgres dbname=postgres host=postgres port=5432 sslmode=disable"
	db, err := sql.Open("pgx", dsn)
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
