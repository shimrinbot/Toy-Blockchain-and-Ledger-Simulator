package cli

import (
	"fmt"
	"time"

	"toy-blockchain/block"
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

		err = storage.SaveBlockchain(bc, "chain.json")
		if err != nil {
			fmt.Println("Error creating blockchain:", err)
		}
	}

	cli := &CLI{
		Blockchain: bc,
		Ledger:     ledger.NewLedger(),
	}

	// Rebuild ledger state
	for _, b := range bc.Blocks {
		for _, tx := range b.Transactions {
			cli.Ledger.Balances[tx.Sender] -= tx.Amount
			cli.Ledger.Balances[tx.Recipient] += tx.Amount
		}
	}
	for _, tx := range bc.PendingTxs {
		cli.Ledger.Balances[tx.Sender] -= tx.Amount
		cli.Ledger.Balances[tx.Recipient] += tx.Amount
	}

	return cli
}

func (c *CLI) PrintBlockchain() {

	fmt.Println("Blockchain:")

	for _, b := range c.Blockchain.Blocks {

		fmt.Println("--------------------------------")
		fmt.Println("Index:", b.Index)
		fmt.Println("Timestamp:", b.Timestamp)
		fmt.Println("Transactions:", b.Transactions)
		fmt.Println("Previous Hash:", b.PreviousHash)
		fmt.Println("Hash:", b.Hash)
		fmt.Println("Nonce:", b.Nonce)
	}

	fmt.Println("--------------------------------")
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

func (c *CLI) AddTransaction(sender, recipient string, amount float64) {

	tx := ledger.Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}

	err := c.Ledger.ApplyTransaction(tx)
	if err != nil {
		fmt.Println("Transaction Failed:", err)
		return
	}

	c.Blockchain.AddTransaction(tx)

	err = storage.SaveBlockchain(c.Blockchain, "chain.json")
	if err != nil {
		fmt.Println("Error saving transaction:", err)
		return
	}

	fmt.Println("Transaction added and saved.")
}

func (c *CLI) Mine() {

	if len(c.Blockchain.PendingTxs) == 0 {
		fmt.Println("No pending transactions to mine.")
		return
	}

	start := time.Now()

	c.Blockchain.MinePendingTransactions()

	duration := time.Since(start)

	err := storage.SaveBlockchain(c.Blockchain, "chain.json")
	if err != nil {
		fmt.Println("Error saving blockchain:", err)
		return
	}

	last := c.Blockchain.Blocks[len(c.Blockchain.Blocks)-1]

	fmt.Println("--------------------------------")
	fmt.Println("Block mined and saved")
	fmt.Println("Difficulty :", block.Difficulty)
	fmt.Println("Nonce      :", last.Nonce)
	fmt.Println("Hash       :", last.Hash)
	fmt.Println("Mining Time:", duration)
	fmt.Println("--------------------------------")
}

func (c *CLI) Faucet(account string, amount float64) {

	tx := ledger.Transaction{
		Sender:    "SYSTEM",
		Recipient: account,
		Amount:    amount,
	}

	c.Blockchain.AddTransaction(tx)

	err := storage.SaveBlockchain(c.Blockchain, "chain.json")
	if err != nil {
		fmt.Println("Error saving faucet transaction:", err)
		return
	}

	fmt.Println("Faucet added and saved. (Pending to be mined)")
}