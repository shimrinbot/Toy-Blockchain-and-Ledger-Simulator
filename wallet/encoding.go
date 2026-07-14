package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"log"
)

// PublicKeyToBytes converts a complex ECDSA public key into a standard byte array.
func PublicKeyToBytes(pubKey *ecdsa.PublicKey) []byte {
	// x509 is an international standard for public key infrastructure.
	bytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		log.Panic("Failed to marshal public key: ", err)
	}
	return bytes
}

// BytesToPublicKey converts a standard byte array back into an ECDSA public key.
func BytesToPublicKey(pubKeyBytes []byte) (*ecdsa.PublicKey, error) {
	// Parse the bytes back into a generic cryptographic key
	genericKey, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key bytes: %v", err)
	}
	
	// "Type assertion": We tell Go to specifically treat this generic key as an ECDSA public key
	pubKey, ok := genericKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("the parsed key is not an ECDSA public key")
	}
	
	return pubKey, nil
}
