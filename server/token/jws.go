package token

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/go-jose/go-jose/v3"
)

func ExampleRsaJWS() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	serialized, err := SignRsaJWS(privateKey, "Lorem ipsum dolor sit amet")
	if err != nil {
		panic(err)
	}

	output, err := VerifyRsaJWS(&privateKey.PublicKey, serialized)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}

func SignRsaJWS(privateKey *rsa.PrivateKey, serialized string) (string, error) {
	// Instantiate a signer using RSASSA-PSS (SHA512) with the given private key.
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.PS512, Key: privateKey}, nil)
	if err != nil {
		return "", err
	}

	// Sign a sample payload. Calling the signer returns a protected JWS object,
	// which can then be serialized for output afterwards. An error would
	// indicate a problem in an underlying cryptographic primitive.
	var payload = []byte("Lorem ipsum dolor sit amet")
	object, err := signer.Sign(payload)
	if err != nil {
		return "", err
	}

	// Serialize the signed object using the JWS JSON Serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	return object.FullSerialize(), nil
}

func VerifyRsaJWS(publicKey *rsa.PublicKey, serialized string) ([]byte, error) {
	// Parse the serialized, protected JWS object. An error would indicate that
	// the given input did not represent a valid message.
	object, err := jose.ParseSigned(serialized)
	if err != nil {
		return nil, err
	}

	// Now we can verify the signature on the payload. An error here would
	// indicate that the message failed to verify, e.g. because the signature was
	// broken or the message was tampered with.
	output, err := object.Verify(&publicKey)
	if err != nil {
		return nil, err
	}

	return output, nil
}
