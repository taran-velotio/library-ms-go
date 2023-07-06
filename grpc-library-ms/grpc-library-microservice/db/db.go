package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "Test@123"
	DB_NAME     = "postgres"
	DB_HOST     = "db"
	DB_PORT     = "5432"
)

func Init() error {
	dbinfo := fmt.Sprintf("host =%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
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
