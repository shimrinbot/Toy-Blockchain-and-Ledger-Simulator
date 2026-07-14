package block

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

var Difficulty = 3

func SetDifficulty(level int) {
	if level > 0 {
		Difficulty = level
	}
}

func MineBlock(block *Block) {
	target := strings.Repeat("0", Difficulty)
	start := time.Now()

	numWorkers := runtime.NumCPU()
	fmt.Printf("Mining with %d concurrent workers...\n", numWorkers)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	resultChan := make(chan *Block)

	nonceStep := 1000000000

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		
		go func(workerID int) {
			defer wg.Done()
			
			startNonce := workerID * nonceStep
			localBlock := *block 

			for nonce := startNonce; ; nonce++ {
				// Check for cancellation every 1000 iterations to save CPU overhead
				if nonce%1000 == 0 {
					select {
					case <-ctx.Done():
						return
					default:
					}
				}

				localBlock.Nonce = nonce
				localBlock.Hash = CalculateHash(localBlock)

				if strings.HasPrefix(localBlock.Hash, target) {
					select {
					case resultChan <- &localBlock:
					case <-ctx.Done():
					}
					return
				}
			}
		}(i)
	}

	minedBlock := <-resultChan
	cancel()

	block.Nonce = minedBlock.Nonce
	block.Hash = minedBlock.Hash

	fmt.Println("--------------------------------")
	fmt.Println("Block mined concurrently!")
	fmt.Println("Nonce:", block.Nonce)
	fmt.Println("Hash :", block.Hash)
	fmt.Println("Time :", time.Since(start))
	fmt.Println("--------------------------------")
	
	wg.Wait()
}
