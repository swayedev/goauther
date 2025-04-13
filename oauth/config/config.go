package config

// Encryption Keys

// // CertificateKey struct
// type CertificateKey struct {
// 	privateKey string
// 	publicKey  string
// 	algorithm  string
// }

// // Certificate interface
// type Certificate interface {
// 	SetPrivateKey(privateKey string)
// 	SetPublicKey(publicKey string)
// 	SetType(t string)
// 	GetPrivateKey() string
// 	GetPublicKey() string
// 	GetType() string
// }

// // Set the private key of the certificate
// func (c *CertificateKey) SetPrivateKey(privateKey string) {
// 	c.privateKey = privateKey
// }

// // Set the public key of the certificate
// func (c *CertificateKey) SetPublicKey(publicKey string) {
// 	c.publicKey = publicKey
// }

// // Set the algorithm type of the certificate
// func (c *CertificateKey) SetType(t string) {
// 	c.algorithm = t
// }

// // Get the private key of the certificate
// func (c *CertificateKey) GetPrivateKey() string {
// 	return c.privateKey
// }

// // Get the public key of the certificate
// func (c *CertificateKey) GetPublicKey() string {
// 	return c.publicKey
// }

// // Get the algorithm type of the certificate
// func (c *CertificateKey) GetType() string {
// 	return c.algorithm
// }

// // path of the private key

// /*
//    |--------------------------------------------------------------------------
//    | Personal Access Client
//    |--------------------------------------------------------------------------
//    |
//    | If you enable client hashing, you should set the personal access client
//    | ID and unhashed secret within your environment file. The values will
//    | get used while issuing fresh personal access tokens to your users.
//    |
// */

// // 'personal_access_client' => [
// //     'id' => env('PASSPORT_PERSONAL_ACCESS_CLIENT_ID'),
// //     'secret' => env('PASSPORT_PERSONAL_ACCESS_CLIENT_SECRET'),
// // ],
