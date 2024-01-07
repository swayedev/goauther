package generate

import (
	"testing"
)

func TestGenerateRsaPemKeys(t *testing.T) {
	privateKey, publicKey, err := GenerateRsaPemKeys()
	if err != nil {
		t.Errorf("GenerateRsaPemKeys() returned an error: %v", err)
	}

	// TODO: Add assertions to validate the generated keys
	// For example, you can check the length of the private and public keys,
	// or decode the keys and verify their format.

	// Example assertion:
	if len(privateKey) == 0 {
		t.Errorf("GenerateRsaPemKeys() returned an empty private key")
	}

	if len(publicKey) == 0 {
		t.Errorf("GenerateRsaPemKeys() returned an empty public key")
	}
}
