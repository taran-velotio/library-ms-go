package main

import (
	"context"
	"library-comp/controller"
	"library-comp/db"
	"library-comp/proto/author/author"
	"library-comp/proto/book/book"
	"library-comp/repository"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	bookRepo := repository.NewBookRepository()
	authorRepo := repository.NewAuthorRepository()

	bookController := controller.NewBookController(bookRepo)
	authorController := controller.NewAuthorController(authorRepo)
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Book and author service to the server
	book.RegisterBookServiceServer(s, bookController)
	author.RegisterBookServiceServer(s, authorController)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = book.RegisterBookServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register book gateway:", err)
	}
	err = author.RegisterBookServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register author gateway", err)
	}
	//grpc Gateway server
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

}
