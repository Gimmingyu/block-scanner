package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/big"
	"scanner/internal/blockchain"
	"scanner/internal/models"
	"time"
)

type App struct {
	container *Container
}

func NewApp(container *Container) *App {
	return &App{container: container}
}

func (a *App) Container() *Container {
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
		client                   *blockchain.KlaytnService
		block                    map[string]interface{}
		marshalled               []byte
		blockData                = new(models.KlaytnBlockData)
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

			if marshalled, err = json.Marshal(block); err != nil {
				fmt.Printf("failed to marshal block: %v\n", err)
				goto Sleep
			}

			if err = json.Unmarshal(marshalled, &blockData); err != nil {
				fmt.Printf("failed to unmarshal block data: %v\n", err)
				goto Sleep
			}

			for k, v := range blockData.Transactions {
				log.Printf("%v: %v\n", k, v)
			}

			//fmt.Printf("block number: %v\n", block.NumberU64())
			// TODO: process block

			prevBlockNumber = currentBlockNumber

		Sleep:
			time.Sleep(time.Second * 5)
		}
	})

	return group.Wait()
}
