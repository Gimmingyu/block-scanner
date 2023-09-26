package blockchain

import (
	"os"
	"scanner/pkg/env"
	"testing"
)

func TestNewEthClient(t *testing.T) {
	if err := env.LoadEnv(".test.env"); err != nil {
		panic(err)
	}

	type args struct {
		endpoint string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNewEthClient",
			args: args{
				endpoint: os.Getenv("ETHEREUM_NODE_ENDPOINT"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewEthClient(tt.args.endpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEthClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
