package v1

import (
	"context"
	"github.com/GGmaz/BookManager/internal/service"
	"github.com/GGmaz/BookManager/pkg/requests"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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

	id := handler.bookCollectionService.Create(handler.ctx, bookCollection)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "There was error while creating book collection."})
		return
	}

	log.WithFields(log.Fields{"service_name": "book-service", "method_name": "Create"}).Info("Book collection successfully created.")
	c.JSON(http.StatusOK, gin.H{"status": "Book collection created successfully!"})
}

func (handler *BookCollectionHandler) GetAll(c *gin.Context) {
	bookCollections := handler.bookCollectionService.GetAll(handler.ctx)

	c.JSON(http.StatusOK, gin.H{"bookCollections": bookCollections})
}

func (handler *BookCollectionHandler) GetBooksForCollection(c *gin.Context) {
	var req requests.CollectionByIdRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	books, errMessage := handler.bookCollectionService.GetBooksForCollection(handler.ctx, req.Id)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func (handler *BookCollectionHandler) AddBookToCollection(c *gin.Context) {
	var req requests.BookCollectionRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error: ": err.Error()})
		return
	}

	errMessage := handler.bookCollectionService.AddBookToCollection(handler.ctx, req.CollectionId, req.BookId)

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

	errMessage := handler.bookCollectionService.Delete(handler.ctx, req.Id)

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

	errMessage := handler.bookCollectionService.RemoveBookFromCollection(handler.ctx, req.CollectionId, req.BookId)

	if errMessage != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMessage})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Book successfully removed from collection!"})
}
