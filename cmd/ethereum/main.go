package main

import (
	"log"
	"os"
	"scanner/cmd/ethereum/internal"
	"scanner/internal/documents"
	blockchain2 "scanner/pkg/blockchain"
	"scanner/pkg/connection"
	"scanner/pkg/env"
	"scanner/pkg/repository"
	"time"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {

	endpoint := os.Getenv("ETHEREUM_NODE_ENDPOINT")
	ethClient, err := blockchain2.NewEthClient(endpoint)
	if err != nil {
		log.Panicf("failed to create ethereum client: %v", err)
	}

	mongoClient := connection.NewMongoConnection(os.Getenv("MONGO_URI"))
	transactionRepository := repository.NewMongoRepository[documents.EthereumTransaction](mongoClient.Database("ethereum").Collection("transactions"))

	container := internal.NewContainer(blockchain2.NewEthereumService(ethClient), transactionRepository)
	app := internal.New(container)
	app.SetInterval(time.Minute)
	if err = app.Scan(); err != nil {
		log.Panicf("failed to run ethereum scanner: %v", err)
	}
}
