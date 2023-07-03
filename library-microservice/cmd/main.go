package main

import (
	"fmt"
	"library-comp/controller"
	"library-comp/db"
	"library-comp/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatal("failed to connect to the db", err)
	}

	bookRepo := repository.NewBookRepository()
	authorRepo := repository.NewAuthorRepository()

	bookController := controller.NewBookController(bookRepo)
	authorController := controller.NewAuthorController(*authorRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/books", func(r chi.Router) {
		r.Get("/all", bookController.GetListOfBooks)
		r.Get("/{id}", bookController.GetBook)
		r.Post("/", bookController.CreateBook)
		r.Put("/{id}", bookController.UpdateBook)
		r.Delete("/{id}", bookController.DeleteBook)
	})

	router.Route("/authors", func(r chi.Router) {
		r.Get("/all", authorController.GetListOfAuthors)
		r.Get("/{id}", authorController.GetAuthor)
		r.Post("/", authorController.CreateAuthor)
		r.Put("/{id}", authorController.UpdateAuthor)
		r.Delete("/{id}", authorController.DeleteAuthor)
	})

	fmt.Println("Server starting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
