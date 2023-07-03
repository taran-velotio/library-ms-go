package controller

import (
	"encoding/json"
	"library-comp/models"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type AuthorController struct {
	authorRepository models.AuthorRepository
}

func NewAuthorController(repo models.AuthorRepository) *AuthorController {
	return &AuthorController{
		authorRepository: repo,
	}
}

func (t *AuthorController) GetAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(authorId)

	if err != nil {
		log.Println("Invalid Author Id", err)
		http.Error(w, "Invalid Author Id", http.StatusBadRequest)
		return
	}

	author, err := t.authorRepository.GetAuthor(id)
	if err != nil {
		log.Println("Failed to get author Id", err)
		http.Error(w, "Failed to get author Id", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(author)
}

func (t *AuthorController) GetListOfAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := t.authorRepository.GetListOfAuthors()
	if err != nil {
		log.Println("Failed to get Authors", err)
		http.Error(w, "Failed to get Authors", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(authors)
}

func (t *AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		log.Println("Invalid author data provided", err)
		http.Error(w, "Invalid author data provided", http.StatusInternalServerError)
		return
	}

	createdAuthor, err := t.authorRepository.CreateAuthor(author)
	if err != nil {
		log.Println("Failed to create author", err)
		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(createdAuthor)
}

func (t *AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		log.Println("Invalid author data", err)
		http.Error(w, "Invalid author data", http.StatusBadRequest)
		return
	}

	// get id from author to update data
	authorID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(authorID)
	if err != nil {
		log.Println("Invalid author Id", err)
		http.Error(w, "Invalid author Id", http.StatusBadRequest)
		return
	}

	author.Id = id
	updatedAuthor, err := t.authorRepository.UpdateAuthor(author)
	if err != nil {
		log.Println("Failed to update author", err)
		http.Error(w, "Failed to update author", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedAuthor)
}

func (t *AuthorController) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(authorId)
	if err != nil {
		log.Println("Invalid author Id", err)
		http.Error(w, "Invalid author Id", http.StatusBadRequest)
		return
	}

	err = t.authorRepository.DeleteAuthor(id)
	if err != nil {
		log.Println("Failed to delete author", err)
		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
