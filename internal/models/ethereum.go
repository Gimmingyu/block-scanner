package models

import (
	"github.com/ethereum/go-ethereum/common"
)

type EthereumBlockData struct {
}

type AccessTuple struct {
	Address     common.Address `json:"address"     bson:"address"`
	StorageKeys []common.Hash  `json:"storageKeys" bson:"storageKeys"`
}

type EthereumTransaction struct {
	Type                 string         `json:"type" bson:"type"`
	ChainID              string         `json:"chainID" bson:"chainID"`
	Nonce                string         `json:"nonce" bson:"nonce"`
	To                   string         `json:"to" bson:"to"`
	Gas                  string         `json:"gas" bson:"gas"`
	GasPrice             string         `json:"gasPrice" bson:"gasPrice"`
	MaxPriorityFeePerGas string         `json:"maxPriorityFeePerGas" bson:"maxPriorityFeePerGas"`
	MaxFeePerGas         string         `json:"maxFeePerGas" bson:"maxFeePerGas"`
	Value                string         `json:"value" bson:"value"`
	Input                string         `json:"input" bson:"input"`
	AccessList           []*AccessTuple `json:"accessList" bson:"accessList"`
	V                    string         `json:"v" bson:"v"`
	R                    string         `json:"r" bson:"r"`
	S                    string         `json:"s" bson:"s"`
	YParity              string         `json:"yParity" bson:"YParity"`
	Hash                 string         `json:"hash" bson:"hash"`
}

type EthereumTransactions []EthereumTransaction
