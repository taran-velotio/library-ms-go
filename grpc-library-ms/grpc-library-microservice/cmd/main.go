package main

import (
	"fmt"
	"library-comp/controller"
	"library-comp/db"
	"library-comp/proto/book/book"
	"library-comp/repository"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	bookRepo := repository.NewBookRepository()
	// authorRepo := repository.NewAuthorRepository()

	bookController := controller.NewBookController(bookRepo)
	// authorController := controller.NewAuthorController(*authorRepo)

	// Create a gRPC server
	server := grpc.NewServer()

	//registering grpc controller with servers
	book.RegisterBookServiceServer(server, bookController)

	address := "localhost:8088"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	fmt.Println("gRPC server starting on", address)

	err = server.Serve(lis)
	if err != nil {
		log.Fatal("Failed to start gRPC server:", err)
	}
}
