package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
)

func Sign(privateKey *ecdsa.PrivateKey, dataHash []byte) ([]byte, error) {

	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, dataHash)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

func Verify(publicKey *ecdsa.PublicKey, dataHash []byte, signature []byte) bool {

	return ecdsa.VerifyASN1(publicKey, dataHash, signature)
}
