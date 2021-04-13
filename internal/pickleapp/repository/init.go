package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func getPostgres() *sql.DB {
	dsn := "postgresql://localhost/database_name"
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

func Init() *sqlx.DB {
	return sqlx.NewDb(getPostgres(), "psx")
}
