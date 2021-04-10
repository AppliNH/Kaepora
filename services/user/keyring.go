package user

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/json"
	"errors"
	"log"
	"primitivofr/kaepora/kvdb"

	"github.com/boltdb/bolt"
)

var dbKeys *bolt.DB

// UserKeys holds the RSA of a user
type UserKeys struct {
	Username      string   `json:"-"`
	EncPrivateKey [][]byte `json:"encPrivateKey"`
	PublicKey     []byte   `json:"publicKey"`
}

// NewUserKeys builds a UserKeys object. It uses a hash512 of the user's password, to encrypt the private key
func (myUser *User) NewUserKeys() (*UserKeys, error) {

	var marshalledPrivKey []byte
	var marshalledPublicKey []byte

	for {
		privKey, publicKey, err := generateRSA()
		if err != nil {
			return nil, err
		}

		marshalledPrivKey = x509.MarshalPKCS1PrivateKey(privKey)
		marshalledPublicKey = x509.MarshalPKCS1PublicKey(publicKey)

		// chunkSlice (down there) MUST return n slices of equal size, here 16.
		// So if len(marshalledPrivKey) is an number that is not fully dividable by 16, it won't work
		// More specifically, encryptAES won't work

		if (len(marshalledPrivKey) % 16) == 0 {
			break
		}

	}

	chunkedPk := chunkSlice(marshalledPrivKey, 16)

	key := sha512.Sum512_256([]byte(myUser.Password))

	// sha256 := sha256.Sum256([]byte(myUser.Password))

	log.Println(len(marshalledPrivKey))

	//key, err := hex.DecodeString(myUser.Password)
	// if err != nil {
	// 	return nil, err
	// }
	var encChunkedPk [][]byte

	for _, chunk := range chunkedPk {

		encChunk, err := encryptAES(key[:], chunk)
		if err != nil {
			return nil, err
		}

		encChunkedPk = append(encChunkedPk, encChunk)

	}

	return &UserKeys{
		Username:      myUser.Username,
		EncPrivateKey: encChunkedPk,
		PublicKey:     marshalledPublicKey,
	}, nil

}

// GetKeys reads keys from db and return a UserKeys object
func (myUser *User) getKeys() (*UserKeys, error) {
	data, err := kvdb.ReadData(dbKeys, "keys", myUser.Username)
	if err != nil {
		return nil, err
	}
	if data == "" {
		return nil, errors.New("No keys have been found for this user " + myUser.Username)
	}

	var keys UserKeys
	if err := json.Unmarshal([]byte(data), &keys); err != nil {
		return nil, err
	}

	return &keys, nil
}

// GetPublicKey return public key
func (myUser *User) GetPublicKey() (*rsa.PublicKey, error) {
	keys, err := myUser.getKeys()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return x509.ParsePKCS1PublicKey(keys.PublicKey)

}

// GetPrivateKey decrypts the private key using the key which has encrypted it, and returns it
func (myUser *User) GetPrivateKey() (*rsa.PrivateKey, error) {

	keys, err := myUser.getKeys()
	if err != nil {
		return nil, err
	}

	//sha256 := sha256.Sum256([]byte(myUser.Password))
	// key, err := hex.DecodeString(myUser.Password)
	// if err != nil {
	// 	return nil, err
	// }

	key := sha512.Sum512_256([]byte(myUser.Password))

	var decryptedPrivKey []byte

	for _, encChunk := range keys.EncPrivateKey {
		decryptedChunk, err := decryptAES(key[:], encChunk)
		if err != nil {
			return nil, err
		}
		for _, chunk := range decryptedChunk {
			decryptedPrivKey = append(decryptedPrivKey, chunk)
		}
	}

	return x509.ParsePKCS1PrivateKey(decryptedPrivKey)
}

// SaveToDB saves keys to db
func (keys *UserKeys) SaveToDB() error {
	jsonStr, err := json.Marshal(keys)
	if err != nil {
		return err
	}
	return kvdb.WriteData(dbKeys, "keys", keys.Username, string(jsonStr))
}

// generateRSA generates a pair of RSA keys
func generateRSA() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	log.Println(privateKey.Size())
	if err != nil {
		return nil, nil, err
	}
	publicKey := privateKey.PublicKey

	return privateKey, &publicKey, nil
}

func encryptAES(key []byte, plaintext []byte) ([]byte, error) {
	// create cipher
	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return out, nil
}

func decryptAES(key []byte, ciphertext []byte) ([]byte, error) {

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	// s := string(pt[:])
	return pt, nil
}

// Will create chunks of several numbers
func chunkSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}