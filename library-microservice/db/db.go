package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() error {
	dbinfo := "postgres://postgre:Test@123@localhost:5432/postgre"
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return err
	}

	log.Println("connected to the database", db)

	return nil
}
func SetupDB() *sql.DB {
	return db
}
