package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"
	"log"
	"math/big"
	"scanner/internal/blockchain"
	"time"
)

type App struct {
	container *Container
}

func New(c *Container) *App {
	return &App{container: c}
}

func (a *App) Scan() error {
	var (
		group                    *errgroup.Group
		_                        context.Context
		prevBlockNumber          uint64
		currentBlockNumber       uint64
		bigIntCurrentBlockNumber *big.Int
		err                      error
		client                   *blockchain.EthereumService
		block                    map[string]interface{}
		marshalled               []byte
		blockData                *types.Block
	)

	client = a.container.Client()
	group, _ = errgroup.WithContext(context.Background())

	group.Go(func() error {
		for {
			currentBlockNumber, err = client.CurrentBlockNumber()
			if err != nil {
				fmt.Printf("failed to get current block number: %v\n", err)
				goto Sleep
			}

			if prevBlockNumber == currentBlockNumber {
				continue
			}

			log.Println(currentBlockNumber)

			bigIntCurrentBlockNumber = big.NewInt(int64(currentBlockNumber))
			block, err = client.GetBlockByNumber(bigIntCurrentBlockNumber)
			if err != nil {
				log.Printf("failed to get block by number: %v\n", err)
				goto Sleep
			}

			log.Println(block)
			time.Sleep(time.Hour)

			if marshalled, err = json.Marshal(block); err != nil {
				log.Printf("failed to marshal block: %v\n", err)
				goto Sleep
			}

			if err = json.Unmarshal(marshalled, &blockData); err != nil {
				log.Printf("failed to unmarshal block: %v\n", err)
				goto Sleep
			}

			log.Println(blockData)

			for k, v := range blockData.Transactions() {
				log.Printf("%v: %v\n", k, v)
			}

			prevBlockNumber = currentBlockNumber
		Sleep:
			time.Sleep(time.Second * 5)
		}
	})

	return group.Wait()
}
