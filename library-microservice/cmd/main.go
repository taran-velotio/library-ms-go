package main

import (
	"fmt"
	"library-comp/db"
	"log"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("failed to connect to the db", err)
	}
	fmt.Println("Library-microservice")
}
