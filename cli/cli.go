package cli

import (
	"fmt"
	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
)

type CLI struct {
	Blockchain *blockchain.Blockchain
	Ledger     *ledger.Ledger
}

func NewCLI() *CLI {

	return &CLI{
		Blockchain: blockchain.NewBlockchain(),
		Ledger:     ledger.NewLedger(),
	}
}
func (c *CLI) PrintBlockchain() {

	fmt.Println("Blockchain:")

	for _, b := range c.Blockchain.Blocks {

		fmt.Println("----------------")
		fmt.Println("Index:", b.Index)
		fmt.Println("Timestamp:", b.Timestamp)
		fmt.Println("Transactions:", b.Transactions)
		fmt.Println("Previous Hash:", b.PreviousHash)
		fmt.Println("Hash:", b.Hash)
	}
}

func (c *CLI) ValidateBlockchain() {

	err := c.Blockchain.Validate()

	if err != nil {
		fmt.Println("Validation Failed:", err)
		return
	}

	fmt.Println("Blockchain is valid ✅")
}

func (c *CLI) PrintBalances() {

	c.Ledger.PrintBalances()

}
