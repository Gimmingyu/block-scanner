package solana

type Service interface {
	GetBalance(address string) (uint64, error)
	CurrentBlock() (uint64, error)
	GetBlockByNumber(blockNumber uint64) (map[string]interface{}, error)
}
