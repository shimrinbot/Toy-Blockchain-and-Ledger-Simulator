package block

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []string
	PreviousHash string
	Nonce        int
	Hash         string
}

func NewBlock(transactions []string, previousHash string, index int) Block {

	block := Block{
		Index:        index,
		Timestamp:    0,
		Transactions: transactions,
		PreviousHash: previousHash,
		Nonce:        0,
	}

	block.Hash = CalculateHash(block)

	return block
}
