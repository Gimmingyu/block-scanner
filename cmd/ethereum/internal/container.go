package internal

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"scanner/internal/blockchain"
	"scanner/internal/models"
	"scanner/pkg/repository"
)

type Container struct {
	client     *blockchain.EthereumService
	repository repository.MongoRepository[models.EthereumTransaction]
}

func (c *Container) Client() *blockchain.EthereumService {
	return c.client
}

func NewContainer(client *ethclient.Client, repo repository.MongoRepository[models.EthereumTransaction]) *Container {
	return &Container{client: blockchain.NewEthereumService(client), repository: repo}
}
