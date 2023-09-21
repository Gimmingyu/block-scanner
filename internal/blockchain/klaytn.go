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
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	//TODO implement me
	panic("implement me")
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

func (k *KlaytnService) GetBlockByHash(hash common.Hash) (*types.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) TransactionCount(ctx context.Context, hash common.Hash) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	//TODO implement me
	panic("implement me")
}

func (k *KlaytnService) TransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	//TODO implement me
	panic("implement me")
}

func NewKlaytnService(client *rpc.Client) *KlaytnService {
	return &KlaytnService{client: client}
}
