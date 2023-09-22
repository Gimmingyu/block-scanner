package models

type KlaytnBlockData struct {
	ReceiptsRoot    string             `json:"receiptsRoot,omitempty" bson:"receiptsRoot"`
	Timestamp       string             `json:"timestamp,omitempty" bson:"timestamp"`
	ExtraData       string             `json:"extraData,omitempty" bson:"extraData"`
	Number          string             `json:"number,omitempty" bson:"number"`
	TotalBlockScore string             `json:"totalBlockScore,omitempty" bson:"totalBlockScore"`
	Reward          string             `json:"reward,omitempty" bson:"reward"`
	BaseFeePerGas   string             `json:"baseFeePerGas,omitempty" bson:"baseFeePerGas"`
	Hash            string             `json:"hash,omitempty" bson:"hash"`
	Size            string             `json:"size,omitempty" bson:"size"`
	StateRoot       string             `json:"stateRoot,omitempty" bson:"stateRoot"`
	ParentHash      string             `json:"parentHash,omitempty" bson:"parentHash"`
	Transactions    KlaytnTransactions `json:"transactions,omitempty" bson:"transactions"`
}

type KlaytnTransaction struct {
	From               string     `json:"from,omitempty" bson:"from"`
	Input              string     `json:"input,omitempty" bson:"input"`
	To                 string     `json:"to,omitempty" bson:"to"`
	TypeInt            float64    `json:"typeInt,omitempty" bson:"typeInt"`
	BlockHash          string     `json:"blockHash,omitempty" bson:"blockHash"`
	SenderTxHash       string     `json:"senderTxHash,omitempty" bson:"senderTxHash"`
	FeePayer           string     `json:"feePayer,omitempty" bson:"feePayer"`
	FeePayerSignatures Signatures `json:"feePayerSignatures,omitempty" bson:"feePayerSignatures"`
	TransactionIndex   string     `json:"transactionIndex,omitempty" bson:"transactionIndex"`
	Nonce              string     `json:"nonce,omitempty" bson:"nonce"`
	Gas                string     `json:"gas,omitempty" bson:"gas"`
	GasPrice           string     `json:"gasPrice,omitempty" bson:"gasPrice"`
	Value              string     `json:"value,omitempty" bson:"value"`
	Signatures         Signatures `json:"signatures,omitempty" bson:"signatures"`
	Type               string     `json:"type,omitempty" bson:"type"`
	BlockNumber        string     `json:"blockNumber,omitempty" bson:"blockNumber"`
}

type KlaytnTransactions []KlaytnTransaction
