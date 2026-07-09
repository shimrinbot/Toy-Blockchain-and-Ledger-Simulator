package block

import "toy-blockchain/ledger"

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []ledger.Transaction
	PreviousHash string
	Nonce        int
	Hash         string
}

func NewBlock(transactions []ledger.Transaction, previousHash string, index int) Block {

	block := Block{
		Index:        index,
		Timestamp:    0,
		Transactions: transactions,
		PreviousHash: previousHash,
		Nonce:        0,
	}

	//block.Hash = CalculateHash(block)
	MineBlock(&block)
	return block
}
