package models

// UserKeys holds the RSA of a user
type UserKeys struct {
	Username      string   `json:"username"`
	EncPrivateKey [][]byte `json:"encPrivateKey"`
	PublicKey     []byte   `json:"publicKey"`
}

// NewUserKeys builds a UserKeys object. It uses a hash512 of the user's password, to encrypt the private key
func NewUserKeys(username string, encPrivateKey [][]byte, publicKey []byte) *UserKeys {

	return &UserKeys{
		Username:      username,
		EncPrivateKey: encPrivateKey,
		PublicKey:     publicKey,
	}

}
