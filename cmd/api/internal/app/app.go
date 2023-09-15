package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
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

func (a *App) Run() error {
	return a.router.Run(fmt.Sprintf(":%s", os.Getenv("MAIN_PORT")))
}

func NewApp(container *container.Container) *App {
	return &App{container: container, router: gin.Default()}
}
