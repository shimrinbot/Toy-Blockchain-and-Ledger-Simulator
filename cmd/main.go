package main

import (
	"fmt"
	"os"
	"strconv"
	"toy-blockchain/cli"
	"toy-blockchain/block"
)

func main() {

	app := cli.NewCLI()

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  print")
		fmt.Println("  validate")
		fmt.Println("  balances")
		return
	}

	command := os.Args[1]

	// Optional difficulty argument
for i := 2; i < len(os.Args)-1; i++ {
	if os.Args[i] == "--difficulty" {

		level, err := strconv.Atoi(os.Args[i+1])

		if err == nil {
			block.SetDifficulty(level)
			fmt.Println("Mining difficulty set to:", level)
		}
	}
}

	switch command {

	case "print":
		app.PrintBlockchain()

	case "validate":
		app.ValidateBlockchain()

	case "balances":
		app.PrintBalances()

	case "faucet":
		if len(os.Args) < 4 {
			fmt.Println("Usage: faucet account amount")
			return
		}
		amount, _ := strconv.ParseFloat(os.Args[3], 64)
		app.Faucet(os.Args[2], amount)

	case "add":

		if len(os.Args) < 4 {
			fmt.Println("Usage: add receiver amount")
			return
		}

		amount, _ := strconv.ParseFloat(os.Args[3], 64)

		app.AddTransaction(
			os.Args[2],
			amount,
		)


	case "mine":
		app.Mine()

	case "sync":
		if len(os.Args) < 3 {
			fmt.Println("Usage: sync <filename>")
			return
		}
		app.Sync(os.Args[2])

	case "help":

		fmt.Println("Available Commands:")
		fmt.Println("  faucet <account> <amount>")
		fmt.Println("  add <receiver> <amount>")
		fmt.Println("  mine")
		fmt.Println("  sync <filename>")
		fmt.Println("  print")
		fmt.Println("  validate")
		fmt.Println("  balances")
		fmt.Println("  help")

	default:
	fmt.Println("Unknown command:", command)
	fmt.Println("Type 'help' to see available commands.")
	}
}
