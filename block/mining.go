package block

import (
	"fmt"
	"strings"
	"time"
)

var Difficulty = 3

func MineBlock(block *Block) {

	target := strings.Repeat("0", Difficulty)

	start := time.Now()

	for {
		block.Hash = CalculateHash(*block)

		if strings.HasPrefix(block.Hash, target) {

			fmt.Println("--------------------------------")
			fmt.Println("Block mined!")
			fmt.Println("Nonce:", block.Nonce)
			fmt.Println("Hash :", block.Hash)
			fmt.Println("Time :", time.Since(start))
			fmt.Println("--------------------------------")

			break
		}

		block.Nonce++
	}
}
