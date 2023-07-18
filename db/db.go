package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnPostgres() *sql.DB {
	connection := "user=postgres dbname=golang-marketplace password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	return db
}