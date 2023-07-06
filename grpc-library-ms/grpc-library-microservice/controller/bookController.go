package controller

import (
	"encoding/json"
	"library-comp/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type BookController struct {
	bookRepository models.BookRepository
}

func NewBookController(repo models.BookRepository) *BookController {
	return &BookController{
		bookRepository: repo,
	}
}

func (t *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")
	if bookId == "" {
		log.Println("Invalid book Id : empty value")
		http.Error(w, "Invalid book Id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(bookId)

	if err != nil {
		log.Println("Invalid book Id", err)
		http.Error(w, "Invalid bookId", http.StatusBadRequest)
		return
	}

	book, err := t.bookRepository.GetBook(id)
	if err != nil {
		log.Println("Failed to get book", err)
		http.Error(w, "Failed to get book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (t *BookController) GetListOfBooks(w http.ResponseWriter, r *http.Request) {
	book, err := t.bookRepository.GetListOfBooks()
	if err != nil {
		log.Println("Failed to get books", err)
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(book)
}

func (t *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Println("Invalid book details", err)
		http.Error(w, "Invalid book details", http.StatusBadRequest)
		return
	}

	createdBook, err := t.bookRepository.CreateBook(book)
	if err != nil {
		log.Println("Failed to create book", err)
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdBook)
}

func (t *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Println("Invalid book data", err)
		http.Error(w, "Invalid book data", http.StatusBadRequest)
		return
	}

	bookId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(bookId)
	if err != nil {
		log.Println("Invalid book Id", err)
		http.Error(w, "Invalid book id", http.StatusBadRequest)
		return
	}

	book.Id = id
	updatedBook, err := t.bookRepository.UpdateBook(book)
	if err != nil {
		log.Println("Failed to update Book", err)
		http.Error(w, "Failed to update Book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedBook)
}

func (t *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(bookId)

	if err != nil {
		log.Println("Invalid book Id", err)
		http.Error(w, "Invalid book Id", http.StatusBadRequest)
		return
	}

	err = t.bookRepository.DeleteBook(id)
	if err != nil {
		log.Println("Failed to Delete book", err)
		http.Error(w, "Failed to Delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
