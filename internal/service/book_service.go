package service

import (
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

func (s *BookService) Create(book requests.CreateBookRequest) int64 {
	return s.bookRepo.Create(book.Title, book.Author, book.PublishedDate, book.Edition, book.Description, book.Genre)
}

func (s *BookService) GetAll(page, pageSize int) ([]model.Book, int64) {
	return s.bookRepo.GetAll(page, pageSize)
}

func (s *BookService) Update(id int64, book requests.UpdateBookRequest) string {
	if s.bookRepo.GetByID(id).ID == 0 {
		return "Book not found"
	}
	updatedBook := s.bookRepo.Update(id, book.Title, book.Author, book.PublishedDate, book.Edition, book.Description, book.Genre)
	if updatedBook.ID == 0 {
		return "There was an error updating the book"
	}
	return ""
}

func (s *BookService) Delete(id int64) string {
	if s.bookRepo.GetByID(id).ID == 0 {
		return "Book not found"
	}
	s.bookRepo.Delete(id)
	return ""

}
