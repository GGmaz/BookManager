package service

import (
	"github.com/GGmaz/BookManager/internal/repo"
)

type BookService struct {
	bookRepo *repo.BookRepository
}

func NewBook() (*BookService, error) {
	bookRepo, err := repo.NewBook()
	if err != nil {
		return nil, err
	}

	s := &BookService{
		bookRepo: bookRepo,
	}

	return s, nil
}
