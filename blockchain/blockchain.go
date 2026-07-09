package blockchain

import "toy-blockchain/block"	
import "toy-blockchain/ledger"

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

	bc.Blocks = append(
		bc.Blocks,
		newBlock,
	)

	bc.PendingTxs = []ledger.Transaction{}
}