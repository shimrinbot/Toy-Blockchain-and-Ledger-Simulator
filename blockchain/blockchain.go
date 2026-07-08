package blockchain

import "toy-blockchain/block"

type Blockchain struct {
	Blocks []block.Block
}

func NewBlockchain() *Blockchain {

	bc := &Blockchain{}

	bc.Blocks = append(bc.Blocks, block.NewGenesisBlock())

	return bc
}
