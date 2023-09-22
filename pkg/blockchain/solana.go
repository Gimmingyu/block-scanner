package blockchain

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SolanaService struct {
	endpoint  string
	rpcClient *http.Client
}

type GetBlockOps struct {
	Encoding                       string `json:"encoding"`
	Commitment                     string `json:"commitment"`
	MaxSupportedTransactionVersion int    `json:"maxSupportedTransactionVersion"`
	TransactionDetails             string `json:"transactionDetails"`
}

func NewSolanaService(endpoint string) *SolanaService {
	return &SolanaService{endpoint: endpoint, rpcClient: &http.Client{}}
}

func (c *SolanaService) Endpoint() string {
	return c.endpoint
}

func (c *SolanaService) SetEndpoint(endpoint string) {
	c.endpoint = endpoint
}

func (c *SolanaService) CurrentBlockNumber() (uint64, error) {
	var (
		resp        []byte
		err         error
		data        map[string]interface{}
		result      map[string]interface{}
		latestBlock uint64
	)

	resp, err = c.callRPCMethod("getEpochInfo", []interface{}{})
	if err != nil {
		return 0, err
	}

	if err = json.Unmarshal(resp, &data); err != nil {
		return 0, err
	}

	if r, ok := data["result"].(map[string]interface{}); ok {
		result = r
	} else {
		return 0, err
	}

	log.Println(result)

	if result["absoluteSlot"] == nil {
		return 0, err
	}

	if f, ok := result["absoluteSlot"].(float64); ok {
		latestBlock = uint64(f)
	} else {
		return 0, err
	}

	return latestBlock, nil
}

func (c *SolanaService) GetBlockByNumber(blockNumber uint64, opts *GetBlockOps) (map[string]interface{}, error) {
	var (
		resp  []byte
		err   error
		block map[string]interface{}
	)

	resp, err = c.callRPCMethod("getBlock", []interface{}{blockNumber, opts})

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp, &block); err != nil {
		return nil, err
	}

	return block, nil
}

func (c *SolanaService) callRPCMethod(method string, params []interface{}) ([]byte, error) {
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

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return io.ReadAll(resp.Body)
}
