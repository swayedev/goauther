package config

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/swayedev/oauth/models"
)

var priKey = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIAJnn7DJde0e92CdzqJYTp3J410FZkDX36AJ366SfiHx
-----END PRIVATE KEY-----`
var pubKey = `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEASe1dPbtZH9fCOKFXt3XRO6Y/xPam+/RJLk0JrUoNUDQ=
-----END PUBLIC KEY-----`

func GetCert() (models.Certificate, error) {
	return CertFromString(priKey, pubKey)
}

func CertFromString(privateKey string, publicKey string) (models.Certificate, error) {
	algoType, err := checkKeyType([]byte(publicKey))
	if err != nil {
		return nil, err
	}

	c := &models.CertificateKey{}
	c.SetType(algoType)
	c.SetPrivateKey([]byte(privateKey))
	c.SetPublicKey([]byte(publicKey))
	return c, nil
}

func CertFromFile(privateKeyPath string, publicKeyPath string) (models.Certificate, error) {
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}

	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	algoType, err := checkKeyType(publicKey)
	if err != nil {
		return nil, err
	}

	c := &models.CertificateKey{}
	c.SetType(algoType)
	c.SetPrivateKey(privateKey)
	c.SetPublicKey(publicKey)
	return c, nil
}

func checkKeyType(pemKey []byte) (string, error) {
	block, _ := pem.Decode(pemKey)
	if block == nil {
		return "", errors.New("failed to parse PEM block containing the key")
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	switch key.(type) {
	case *rsa.PublicKey:
		return "RSA", nil
	case ed25519.PublicKey:
		return "ED25519", nil
	default:
		return "", errors.New("unknown key type")
	}
}
