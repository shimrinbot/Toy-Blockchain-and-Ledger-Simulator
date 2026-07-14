package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func CalculateHash(block Block) string {

	record := fmt.Sprintf(
		"%d%d%s%s%d",
		block.Index,
		block.Timestamp,
		block.MerkleRoot,
		block.PreviousHash,
		block.Nonce,
	)

	hash := sha256.Sum256([]byte(record))

	return hex.EncodeToString(hash[:])
}
