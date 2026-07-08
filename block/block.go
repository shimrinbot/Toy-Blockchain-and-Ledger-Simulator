package block

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []string // Placeholder for now
	PreviousHash string
	Nonce        int
	Hash         string
}
