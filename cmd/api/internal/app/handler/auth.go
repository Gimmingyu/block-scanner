package handler

import (
	"github.com/gin-gonic/gin"
	"scanner/cmd/api/internal/app/dto"
	"scanner/cmd/api/internal/app/presenter"
	"scanner/cmd/api/internal/container"
)

type AuthHandler struct {
	container *container.Container
}

func NewAuthHandler(container *container.Container) *AuthHandler {
	return &AuthHandler{container: container}
}

func (a *AuthHandler) Index(router *gin.Engine) {
	group := router.Group("auth")

	group.POST("/login", a.Login)
	group.POST("/register", a.Register)
	group.POST("/logout", a.Logout)
	group.POST("/refresh", a.Refresh)
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	req, err := presenter.BindJSON[dto.LoginRequest](ctx)
	if err != nil {
		presenter.Error(ctx, err)
		return
	}

	presenter.Success(ctx, req)
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	req, err := presenter.BindJSON[dto.RegisterRequest](ctx)
	if err != nil {
		presenter.Error(ctx, err)
		return
	}

	presenter.Success(ctx, req)
}

func (a *AuthHandler) Logout(ctx *gin.Context) {
	req, err := presenter.BindJSON[dto.LogoutRequest](ctx)
	if err != nil {
		presenter.Error(ctx, err)
		return
	}

	presenter.Success(ctx, req)
}

func (a *AuthHandler) Refresh(ctx *gin.Context) {
	req, err := presenter.BindJSON[dto.RefreshRequest](ctx)
	if err != nil {
		presenter.Error(ctx, err)
		return
	}

	presenter.Success(ctx, req)
}