package block

import (
	"strings"
)
const Difficulty = 3

func MineBlock(block *Block) {

	target := strings.Repeat("0", Difficulty)

	for {

		block.Hash = CalculateHash(*block)

		if strings.HasPrefix(block.Hash, target) {
			break
		}

		block.Nonce++
	}
}