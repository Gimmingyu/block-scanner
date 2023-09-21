package main

import (
	"log"
	"os"
	"scanner/cmd/klaytn/internal"
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

	endpoint := os.Getenv("KLAYTN_NODE_ENDPOINT")
	ethClient, err := blockchain.NewEthClient(endpoint)
	if err != nil {
		log.Panicf("failed to create klaytn client: %v", err)
	}

	container := internal.NewContainer(blockchain.NewKlaytnService(ethClient.Client()))
	app := internal.NewApp(container)
	app.SetInterval(time.Minute)

	if err := app.Scan(); err != nil {
		log.Panicf("failed to run klaytn scanner: %v", err)
	}
}
