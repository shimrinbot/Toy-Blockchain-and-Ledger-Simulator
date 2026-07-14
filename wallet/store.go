package wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io"
	"os"
)

type WalletData struct {
	PrivateKeyBytes []byte
}

// deriveKey creates a 32-byte AES key from a password using SHA-256
func deriveKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

func SaveWallet(privKey *ecdsa.PrivateKey, filename string, password string) error {
	bytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil {
		return err
	}
	
	data := WalletData{PrivateKeyBytes: bytes}
	fileBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	
	key := deriveKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, fileBytes, nil)
	return os.WriteFile(filename, ciphertext, 0600)
}

func LoadWallet(filename string, password string) (*ecdsa.PrivateKey, error) {
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	key := deriveKey(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	nonce, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("invalid password or corrupted wallet")
	}

	var data WalletData
	if err := json.Unmarshal(plaintext, &data); err != nil {
		return nil, err
	}
	
	return x509.ParseECPrivateKey(data.PrivateKeyBytes)
}
