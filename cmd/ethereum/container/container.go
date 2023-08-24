package container

import "github.com/ethereum/go-ethereum/ethclient"

type Container struct {
	client *ethclient.Client
}

func NewContainer(client *ethclient.Client) *Container {
	return &Container{client: client}
}
