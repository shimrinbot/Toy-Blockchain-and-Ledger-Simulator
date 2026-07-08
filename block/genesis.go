package block

const GenesisPreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"

func NewGenesisBlock() Block {

	block := Block{
		Index:        0,
		Timestamp:    0,
		Transactions: []string{"Genesis Block"},
		PreviousHash: GenesisPreviousHash,
		Nonce:        0,
	}

	block.Hash = CalculateHash(block)

	return block
}
