package blockchain

import (
	"errors"
	"fmt"
	"strings"

	"toy-blockchain/block"
)
func (bc *Blockchain) Validate() error {

	if len(bc.Blocks) == 0 {
		return errors.New("blockchain is empty")
	}

	target := strings.Repeat("0", block.Difficulty)

	for i := 0; i < len(bc.Blocks); i++ {

		current := bc.Blocks[i]

		// Check if the stored hash is correct
		calculatedHash := block.CalculateHash(current)

		if current.Hash != calculatedHash {
			return fmt.Errorf("invalid hash at block %d", i)
		}

		// Skip previous hash check for Genesis Block
		if i == 0 {
			continue
		}

		previous := bc.Blocks[i-1]

		// Check blockchain links
		if current.PreviousHash != previous.Hash {
			return fmt.Errorf("broken previous hash link at block %d", i)
		}

		// Check Proof of Work
		if !strings.HasPrefix(current.Hash, target) {
			return fmt.Errorf("proof of work failed at block %d", i)
		}
	}

	return nil
}