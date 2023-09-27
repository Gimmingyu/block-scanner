package internal

import (
	"scanner/internal/documents"
	"scanner/pkg/blockchain"
	"scanner/pkg/repository"
)

type Container struct {
	client     *blockchain.KlaytnService
	repository repository.MongoRepository[documents.KlaytnTransaction]
}

func (c *Container) Client() *blockchain.KlaytnService {
	return c.client
}

func NewContainer(client *blockchain.KlaytnService, repo repository.MongoRepository[documents.KlaytnTransaction]) *Container {
	return &Container{client: client, repository: repo}
}
