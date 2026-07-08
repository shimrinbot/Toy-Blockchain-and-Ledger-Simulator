package block

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []string
	PreviousHash string
	Nonce        int
	Hash         string
}
