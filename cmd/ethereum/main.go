package main

import (
	"log"
	"os"
	"scanner/cmd/ethereum/container"
	"scanner/cmd/ethereum/internal"
	"scanner/internal/env"
	"scanner/internal/evm"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {

	endpoint := os.Getenv("ETHEREUM_NODE_ENDPOINT")
	ethClient, err := evm.NewEthClient(endpoint)
	if err != nil {
		log.Panicf("failed to create ethereum client: %v", err)
	}

	c := container.NewContainer(ethClient)
	app := internal.New(c)

	if err := app.Run(); err != nil {
		log.Panicf("failed to run ethereum scanner: %v", err)
	}
}
