package users

import (
	"encoding/json"
	"log"
	"net/http"

	usersService "github.com/applinh/kaepora/services/users"
	utilserrors "github.com/applinh/kaepora/utils/errors"

	"github.com/gorilla/mux"
)

// GetUserPubKey returns pub key the requested user
func GetUserPubKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]
	res, err := usersService.GetUserPubKey(username)

	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, err.Error())
	}

	responseJSON, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
