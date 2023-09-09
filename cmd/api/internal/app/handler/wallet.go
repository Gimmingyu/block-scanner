package handler

import (
	"github.com/gin-gonic/gin"
	"scanner/cmd/api/internal/container"
)

type WalletHandler struct {
	container container.Container
}

func (w *WalletHandler) Index(router *gin.Engine) {
	//TODO implement me
	panic("implement me")
}
