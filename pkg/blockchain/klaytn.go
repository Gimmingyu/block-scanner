package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strconv"
)

type KlaytnService struct {
	client *rpc.Client
}

type KlaytnTransaction struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
}

type KlaytnBlock struct {
	Number           string               `json:"number"`
	Hash             string               `json:"hash"`
	ParentHash       string               `json:"parentHash"`
	Nonce            string               `json:"nonce"`
	Sha3Uncles       string               `json:"sha3Uncles"`
	LogsBloom        string               `json:"logsBloom"`
	TransactionsRoot string               `json:"transactionsRoot"`
	StateRoot        string               `json:"stateRoot"`
	ReceiptsRoot     string               `json:"receiptsRoot"`
	Miner            string               `json:"miner"`
	Difficulty       string               `json:"difficulty"`
	TotalDifficulty  string               `json:"totalDifficulty"`
	ExtraData        string               `json:"extraData"`
	Size             string               `json:"size"`
	GasLimit         string               `json:"gasLimit"`
	GasUsed          string               `json:"gasUsed"`
	Timestamp        string               `json:"timestamp"`
	Transactions     []*KlaytnTransaction `json:"transactions"`
	Uncles           []string             `json:"uncles"`
}

func (k *KlaytnService) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var balanceStr string
	if err := k.client.Call(&balanceStr, "klay_getBalance", account.Hex(), fmt.Sprintf("0x%x", blockNumber.Uint64())); err != nil {
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
	if err := k.client.Call(&nonceStr, "klay_getTransactionCount", account.Hex(), fmt.Sprintf("0x%x", blockNumber.Uint64())); err != nil {
		return 0, err
	}

	nonce, err := strconv.ParseUint(nonceStr[2:], 16, 64)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

func (k *KlaytnService) CurrentBlockNumber() (uint64, error) {
	var (
		stringNumber       string
		currentBlockNumber uint64
		err                error
	)

	err = k.client.Call(&stringNumber, "klay_blockNumber")
	if err != nil {
		return 0, err
	}

	currentBlockNumber, err = strconv.ParseUint(stringNumber, 0, 64)
	if err != nil {
		return 0, err
	}

	return currentBlockNumber, nil
}

func (k *KlaytnService) GetBlockByNumber(number *big.Int) (*KlaytnBlock, error) {
	var (
		block = new(KlaytnBlock)
		err   error
	)
	err = k.client.Call(&block, "klay_getBlockByNumber", fmt.Sprintf("0x%x", number.Uint64()), true)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (k *KlaytnService) GetBlockByHash(hash common.Hash) (*KlaytnBlock, error) {
	var block = new(KlaytnBlock)
	if err := k.client.Call(&block, "klay_getBlockByHash", hash.Hex(), true); err != nil {
		return nil, err
	}
	return block, nil
}

func (k *KlaytnService) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var gasPriceStr string
	if err := k.client.Call(&gasPriceStr, "klay_gasPrice"); err != nil {
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
	if err := k.client.Call(&gasLimitStr, "klay_estimateGas", call); err != nil {
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
	if err := k.client.Call(&txHash, "klay_sendRawTransaction", data); err != nil {
		return err
	}
	return nil
}

func (k *KlaytnService) TransactionCount(ctx context.Context, hash common.Hash) (uint, error) {
	var txCount uint
	if err := k.client.Call(&txCount, "klay_getBlockTransactionCountByHash", hash.Hex()); err != nil {
		return 0, err
	}
	return txCount, nil
}

func (k *KlaytnService) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	var tx *types.Transaction
	if err := k.client.Call(&tx, "klay_getTransactionByHash", hash.Hex()); err != nil {
		return nil, false, err
	}
	return tx, true, nil
}

func (k *KlaytnService) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	if err := k.client.Call(&receipt, "klay_getTransactionReceipt", hash.Hex()); err != nil {
		return nil, err
	}
	return receipt, nil
}

func NewKlaytnService(client *rpc.Client) *KlaytnService {
	return &KlaytnService{client: client}
}
