package main

import (
	"github.com/swayedev/oauth/config"
)

func main() {
	cert, err := config.GetCert()
	if err != nil {
		panic(err)
	}

	println("Certificate Type:", cert.GetType())
	println("Private Key:", cert.GetPrivateKeyString())
	println("Public Key:", cert.GetPublicKeyString())
}
