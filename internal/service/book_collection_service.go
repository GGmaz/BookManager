package service

import (
	"context"
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

func (s *BookCollectionService) Create(ctx context.Context, collection requests.CreateBookCollectionRequest) int64 {
	return s.bookCollectionRepo.Create(ctx, collection.Name)
}

func (s *BookCollectionService) GetAll(ctx context.Context) []model.BookCollection {
	return s.bookCollectionRepo.GetAll(ctx)
}

func (s *BookCollectionService) GetBooksForCollection(ctx context.Context, id int64) ([]model.Book, string) {
	if s.bookCollectionRepo.GetByID(ctx, id).ID == 0 {
		return nil, "Collection not found"
	}
	return s.bookCollectionRepo.GetBooksForCollection(ctx, id), ""
}

func (s *BookCollectionService) AddBookToCollection(ctx context.Context, collectionId, bookId int64) string {
	if s.bookCollectionRepo.GetByID(ctx, collectionId).ID == 0 {
		return "Collection not found"
	}
	if s.bookRepo.GetByID(ctx, bookId).ID == 0 {
		return "Book not found"
	}

	s.bookCollectionRepo.AddBookToCollection(ctx, collectionId, bookId)
	return ""
}

func (s *BookCollectionService) Delete(ctx context.Context, id int64) string {
	if s.bookCollectionRepo.GetByID(ctx, id).ID == 0 {
		return "Collection not found"
	}
	s.bookCollectionRepo.Delete(ctx, id)
	return ""
}

func (s *BookCollectionService) RemoveBookFromCollection(ctx context.Context, collectionId, bookId int64) string {
	if s.bookCollectionRepo.GetByID(ctx, collectionId).ID == 0 {
		return "Collection not found"
	}
	if s.bookRepo.GetByID(ctx, bookId).ID == 0 {
		return "Book not found"
	}
	s.bookCollectionRepo.RemoveBookFromCollection(ctx, collectionId, bookId)
	return ""
}
