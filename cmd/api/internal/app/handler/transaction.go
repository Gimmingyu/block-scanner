package handler

import (
	"github.com/gin-gonic/gin"
	"scanner/cmd/api/internal/container"
)

type TransactionHandler struct {
	container container.Container
}

func (t *TransactionHandler) Index(router *gin.Engine) {
	//TODO implement me
	panic("implement me")
}
