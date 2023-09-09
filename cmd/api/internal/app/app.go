package app

import (
	"github.com/gin-gonic/gin"
	"scanner/cmd/api/internal/app/handler"
	"scanner/cmd/api/internal/container"
)

type App struct {
	container *container.Container
	handlers  []handler.Handler
	router    *gin.Engine
}

func (a *App) AppendHandler(handler ...handler.Handler) {
	a.handlers = append(a.handlers, handler...)
}

func (a *App) SetRouter() {
	for _, h := range a.handlers {
		h.Index(a.router)
	}
}

func NewApp(container *container.Container) *App {
	return &App{container: container, router: gin.Default()}
}
