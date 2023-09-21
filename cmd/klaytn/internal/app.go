package internal

import (
	"fmt"
	"log"
	"math/big"
	"scanner/internal/blockchain"
	"time"
)

type App struct {
	container *Container
	interval  time.Duration
}

func NewApp(container *Container) *App {
	return &App{container: container}
}

func (a *App) SetInterval(interval time.Duration) {
	a.interval = interval
}

func (a *App) Container() *Container {
	return a.container
}

func (a *App) Scan() error {

	var client = a.container.Client()
	var err = make(chan error)
	defer close(err)

	go func(ch chan error) {
		var retryCount = 0
		for {
			if retryCount > 10 {
				ch <- fmt.Errorf("failed to get current block number after 10 retries")
				return
			}

			var (
				prevBlockNumber          = uint64(0)
				currentBlockNumber, _err = client.CurrentBlockNumber()
			)

			if _err != nil {
				fmt.Printf("failed to get current block number: %v\n", _err)
				retryCount++
				continue
			}

			if prevBlockNumber == currentBlockNumber {
				retryCount++
				continue
			}

			var (
				bigIntCurrentBlockNumber = big.NewInt(int64(currentBlockNumber))
				block                    *blockchain.KlaytnBlock
			)
			block, _err = client.GetBlockByNumber(bigIntCurrentBlockNumber)
			if _err != nil {
				log.Printf("failed to get block by number: %v\n", _err)
				retryCount++
				continue
			}

			for _, v := range block.Transactions {
				log.Println("HASH", v.Hash)
				log.Println("BLOCK NUMBER", v.BlockNumber)
				log.Println("BLOCK HASH", v.BlockHash)
				log.Println("TRANSACTION INDEX", v.TransactionIndex)
				log.Println("GAS", v.Gas)
				log.Println("GASPRICE", v.GasPrice)
				log.Println("VALUE", v.Value)
				log.Println("INPUT", v.Input)
				log.Println("FROM", v.From)
				log.Println("TO", v.To)
				log.Println("NONCE", v.Nonce)
			}

			prevBlockNumber = currentBlockNumber
			retryCount = 0
			time.Sleep(a.interval)
		}
	}(err)

	return <-err
}
