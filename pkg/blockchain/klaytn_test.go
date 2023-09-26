package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"os"
	"scanner/pkg/env"
	"testing"
)

func init() {
	if err := env.LoadEnv(".test.env"); err != nil {
		panic(err)
	}

	client, err = ethclient.Dial(os.Getenv("KLAYTN_NODE_ENDPOINT"))
	if err != nil {
		panic(err)
	}

	testAccount = common.HexToAddress(os.Getenv("TEST_ACCOUNT"))
	testContract = common.HexToAddress(os.Getenv("TEST_CONTRACT"))
}

func TestKlaytnService_TransactionReceipt(t *testing.T) {
	type fields struct {
		client *rpc.Client
	}
	type args struct {
		ctx  context.Context
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
			name: "TestKlaytnService_TransactionReceipt",
			fields: fields{
				client: client.Client(),
			},
			args: args{
				ctx:  context.TODO(),
				hash: common.HexToHash("0xf65fe2b671968f3a01e4cc3117e5fc839a0ba7d7931386119da7b11f621c0053"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &KlaytnService{
				client: tt.fields.client,
			}
			_, err := k.TransactionReceipt(tt.args.ctx, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionReceipt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
