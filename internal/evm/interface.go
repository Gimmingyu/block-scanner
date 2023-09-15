package evm

import (
	"context"
	"math/big"
)

type Block interface {
}

type Transaction interface{}

type Receipt interface{}

type Message interface {
}

type Transactions []Transaction

type Service[B Block, T Transaction, R Receipt, M Message] interface {
	BalanceAt(ctx context.Context, account string, blockNumber *big.Int) (*big.Int, error)
	NonceAt(ctx context.Context, account string, blockNumber *big.Int) (uint64, error)
	CurrentBlockNumber() (uint64, error)
	GetBlockByNumber(number *big.Int) (*B, error)
	GetBlockByHash(hash string) (*B, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, call M) (uint64, error)
	SendTransaction(ctx context.Context, tx *T) error
	TransactionCount(ctx context.Context, hash string) (uint, error)
	TransactionByHash(ctx context.Context, hash string) (*T, bool, error)
	TransactionReceipt(ctx context.Context, hash string) (*R, error)
}
