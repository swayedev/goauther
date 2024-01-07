package generate

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

// Generate ED25519 Key
func GenerateEd25519Keys() (privateKey []byte, publicKey []byte, err error) {
	// Generate the private and public keys
	pubKey, priKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Encode the private key to PEM format
	privKeyBytes, err := x509.MarshalPKCS8PrivateKey(priKey)
	if err != nil {
		return nil, nil, err
	}
	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privKeyBytes,
	})

	// Encode the public key to PEM format
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, nil, err
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	return privKeyPEM, pubKeyPEM, nil
}
