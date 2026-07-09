package main

import (
	"fmt"

	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
)

func main() {

	bc := blockchain.NewBlockchain()

	bc.AddBlock([]ledger.Transaction{
		{
			Sender:    "Alice",
			Recipient: "Bob",
			Amount:    50,
		},
	})

	bc.AddBlock([]ledger.Transaction{
		{
			Sender:    "Bob",
			Recipient: "Charlie",
			Amount:    20,
		},
	})

	fmt.Println("Blockchain created!")
	fmt.Println()

	for _, b := range bc.Blocks {
		fmt.Println("Index:", b.Index)
		fmt.Println("Timestamp:", b.Timestamp)
		fmt.Println("Transactions:", b.Transactions)
		fmt.Println("Previous Hash:", b.PreviousHash)
		fmt.Println("Hash:", b.Hash)
		fmt.Println("--------------------------------")
	}

	// ===== Ledger Test =====

	l := ledger.NewLedger()

	l.Faucet("Alice", 100)

	err := l.ApplyTransaction(ledger.Transaction{
	Sender:    "Alice",
	Recipient: "Bob",
	Amount:    50,
})

if err != nil {
	fmt.Println(err)
}

		err = l.ApplyTransaction(ledger.Transaction{
		Sender:    "Bob",
		Recipient: "Charlie",
		Amount:    20,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = l.ApplyTransaction(ledger.Transaction{
	Sender:    "Alice",
	Recipient: "Charlie",
	Amount:    500,
})

if err != nil {
	fmt.Println(err)
}

err = l.ApplyTransaction(ledger.Transaction{
	Sender:    "Bob",
	Recipient: "Charlie",
	Amount:    -10,
})

if err != nil {
	fmt.Println(err)
}

	l.PrintBalances()

	l.PrintBalances()

bc.Blocks[1].Transactions[0].Amount = 5000

bc.Blocks[1].Transactions[0].Amount = 5000

// Validate the blockchain
err = bc.Validate()

if err != nil {
	fmt.Println("Validation Failed:", err)
} else {
	fmt.Println("Blockchain is valid ✅")
}
}