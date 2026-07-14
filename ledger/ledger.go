package ledger

import (
	"encoding/hex"
	"errors"
	"fmt"
	"toy-blockchain/wallet" 
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
		// --- NEW SECURITY ENFORCEMENT ---
		
		// 1. Ensure the transaction actually has a public key attached
		if len(tx.PublicKey) == 0 {
			return errors.New("transaction rejected: missing public key")
		}

		// 2. Decode the raw bytes back into a cryptographic key object using the tool we built in Lesson 4
		pubKey, err := wallet.BytesToPublicKey(tx.PublicKey)
		if err != nil {
			return fmt.Errorf("transaction rejected: invalid public key format: %v", err)
		}

		// 3. Verify the signature against the transaction's hash!
		isValid := wallet.Verify(pubKey, tx.Hash(), tx.Signature)
		if !isValid {
			return errors.New("transaction rejected: invalid cryptographic signature")
		}

		// 4. Ensure the Sender string exactly matches the hex representation of the public key.
		// CRITICAL: If we didn't do this, Eve could sign a transaction with HER private key, 
		// but maliciously put "Alice" in the Sender field to steal Alice's funds!
		address := hex.EncodeToString(tx.PublicKey)
		if tx.Sender != address {
			return errors.New("transaction rejected: sender address does not match public key")
		}
		// --- END SECURITY ENFORCEMENT ---

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
