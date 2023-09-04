package solana

import "net/http"

type Client struct {
	endpoint  string
	rpcClient *http.Client
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint: endpoint, rpcClient: &http.Client{}}
}

func (c *Client) Endpoint() string {
	return c.endpoint
}

func (c *Client) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}
