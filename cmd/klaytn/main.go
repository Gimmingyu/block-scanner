package main

import (
	"log"
	"os"
	"scanner/cmd/klaytn/internal"
	"scanner/internal/blockchain"
	"scanner/internal/env"
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

	c := internal.NewContainer(blockchain.NewKlaytnService(ethClient.Client()))
	a := internal.NewApp(c)

	if err := a.Run(); err != nil {
		log.Panicf("failed to run klaytn scanner: %v", err)
	}
}
