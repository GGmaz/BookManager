package service

import (
	"context"
	"github.com/GGmaz/BookManager/internal/model"
	"github.com/GGmaz/BookManager/internal/repo"
	"github.com/GGmaz/BookManager/pkg/requests"
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

func (s *BookService) Create(ctx context.Context, book requests.CreateBookRequest) int64 {
	return s.bookRepo.Create(ctx, book.Title, book.Author, book.PublishedDate, book.Edition, book.Description, book.Genre)
}

func (s *BookService) GetAll(ctx context.Context) []model.Book {
	return s.bookRepo.GetAll(ctx)
}

func (s *BookService) Update(ctx context.Context, id int64, book requests.UpdateBookRequest) string {
	if s.bookRepo.GetByID(ctx, id).ID == 0 {
		return "Book not found"
	}
	updatedBook := s.bookRepo.Update(ctx, id, book.Title, book.Author, book.PublishedDate, book.Edition, book.Description, book.Genre)
	if updatedBook.ID == 0 {
		return "There was an error updating the book"
	}
	return ""
}

func (s *BookService) Delete(ctx context.Context, id int64) string {
	if s.bookRepo.GetByID(ctx, id).ID == 0 {
		return "Book not found"
	}
	s.bookRepo.Delete(ctx, id)
	return ""

}
