package blockchain

import (
	"fmt"
	"toy-blockchain/block"
	"toy-blockchain/ledger"
)

type Blockchain struct {
	Blocks       []block.Block
	PendingTxs   []ledger.Transaction
}

func NewBlockchain() *Blockchain {

	genesis := block.NewGenesisBlock()

	return &Blockchain{
	Blocks:     []block.Block{genesis},
	PendingTxs: []ledger.Transaction{},
}
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

func (bc *Blockchain) AddTransaction(tx ledger.Transaction) {

	bc.PendingTxs = append(bc.PendingTxs, tx)

}

func (bc *Blockchain) MinePendingTransactions() {

	if len(bc.PendingTxs) == 0 {
		return
	}

	previousBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := block.NewBlock(
		bc.PendingTxs,
		previousBlock.Hash,
		previousBlock.Index+1,
	)

	timeDiff := newBlock.Timestamp - previousBlock.Timestamp
	fmt.Printf("Time since last block: %ds (Target: ~5s)\n", timeDiff)

	if timeDiff < 5 {
		block.Difficulty++
		fmt.Printf("Adjusting difficulty UP to %d for the next block.\n", block.Difficulty)
	} else if timeDiff > 10 && block.Difficulty > 1 {
		block.Difficulty--
		fmt.Printf("Adjusting difficulty DOWN to %d for the next block.\n", block.Difficulty)
	}

	bc.Blocks = append(
		bc.Blocks,
		newBlock,
	)

	bc.PendingTxs = []ledger.Transaction{}
}