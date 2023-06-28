package repository

import (
	"library-comp/db"
	"library-comp/models"
	"log"
)

type AuthorRepository struct{}

func NewAuthorRepository() models.AuthorRepository {
	return &AuthorRepository{}
}

func (repo *AuthorRepository) GetAuthor(id int) (models.Author, error) {

	db := db.SetupDB()
	var author models.Author
	err := db.QueryRow("SELECT * FROM authors WHERE id = $1").Scan(&author.Id, &author.Name)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil

}

func (repo *AuthorRepository) GetListOfAuthors() ([]models.Author, error) {

	db := db.SetupDB()
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}

	var authors []models.Author
	for rows.Next() {

		var author models.Author
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			log.Println("Failed to enter author details", err)
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (repo *AuthorRepository) CreateAuthor(author models.Author) (models.Author, error) {
	db := db.SetupDB()

	err := db.QueryRow("INSERT INTO authors (name) VALUES ($1) RETURNING id", author.Name).Scan(&author.Id)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

func (repo *AuthorRepository) UpdateAuthor(author models.Author) (models.Author, error) {
	db := db.SetupDB()

	_, err := db.Exec("UPDATE authors SET name = $1, WHERE id = $2", author.Name, author.Id)
	if err != nil {
		return models.Author{}, err
	}

	return author, nil
}

func (repo *AuthorRepository) DeleteAuthor(id int) error {
	db := db.SetupDB()

	_, err := db.Exec("DELETE FROM authors WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
