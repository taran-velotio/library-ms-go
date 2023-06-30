package repository

import (
	"library-comp/db"
	"library-comp/models"
	"log"
)

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (repo *BookRepository) GetBook(id int) (models.Book, error) {

	db := db.SetupDB()
	var book models.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(&book.Id, &book.Name, &book.Author, &book.Price)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil

}

func (repo *BookRepository) GetListOfBooks() ([]models.Book, error) {
	db := db.SetupDB()
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []models.Book
	for rows.Next() {

		var book models.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Price)
		if err != nil {
			log.Println("Failed to enter book details", err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (repo *BookRepository) CreateBook(book models.Book) (models.Book, error) {
	db := db.SetupDB()

	err := db.QueryRow("INSERT INTO books (name,author,price) VALUES ($1,$2,$3) RETURNING id", book.Name, book.Author, book.Price).Scan(&book.Id)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (repo *BookRepository) UpdateBook(book models.Book) (models.Book, error) {
	db := db.SetupDB()

	_, err := db.Exec("UPDATE books SET name = $1, author = $2, price = $3 WHERE id = $4", book.Name, book.Author, book.Price, book.Id)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (repo *BookRepository) DeleteBook(id int) error {
	db := db.SetupDB()

	_, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
