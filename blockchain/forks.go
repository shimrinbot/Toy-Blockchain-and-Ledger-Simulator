package blockchain

import "fmt"

// ResolveConflict implements the "Longest Valid Chain" rule.
// It compares the current chain with a competing chain. If the competing chain is longer
// and mathematically valid, it replaces our current chain.
func (bc *Blockchain) ResolveConflict(competingChain *Blockchain) bool {
	// 1. Check if it's strictly longer
	if len(competingChain.Blocks) <= len(bc.Blocks) {
		fmt.Println("Competing chain is not longer than our chain. Rejecting.")
		return false
	}

	// 2. Validate the competing chain's integrity
	if err := competingChain.Validate(); err != nil {
		fmt.Printf("Competing chain is longer but INVALID (%v). Rejecting.\n", err)
		return false
	}

	// 3. Adopt the competing chain
	fmt.Println("Competing chain is valid and longer! Resolving fork by replacing our chain.")
	bc.Blocks = competingChain.Blocks
	
	// Note: In a real blockchain, we would also need to return the unmined transactions 
	// from our discarded blocks back to the mempool (PendingTxs).
	// For this toy, we simply accept the new chain.

	return true
}
