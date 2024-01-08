package token

import (
	"crypto/rand"
	"crypto/rsa"
	"log"

	"github.com/go-jose/go-jose/v3"
)

func ExampleRsaJWE() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed to generate key: %v", err)
	}
	plaintext := []byte("Lorem ipsum dolor sit amet")
	serialized, err := EncryptRsaJWE(&privateKey.PublicKey, plaintext)
	if err != nil {
		log.Fatalf("failed to encrypt: %v", err)
	}

	decrypted, err := DecryptRsaJWE(privateKey, serialized)
	if err != nil {
		panic(err)
	}

	if string(decrypted) != string(plaintext) {
		panic("decrypted != plaintext")
	}
}

func EncryptRsaJWE(publicKey *rsa.PublicKey, data []byte) (string, error) {
	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	encrypter, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: publicKey}, nil)
	if err != nil {
		return "", err
	}

	// Encrypt a sample plaintext. Calling the encrypter returns an encrypted
	// JWE object, which can then be serialized for output afterwards. An error
	// would indicate a problem in an underlying cryptographic primitive.
	object, err := encrypter.Encrypt(data)
	if err != nil {
		return "", err
	}

	// Serialize the encrypted object using the JWE JSON Serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	return object.FullSerialize(), nil
}

func DecryptRsaJWE(privateKey *rsa.PrivateKey, serialized string) ([]byte, error) {
	// Parse the serialized, encrypted JWE object. An error would indicate that
	// the given input did not represent a valid message.
	object, err := jose.ParseEncrypted(serialized)
	if err != nil {
		return nil, err
	}

	// Now we can decrypt and get back our original plaintext. An error here
	// would indicate that the message failed to decrypt, e.g. because the auth
	// tag was broken or the message was tampered with.
	decrypted, err := object.Decrypt(privateKey)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
