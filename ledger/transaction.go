package ledger

import (
	"crypto/sha256"
	"encoding/json"
	"log"
)

type Transaction struct {
	Sender    string  
	Recipient string
	Amount    float64
	Sequence  int     // NEW: Prevents replay attacks!
	
	// --- NEW SECURITY FIELDS ---
	PublicKey []byte 
	Signature []byte 
}

func (t *Transaction) Hash() []byte {
	// Create an anonymous, temporary struct that mirrors Transaction but omits the Signature.
	data := struct {
		Sender    string
		Recipient string
		Amount    float64
		Sequence  int
		PublicKey []byte
	}{
		Sender:    t.Sender,
		Recipient: t.Recipient,
		Amount:    t.Amount,
		Sequence:  t.Sequence,
		PublicKey: t.PublicKey,
	}

	// Serialize our data to JSON bytes
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Panic("Failed to marshal transaction for hashing: ", err)
	}

	// Generate the SHA-256 hash
	hash := sha256.Sum256(bytes)
	
	// Convert the fixed array [32]byte into a flexible slice []byte
	return hash[:]
}
