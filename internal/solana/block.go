package solana

import (
	"encoding/json"
	"log"
)

func (c *Client) CurrentBlock() (uint64, error) {
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

	log.Println(string(resp))

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

func (c *Client) GetBlockByNumber(blockNumber uint64, opts *GetBlockOps) (map[string]interface{}, error) {
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
