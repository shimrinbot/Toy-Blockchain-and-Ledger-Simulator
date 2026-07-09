package block

import "testing"


func TestHashDeterminism(t *testing.T) {

	block := Block{
		Index: 1,
		Timestamp: 12345,
		Transactions: nil,
		PreviousHash: "abc",
		Nonce: 0,
	}


	hash1 := CalculateHash(block)
	hash2 := CalculateHash(block)


	if hash1 != hash2 {
		t.Error("Hash is not deterministic")
	}
}
func TestMiningDifficulty(t *testing.T) {

	b := Block{
		Index: 1,
		Timestamp: 123,
	}

	MineBlock(&b)


	if len(b.Hash) < Difficulty {
		t.Error("Invalid hash length")
	}


	for i := 0; i < Difficulty; i++ {

		if b.Hash[i] != '0' {
			t.Error("Mining did not satisfy difficulty")
		}
	}
}