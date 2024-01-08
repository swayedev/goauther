package config_test

import (
	"testing"

	"github.com/swayedev/goauther/config"
	"github.com/swayedev/goauther/models"
)

func TestCertFromString(t *testing.T) {
	privateKey := `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIAJnn7DJde0e92CdzqJYTp3J410FZkDX36AJ366SfiHx
-----END PRIVATE KEY-----`
	publicKey := `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEASe1dPbtZH9fCOKFXt3XRO6Y/xPam+/RJLk0JrUoNUDQ=
-----END PUBLIC KEY-----`

	expectedCert := &models.CertificateKey{}
	expectedCert.SetType("ED25519")
	expectedCert.SetPrivateKey([]byte(privateKey))
	expectedCert.SetPublicKey([]byte(publicKey))

	cert, err := config.CertFromString(privateKey, publicKey)
	if err != nil {
		t.Errorf("CertFromString() returned an error: %v", err)
	}

	if cert.GetType() != expectedCert.GetType() {
		t.Errorf("CertFromString() returned a certificate with type %s, expected %s", cert.GetType(), expectedCert.GetType())
	}

	if string(cert.GetPrivateKey()) != string(expectedCert.GetPrivateKey()) {
		t.Errorf("CertFromString() returned a certificate with private key %s, expected %s", cert.GetPrivateKey(), expectedCert.GetPrivateKey())
	}

	if string(cert.GetPublicKey()) != string(expectedCert.GetPublicKey()) {
		t.Errorf("CertFromString() returned a certificate with public key %s, expected %s", cert.GetPublicKey(), expectedCert.GetPublicKey())
	}
}

func TestCertFromFile(t *testing.T) {
	privateKeyPath := "../cert/ed25519"
	publicKeyPath := "../cert/ed25519.pub"

	expectedPrivateKey := []byte(`-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIAJnn7DJde0e92CdzqJYTp3J410FZkDX36AJ366SfiHx
-----END PRIVATE KEY-----`)
	expectedPublicKey := []byte(`-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEASe1dPbtZH9fCOKFXt3XRO6Y/xPam+/RJLk0JrUoNUDQ=
-----END PUBLIC KEY-----`)

	expectedCert := &models.CertificateKey{}
	expectedCert.SetType("ED25519")
	expectedCert.SetPrivateKey(expectedPrivateKey)
	expectedCert.SetPublicKey(expectedPublicKey)

	cert, err := config.CertFromFile(privateKeyPath, publicKeyPath)
	if err != nil {
		t.Errorf("CertFromFile() returned an error: %v", err)
	}

	if cert.GetType() != expectedCert.GetType() {
		t.Errorf("CertFromFile() returned a certificate with type %s, expected %s", cert.GetType(), expectedCert.GetType())
	}

	if string(cert.GetPrivateKey()) != string(expectedCert.GetPrivateKey()) {
		t.Errorf("CertFromFile() returned a certificate with private key %s, expected %s", cert.GetPrivateKey(), expectedCert.GetPrivateKey())
	}

	if string(cert.GetPublicKey()) != string(expectedCert.GetPublicKey()) {
		t.Errorf("CertFromFile() returned a certificate with public key %s, expected %s", cert.GetPublicKey(), expectedCert.GetPublicKey())
	}
}
