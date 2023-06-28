package models

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// interface
type BookRepository interface {
	GetBook(id int) (Book, error)
	GetListOfBooks() ([]Book, error)
	CreateBook(bookDetails Book) (Book, error)
	UpdateBook(bookDetails Book) (Book, error)
	DeleteBook(id int) error
}
