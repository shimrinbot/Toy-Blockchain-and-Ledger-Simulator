package ledger

import (
	"crypto/ecdsa"
	"encoding/hex"
	"testing"
	"toy-blockchain/wallet"
)

// createSecureTx is a new helper function just for our tests.
// It automatically generates a mathematically valid transaction so we don't have to copy-paste this logic.
func createSecureTx(senderKey *ecdsa.PrivateKey, recipient string, amount float64, sequence int) Transaction {
	pubKeyBytes := wallet.PublicKeyToBytes(&senderKey.PublicKey)
	senderAddress := hex.EncodeToString(pubKeyBytes)

	tx := Transaction{
		Sender:    senderAddress,
		Recipient: recipient,
		Amount:    amount,
		Sequence:  sequence,
		PublicKey: pubKeyBytes,
	}

	// Sign the hash of the transaction
	sig, _ := wallet.Sign(senderKey, tx.Hash())
	tx.Signature = sig

	return tx
}

func TestTransactionSuccess(t *testing.T) {
	l := NewLedger()
	
	// 1. Generate Alice's cryptographic identity
	aliceKey, _ := wallet.GenerateKeyPair()
	aliceAddress := hex.EncodeToString(wallet.PublicKeyToBytes(&aliceKey.PublicKey))

	// 2. Fund Alice's new hexadecimal address (not the string "Alice")
	l.Faucet(aliceAddress, 100)

	// 3. Create and apply a secure transaction
	tx := createSecureTx(aliceKey, "Bob", 50, 1)
	err := l.ApplyTransaction(tx)
	
	if err != nil {
		t.Errorf("Expected transaction to succeed, but it failed: %v", err)
	}
	
	// 4. Verify balances
	if l.Balances[aliceAddress] != 50 {
		t.Error("Alice balance incorrect")
	}
}

func TestRejectOverspending(t *testing.T) {
	l := NewLedger()
	
	aliceKey, _ := wallet.GenerateKeyPair()
	aliceAddress := hex.EncodeToString(wallet.PublicKeyToBytes(&aliceKey.PublicKey))

	l.Faucet(aliceAddress, 20)

	tx := createSecureTx(aliceKey, "Bob", 100, 1)
	err := l.ApplyTransaction(tx)
	
	// NEW: We now strictly check that it failed for the RIGHT reason!
	if err == nil || err.Error() != "insufficient balance" {
		t.Errorf("Expected 'insufficient balance' error, but got: %v", err)
	}
}

func TestRejectNegativeAmount(t *testing.T) {
	l := NewLedger()
	
	aliceKey, _ := wallet.GenerateKeyPair()
	
	tx := createSecureTx(aliceKey, "Bob", -10, 1)
	err := l.ApplyTransaction(tx)
	
	// NEW: Strict error checking
	if err == nil || err.Error() != "amount must be greater than zero" {
		t.Errorf("Expected 'amount must be greater than zero' error, but got: %v", err)
	}
}
