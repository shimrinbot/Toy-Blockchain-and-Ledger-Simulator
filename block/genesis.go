package block

import "toy-blockchain/ledger"

const GenesisPreviousHash = "0000000000000000000000000000000000000000000000000000000000000000"

func NewGenesisBlock() Block {

	block := Block{
		Index:        0,
		Timestamp:    0,
		Transactions: []ledger.Transaction{},
		PreviousHash: GenesisPreviousHash,
		Nonce:        0,
	}

	block.Hash = CalculateHash(block)

	return block
}
