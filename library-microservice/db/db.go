package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	DB_USER     = "-"
	DB_PASSWORD = "-"
	DB_NAME     = "-"
)

func Init() error {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println("issue in connecting to database")
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("connected to the database", db)

	return nil
}
func SetupDB() *sql.DB {
	return db
}
