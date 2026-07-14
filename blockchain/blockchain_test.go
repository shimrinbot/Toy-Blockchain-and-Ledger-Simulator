package blockchain

import (
	"testing"

	"toy-blockchain/ledger"
)


func TestValidBlockchain(t *testing.T) {

	bc := NewBlockchain()

	tx := ledger.Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    50,
	}

	bc.AddTransaction(tx)

	bc.MinePendingTransactions("Miner")


	err := bc.Validate()

	if err != nil {
		t.Error("Expected blockchain to be valid")
	}
}



func TestTamperedBlockchain(t *testing.T) {

	bc := NewBlockchain()


	tx := ledger.Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    50,
	}


	bc.AddTransaction(tx)

	bc.MinePendingTransactions("Miner")


	// Modify block data after mining
	bc.Blocks[1].Transactions[0].Amount = 5000


	err := bc.Validate()


	if err == nil {
		t.Error("Expected tampered blockchain to fail")
	}
}