package evm

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type App struct {
	client *ethclient.Client
}

func New(client *ethclient.Client) *App {
	return &App{client: client}
}
