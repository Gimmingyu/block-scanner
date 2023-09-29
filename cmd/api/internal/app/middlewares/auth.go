package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Authenticated(issuer, signature string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
