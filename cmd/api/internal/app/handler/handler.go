package handler

import "scanner/cmd/api/internal/container"

func Handlers(container *container.Container) []Handler {
	return []Handler{
		NewAuthHandler(container),
	}
}
