package users

import (
	"crypto/x509"
	"errors"
	"log"

	"github.com/applinh/kaepora/lib/auth"
)

// SignIn is the route handler to sign a user up
func GetUserPubKey(username string) (map[string]interface{}, error) {
	targetUser, err := auth.NewUser(username, "pass")

	if err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())

	}

	pubKey, err := targetUser.GetPublicKey()
	if err != nil {
		log.Println(err)
		return nil, errors.New("No data found for : " + username)
	}

	marshalledPublicKey := x509.MarshalPKCS1PublicKey(pubKey)

	return map[string]interface{}{
		"publicKey": marshalledPublicKey,
	}, nil

}
