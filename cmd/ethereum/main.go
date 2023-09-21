package main

import (
	"log"
	"os"
	"scanner/cmd/ethereum/internal"
	"scanner/internal/blockchain"
	"scanner/internal/env"
	"time"
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

	container := internal.NewContainer(ethClient)
	app := internal.New(container)
	app.SetInterval(time.Minute)
	if err = app.Scan(); err != nil {
		log.Panicf("failed to run ethereum scanner: %v", err)
	}
}
