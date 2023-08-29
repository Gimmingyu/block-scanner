package container

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"scanner/internal/evm"
)

type Container struct {
	client evm.Service
}

func (c *Container) Client() evm.Service {
	return c.client
}

func NewContainer(client *ethclient.Client) *Container {
	return &Container{client: evm.New(client)}
}
