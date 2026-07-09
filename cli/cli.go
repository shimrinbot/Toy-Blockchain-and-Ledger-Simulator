package cli

import (
	"fmt"
	"toy-blockchain/blockchain"
	"toy-blockchain/ledger"
	"toy-blockchain/storage"
)

type CLI struct {
	Blockchain *blockchain.Blockchain
	Ledger     *ledger.Ledger
}

func NewCLI() *CLI {

	bc, err := storage.LoadBlockchain("chain.json")

	if err != nil {

		bc = blockchain.NewBlockchain()

		storage.SaveBlockchain(
			bc,
			"chain.json",
		)
	}


	return &CLI{
		Blockchain: bc,
		Ledger: ledger.NewLedger(),
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

func (c *CLI) AddTransaction(
	sender string,
	recipient string,
	amount float64,
) {

	tx := ledger.Transaction{
		Sender: sender,
		Recipient: recipient,
		Amount: amount,
	}

	c.Blockchain.AddTransaction(tx)

	fmt.Println("Transaction added")
}
func (c *CLI) Mine() {

	c.Blockchain.MinePendingTransactions()


	err := storage.SaveBlockchain(
		c.Blockchain,
		"chain.json",
	)


	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println("Block mined and saved")
}