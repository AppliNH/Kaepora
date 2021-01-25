package routesusers

import (
	"encoding/json"
	"log"
	"net/http"

	user "primitivofr/kaepora/services/user"
	utilserrors "primitivofr/kaepora/utils/errors"
)

// SignIn is the route handler to sign a user up
func SignIn(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while reading incoming request")
		return
	}

	myUser, _ := user.NewUser(data["username"].(string), data["password"].(string))

	isAuth, err := myUser.Authenticate()
	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while authenticating user")
		return
	}

	if isAuth {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
