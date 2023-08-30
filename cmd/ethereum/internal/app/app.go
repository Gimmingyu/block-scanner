package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/big"
	"scanner/cmd/ethereum/internal/container"
	"scanner/internal/evm"
)

type App struct {
	container *container.Container
}

func New(c *container.Container) *App {
	return &App{container: c}
}

func (a *App) Run() error {
	return a.Scan()
}

func (a *App) Scan() error {
	var (
		group                    *errgroup.Group
		_                        context.Context
		prevBlockNumber          uint64
		currentBlockNumber       uint64
		bigIntCurrentBlockNumber *big.Int
		err                      error
		client                   evm.Service
		block                    map[string]interface{}
	)

	client = a.container.Client()
	group, _ = errgroup.WithContext(context.Background())

	group.Go(func() error {
		for {
			currentBlockNumber, err = client.CurrentBlockNumber()
			if err != nil {
				return fmt.Errorf("failed to get current block number: %v", err)
			}

			if prevBlockNumber == currentBlockNumber {
				continue
			}

			bigIntCurrentBlockNumber = big.NewInt(int64(currentBlockNumber))
			block, err = client.GetBlockByNumber(bigIntCurrentBlockNumber)
			if err != nil {
				return fmt.Errorf("failed to get block by number: %v", err)
			}

			//fmt.Printf("block number: %v\n", block.NumberU64())
			log.Println(block)
			// TODO: process block

			prevBlockNumber = currentBlockNumber
		}
	})

	return group.Wait()
}
