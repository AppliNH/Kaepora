package routesusers

import (
	"encoding/json"
	"log"
	"net/http"
	utilserrors "primitivofr/kaepora/utils/errors"

	user "primitivofr/kaepora/services/user"
)

// SignUp is the route handler to sign a user up
func SignUp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusBadRequest, "User already exist with this username")
		return

	}

	myUser, _ := user.NewUser(data["username"].(string), data["password"].(string))

	exist, err := myUser.UserExist()
	if err != nil {

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exist {
		utilserrors.SendHTTPError(w, http.StatusBadRequest, "User already exist with this username")
		return
	}

	if err := myUser.SignUp(); err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while signing user up")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
