/*
An example of adding OpenID Connect support to oauth.
*/
package main

// var (
// 	issuer = "http://127.0.0.1:14001"
// 	server = oauth.NewServer(oauth.NewServerConfig(), example.NewTestStorage())

// 	jwtSigner  jose.Signer
// 	publicKeys *jose.JsonWebKeySet
// )

// func main() {
// 	server.Certificate, _ = config.GetCert()
// 	// Load signing key.
// 	block, _ := pem.Decode(server.Certificate.GetPrivateKey())
// 	if block == nil {
// 		log.Fatalf("no private key found")
// 	}
// 	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
// 	if err != nil {
// 		log.Fatalf("failed to parse key: %v", err)
// 	}

// 	// Configure jwtSigner and public keys.
// 	privateKey := &jose.JsonWebKey{
// 		Key:       key,
// 		Algorithm: "RS256",
// 		Use:       "sig",
// 		KeyID:     "1", // KeyID should use the key thumbprint.
// 	}

// 	jwtSigner, err = jose.NewSigner(jose.RS256, privateKey)
// 	if err != nil {
// 		log.Fatalf("failed to create jwtSigner: %v", err)
// 	}
// 	publicKeys = &jose.JsonWebKeySet{
// 		Keys: []jose.JsonWebKey{
// 			jose.JsonWebKey{Key: &key.PublicKey,
// 				Algorithm: "RS256",
// 				Use:       "sig",
// 				KeyID:     "1",
// 			},
// 		},
// 	}

// 	// Register the four manditory OpenID Connect endpoints: discovery, public keys, auth, and token.
// 	http.HandleFunc("/.well-known/openid-configuration", handleDiscovery)
// 	http.HandleFunc("/publickeys", handlePublicKeys)
// 	http.HandleFunc("/authorize", handleAuthorization)
// 	http.HandleFunc("/token", handleToken)

// 	log.Fatal(http.ListenAndServe("127.0.0.1:14001", nil))
// }

// // The ID Token represents a JWT passed to the client as part of the token response.
// //
// // https://openid.net/specs/openid-connect-core-1_0.html#IDToken
// type IDToken struct {
// 	Issuer     string `json:"iss"`
// 	UserID     string `json:"sub"`
// 	ClientID   string `json:"aud"`
// 	Expiration int64  `json:"exp"`
// 	IssuedAt   int64  `json:"iat"`

// 	Nonce string `json:"nonce,omitempty"` // Non-manditory fields MUST be "omitempty"

// 	// Custom claims supported by this server.
// 	//
// 	// See: https://openid.net/specs/openid-connect-core-1_0.html#StandardClaims

// 	Email         string `json:"email,omitempty"`
// 	EmailVerified *bool  `json:"email_verified,omitempty"`

// 	Name       string `json:"name,omitempty"`
// 	FamilyName string `json:"family_name,omitempty"`
// 	GivenName  string `json:"given_name,omitempty"`
// 	Locale     string `json:"locale,omitempty"`
// }

// // handleDiscovery returns the OpenID Connect discovery object, allowing clients
// // to discover OAuth2 resources.
// func handleDiscovery(w http.ResponseWriter, r *http.Request) {
// 	// For other example see: https://accounts.google.com/.well-known/openid-configuration
// 	data := map[string]interface{}{
// 		"issuer":                                issuer,
// 		"authorization_endpoint":                issuer + "/authorize",
// 		"token_endpoint":                        issuer + "/token",
// 		"jwks_uri":                              issuer + "/publickeys",
// 		"response_types_supported":              []string{"code"},
// 		"subject_types_supported":               []string{"public"},
// 		"id_token_signing_alg_values_supported": []string{"RS256"},
// 		"scopes_supported":                      []string{"openid", "email", "profile"},
// 		"token_endpoint_auth_methods_supported": []string{"client_secret_basic"},
// 		"claims_supported": []string{
// 			"aud", "email", "email_verified", "exp",
// 			"family_name", "given_name", "iat", "iss",
// 			"locale", "name", "sub",
// 		},
// 	}

// 	raw, err := json.MarshalIndent(data, "", "  ")
// 	if err != nil {
// 		log.Printf("failed to marshal data: %v", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Content-Length", strconv.Itoa(len(raw)))
// 	w.Write(raw)
// }

// // handlePublicKeys publishes the public part of this server's signing keys.
// // This allows clients to verify the signature of ID Tokens.
// func handlePublicKeys(w http.ResponseWriter, r *http.Request) {
// 	raw, err := json.MarshalIndent(publicKeys, "", "  ")
// 	if err != nil {
// 		log.Printf("failed to marshal data: %v", err)
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Content-Length", strconv.Itoa(len(raw)))
// 	w.Write(raw)
// }

// func handleAuthorization(w http.ResponseWriter, r *http.Request) {
// 	resp := server.NewResponse()
// 	defer resp.Close()

// 	if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
// 		if !example.HandleLoginPage(ar, w, r) {
// 			return
// 		}

// 		ar.Authorized = true
// 		scopes := make(map[string]bool)
// 		for _, s := range strings.Fields(ar.Scope) {
// 			scopes[s] = true
// 		}

// 		// If the "openid" connect scope is specified, attach an ID Token to the
// 		// authorization response.
// 		//
// 		// The ID Token will be serialized and signed during the code for token exchange.
// 		if scopes["openid"] {

// 			// These values would be tied to the end user authorizing the client.
// 			now := time.Now()
// 			idToken := IDToken{
// 				Issuer:     issuer,
// 				UserID:     "id-of-test-user",
// 				ClientID:   ar.Client.GetId(),
// 				Expiration: now.Add(time.Hour).Unix(),
// 				IssuedAt:   now.Unix(),
// 				Nonce:      r.URL.Query().Get("nonce"),
// 			}

// 			if scopes["profile"] {
// 				idToken.Name = "Jane Doe"
// 				idToken.GivenName = "Jane"
// 				idToken.FamilyName = "Doe"
// 				idToken.Locale = "us"
// 			}

// 			if scopes["email"] {
// 				t := true
// 				idToken.Email = "jane.doe@example.com"
// 				idToken.EmailVerified = &t
// 			}
// 			// NOTE: The storage must be able to encode and decode this object.
// 			ar.UserData = &idToken
// 		}
// 		server.FinishAuthorizeRequest(resp, r, ar)
// 	}

// 	if resp.IsError && resp.InternalError != nil {
// 		log.Printf("internal error: %v", resp.InternalError)
// 	}
// 	example.OutputJSON(resp, w, r)
// }

// func handleToken(w http.ResponseWriter, r *http.Request) {
// 	resp := server.NewResponse()
// 	defer resp.Close()

// 	if ar := server.HandleAccessRequest(resp, r); ar != nil {
// 		ar.Authorized = true
// 		server.FinishAccessRequest(resp, r, ar)

// 		// If an ID Token was encoded as the UserData, serialize and sign it.
// 		if idToken, ok := ar.UserData.(*IDToken); ok && idToken != nil {
// 			encodeIDToken(resp, idToken, jwtSigner)
// 		}
// 	}
// 	if resp.IsError && resp.InternalError != nil {
// 		fmt.Printf("ERROR: %s\n", resp.InternalError)
// 	}
// 	example.OutputJSON(resp, w, r)
// }

// // encodeIDToken serializes and signs an ID Token then adds a field to the token response.
// func encodeIDToken(resp *oauth.Response, idToken *IDToken, singer jose.Signer) {
// 	resp.InternalError = func() error {
// 		payload, err := json.Marshal(idToken)
// 		if err != nil {
// 			return fmt.Errorf("failed to marshal token: %v", err)
// 		}
// 		jws, err := jwtSigner.Sign(payload)
// 		if err != nil {
// 			return fmt.Errorf("failed to sign token: %v", err)
// 		}
// 		raw, err := jws.CompactSerialize()
// 		if err != nil {
// 			return fmt.Errorf("failed to serialize token: %v", err)
// 		}
// 		resp.Output["id_token"] = raw
// 		return nil
// 	}()

// 	// Record errors as internal server errors.
// 	if resp.InternalError != nil {
// 		resp.IsError = true
// 		resp.ErrorId = oauth.E_SERVER_ERROR
// 	}
// }
