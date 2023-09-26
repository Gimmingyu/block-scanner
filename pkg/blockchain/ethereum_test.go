package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"scanner/pkg/env"
	"testing"
)

var (
	client       *ethclient.Client
	testAccount  common.Address
	testContract common.Address
	err          error
)

func init() {
	if err := env.LoadEnv(".test.env"); err != nil {
		panic(err)
	}

	client, err = ethclient.Dial(os.Getenv("ETHEREUM_NODE_ENDPOINT"))
	if err != nil {
		panic(err)
	}

	testAccount = common.HexToAddress(os.Getenv("TEST_ACCOUNT"))
	testContract = common.HexToAddress(os.Getenv("TEST_CONTRACT"))
}

func TestEthereumService_BalanceAt(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx         context.Context
		account     common.Address
		blockNumber *big.Int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_BalanceAt",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:         context.Background(),
				account:     testAccount,
				blockNumber: new(big.Int).SetUint64(0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err = s.BalanceAt(tt.args.ctx, tt.args.account, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("BalanceAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEthereumService_CodeAt(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx         context.Context
		contract    common.Address
		blockNumber *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_CodeAt",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:         context.Background(),
				contract:    testContract,
				blockNumber: new(big.Int).SetUint64(18220110),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err = s.CodeAt(tt.args.ctx, tt.args.contract, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("CodeAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEthereumService_CurrentBlockNumber(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_CurrentBlockNumber",
			fields: fields{
				client: client,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err := s.CurrentBlockNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrentBlockNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEthereumService_EstimateGas(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx  context.Context
		call ethereum.CallMsg
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_EstimateGas",
			fields: fields{
				client: client,
			},
			args: args{
				ctx: context.Background(),
				call: ethereum.CallMsg{
					From: testAccount,
					To:   &testContract,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err := s.EstimateGas(tt.args.ctx, tt.args.call)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstimateGas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEthereumService_GetBlockByHash(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		hash common.Hash
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_GetBlockByHash",
			fields: fields{
				client: client,
			},
			args: args{
				hash: common.HexToHash("0x41be2acf31b6b5265da905943ff8e67c392e210dc421ab66ba9856a937be281f"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err := s.GetBlockByHash(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEthereumService_GetBlockByNumber(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		number *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_GetBlockByNumber",
			fields: fields{
				client: client,
			},
			args: args{
				number: new(big.Int).SetUint64(18220110),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err := s.GetBlockByNumber(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEthereumService_NonceAt(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx         context.Context
		account     common.Address
		blockNumber *big.Int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestEthereumService_NonceAt",
			fields: fields{
				client: client,
			},
			args: args{
				ctx:         context.Background(),
				account:     testAccount,
				blockNumber: new(big.Int).SetUint64(18220110),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			_, err := s.NonceAt(tt.args.ctx, tt.args.account, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NonceAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
