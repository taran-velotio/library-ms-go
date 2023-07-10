package tests

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"library-comp/controller"
// 	"library-comp/models"
// 	"library-comp/proto/book/book"
// 	"library-comp/tests/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-chi/chi"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestGetBook_Success(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	bookInfo := &book.Book{Id: 1, Name: "Book1", Author: "Author1", Price: 99}

// 	bookRepository.On("GetBook", context.Background(), mock.AnythingOfType("*book.GetBookRequest")).Return(&book.GetBookResonse{
// 		Book: bookInfo,
// 	}, nil)

// 	req := httptest.NewRequest("GET", "http://localhost:8080/books/1", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()
// 	rctx := chi.NewRouteContext()

// 	rctx.URLParams.Add("id", "1")

// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Get("/books/{id}", bookController.GetBook)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestGetListOfBooks_Success(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	books := []models.Book{
// 		{Id: 1, Name: "Book1", Author: "Author1", Price: 99},
// 		{Id: 2, Name: "Book2", Author: "Author2", Price: 79},
// 	}

// 	bookRepository.On("GetListOfBooks").Return(books, nil)

// 	req := httptest.NewRequest("GET", "http://localhost:8080/books/all", nil)
// 	w := httptest.NewRecorder()

// 	bookController.GetListOfBooks(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestCreateBook_Success(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	newBook := models.Book{Id: 1, Name: "Book1", Author: "Author1", Price: 99}
// 	createdBook := models.Book{Id: 1, Name: "Book1", Author: "Author1", Price: 99}

// 	bookRepository.On("CreateBook", newBook).Return(createdBook, nil)

// 	reqBody, _ := json.Marshal(newBook)
// 	req := httptest.NewRequest("POST", "http://localhost:8080/books", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	bookController.CreateBook(w, req)

// 	resp := w.Result()

// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestUpdateBook_Success(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	updatedBook := models.Book{Id: 1, Name: "Book1Updated", Author: "Author1", Price: 99}

// 	bookRepository.On("UpdateBook", updatedBook).Return(updatedBook, nil)

// 	reqBody, _ := json.Marshal(updatedBook)
// 	req := httptest.NewRequest("PUT", "http://localhost:8080/books/1", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("id", "1")
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Put("/books/{id}", bookController.UpdateBook)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// }
// func TestDeleteBook_Success(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	bookId := 1

// 	bookRepository.On("DeleteBook", bookId).Return(nil)

// 	req := httptest.NewRequest("DELETE", "http://localhost:8080/books/1", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()
// 	rctx := chi.NewRouteContext()

// 	rctx.URLParams.Add("id", "1")

// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Delete("/books/{id}", bookController.DeleteBook)

// 	router.ServeHTTP(w, req)

// 	bookController.DeleteBook(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
// }
// func TestGetBook_Failure(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	bookRepository.On("GetBook", 1).Return(models.Book{}, errors.New("error"))

// 	req := httptest.NewRequest("GET", "http://localhost:8080/books/1", nil)
// 	w := httptest.NewRecorder()
// 	router := chi.NewRouter()
// 	rctx := chi.NewRouteContext()

// 	rctx.URLParams.Add("id", "1")

// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Get("/books/{id}", bookController.GetBook)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestGetListOfBooks_Failure(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	bookRepository.On("GetListOfBooks").Return(nil, errors.New("error"))

// 	req := httptest.NewRequest("GET", "http://localhost:8080/books/all", nil)
// 	w := httptest.NewRecorder()

// 	bookController.GetListOfBooks(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }

// func TestCreateBook_Failure(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	newBook := models.Book{Id: 1, Name: "Book1", Author: "Author1", Price: 99}

// 	bookRepository.On("CreateBook", newBook).Return(models.Book{}, errors.New("error"))

// 	reqBody, _ := json.Marshal(newBook)
// 	req := httptest.NewRequest("POST", "http://localhost:8080/books", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	bookController.CreateBook(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestUpdateBook_Failure(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	updatedBook := models.Book{Id: 1, Name: "Book1Updated", Author: "Author1", Price: 99}

// 	bookRepository.On("UpdateBook", updatedBook).Return(models.Book{}, errors.New("error"))

// 	reqBody, _ := json.Marshal(updatedBook)
// 	req := httptest.NewRequest("PUT", "http://localhost:8080/books/1", bytes.NewReader(reqBody))
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()

// 	rctx := chi.NewRouteContext()
// 	rctx.URLParams.Add("id", "1")
// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Put("/books/{id}", bookController.UpdateBook)

// 	router.ServeHTTP(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
// func TestDeleteBook_Failure(t *testing.T) {
// 	bookRepository := new(mocks.BookRepository)
// 	bookController := controller.NewBookController(bookRepository)

// 	bookId := 1

// 	bookRepository.On("DeleteBook", bookId).Return(errors.New("error"))

// 	req := httptest.NewRequest("DELETE", "http://localhost:8080/books/1", nil)
// 	w := httptest.NewRecorder()

// 	router := chi.NewRouter()
// 	rctx := chi.NewRouteContext()

// 	rctx.URLParams.Add("id", "1")

// 	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

// 	router.Delete("/books/{id}", bookController.DeleteBook)

// 	router.ServeHTTP(w, req)

// 	bookController.DeleteBook(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
// }
