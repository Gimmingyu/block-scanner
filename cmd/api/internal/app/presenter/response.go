package presenter

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"status": "success",
		"result": data,
	})
}

func Error(ctx *gin.Context, err error) {
	ctx.JSON(200, gin.H{
		"status": "error",
		"result": err.Error(),
	})
}
