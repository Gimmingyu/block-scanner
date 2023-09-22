package internal

import (
	"scanner/internal/models"
	"scanner/pkg/blockchain"
	"scanner/pkg/repository"
)

type Container struct {
	client     *blockchain.EthereumService
	repository repository.MongoRepository[models.EthereumTransaction]
}

func (c *Container) Client() *blockchain.EthereumService {
	return c.client
}

func NewContainer(client *blockchain.EthereumService, repo repository.MongoRepository[models.EthereumTransaction]) *Container {
	return &Container{client: client, repository: repo}
}
