package generate

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestGenerateEd25519Keys(t *testing.T) {
	privateKey, publicKey, err := GenerateEd25519Keys()
	if err != nil {
		t.Errorf("GenerateEd25519Keys() returned an error: %v", err)
	}

	// Decode the private key from PEM format
	block, _ := pem.Decode(privateKey)
	if block == nil {
		t.Errorf("private_key: pem.Decode() returned nil")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		t.Errorf("private_key: x509.ParsePKCS8PrivateKey() returned an error: %v", err)
	}

	// Decode the public key from PEM format
	block, _ = pem.Decode(publicKey)
	if block == nil {
		t.Errorf("public_key: pem.Decode() returned nil")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		t.Errorf("public_key: x509.ParsePKIXPublicKey() returned an error: %v", err)
	}

	// Verify that the private key and public key are valid
	_, ok := priKey.(ed25519.PrivateKey)
	if !ok {
		t.Errorf("private_key: x509.ParsePKCS8PrivateKey() returned a private key of type %T, expected ed25519.PrivateKey", priKey)
	}
	_, ok = pubKey.(ed25519.PublicKey)
	if !ok {
		t.Errorf("public_key: x509.ParsePKIXPublicKey() returned a public key of type %T, expected ed25519.PublicKey", pubKey)
	}
}
