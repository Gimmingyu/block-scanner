package internal

import (
	"scanner/internal/blockchain"
)

type Container struct {
	client *blockchain.KlaytnService
}

func NewContainer(client *blockchain.KlaytnService) *Container {
	return &Container{client: client}
}

func (c *Container) Client() *blockchain.KlaytnService {
	return c.client
}
