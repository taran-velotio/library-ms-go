package controller

import (
	"context"
	"library-comp/proto/author/author"
	"library-comp/repository"
	"log"
)

type AuthorController struct {
	author.UnimplementedAuthorServiceServer
	authorRepository *repository.AuthorRepository
}

func NewAuthorController(repo *repository.AuthorRepository) *AuthorController {
	return &AuthorController{
		authorRepository: repo,
	}
}

func (t *AuthorController) GetAuthor(ctx context.Context, req *author.GetAuthorRequest) (*author.GetAuthorResonse, error) {
	authorInfo, err := t.authorRepository.GetAuthor(ctx, req)
	if err != nil {
		log.Println("Failed to get author", err)
		return nil, err
	}

	response := authorInfo
	return response, nil
}

func (t *AuthorController) GetListOfAuthors(ctx context.Context, req *author.GetListOfAuthorsRequest) (*author.GetListOfAuthorsResponse, error) {
	authorsInfo, err := t.authorRepository.GetListOfAuthors(ctx, req)
	if err != nil {
		log.Println("failed to get list of authors", err)
		return nil, err
	}

	response := authorsInfo

	return response, nil
}

func (t *AuthorController) CreateAuthor(ctx context.Context, req *author.CreateAuthorRequest) (*author.CreateAuthorResponse, error) {
	createdAuthor, err := t.authorRepository.CreateAuthor(ctx, req)
	if err != nil {
		log.Println("failed to create author", err)
		return nil, err
	}

	response := createdAuthor

	return response, nil
}

func (t *AuthorController) UpdateAuthor(ctx context.Context, req *author.UpdateAuthorRequest) (*author.UpdateAuthorResponse, error) {
	updatedAuthor, err := t.authorRepository.UpdateAuthor(ctx, req)
	if err != nil {
		log.Println("Failed to update author", err)
		return nil, err
	}

	response := updatedAuthor

	return response, nil
}

func (t *AuthorController) DeleteAuthor(ctx context.Context, req *author.DeleteAuthorRequest) (*author.DeleteAuthorResponse, error) {

	deletedAuthor, err := t.authorRepository.DeleteAuthor(ctx, req)
	if err != nil {
		log.Println("Failed to delete an author", err)
		return nil, err
	}

	response := deletedAuthor

	return response, nil
}
