package repository

import (
	"context"
	"library-comp/db"
	"library-comp/proto/book"
	"log"
)

type BookRepository struct {
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (repo *BookRepository) GetBook(ctx context.Context, req *book.GetBookRequest) (*book.GetBookResonse, error) {
	db, err1 := db.Init()
	if err1 != nil {
		log.Fatalf("Failed to connect to database: %v", err1)
	}
	var bookInfo book.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = $1", req.GetId()).Scan(&bookInfo.Id, &bookInfo.Name, &bookInfo.Author, &bookInfo.Price)
	if err != nil {
		return &book.GetBookResonse{}, nil
	}

	response := &book.GetBookResonse{
		Book: &bookInfo,
	}
	return response, nil
}

func (repo *BookRepository) GetListOfBooks(ctx context.Context, req *book.GetListOfBooksRequest) (*book.GetListOfBooksResponse, error) {
	db, err := db.Init()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	rows, err := db.Query("SELECT id, name,author,price FROM books")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []*book.Book
	for rows.Next() {

		var bookInfo book.Book
		err := rows.Scan(&bookInfo.Id, &bookInfo.Name, &bookInfo.Author, &bookInfo.Price)
		if err != nil {
			log.Println("Failed to enter book details", err)
			return nil, err
		}
		books = append(books, &bookInfo)
	}

	response := &book.GetListOfBooksResponse{
		Books: books,
	}

	return response, nil
}

func (repo *BookRepository) CreateBook(ctx context.Context, req *book.CreateBookRequest) (*book.CreateBookResponse, error) {
	db, err1 := db.Init()
	if err1 != nil {
		log.Fatalf("Failed to connect to database: %v", err1)
	}
	var bookID int32
	err := db.QueryRow("INSERT INTO books (name,author,price) VALUES ($1,$2,$3) RETURNING id", req.GetName(), req.GetAuthor(), req.GetPrice()).Scan(&bookID)
	if err != nil {
		return &book.CreateBookResponse{}, err
	}

	response := &book.CreateBookResponse{
		Book: &book.Book{
			Id:     bookID,
			Name:   req.Name,
			Author: req.Author,
			Price:  req.Price,
		},
	}

	return response, nil
}

func (repo *BookRepository) UpdateBook(ctx context.Context, req *book.UpdateBookRequest) (*book.UpdateBookResponse, error) {
	db, err1 := db.Init()
	if err1 != nil {
		log.Fatalf("Failed to connect to database: %v", err1)
	}
	_, err := db.Exec("UPDATE books SET name = $1, author = $2, price = $3 WHERE id = $4", req.GetName(), req.GetAuthor(), req.GetPrice(), req.GetId())
	if err != nil {
		return &book.UpdateBookResponse{}, err
	}

	response := &book.UpdateBookResponse{
		Book: &book.Book{
			Id:     req.Id,
			Name:   req.Name,
			Author: req.Author,
			Price:  req.Price,
		},
	}

	return response, nil
}

func (repo *BookRepository) DeleteBook(ctx context.Context, req *book.DeleteBookRequest) (*book.DeleteBookResponse, error) {
	db, err1 := db.Init()
	if err1 != nil {
		log.Fatalf("Failed to connect to database: %v", err1)
	}
	_, err := db.Exec("DELETE FROM books WHERE id = $1", req.GetId())
	if err != nil {
		return nil, err
	}

	response := &book.DeleteBookResponse{}

	return response, nil
}
