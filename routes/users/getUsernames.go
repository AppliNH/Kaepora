package users

import (
	"encoding/json"
	"log"
	"net/http"
	"primitivofr/kaepora/services/user"
	utilserrors "primitivofr/kaepora/utils/errors"
)

// SignIn is the route handler to sign a user up
func GetUsernames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := user.GetUsernamesFromDB()

	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while reading DB")
	}

	response := map[string]interface{}{
		"users": res,
	}

	responseJSON, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
