package auth

import (
	"errors"
	"log"

	"github.com/applinh/kaepora/lib/auth"
)

// SignUp is the route handler to sign a user up
func SignUp(username string, password string) (map[string]string, error) {

	myUser, err := auth.NewUser(username, password)
	if err != nil {
		return nil, err
	}

	exist, err := myUser.UserExist()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if exist {
		return nil, errors.New("User already exist with this username")
	}

	if err := myUser.SignUp(); err != nil {
		log.Println(err)
		return nil, errors.New("Internal error occured while signing user up : " + err.Error())
	}

	keys, err := myUser.NewUserKeys()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Internal error occured while generating keys : " + err.Error())
	}

	if err := keys.SaveToDB(); err != nil {
		log.Println(err)
		return nil, errors.New("Internal error occured while saving encrypted keys to db : " + err.Error())
	}

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
