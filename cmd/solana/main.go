package main

import (
	"log"
	"os"
	"scanner/internal/env"
	"scanner/internal/solana"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {
	client := solana.NewClient(os.Getenv("SOLANA_NODE_ENDPOINT"))

	blockNumber, err := client.CurrentBlock()
	if err != nil {
		panic(err)
	}

	log.Println("Current block:", blockNumber)
	log.Println(err)

	block, err := client.GetBlockByNumber(blockNumber, &solana.GetBlockOps{
		Encoding:                       "jsonParsed",
		MaxSupportedTransactionVersion: 0,
		TransactionDetails:             "full",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Block:", block["result"].(map[string]interface{})["blockhash"])
}
