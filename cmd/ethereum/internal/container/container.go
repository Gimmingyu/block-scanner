package container

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"scanner/internal/blockchain"
)

type Container struct {
	client blockchain.Service
}

func (c *Container) Client() blockchain.Service {
	return c.client
}

func NewContainer(client *ethclient.Client) *Container {
	return &Container{client: blockchain.New(client)}
}
