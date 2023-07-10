package tests

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"library-comp/controller"
// 	"library-comp/models"
// 	"library-comp/tests/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-chi/chi"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAuthor_Success(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	author := models.Author{Id: 1, Name: "Author1"}

// 	authorRepository.On("GetAuthor", 1).Return(author, nil)

// 	req := httptest.NewRequest("GET", "http://localhost:8080/authors/1", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()
// 	rctx := chi.NewRouteContext()

// 	rctx.URLParams.Add("id", "1")

// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Get("/authors/{id}", authorController.GetAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestGetListOfAuthors_Success(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	authors := []models.Author{
// 		{Id: 1, Name: "Author1"},
// 		{Id: 2, Name: "Author2"},
// 	}

// 	authorRepository.On("GetListOfAuthors").Return(authors, nil)

// 	req := httptest.NewRequest("GET", "http://localhost:8080/authors/all", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()

// 	router.Get("/authors/all", authorController.GetListOfAuthors)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestCreateAuthor_Success(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	newAuthor := models.Author{Id: 1, Name: "Author1"}
// 	createdAuthor := models.Author{Id: 1, Name: "Author1"}

// 	authorRepository.On("CreateAuthor", newAuthor).Return(createdAuthor, nil)

// 	reqBody, _ := json.Marshal(newAuthor)
// 	req := httptest.NewRequest("POST", "http://localhost:8080/authors/", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()

// 	router.Post("/authors/", authorController.CreateAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestUpdateAuthor_Success(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	updatedAuthor := models.Author{Id: 1, Name: "Author1Updated"}

// 	authorRepository.On("UpdateAuthor", updatedAuthor).Return(updatedAuthor, nil)

// 	reqBody, _ := json.Marshal(updatedAuthor)
// 	req := httptest.NewRequest("PUT", "http://localhost:8080/authors/1", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("id", "1")
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router := chi.NewRouter()
// 	router.Put("/authors/{id}", authorController.UpdateAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestDeleteAuthor_Success(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	authorId := 1

// 	authorRepository.On("DeleteAuthor", authorId).Return(nil)

// 	req := httptest.NewRequest("DELETE", "http://localhost:8080/authors/1", nil)
// 	w := httptest.NewRecorder()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("id", "1")
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router := chi.NewRouter()
// 	router.Delete("/authors/{id}", authorController.DeleteAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
// }

// func TestGetAuthor_Failure(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	author := models.Author{}

// 	authorRepository.On("GetAuthor", 1).Return(author, errors.New("error"))

// 	req := httptest.NewRequest("GET", "http://localhost:8080/authors/1", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()
// 	rctx := chi.NewRouteContext()

// 	rctx.URLParams.Add("id", "1")

// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Get("/authors/{id}", authorController.GetAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestGetListOfAuthors_Failure(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	authors := []models.Author{}

// 	authorRepository.On("GetListOfAuthors").Return(authors, errors.New("error"))

// 	req := httptest.NewRequest("GET", "http://localhost:8080/authors", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()

// 	router.Get("/authors", authorController.GetListOfAuthors)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestCreateAuthor_Failure(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	newAuthor := models.Author{Id: 1, Name: "Author1"}
// 	invalidAuthor := models.Author{}

// 	authorRepository.On("CreateAuthor", newAuthor).Return(invalidAuthor, errors.New("error"))

// 	reqBody, _ := json.Marshal(newAuthor)
// 	req := httptest.NewRequest("POST", "http://localhost:8080/authors/", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()

// 	router.Post("/authors/", authorController.CreateAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestUpdateAuthor_Failure(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	updatedAuthor := models.Author{Id: 1, Name: "Author1Updated"}

// 	authorRepository.On("UpdateAuthor", updatedAuthor).Return(models.Author{}, errors.New("error"))

// 	reqBody, _ := json.Marshal(updatedAuthor)
// 	req := httptest.NewRequest("PUT", "http://localhost:8080/authors/1", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("id", "1")
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router := chi.NewRouter()
// 	router.Put("/authors/{id}", authorController.UpdateAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestDeleteAuthor_Failure(t *testing.T) {
// 	authorRepository := new(mocks.AuthorRepository)
// 	authorController := controller.NewAuthorController(authorRepository)

// 	authorId := 1

// 	authorRepository.On("DeleteAuthor", authorId).Return(errors.New("error"))

// 	req := httptest.NewRequest("DELETE", "http://localhost:8080/authors/1", nil)
// 	w := httptest.NewRecorder()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("id", "1")
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router := chi.NewRouter()
// 	router.Delete("/authors/{id}", authorController.DeleteAuthor)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
