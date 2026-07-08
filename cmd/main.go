package main

import (
	"fmt"

	"toy-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	fmt.Println("Blockchain created!")
	fmt.Println()

	for _, b := range bc.Blocks {
		fmt.Println("Index:", b.Index)
		fmt.Println("Timestamp:", b.Timestamp)
		fmt.Println("Transactions:", b.Transactions)
		fmt.Println("Previous Hash:", b.PreviousHash)
		fmt.Println("Hash:", b.Hash)
	}
}
