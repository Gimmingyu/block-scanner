package internal

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"
	"math/big"
	"scanner/internal/evm"
)

func (a *App) Scan() error {
	var (
		group                    *errgroup.Group
		_                        context.Context
		currentBlockNumber       uint64
		bigIntCurrentBlockNumber *big.Int
		err                      error
		client                   evm.Service
		block                    *types.Block
	)

	client = a.container.Client()
	group, _ = errgroup.WithContext(context.Background())

	group.Go(func() error {
		for {
			currentBlockNumber, err = client.CurrentBlockNumber()
			if err != nil {
				return fmt.Errorf("failed to get current block number: %v", err)
			}

			bigIntCurrentBlockNumber = big.NewInt(int64(currentBlockNumber))
			block, err = client.GetBlockByNumber(bigIntCurrentBlockNumber)
			if err != nil {
				return fmt.Errorf("failed to get block by number: %v", err)
			}

			fmt.Printf("block number: %v\n", block.NumberU64())
			// TODO: process block
		}
	})

	return group.Wait()
}
