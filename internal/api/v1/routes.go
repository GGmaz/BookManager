package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
)

func RegisterVersion(router *gin.Engine, ctx context.Context) {
	if len(os.Args) < 2 {
		registerBookREST(router, ctx)
		registerBookCollection(router, ctx)
	}
	registerBookCLI(ctx)
}
