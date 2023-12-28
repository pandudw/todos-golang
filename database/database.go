package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	dsn := "root@tcp(localhost:3307)/todos"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db
}