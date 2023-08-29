package container

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"scanner/internal/evm"
)

type Container struct {
	client *ethclient.Client
	evm.Service
}

func (c *Container) Client() *ethclient.Client {
	return c.client
}

func NewContainer(client *ethclient.Client) *Container {
	return &Container{client: client}
}
