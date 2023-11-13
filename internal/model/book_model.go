package model

import "time"

type Book struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
	Edition       string    `json:"edition"`
	Description   string    `json:"description"`
	Genre         string    `json:"genre"`
	CollectionId  int64     `json:"collectionId"`
}

type BookCollection struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books" gorm:"ForeignKey:CollectionId"`
}
