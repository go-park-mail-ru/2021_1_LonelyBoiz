package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type RepoSqlx struct {
	DB *sqlx.DB
}

func getPostgres() *sql.DB {
	//jdbc:postgresql://localhost:5432/postgres
	dsn := "dbname=postgres host=127.0.0.1 port=5432 sslmode=disable"
	db, err := sql.Open("pgx", dsn)
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

func Init() RepoSqlx {
	var repo RepoSqlx
	repo.DB = sqlx.NewDb(getPostgres(), "psx")

	return repo
}
