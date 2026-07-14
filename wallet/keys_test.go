package wallet

import (
	"fmt"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	privateKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	fmt.Printf("Private Key: %x\n", privateKey.D)
	fmt.Printf("Public Key (X): %x\n", privateKey.PublicKey.X)
	fmt.Printf("Public Key (Y): %x\n", privateKey.PublicKey.Y)
}
