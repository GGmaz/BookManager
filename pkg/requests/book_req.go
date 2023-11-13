package requests

import "time"

type CreateBookRequest struct {
	Id            int64     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
	Edition       string    `json:"edition"`
	Description   string    `json:"description"`
	Genre         string    `json:"genre"`
}

type UpdateBookRequest struct {
	Id            int64     `uri:"id" binding:"required"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
	Edition       string    `json:"edition"`
	Description   string    `json:"description"`
	Genre         string    `json:"genre"`
}

type DeleteBookRequest struct {
	Id int64 `uri:"id" binding:"required"`
}
