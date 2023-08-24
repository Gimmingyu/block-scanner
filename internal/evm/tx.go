package evm

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (a *App) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return a.client.SuggestGasPrice(ctx)
}

func (a *App) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return a.client.SuggestGasTipCap(ctx)
}

func (a *App) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return a.client.EstimateGas(ctx, call)
}

func (a *App) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return a.client.SendTransaction(ctx, tx)
}

func (a *App) TransactionCount(ctx context.Context, hash common.Hash) (uint, error) {
	return a.client.TransactionCount(ctx, hash)
}

func (a *App) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return a.client.TransactionByHash(ctx, hash)
}

func (a *App) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	return a.client.TransactionReceipt(ctx, hash)
}
