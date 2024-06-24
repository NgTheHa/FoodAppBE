package Utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserId(ctx context.Context) int {
	ginContext, _ := ctx.(*gin.Context)
	userId, _ := strconv.Atoi(ginContext.Request.Header.Get("UserId"))
	return userId
}
