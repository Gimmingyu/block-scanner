package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (s *service) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return s.client.CodeAt(ctx, contract, blockNumber)
}
