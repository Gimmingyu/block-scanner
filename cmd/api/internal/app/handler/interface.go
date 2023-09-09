package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	Index(router *gin.Engine)
}
