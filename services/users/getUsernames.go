package users

import (
	"log"

	"github.com/applinh/kaepora/lib/user"
)

// GetUsernames returns all usernames of the db
func GetUsernames() (map[string]interface{}, error) {

	res, err := user.GetUsernamesFromDB()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return map[string]interface{}{
		"users": res,
	}, nil

}
