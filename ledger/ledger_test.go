package ledger

import "testing"


func TestTransactionSuccess(t *testing.T) {

	l := NewLedger()

	l.Faucet("Alice", 100)


	err := l.ApplyTransaction(Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    50,
	})


	if err != nil {
		t.Error("Expected transaction to succeed")
	}


	if l.Balances["Alice"] != 50 {
		t.Error("Alice balance incorrect")
	}


	if l.Balances["Bob"] != 50 {
		t.Error("Bob balance incorrect")
	}
}



func TestRejectOverspending(t *testing.T) {

	l := NewLedger()

	l.Faucet("Alice", 20)


	err := l.ApplyTransaction(Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    100,
	})


	if err == nil {
		t.Error("Expected overspending to fail")
	}


	if l.Balances["Alice"] != 20 {
		t.Error("Balance changed after failed transaction")
	}
}



func TestRejectNegativeAmount(t *testing.T) {

	l := NewLedger()


	err := l.ApplyTransaction(Transaction{
		Sender:    "Alice",
		Recipient: "Bob",
		Amount:    -10,
	})


	if err == nil {
		t.Error("Expected negative amount to fail")
	}
}