package main

import (
	"fmt"
	"os"
	"strconv"
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
		case "add":

	if len(os.Args) < 5 {
		fmt.Println("Usage: add sender receiver amount")
		return
	}

	amount, _ := strconv.ParseFloat(os.Args[4],64)

	app.AddTransaction(
		os.Args[2],
		os.Args[3],
		amount,
	)


case "mine":

	app.Mine()

case "help":

	fmt.Println("Available Commands:")
	fmt.Println("  add <sender> <receiver> <amount>")
	fmt.Println("  mine")
	fmt.Println("  print")
	fmt.Println("  validate")
	fmt.Println("  balances")
	fmt.Println("  help")

	default:
	fmt.Println("Unknown command:", command)
	fmt.Println("Type 'help' to see available commands.")
	}
}
