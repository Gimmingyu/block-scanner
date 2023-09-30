package presenter

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"status": "success",
		"result": data,
	})
}

func Error(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"status": "error",
		"result": err.Error(),
	})
}
