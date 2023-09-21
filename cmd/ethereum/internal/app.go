package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"scanner/internal/models"
	"time"
)

type App struct {
	container *Container
	interval  time.Duration
}

func New(c *Container) *App {
	return &App{container: c}
}

func (a *App) SetInterval(interval time.Duration) {
	a.interval = interval
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
				block                    *types.Block
			)
			block, _err = client.GetBlockByNumber(bigIntCurrentBlockNumber)
			if _err != nil {
				log.Printf("failed to get block by number: %v\n", _err)
				retryCount++
				continue
			}

			var marshalled []byte
			var transactions []*models.EthereumTransaction
			for _, transaction := range block.Transactions() {
				marshalled, _err = transaction.MarshalJSON()
				if _err != nil {
					continue
				}

				document := new(models.EthereumTransaction)
				_err = json.Unmarshal(marshalled, &document)
				if _err != nil {
					continue
				}

				transactions = append(transactions, document)
			}

			if _err = a.container.repository.InsertMany(context.TODO(), transactions); _err != nil {
				log.Printf("failed to insert transactions: %v\n", _err)
				retryCount++
				continue
			}

			prevBlockNumber = currentBlockNumber
			retryCount = 0
			time.Sleep(a.interval)
		}
	}(err)

	return <-err
}
