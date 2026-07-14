package block

import (
	"time"
	"toy-blockchain/ledger"
)

type Block struct {
	Index        int
	Timestamp    int64
	Transactions []ledger.Transaction
	MerkleRoot   string
	PreviousHash string
	Nonce        int
	Hash         string
}

func NewBlock(transactions []ledger.Transaction, previousHash string, index int) Block {

	block := Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		MerkleRoot:   CalculateMerkleRoot(transactions),
		PreviousHash: previousHash,
		Nonce:        0,
	}

	//block.Hash = CalculateHash(block)
	MineBlock(&block)
	return block
}
