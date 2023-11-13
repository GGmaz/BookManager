package v1

import (
	"context"

	"github.com/gin-gonic/gin"
)

func RegisterVersion(router *gin.Engine, ctx context.Context) {
	registerBook(router, ctx)
}
