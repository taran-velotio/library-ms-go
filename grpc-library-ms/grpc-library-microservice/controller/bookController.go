package controller

import (
	"context"
	"library-comp/proto/book"
	"library-comp/repository"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookController struct {
	book.UnimplementedBookServiceServer
	bookRepository *repository.BookRepository
}

func NewBookController(repo *repository.BookRepository) *BookController {
	return &BookController{
		bookRepository: repo,
	}
}

func (t *BookController) GetBook(ctx context.Context, req *book.GetBookRequest) (*book.GetBookResonse, error) {
	bookInfo, err := t.bookRepository.GetBook(ctx, req)
	if err != nil {
		st := status.New(codes.Internal, "Internal server error")
		log.Println("failed to get book", err)
		return nil, st.Err()
	}
	response := &book.GetBookResonse{
		Book: bookInfo.Book,
	}
	return response, nil
}

func (t *BookController) GetListOfBooks(ctx context.Context, req *book.GetListOfBooksRequest) (*book.GetListOfBooksResponse, error) {
	booksInfo, err := t.bookRepository.GetListOfBooks(ctx, req)
	if err != nil {
		st := status.New(codes.Internal, "Internal server error")
		log.Println("Failed to get list of books", err)
		return nil, st.Err()
	}

	//list of books to return
	var books []*book.Book
	for _, val := range booksInfo.Books {
		books = append(books, &book.Book{
			Id:     val.Id,
			Name:   val.Name,
			Author: val.Author,
			Price:  val.Price,
		})
	}
	response := &book.GetListOfBooksResponse{
		Books: books,
	}

	return response, nil
}

func (t *BookController) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {

	createdBook, err := t.bookRepository.CreateBook(ctx, req)
	if err != nil {
		st := status.New(codes.Internal, "Internal server error")
		log.Println("Failed to create book", err)
		return nil, st.Err()
	}

	response := createdBook

	return response, nil
}

func (t *BookController) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (*book.UpdateBookResponse, error) {

	updatedBook, err := t.bookRepository.UpdateBook(ctx, req)
	if err != nil {
		st := status.New(codes.Internal, "Internal server error")
		log.Println("Failed to update book", err)
		return nil, st.Err()
	}

	response := updatedBook

	return response, nil
}

func (t *BookController) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*book.DeleteBookResponse, error) {
	deletedBook, err := t.bookRepository.DeleteBook(ctx, req)
	if err != nil {
		st := status.New(codes.Internal, "Internal server error")
		log.Println("failed to delete the book", err)
		return nil, st.Err()
	}

	response := deletedBook
	return response, nil
}
