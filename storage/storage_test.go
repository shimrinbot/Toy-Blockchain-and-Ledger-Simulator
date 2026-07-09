package storage

import (
	"os"
	"testing"

	"toy-blockchain/blockchain"
)


func TestSaveAndLoadBlockchain(t *testing.T) {

	filename := "test_chain.json"


	bc := blockchain.NewBlockchain()


	err := SaveBlockchain(
		bc,
		filename,
	)

	if err != nil {
		t.Fatal(err)
	}


	loaded, err := LoadBlockchain(filename)

	if err != nil {
		t.Fatal(err)
	}


	if len(loaded.Blocks) != len(bc.Blocks) {
		t.Error("Loaded blockchain does not match")
	}


	os.Remove(filename)
}