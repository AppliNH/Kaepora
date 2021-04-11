package users

import (
	"encoding/json"
	"log"
	"net/http"

	usersService "github.com/applinh/kaepora/services/users"
	utilserrors "github.com/applinh/kaepora/utils/errors"
)

// SignIn is the route handler to sign a user up
func GetUsernames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := usersService.GetUsernames()

	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while reading DB : "+err.Error())
	}

	responseJSON, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
