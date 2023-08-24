package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (a *App) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return a.client.CodeAt(ctx, contract, blockNumber)
}
