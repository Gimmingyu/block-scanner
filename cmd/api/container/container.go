package container

import "github.com/ethereum/go-ethereum/ethclient"

type Container struct {
	client *ethclient.Client
}
