package ledger

import (
	"errors"
	"fmt"
)

type Ledger struct {
	Balances map[string]float64
}

func NewLedger() *Ledger {
	return &Ledger{
		Balances: make(map[string]float64),
	}
}

func (l *Ledger) Faucet(account string, amount float64) {
	l.Balances[account] += amount
}

func (l *Ledger) ApplyTransaction(tx Transaction) error {

	if tx.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if tx.Sender != "SYSTEM" {
		if l.Balances[tx.Sender] < tx.Amount {
			return errors.New("insufficient balance")
		}
		l.Balances[tx.Sender] -= tx.Amount
	}

	l.Balances[tx.Recipient] += tx.Amount

	return nil
}

func (l *Ledger) PrintBalances() {

	fmt.Println("\nCurrent Balances")

	for account, balance := range l.Balances {
		fmt.Printf("%s : %.2f\n", account, balance)
	}
}
