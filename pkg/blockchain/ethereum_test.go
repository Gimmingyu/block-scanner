package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"reflect"
	"testing"
)

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
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.BalanceAt(tt.args.ctx, tt.args.account, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("BalanceAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BalanceAt() got = %v, want %v", got, tt.want)
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
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.CodeAt(tt.args.ctx, tt.args.contract, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("CodeAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CodeAt() got = %v, want %v", got, tt.want)
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
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.CurrentBlockNumber()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrentBlockNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CurrentBlockNumber() got = %v, want %v", got, tt.want)
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
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.EstimateGas(tt.args.ctx, tt.args.call)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstimateGas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EstimateGas() got = %v, want %v", got, tt.want)
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
		want    *types.Block
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.GetBlockByHash(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockByHash() got = %v, want %v", got, tt.want)
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
		want    *types.Block
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.GetBlockByNumber(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBlockByNumber() got = %v, want %v", got, tt.want)
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
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.NonceAt(tt.args.ctx, tt.args.account, tt.args.blockNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NonceAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NonceAt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthereumService_SendTransaction(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx context.Context
		tx  *types.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			if err := s.SendTransaction(tt.args.ctx, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("SendTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEthereumService_SuggestGasPrice(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.SuggestGasPrice(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SuggestGasPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuggestGasPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthereumService_SuggestGasTipCap(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.SuggestGasTipCap(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SuggestGasTipCap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuggestGasTipCap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthereumService_TransactionByHash(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx  context.Context
		hash common.Hash
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Transaction
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, got1, err := s.TransactionByHash(tt.args.ctx, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionByHash() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TransactionByHash() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEthereumService_TransactionCount(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx  context.Context
		hash common.Hash
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.TransactionCount(tt.args.ctx, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TransactionCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthereumService_TransactionReceipt(t *testing.T) {
	type fields struct {
		client *ethclient.Client
	}
	type args struct {
		ctx  context.Context
		hash common.Hash
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Receipt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EthereumService{
				client: tt.fields.client,
			}
			got, err := s.TransactionReceipt(tt.args.ctx, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionReceipt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionReceipt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEthereumService(t *testing.T) {
	type args struct {
		client *ethclient.Client
	}
	tests := []struct {
		name string
		args args
		want *EthereumService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEthereumService(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEthereumService() = %v, want %v", got, tt.want)
			}
		})
	}
}
