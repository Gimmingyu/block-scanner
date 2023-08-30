package main

import (
	"log"
	"os"
	"scanner/cmd/klaytn/internal/app"
	"scanner/cmd/klaytn/internal/container"

	"scanner/internal/env"
	"scanner/internal/evm"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {

	endpoint := os.Getenv("KLAYTN_NODE_ENDPOINT")
	ethClient, err := evm.NewEthClient(endpoint)
	if err != nil {
		log.Panicf("failed to create klaytn client: %v", err)
	}

	c := container.NewContainer(evm.NewKlaytnService(ethClient.Client()))
	a := app.NewApp(c)

	if err := a.Run(); err != nil {
		log.Panicf("failed to run klaytn scanner: %v", err)
	}
}
