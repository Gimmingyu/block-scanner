package main

import (
	"log"
	"os"
	"scanner/internal/blockchain"
	"scanner/internal/env"
)

func init() {
	if err := env.LoadEnv(".env"); err != nil {
		panic(err)
	}
}

func main() {
	client := blockchain.NewSolanaService(os.Getenv("SOLANA_NODE_ENDPOINT"))

	blockNumber, err := client.CurrentBlockNumber()
	if err != nil {
		panic(err)
	}

	log.Println("Current block:", blockNumber)
	log.Println(err)

	block, err := client.GetBlockByNumber(blockNumber, &blockchain.GetBlockOps{
		Encoding:                       "jsonParsed",
		MaxSupportedTransactionVersion: 0,
		TransactionDetails:             "full",
	})
	if err != nil {
		panic(err)
	}

	log.Println("Block:", block["result"].(map[string]interface{})["blockhash"])
}
