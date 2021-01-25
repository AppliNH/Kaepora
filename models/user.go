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
	dbUsers, err = kvdb.InitDB("kaeporadb", "users")
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

	hashedPwd := sha256.Sum256([]byte(password))
	return &User{
		Username: username,
		Password: hex.EncodeToString(hashedPwd[:]),
	}, nil

}

// SignUp user into db
func (myUser *User) SignUp() error {

	if err := kvdb.WriteData(dbUsers, "users", myUser.Username, myUser.Password); err != nil {
		return err
	}

	return nil
}

// Authenticate allows to auth a user
func (myUser *User) Authenticate() (bool, error) {

	actualHashedPwd, err := kvdb.ReadData(dbUsers, "users", myUser.Username)

	if err != nil || actualHashedPwd == "" {
		return false, errors.New("Couldn't find user " + myUser.Username)
	}

	if actualHashedPwd == myUser.Password {
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
