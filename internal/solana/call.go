package solana

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) callRPCMethod(method string, params []interface{}) ([]byte, error) {
	var (
		reqBody      map[string]interface{}
		reqBodyBytes []byte
		resp         *http.Response
		err          error
	)

	reqBody = map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}

	reqBodyBytes, err = json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err = http.Post(c.Endpoint(), "application/json", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
