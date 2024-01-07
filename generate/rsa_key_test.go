package generate

import (
	"testing"
)

func TestGenerateRsaKeys(t *testing.T) {
	privateKey, publicKey, err := GenerateRsaKeys()
	if err != nil {
		t.Errorf("GenerateRsaKeys() returned an error: %v", err)
	}

	// TODO: Add assertions to validate the generated keys
	// For example, you can check the length of the private and public keys,
	// or decode the keys and verify their format.

	// Example assertion:
	if len(privateKey) == 0 {
		t.Errorf("GenerateRsaKeys() returned an empty private key")
	}

	if len(publicKey) == 0 {
		t.Errorf("GenerateRsaKeys() returned an empty public key")
	}
}
