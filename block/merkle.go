package block

import (
	"crypto/sha256"
	"encoding/hex"
	"toy-blockchain/ledger"
)

// CalculateMerkleRoot computes the Merkle Root hash of a list of transactions.
func CalculateMerkleRoot(transactions []ledger.Transaction) string {
	if len(transactions) == 0 {
		return ""
	}

	// Extract the hash of each transaction as the leaves of the tree
	var hashes []string
	for _, tx := range transactions {
		hashBytes := tx.Hash()
		hashes = append(hashes, hex.EncodeToString(hashBytes))
	}

	// Recursively hash pairs until only one hash remains
	return buildMerkleTree(hashes)
}

func buildMerkleTree(hashes []string) string {
	if len(hashes) == 1 {
		return hashes[0]
	}

	var newLevel []string

	// Process in pairs
	for i := 0; i < len(hashes); i += 2 {
		// If there is an odd number of hashes, duplicate the last one
		var left, right string
		left = hashes[i]
		if i+1 < len(hashes) {
			right = hashes[i+1]
		} else {
			right = left
		}

		combined := left + right
		hash := sha256.Sum256([]byte(combined))
		newLevel = append(newLevel, hex.EncodeToString(hash[:]))
	}

	return buildMerkleTree(newLevel)
}
