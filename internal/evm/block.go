package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (a *App) CurrentBlockNumber() (uint64, error) {
	return a.client.BlockNumber(context.Background())
}

func (a *App) GetBlockByNumber(number *big.Int) (*types.Block, error) {
	return a.client.BlockByNumber(context.Background(), number)
}

func (a *App) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	return a.client.BlockByHash(context.Background(), hash)
}
