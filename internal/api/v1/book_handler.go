package v1

import (
	"context"
	"github.com/gin-gonic/gin"
)

func registerBook(router *gin.Engine, ctx context.Context) {
	v1 := router.Group("/v1")
	{
		v1.POST("/books", Create)
	}
}

func Create(c *gin.Context) {
}
