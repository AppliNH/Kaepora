package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"

	"primitivofr/kaepora/models"
	"primitivofr/kaepora/services/db"
)

type User struct {
	*models.User
}

// NewUser instanciates a user object
func NewUser(username string, password string) (*User, error) {

	// hashedPwd := sha256.Sum256([]byte(password))

	user, err := models.NewUser(username, password)
	if err != nil {
		return nil, err
	}

	return &User{user}, nil

}

// SignUp user into db
func (myUser *User) SignUp() error {

	hashedPwd := sha256.Sum256([]byte(myUser.Password))
	dbUsers, err := db.NewUsersDb()
	if err != nil {
		return err
	}

	if err := dbUsers.Write(myUser.Username, hex.EncodeToString(hashedPwd[:])); err != nil {
		return err
	}

	return nil
}

// Authenticate allows to auth a user
func (myUser *User) Authenticate() (bool, error) {

	dbUsers, err := db.NewUsersDb()
	if err != nil {
		return false, err
	}

	actualHashedPwd, err := dbUsers.Read(myUser.Username)

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
	dbUsers, err := db.NewUsersDb()
	if err != nil {
		return false, err
	}

	data, err := dbUsers.Read(myUser.Username)

	if err != nil {
		log.Println(err)
		return false, err
	}

	if data != "" {
		return true, nil
	}

	return false, nil
}
