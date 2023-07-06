package repository

import (
	"context"
	"library-comp/db"
	"library-comp/proto/author/author"
	"log"
)

type AuthorRepository struct{}

func NewAuthorRepository() *AuthorRepository {
	return &AuthorRepository{}
}

func (repo *AuthorRepository) GetAuthor(ctx context.Context, req *author.GetAuthorRequest) (*author.GetAuthorResonse, error) {

	db := db.SetupDB()
	var authorInfo author.Author
	err := db.QueryRow("SELECT * FROM authors WHERE id = $1", req.GetId()).Scan(&authorInfo.Id, &authorInfo.Name)
	if err != nil {
		return &author.GetAuthorResonse{}, err
	}

	response := &author.GetAuthorResonse{
		Author: &authorInfo,
	}

	return response, nil

}

func (repo *AuthorRepository) GetListOfAuthors(ctx context.Context, req *author.GetListOfAuthorsRequest) (*author.GetListOfAuthorsResponse, error) {

	db := db.SetupDB()
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}

	var authors []*author.Author
	for rows.Next() {

		var authorInfo author.Author
		err := rows.Scan(&authorInfo.Id, &authorInfo.Name)
		if err != nil {
			log.Println("Failed to enter author details", err)
			return nil, err
		}
		authors = append(authors, &authorInfo)
	}

	response := &author.GetListOfAuthorsResponse{
		Authors: authors,
	}

	return response, nil
}

func (repo *AuthorRepository) CreateAuthor(ctx context.Context, req *author.CreateAuthorRequest) (*author.CreateAuthorResponse, error) {
	db := db.SetupDB()

	var authorID int32
	err := db.QueryRow("INSERT INTO authors (name) VALUES ($1) RETURNING id", req.GetName()).Scan(&authorID)
	if err != nil {
		return &author.CreateAuthorResponse{}, err
	}

	response := &author.CreateAuthorResponse{
		Author: &author.Author{
			Id:   authorID,
			Name: req.Name,
		},
	}

	return response, nil
}

func (repo *AuthorRepository) UpdateAuthor(ctx context.Context, req *author.UpdateAuthorRequest) (*author.UpdateAuthorResponse, error) {
	db := db.SetupDB()

	_, err := db.Exec("UPDATE authors SET name = $1, WHERE id = $2", req.GetName(), req.GetId())
	if err != nil {
		return &author.UpdateAuthorResponse{}, err
	}

	response := &author.UpdateAuthorResponse{
		Author: &author.Author{
			Id:   req.Id,
			Name: req.Name,
		},
	}

	return response, nil
}

func (repo *AuthorRepository) DeleteAuthor(ctx context.Context, req *author.DeleteAuthorRequest) (*author.DeleteAuthorResponse, error) {
	db := db.SetupDB()

	_, err := db.Exec("DELETE FROM authors WHERE id = $1", req.GetId())
	if err != nil {
		return nil, err
	}

	response := &author.DeleteAuthorResponse{}

	return response, nil
}
