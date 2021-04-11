package auth

import (
	"errors"
	"log"

	auth "github.com/applinh/kaepora/lib/auth"
	"github.com/applinh/kaepora/models"
)

// SignIn sign a user in and sends back the keys
func SignIn(username string, password string) (*models.PlainKeyPair, error) {

	myUser, _ := auth.NewUser(username, password)

	isAuth, err := myUser.Authenticate()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if isAuth {

		return myUser.GeneratePlainKeyPair()

	}

	return nil, errors.New("invalid auth")

}
