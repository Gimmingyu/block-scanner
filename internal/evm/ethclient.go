package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func NewEthClient(endpoint string) (*ethclient.Client, error) {
	conn, err := rpc.Dial(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum node: %v", err)
	}

	return ethclient.NewClient(conn), nil
}
