package models

import "errors"

// User defines a user in db
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUser instanciates a user object
func NewUser(username string, password string) (*User, error) {

	// hashedPwd := sha256.Sum256([]byte(password))
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("missing username or password")
	}
	return &User{
		Username: username,
		Password: password,
	}, nil

}
