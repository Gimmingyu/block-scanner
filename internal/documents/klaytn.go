package documents

type KlaytnTransaction struct {
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
}

type KlaytnBlock struct {
	Number           string               `json:"number"`
	Hash             string               `json:"hash"`
	ParentHash       string               `json:"parentHash"`
	Nonce            string               `json:"nonce"`
	Sha3Uncles       string               `json:"sha3Uncles"`
	LogsBloom        string               `json:"logsBloom"`
	TransactionsRoot string               `json:"transactionsRoot"`
	StateRoot        string               `json:"stateRoot"`
	ReceiptsRoot     string               `json:"receiptsRoot"`
	Miner            string               `json:"miner"`
	Difficulty       string               `json:"difficulty"`
	TotalDifficulty  string               `json:"totalDifficulty"`
	ExtraData        string               `json:"extraData"`
	Size             string               `json:"size"`
	GasLimit         string               `json:"gasLimit"`
	GasUsed          string               `json:"gasUsed"`
	Timestamp        string               `json:"timestamp"`
	Transactions     []*KlaytnTransaction `json:"transactions"`
	Uncles           []string             `json:"uncles"`
}

type KlaytnTransactions []KlaytnTransaction
