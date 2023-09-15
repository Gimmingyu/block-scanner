package evm

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type EthereumService struct {
	client *ethclient.Client
}

func (s *EthereumService) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return s.client.BalanceAt(ctx, account, blockNumber)
}

func (s *EthereumService) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return s.client.NonceAt(ctx, account, blockNumber)
}

func (s *EthereumService) CurrentBlockNumber() (uint64, error) {
	return s.client.BlockNumber(context.Background())
}

func (s *EthereumService) GetBlockByNumber(number *big.Int) (map[string]interface{}, error) {
	var (
		block        *types.Block
		unmarshalled map[string]interface{}
		marshalled   []byte
		err          error
	)

	block, err = s.client.BlockByNumber(context.Background(), number)
	if err != nil {
		return nil, err
	}

	log.Println(block)

	marshalled, err = json.Marshal(block)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(marshalled, &unmarshalled)

	log.Println(unmarshalled)
	if err != nil {
		return nil, err
	}

	return unmarshalled, nil
}

func (s *EthereumService) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	return s.client.BlockByHash(context.Background(), hash)
}

func (s *EthereumService) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return s.client.CodeAt(ctx, contract, blockNumber)
}

func (s *EthereumService) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return s.client.SuggestGasPrice(ctx)
}

func (s *EthereumService) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return s.client.SuggestGasTipCap(ctx)
}

func (s *EthereumService) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return s.client.EstimateGas(ctx, call)
}

func (s *EthereumService) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return s.client.SendTransaction(ctx, tx)
}

func (s *EthereumService) TransactionCount(ctx context.Context, hash common.Hash) (uint, error) {
	return s.client.TransactionCount(ctx, hash)
}

func (s *EthereumService) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return s.client.TransactionByHash(ctx, hash)
}

func (s *EthereumService) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	return s.client.TransactionReceipt(ctx, hash)
}

func New(client *ethclient.Client) Service {
	return &EthereumService{client: client}
}
