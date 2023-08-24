package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (a *App) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return a.client.BalanceAt(ctx, account, blockNumber)
}

func (a *App) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return a.client.NonceAt(ctx, account, blockNumber)
}
