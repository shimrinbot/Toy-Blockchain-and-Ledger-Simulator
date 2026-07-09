package blockchain

import "toy-blockchain/block"
import "toy-blockchain/ledger"

type Blockchain struct {
	Blocks []block.Block
}

func NewBlockchain() *Blockchain {

	bc := &Blockchain{}

	bc.Blocks = append(bc.Blocks, block.NewGenesisBlock())

	return bc
}

func (bc *Blockchain) AddBlock(transactions []ledger.Transaction) {

	lastBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := block.NewBlock(
		transactions,
		lastBlock.Hash,
		lastBlock.Index+1,
	)

	bc.Blocks = append(bc.Blocks, newBlock)
}
