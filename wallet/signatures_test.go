package wallet

import (
	"crypto/sha256"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	// 1. Generate a new key pair for our test user (Alice)
	privateKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	// 2. Create some dummy data and hash it. In our real blockchain, this will be a Transaction.
	dummyData := "Send 5 coins to Bob"
	dataHash := sha256.Sum256([]byte(dummyData))

	// 3. Alice signs the hash with her Private Key
	// Note: sha256.Sum256 returns an array [32]byte, but our function expects a slice []byte. 
	// The [:] syntax converts the fixed array into a slice.
	signature, err := Sign(privateKey, dataHash[:])
	if err != nil {
		t.Fatalf("Failed to sign data: %v", err)
	}

	// 4. Verify the signature using Alice's Public Key. This should SUCCEED.
	isValid := Verify(&privateKey.PublicKey, dataHash[:], signature)
	if !isValid {
		t.Errorf("Signature verification failed! It should be valid.")
	}

	// 5. Try to verify with fake/tampered data to ensure our security works.
	tamperedData := "Send 500 coins to Bob"
	tamperedHash := sha256.Sum256([]byte(tamperedData))
	
	isTamperedValid := Verify(&privateKey.PublicKey, tamperedHash[:], signature)
	if isTamperedValid {
		// If this succeeds, our blockchain is broken!
		t.Errorf("Security Flaw: Signature verification succeeded on tampered data!")
	}
}
