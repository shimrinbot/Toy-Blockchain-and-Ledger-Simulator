package storage

import (
	"encoding/json"
	"os"

	"toy-blockchain/blockchain"
)


func SaveBlockchain(
	bc *blockchain.Blockchain,
	filename string,
) error {

	data, err := json.MarshalIndent(
		bc,
		"",
		"  ",
	)

	if err != nil {
		return err
	}


	return os.WriteFile(
		filename,
		data,
		0644,
	)
}

func LoadBlockchain(
	filename string,
) (*blockchain.Blockchain,error){

	data,err := os.ReadFile(filename)

	if err != nil {
		return nil,err
	}


	var bc blockchain.Blockchain


	err=json.Unmarshal(
		data,
		&bc,
	)

	if err != nil {
		return nil,err
	}


	return &bc,nil
}