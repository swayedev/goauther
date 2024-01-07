package generate

import "github.com/swayedev/oauth/models"

func GenerateCertificate(algorithm string) (cert models.Certificate, err error) {
	switch algorithm {
	case "RSA":
		// Generate RSA key
		privateKey, publicKey, err := GenerateRsaKeys()
		if err != nil {
			return nil, err
		}

		// Create certificate
		cert = &models.CertificateKey{}
		cert.SetPrivateKey(privateKey)
		cert.SetPublicKey(publicKey)
		cert.SetType(algorithm)

		return cert, nil
	case "ED25519":
		privateKey, publicKey, err := GenerateEd25519Keys()
		if err != nil {
			return nil, err
		}

		// Create certificate
		cert = &models.CertificateKey{}
		cert.SetPrivateKey(privateKey)
		cert.SetPublicKey(publicKey)
		cert.SetType(algorithm)

		return cert, nil
	default:
		return nil, nil
	}
}
