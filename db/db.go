package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Dbconnect() *sql.DB {
	connection := "user=postgres dbname=Go_Store password=123789 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
