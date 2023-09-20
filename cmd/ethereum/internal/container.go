package internal

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"scanner/internal/blockchain"
)

type Container struct {
	client *blockchain.EthereumService
}

func (c *Container) Client() *blockchain.EthereumService {
	return c.client
}

func NewContainer(client *ethclient.Client) *Container {
	return &Container{client: blockchain.NewEthereumService(client)}
}
