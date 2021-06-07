package auth

import (
	"errors"
	"log"

	auth "github.com/applinh/kaepora/lib/auth"
)

// SignIn sign a user in and sends back the keys
func SignIn(username string, password string) (map[string]string, error) {

	myUser, _ := auth.NewUser(username, password)

	isAuth, err := myUser.Authenticate()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if isAuth {
		tokenString, err := auth.GenerateJWT(username)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return map[string]string{
			"token": tokenString,
		}, nil
		//return myUser.GeneratePlainKeyPair()

	}

	return nil, errors.New("invalid auth")

}
