package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/big"
	"scanner/cmd/klaytn/internal/container"
	"scanner/internal/evm"
	"time"
)

type App struct {
	container *container.Container
}

func NewApp(container *container.Container) *App {
	return &App{container: container}
}

func (a *App) Container() *container.Container {
	return a.container
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

	client = a.Container().Client()
	group, _ = errgroup.WithContext(context.Background())

	group.Go(func() error {
		for {
			currentBlockNumber, err = client.CurrentBlockNumber()
			if err != nil {
				fmt.Printf("failed to get current block number: %v\n", err)
				goto Sleep
			}

			log.Println(currentBlockNumber)
			if prevBlockNumber == currentBlockNumber {
				goto Sleep
			}

			bigIntCurrentBlockNumber = big.NewInt(int64(currentBlockNumber))
			block, err = client.GetBlockByNumber(bigIntCurrentBlockNumber)
			if err != nil {
				fmt.Printf("failed to get block by number: %v\n", err)
				goto Sleep
			}

			log.Println(block)
			//fmt.Printf("block number: %v\n", block.NumberU64())
			// TODO: process block

			prevBlockNumber = currentBlockNumber

		Sleep:
			time.Sleep(time.Second * 5)
		}
	})

	return group.Wait()
}
