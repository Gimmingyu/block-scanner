package main

import (
	"log"
	"os"
	"scanner/cmd/klaytn/internal"
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

	endpoint := os.Getenv("KLAYTN_NODE_ENDPOINT")
	ethClient, err := blockchain2.NewEthClient(endpoint)
	if err != nil {
		log.Panicf("failed to create klaytn client: %v", err)
	}

	mongoClient := connection.NewMongoConnection(os.Getenv("MONGO_URI"))
	transactionRepository := repository.NewMongoRepository[documents.KlaytnTransaction](mongoClient.Database("klaytn").Collection("transactions"))

	container := internal.NewContainer(blockchain2.NewKlaytnService(ethClient.Client()), transactionRepository)
	app := internal.NewApp(container)
	app.SetInterval(time.Minute)
	if err = app.Scan(); err != nil {
		log.Panicf("failed to run klaytn scanner: %v", err)
	}
}
