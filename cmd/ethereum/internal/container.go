package internal

import (
	"scanner/internal/documents"
	"scanner/pkg/blockchain"
	"scanner/pkg/repository"
)

type Container struct {
	client     *blockchain.EthereumService
	repository repository.MongoRepository[documents.EthereumTransaction]
}

func (c *Container) Client() *blockchain.EthereumService {
	return c.client
}

func NewContainer(client *blockchain.EthereumService, repo repository.MongoRepository[documents.EthereumTransaction]) *Container {
	return &Container{client: client, repository: repo}
}
