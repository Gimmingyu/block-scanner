package internal

type Scanner interface {
	BlockNumber() (uint64, error)
}

type scanner struct {
}
