package container

import (
	"scanner/internal/blockchain"
)

type Container struct {
	client *blockchain.EthereumService
}

func NewContainer(client *blockchain.EthereumService) *Container {
	return &Container{client: client}
}

func (c *Container) Client() *blockchain.EthereumService {
	return c.client
}
