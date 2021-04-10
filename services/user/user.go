package user

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"

	"primitivofr/kaepora/kvdb"

	"github.com/boltdb/bolt"
)

var dbUsers *bolt.DB

func init() {
	var err error
	dbUsers, err = kvdb.InitDB("kaepora-users-db", "users")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	dbKeys, err = kvdb.InitDB("kaepora-keys-db", "keys")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

// User defines a user in db
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUser instanciates a user object
func NewUser(username string, password string) (*User, error) {

	// hashedPwd := sha256.Sum256([]byte(password))

	return &User{
		Username: username,
		Password: password,
	}, nil

}

// SignUp user into db
func (myUser *User) SignUp() error {
	hashedPwd := sha256.Sum256([]byte(myUser.Password))
	if err := kvdb.WriteData(dbUsers, "users", myUser.Username, hex.EncodeToString(hashedPwd[:])); err != nil {
		return err
	}

	return nil
}

// Authenticate allows to auth a user
func (myUser *User) Authenticate() (bool, error) {

	actualHashedPwd, err := kvdb.ReadData(dbUsers, "users", myUser.Username)

	if err != nil {
		return false, errors.New("Error occured while looking inside the db")
	}

	if actualHashedPwd == "" {
		return false, errors.New("Couldn't find user " + myUser.Username)
	}
	hashedPwd := sha256.Sum256([]byte(myUser.Password))

	if actualHashedPwd == hex.EncodeToString(hashedPwd[:]) {
		return true, nil
	}

	return false, nil

}

// UserExist check if username exist in db
func (myUser *User) UserExist() (bool, error) {

	data, err := kvdb.ReadData(dbUsers, "users", myUser.Username)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if data != "" {
		return true, nil
	}

	return false, nil
}
