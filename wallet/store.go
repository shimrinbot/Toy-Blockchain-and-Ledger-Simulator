package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"os"
)

type WalletData struct {
	PrivateKeyBytes []byte
}

func SaveWallet(privKey *ecdsa.PrivateKey, filename string) error {
	bytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return err
	}
	
	data := WalletData{PrivateKeyBytes: bytes}
	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(filename, fileBytes, 0600)
}

func LoadWallet(filename string) (*ecdsa.PrivateKey, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	var data WalletData
	if err := json.Unmarshal(fileBytes, &data); err != nil {
		return nil, err
	}
	
	return x509.ParseECPrivateKey(data.PrivateKeyBytes)
}
