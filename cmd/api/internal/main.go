package main

import (
	"scanner/cmd/api/internal/app"
	"scanner/cmd/api/internal/container"
	"scanner/internal/env"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {
	c := container.NewContainer()
	a := app.NewApp(c)

	a.AppendHandler()
}
