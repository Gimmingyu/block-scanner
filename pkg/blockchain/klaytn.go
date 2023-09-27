package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"scanner/internal/documents"
	"strconv"
)

type KlaytnService struct {
	client *rpc.Client
}

func (k *KlaytnService) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var balanceStr string
	if err := k.client.CallContext(ctx, &balanceStr, "klay_getBalance", account.Hex(), fmt.Sprintf("0x%x", blockNumber.Uint64())); err != nil {
		return nil, err
	}

	balance, ok := new(big.Int).SetString(balanceStr[2:], 16)
	if !ok {
		return nil, fmt.Errorf("failed to convert balance to big.Int")
	}

	return balance, nil
}

func (k *KlaytnService) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var nonceStr string
	if err := k.client.CallContext(ctx, &nonceStr, "klay_getTransactionCount", account.Hex(), fmt.Sprintf("0x%x", blockNumber.Uint64())); err != nil {
		return 0, err
	}

	nonce, err := strconv.ParseUint(nonceStr[2:], 16, 64)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (k *KlaytnService) CurrentBlockNumber(ctx context.Context) (uint64, error) {
	var (
		stringNumber       string
		currentBlockNumber uint64
		err                error
	)

	err = k.client.CallContext(ctx, &stringNumber, "klay_blockNumber")
	if err != nil {
		return 0, err
	}

	currentBlockNumber, err = strconv.ParseUint(stringNumber, 0, 64)
	if err != nil {
		return 0, err
	}

	return currentBlockNumber, nil
}

func (k *KlaytnService) GetBlockByNumber(ctx context.Context, number *big.Int) (*documents.KlaytnBlock, error) {
	var (
		block = new(documents.KlaytnBlock)
		err   error
	)
	err = k.client.CallContext(ctx, &block, "klay_getBlockByNumber", fmt.Sprintf("0x%x", number.Uint64()), true)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (k *KlaytnService) GetBlockByHash(ctx context.Context, hash common.Hash) (*documents.KlaytnBlock, error) {
	var block = new(documents.KlaytnBlock)
	if err := k.client.CallContext(ctx, &block, "klay_getBlockByHash", hash.Hex(), true); err != nil {
		return nil, err
	}
	return block, nil
}

func (k *KlaytnService) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var gasPriceStr string
	if err := k.client.CallContext(ctx, &gasPriceStr, "klay_gasPrice"); err != nil {
		return nil, err
	}

	gasPrice, ok := new(big.Int).SetString(gasPriceStr[2:], 16)
	if !ok {
		return nil, fmt.Errorf("Failed to convert gas price to big.Int")
	}

	return gasPrice, nil
}

func (k *KlaytnService) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	var gasLimitStr string
	if err := k.client.CallContext(ctx, &gasLimitStr, "klay_estimateGas", call); err != nil {
		return 0, err
	}

	gasLimit, err := strconv.ParseUint(gasLimitStr[2:], 16, 64)
	if err != nil {
		return 0, err
	}

	return gasLimit, nil
}

func (k *KlaytnService) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := tx.MarshalJSON()
	if err != nil {
		return err
	}

	var txHash common.Hash
	if err := k.client.CallContext(ctx, &txHash, "klay_sendRawTransaction", data); err != nil {
		return err
	}
	return nil
}

func (k *KlaytnService) TransactionCount(ctx context.Context, hash common.Hash) (uint, error) {
	var txCount uint
	if err := k.client.CallContext(ctx, &txCount, "klay_getBlockTransactionCountByHash", hash.Hex()); err != nil {
		return 0, err
	}
	return txCount, nil
}

func (k *KlaytnService) TransactionByHash(ctx context.Context, hash common.Hash) (*documents.KlaytnTransaction, bool, error) {
	var tx = new(documents.KlaytnTransaction)
	if err := k.client.CallContext(ctx, &tx, "klay_getTransactionByHash", hash.Hex()); err != nil {
		return nil, false, err
	}
	return tx, true, nil
}

func (k *KlaytnService) TransactionReceipt(ctx context.Context, hash common.Hash) (*documents.KlaytnTransaction, error) {
	var receipt = new(documents.KlaytnTransaction)
	if err := k.client.CallContext(ctx, &receipt, "klay_getTransactionReceipt", hash.Hex()); err != nil {
		return nil, err
	}
	return receipt, nil
}

func NewKlaytnService(client *rpc.Client) *KlaytnService {
	return &KlaytnService{client: client}
}
