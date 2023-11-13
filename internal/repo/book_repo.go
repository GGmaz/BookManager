package repo

import (
	"context"
	"github.com/GGmaz/BookManager/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBook() (*BookRepository, error) {
	repo := &BookRepository{}

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	repo.db = db
	err = repo.db.AutoMigrate(&model.BookCollection{}, &model.Book{})
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (repo *BookRepository) Close() error {
	db, err := repo.db.DB()
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func (repo *BookRepository) Create(ctx context.Context, title string, author string, date time.Time, edition string, description string, genre string) int64 {
	book := model.Book{
		Title:         title,
		Author:        author,
		PublishedDate: date,
		Edition:       edition,
		Description:   description,
		Genre:         genre,
	}

	repo.db.Create(&book)
	return book.ID
}

func (repo *BookRepository) GetAll(ctx context.Context) []model.Book {
	var books []model.Book
	repo.db.Find(&books)

	return books
}

func (repo *BookRepository) GetByID(ctx context.Context, id int64) model.Book {
	var book model.Book
	repo.db.First(&book, id)

	return book
}

func (repo *BookRepository) Update(ctx context.Context, id int64, title string, author string, date time.Time, edition string, description string, genre string) model.Book {
	var book model.Book
	repo.db.First(&book, id)

	book.Title = title
	book.Author = author
	book.PublishedDate = date
	book.Edition = edition
	book.Description = description
	book.Genre = genre

	repo.db.Save(&book)
	return book
}

func (repo *BookRepository) Delete(ctx context.Context, id int64) {
	var book model.Book
	repo.db.First(&book, id)

	repo.db.Delete(&book)
}
