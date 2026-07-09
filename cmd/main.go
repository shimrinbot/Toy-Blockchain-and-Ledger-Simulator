package main

import (
	"fmt"
	"os"

	"toy-blockchain/cli"
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

	switch command {

	case "print":
		app.PrintBlockchain()

	case "validate":
		app.ValidateBlockchain()

	case "balances":
		app.PrintBalances()

	default:
		fmt.Println("Unknown command:", command)
	}
}
