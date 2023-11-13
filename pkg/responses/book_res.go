package responses

import "time"

type GetAllBooksResponse struct {
	Id            int64     `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
	Edition       string    `json:"edition"`
	Description   string    `json:"description"`
	Genre         string    `json:"genre"`
}
