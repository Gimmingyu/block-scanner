package evm

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (s *service) CurrentBlockNumber() (uint64, error) {
	return s.client.BlockNumber(context.Background())
}

func (s *service) GetBlockByNumber(number *big.Int) (*types.Block, error) {
	return s.client.BlockByNumber(context.Background(), number)
}

func (s *service) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	return s.client.BlockByHash(context.Background(), hash)
}
