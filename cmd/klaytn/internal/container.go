package internal

import (
	"scanner/internal/models"
	"scanner/pkg/blockchain"
	"scanner/pkg/repository"
)

type Container struct {
	client     *blockchain.KlaytnService
	repository repository.MongoRepository[models.KlaytnTransaction]
}

func (c *Container) Client() *blockchain.KlaytnService {
	return c.client
}

func NewContainer(client *blockchain.KlaytnService, repo repository.MongoRepository[models.KlaytnTransaction]) *Container {
	return &Container{client: client, repository: repo}
}
