package user

import (
	"github.com/applinh/kaepora/lib/db"
)

func GetUsernamesFromDB() ([]string, error) {
	userKeys, err := db.ReadAllKeys()
	if err != nil {
		return nil, err
	}
	var usersList = make([]string, 0)

	for _, user := range userKeys {

		usersList = append(usersList, user.Username)
	}

	return usersList, nil

}
