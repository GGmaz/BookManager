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

type BookCollectionHandler struct {
	bookCollectionService *service.BookCollectionService
	ctx                   context.Context
}

func registerBookCollection(router *gin.Engine, ctx context.Context) {
	bookService, _ := service.NewBookCollection()

	h := &BookCollectionHandler{
		bookCollectionService: bookService,
		ctx:                   ctx,
	}

	v1 := router.Group("/api/v1")
	{
		v1.POST("/collections", h.Create)
		v1.GET("/collections", h.GetAll)
		v1.DELETE("/collections/:collectionId", h.Delete)
		v1.GET("/collections/:collectionId/books", h.GetBooksForCollection)
		v1.DELETE("/collections/:collectionId/books/:bookId", h.RemoveBookFromCollection)
		v1.POST("/collections/:collectionId/books/:bookId", h.AddBookToCollection)
	}
}

func (handler *BookCollectionHandler) Create(c *gin.Context) {
	var bookCollection requests.CreateBookCollectionRequest
	if err := c.ShouldBindJSON(&bookCollection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	id := handler.bookCollectionService.Create(bookCollection)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was error while creating book collection."})
		return
	}

	log.WithFields(log.Fields{"service_name": "book-service", "method_name": "Create"}).Info("Book collection successfully created.")
	c.JSON(http.StatusOK, gin.H{"status": "Book collection created successfully!"})
}

func (handler *BookCollectionHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	bookCollections, total := handler.bookCollectionService.GetAll(page, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"bookCollections": bookCollections,
		"total":           total,
		"page":            page,
	})
}

func (handler *BookCollectionHandler) GetBooksForCollection(c *gin.Context) {
	var req requests.CollectionByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

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

	books, total, errMessage := handler.bookCollectionService.GetBooksForCollection(req.Id, page, pageSize, author, genre, startDate, endDate)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bookCollections": books,
		"total":           total,
		"page":            page,
	})
}

func (handler *BookCollectionHandler) AddBookToCollection(c *gin.Context) {
	var req requests.BookCollectionRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	errMessage := handler.bookCollectionService.AddBookToCollection(req.CollectionId, req.BookId)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book successfully added to collection!"})
}

func (handler *BookCollectionHandler) Delete(c *gin.Context) {
	var req requests.CollectionByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	errMessage := handler.bookCollectionService.Delete(req.Id)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book collection successfully deleted!"})
}

func (handler *BookCollectionHandler) RemoveBookFromCollection(c *gin.Context) {
	var req requests.BookCollectionRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	errMessage := handler.bookCollectionService.RemoveBookFromCollection(req.CollectionId, req.BookId)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book successfully removed from collection!"})
}
