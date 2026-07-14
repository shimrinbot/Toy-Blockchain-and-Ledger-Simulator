package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

// GenerateKeyPair creates a new Elliptic Curve Digital Signature Algorithm (ECDSA) private and public key pair.
func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	// 1. Choose the elliptic curve. P-256 is a standard, highly secure curve.
	curve := elliptic.P256()

	// 2. Generate the private key using cryptographically secure random numbers.
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	// 3. Return the generated private key (which intrinsically contains the public key).
	return privateKey, nil
}
