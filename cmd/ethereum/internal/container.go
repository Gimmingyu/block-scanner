package internal

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"scanner/internal/blockchain"
	"scanner/pkg/repository"
)

type Container struct {
	client *blockchain.EthereumService
	repository.MongoRepository[any]
}

func (c *Container) Client() *blockchain.EthereumService {
	return c.client
}

func NewContainer(client *ethclient.Client) *Container {
	return &Container{client: blockchain.NewEthereumService(client)}
}
