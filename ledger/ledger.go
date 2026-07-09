
package ledger
import "fmt"

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

func (l *Ledger) ApplyTransaction(tx Transaction) {

	l.Balances[tx.Sender] -= tx.Amount

	l.Balances[tx.Recipient] += tx.Amount
}
func (l *Ledger) PrintBalances() {

	fmt.Println("\nCurrent Balances")

	for account, balance := range l.Balances {
		fmt.Printf("%s : %.2f\n", account, balance)
	}
}