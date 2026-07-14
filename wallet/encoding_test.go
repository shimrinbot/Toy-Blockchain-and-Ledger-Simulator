package wallet

import (
	"testing"
)

func TestSerialization(t *testing.T) {
	// 1. Generate a fresh key pair
	privKey, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}
	
	// 2. Serialize the Public Key to raw bytes
	pubKeyBytes := PublicKeyToBytes(&privKey.PublicKey)
	
	// 3. Deserialize the bytes back into a complex object
	parsedPubKey, err := BytesToPublicKey(pubKeyBytes)
	if err != nil {
		t.Fatalf("Failed to parse public key: %v", err)
	}
	
	// 4. Compare the mathematical X and Y coordinates to ensure no data was lost
	if privKey.PublicKey.X.Cmp(parsedPubKey.X) != 0 || privKey.PublicKey.Y.Cmp(parsedPubKey.Y) != 0 {
		t.Errorf("CRITICAL FAILURE: Parsed public key does not match original!")
	}
}
