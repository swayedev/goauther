package config

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"

	"github.com/swayedev/goauther/models"
)

func GetCert() (models.Certificate, error) {
	priKey := `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIAJnn7DJde0e92CdzqJYTp3J410FZkDX36AJ366SfiHx
-----END PRIVATE KEY-----`
	pubKey := `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEASe1dPbtZH9fCOKFXt3XRO6Y/xPam+/RJLk0JrUoNUDQ=
-----END PUBLIC KEY-----`
	return CertFromString(priKey, pubKey)
}

// CertFromString parses the given private and public PEM-encoded keys and returns a Certificate object.
// The private key and public key are provided as strings.
// It checks the key type and sets the appropriate algorithm type for the certificate.
// It returns the Certificate object and any error encountered during the process.
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

// CertFromFile reads the private and public PEM-encoded keys from the specified file paths
// and returns a models.Certificate object containing the keys.
// It checks the key type of the public key and sets it in the certificate.
// If any error occurs during reading the files or checking the key type, it returns the error.
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

// checkKeyType checks the type of a PEM-encoded key.
// It takes a byte slice containing the PEM-encoded key as input.
// It returns the key type as a string and an error if the key type is unknown or if there is an error parsing the key.
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
		return "rsa", nil
	case ed25519.PublicKey:
		return "ed25519", nil
	default:
		return "", errors.New("unknown key type")
	}
}
