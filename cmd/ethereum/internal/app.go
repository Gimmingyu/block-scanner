package internal

import "scanner/cmd/ethereum/container"

type App struct {
	container *container.Container
}

func New(c *container.Container) *App {
	return &App{container: c}
}

func (a *App) Run() error {
	return nil
}
