package generate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// Generate RSA Key
func GenerateRsaKeys() (privateKey []byte, publicKey []byte, err error) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Private Key to PEM
	privKeyBytes := x509.MarshalPKCS1PrivateKey(rsaKey)
	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	})

	// Public Key to PEM
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&rsaKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	return privKeyPEM, pubKeyPEM, nil
}
