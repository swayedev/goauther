package generate

import (
	"testing"
)

func TestGenerateCertificate(t *testing.T) {
	// Test case 1: RSA algorithm
	algorithm := "RSA"
	cert, err := GenerateCertificate(algorithm)
	if err != nil {
		t.Errorf("GenerateCertificate() returned an error: %v", err)
	}

	if cert == nil {
		t.Errorf("GenerateCertificate() returned nil certificate for algorithm %s", algorithm)
	} else {
		if cert.GetType() != algorithm {
			t.Errorf("GenerateCertificate() returned a certificate with type %s, expected %s", cert.GetType(), algorithm)
		}

		// Additional assertions for RSA algorithm

	}

	// Test case 2: ED25519 algorithm
	algorithm = "ED25519"
	cert, err = GenerateCertificate(algorithm)
	if err != nil {
		t.Errorf("GenerateCertificate() returned an error: %v", err)
	}

	if cert == nil {
		t.Errorf("GenerateCertificate() returned nil certificate for algorithm %s", algorithm)
	} else {
		if cert.GetType() != algorithm {
			t.Errorf("GenerateCertificate() returned a certificate with type %s, expected %s", cert.GetType(), algorithm)
		}

		// Additional assertions for ED25519 algorithm

	}

	// Test case 3: Invalid algorithm
	algorithm = "InvalidAlgorithm"
	cert, err = GenerateCertificate(algorithm)
	if err != nil {
		t.Errorf("GenerateCertificate() returned an error: %v", err)
	}

	if cert != nil {
		t.Errorf("GenerateCertificate() returned non-nil certificate for invalid algorithm %s", algorithm)
	}
}
