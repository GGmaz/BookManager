package repo

import (
	"github.com/GGmaz/BookManager/internal/model"
	"gorm.io/gorm"
)

type BookCollectionRepository struct {
	db       *gorm.DB
	bookRepo *BookRepository
}

func NewBookCollection() (*BookCollectionRepository, error) {
	bookRepo, err := NewBook()
	if err != nil {
		return nil, err
	}

	repo := &BookCollectionRepository{
		bookRepo: bookRepo,
		db:       bookRepo.db,
	}

	return repo, nil
}

func (r *BookCollectionRepository) Create(name string) int64 {
	collection := model.BookCollection{
		Name: name,
	}
	r.db.Create(&collection)
	return collection.ID
}

func (r *BookCollectionRepository) GetAll() []model.BookCollection {
	var collections []model.BookCollection
	r.db.Find(&collections)
	return collections
}

func (r *BookCollectionRepository) GetByID(id int64) model.BookCollection {
	var collection model.BookCollection
	r.db.Where("id = ?", id).First(&collection)
	return collection
}

func (r *BookCollectionRepository) GetBooksForCollection(id int64) []model.Book {
	var books []model.Book
	r.db.Model(&model.Book{}).Where("collection_id = ?", id).Find(&books)
	return books
}

func (r *BookCollectionRepository) AddBookToCollection(collectionId, bookId int64) {
	r.db.Model(&model.Book{}).Where("id = ?", bookId).Update("collection_id", collectionId)
}

func (r *BookCollectionRepository) Delete(id int64) {
	var collection model.BookCollection
	r.db.First(&collection, id)

	r.db.Delete(&collection)
}

func (r *BookCollectionRepository) RemoveBookFromCollection(bookId int64) {
	r.db.Model(&model.Book{}).Where("id = ?", bookId).Update("collection_id", nil)
}
