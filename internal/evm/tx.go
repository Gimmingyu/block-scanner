package evm

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (s *service) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return s.client.SuggestGasPrice(ctx)
}

func (s *service) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return s.client.SuggestGasTipCap(ctx)
}

func (s *service) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return s.client.EstimateGas(ctx, call)
}

func (s *service) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return s.client.SendTransaction(ctx, tx)
}

func (s *service) TransactionCount(ctx context.Context, hash common.Hash) (uint, error) {
	return s.client.TransactionCount(ctx, hash)
}

func (s *service) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return s.client.TransactionByHash(ctx, hash)
}

func (s *service) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	return s.client.TransactionReceipt(ctx, hash)
}
