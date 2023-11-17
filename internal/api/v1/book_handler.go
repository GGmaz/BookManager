package v1

import (
	"context"
	"github.com/GGmaz/BookManager/internal/service"
	"github.com/GGmaz/BookManager/pkg/requests"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type BookHandler struct {
	bookService *service.BookService
	ctx         context.Context
}

func registerBook(router *gin.Engine, ctx context.Context) {
	bookService, err := service.NewBook()
	if err != nil {
		log.WithFields(log.Fields{"service_name": "book-service", "method_name": "registerBook"}).Error("Error creating book service.")
		panic("Error creating book service.")
	}
	log.WithFields(log.Fields{"service_name": "book-service", "method_name": "registerBook"}).Info("Successfully created book handler.")
	h := &BookHandler{
		bookService: bookService,
		ctx:         ctx,
	}

	v1 := router.Group("/api/v1")
	{
		v1.POST("/books", h.Create)
		v1.GET("/books", h.GetAll)
		v1.PUT("/books/:id", h.Update)
		v1.DELETE("/books/:id", h.Delete)
	}
}

func (handler *BookHandler) Create(c *gin.Context) {
	var book requests.CreateBookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	id := handler.bookService.Create(book)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was error while creating book."})
		return
	}

	log.WithFields(log.Fields{"service_name": "book-service", "method_name": "Create"}).Info("Book successfully created.")
	c.JSON(http.StatusOK, gin.H{"status": "Book created successfully!"})
}

func (handler *BookHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	author := c.Query("author")
	genre := c.Query("genre")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	books, total := handler.bookService.GetAll(page, pageSize, author, genre, startDate, endDate)

	c.JSON(http.StatusOK, gin.H{
		"books": books,
		"total": total,
		"page":  page,
	})
}

func (handler *BookHandler) Update(c *gin.Context) {
	var book requests.UpdateBookRequest
	if err := c.ShouldBindUri(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	errMessage := handler.bookService.Update(book.Id, book)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book updated successfully!"})
}

func (handler *BookHandler) Delete(c *gin.Context) {
	var req requests.DeleteBookRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	errMessage := handler.bookService.Delete(req.Id)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book deleted successfully!"})
}
