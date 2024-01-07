package models

import "github.com/swayedev/oauth/parser"

type Certificate interface {
	SetPrivateKey(privateKey []byte)
	GetPrivateKey() []byte
	GetPrivateKeyString() string
	ParsePrivateKey() (any, error)
	SetPublicKey(publicKey []byte)
	GetPublicKey() []byte
	GetPublicKeyString() string
	ParsePublicKey() (any, error)
	SetType(t string)
	GetType() string
}

type CertificateKey struct {
	privateKey []byte
	publicKey  []byte
	algorithm  string
}

func (c *CertificateKey) SetPrivateKey(privateKey []byte) {
	c.privateKey = privateKey
}
func (c *CertificateKey) GetPrivateKey() []byte {
	return c.privateKey
}
func (c *CertificateKey) GetPublicKeyString() string {
	return string(c.publicKey)
}

func (c *CertificateKey) SetPublicKey(publicKey []byte) {
	c.publicKey = publicKey
}
func (c *CertificateKey) GetPublicKey() []byte {
	return c.publicKey
}
func (c *CertificateKey) GetPrivateKeyString() string {
	return string(c.privateKey)
}

func (c *CertificateKey) SetType(t string) {
	c.algorithm = t
}
func (c *CertificateKey) GetType() string {
	return c.algorithm
}

func (c *CertificateKey) ParsePrivateKey() (any, error) {
	return parser.ParsePemPrivateKey(c.privateKey)
}

func (c *CertificateKey) ParsePublicKey() (any, error) {
	return parser.ParsePemPublicKey(c.publicKey)
}
