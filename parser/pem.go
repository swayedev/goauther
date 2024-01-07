package parser

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func ParsePemPrivateKey(pemPrivateKey []byte) (any, error) {
	block, _ := pem.Decode(pemPrivateKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

func ParsePemPublicKey(pemPublicKey []byte) (any, error) {
	// Decode the public key from PEM format
	block, _ := pem.Decode(pemPublicKey)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}
