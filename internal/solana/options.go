package solana

type GetBlockOps struct {
	Encoding                       string `json:"encoding"`
	Commitment                     string `json:"commitment"`
	MaxSupportedTransactionVersion int    `json:"maxSupportedTransactionVersion"`
	TransactionDetails             string `json:"transactionDetails"`
}
