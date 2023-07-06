package models

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// interface
type AuthorRepository interface {
	GetAuthor(id int) (Author, error)
	GetListOfAuthors() ([]Author, error)
	CreateAuthor(authorDetails Author) (Author, error)
	UpdateAuthor(authorDetails Author) (Author, error)
	DeleteAuthor(id int) error
}
