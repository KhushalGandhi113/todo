package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

type DB struct {
	*sql.DB
}

func Init() *sql.DB {
	var err error
	connStr := "user=postgres dbname=todo password=Gandhi@123 host=localhost sslmode=disable port=5432"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
