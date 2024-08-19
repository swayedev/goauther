package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/swayedev/goauther/mtls"
)

func main() {
	// Server configuration
	serverConfig := &mtls.Config{
		CertFile: "path/to/server.crt",
		KeyFile:  "path/to/server.key",
		CAFile:   "path/to/ca.crt",
	}

	// Create a new mTLS server
	server, err := mtls.NewServer(":8443", serverConfig, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, mTLS!")
	}))
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	log.Println("Starting mTLS server on https://localhost:8443")
	if err := server.ListenAndServeTLS(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
