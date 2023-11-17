package service

import (
	"github.com/GGmaz/BookManager/internal/model"
	"github.com/GGmaz/BookManager/internal/repo"
	"github.com/GGmaz/BookManager/pkg/requests"
)

type BookCollectionService struct {
	bookCollectionRepo *repo.BookCollectionRepository
	bookRepo           *repo.BookRepository
}

func NewBookCollection() (*BookCollectionService, error) {
	bookCollectionRepo, err := repo.NewBookCollection()
	if err != nil {
		return nil, err
	}
	bookRepo, err := repo.NewBook()
	if err != nil {
		return nil, err
	}

	s := &BookCollectionService{
		bookCollectionRepo: bookCollectionRepo,
		bookRepo:           bookRepo,
	}

	return s, nil
}

func (s *BookCollectionService) Create(collection requests.CreateBookCollectionRequest) int64 {
	return s.bookCollectionRepo.Create(collection.Name)
}

func (s *BookCollectionService) GetAll() []model.BookCollection {
	return s.bookCollectionRepo.GetAll()
}

func (s *BookCollectionService) GetBooksForCollection(id int64) ([]model.Book, string) {
	if s.bookCollectionRepo.GetByID(id).ID == 0 {
		return nil, "Collection not found"
	}
	return s.bookCollectionRepo.GetBooksForCollection(id), ""
}

func (s *BookCollectionService) AddBookToCollection(collectionId, bookId int64) string {
	if s.bookCollectionRepo.GetByID(collectionId).ID == 0 {
		return "Collection not found"
	}
	if s.bookRepo.GetByID(bookId).ID == 0 {
		return "Book not found"
	}

	s.bookCollectionRepo.AddBookToCollection(collectionId, bookId)
	return ""
}

func (s *BookCollectionService) Delete(id int64) string {
	if s.bookCollectionRepo.GetByID(id).ID == 0 {
		return "Collection not found"
	}

	for _, book := range s.bookCollectionRepo.GetBooksForCollection(id) {
		s.bookCollectionRepo.RemoveBookFromCollection(book.ID)
	}

	s.bookCollectionRepo.Delete(id)
	return ""
}

func (s *BookCollectionService) RemoveBookFromCollection(collectionId, bookId int64) string {
	if s.bookCollectionRepo.GetByID(collectionId).ID == 0 {
		return "Collection not found"
	}
	if s.bookRepo.GetByID(bookId).ID == 0 {
		return "Book not found"
	}
	s.bookCollectionRepo.RemoveBookFromCollection(bookId)
	return ""
}
