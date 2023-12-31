package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"scanner/internal/documents"
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
		var currentBlockNumber uint64 = 0
		var _err error
		for {
			if retryCount > 10 {
				ch <- fmt.Errorf("failed to get current block number after 10 retries")
				return
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
			var transactions []*documents.EthereumTransaction
			for _, transaction := range block.Transactions() {
				marshalled, _err = transaction.MarshalJSON()
				if _err != nil {
					continue
				}

				document := new(documents.EthereumTransaction)
				_err = json.Unmarshal(marshalled, &document)
				if _err != nil {
					continue
				}

				transactions = append(transactions, document)
			}

			if len(transactions) == 0 {
				log.Println(currentBlockNumber)
				currentBlockNumber++
				continue
			}

			if _err = a.container.repository.InsertMany(context.TODO(), transactions); _err != nil {
				log.Printf("failed to insert transactions: %v\n", _err)
				retryCount++
				continue
			}

			retryCount = 0
			log.Println(currentBlockNumber)
			currentBlockNumber++
			time.Sleep(a.interval)
		}
	}(err)

	return <-err
}
