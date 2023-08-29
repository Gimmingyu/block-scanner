package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (s *service) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return s.client.BalanceAt(ctx, account, blockNumber)
}

func (s *service) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return s.client.NonceAt(ctx, account, blockNumber)
}
