package cli

import "testing"

func TestNewCLI(t *testing.T) {

	c := NewCLI()

	if c == nil {
		t.Fatal("CLI should not be nil")
	}

	if c.Blockchain == nil {
		t.Fatal("Blockchain should not be nil")
	}

	if c.Ledger == nil {
		t.Fatal("Ledger should not be nil")
	}
}