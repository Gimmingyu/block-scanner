package handler

import (
	"github.com/gin-gonic/gin"
	"scanner/cmd/api/internal/container"
)

type FaucetHandler struct {
	container container.Container
}

func (f *FaucetHandler) Index(router *gin.Engine) {

}
