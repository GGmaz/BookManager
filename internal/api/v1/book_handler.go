package v1

import (
	"context"
	"github.com/GGmaz/BookManager/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	v1 := router.Group("/v1")
	{
		v1.POST("/books", h.Create)
	}
}

func (handler *BookHandler) Create(c *gin.Context) {
}
