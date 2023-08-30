package container

import "scanner/internal/evm"

type Container struct {
	client evm.Service
}

func NewContainer(client evm.Service) *Container {
	return &Container{client: client}
}

func (c *Container) Client() evm.Service {
	return c.client
}
