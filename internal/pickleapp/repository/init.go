package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func getPostgres() *sql.DB {
	//dsn := "user=postgres-sniki dbname=tinder password=postgres-sniki host=127.0.0.1 port=5432 sslmode=disable"
	//db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	dsn := "dbname=postgres host=127.0.0.1 port=5432 sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	defer db.Close()
	if err != nil {
		fmt.Println("cant parse config", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("can`t ping db", err)
		return nil
	}

	db.SetMaxOpenConns(10)

	return db
}

func Init() *sqlx.DB {
	return sqlx.NewDb(getPostgres(), "psx")
}
