package main

import (
	"log"
	"os"
	"scanner/cmd/ethereum/internal/app"
	"scanner/cmd/ethereum/internal/container"
	"scanner/internal/blockchain"
	"scanner/internal/env"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {

	endpoint := os.Getenv("ETHEREUM_NODE_ENDPOINT")
	ethClient, err := blockchain.NewEthClient(endpoint)
	if err != nil {
		log.Panicf("failed to create ethereum client: %v", err)
	}

	c := container.NewContainer(ethClient)
	a := app.New(c)

	if err := a.Run(); err != nil {
		log.Panicf("failed to run ethereum scanner: %v", err)
	}
}
