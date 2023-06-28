package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	dbinfo := "postgres://user:password@localhost:5432/{database_name}"
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Postgre DB")
	return db
}
