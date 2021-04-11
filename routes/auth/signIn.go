package auth

import (
	"encoding/json"
	"log"
	"net/http"

	authService "github.com/applinh/kaepora/services/auth"
	utilserrors "github.com/applinh/kaepora/utils/errors"
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

	res, err := authService.SignIn(data["username"].(string), data["password"].(string))
	if err != nil {
		log.Println(err)
		// TODO: Create error models and adapt the http return code
		utilserrors.SendHTTPError(w, http.StatusUnauthorized, "Error occured while authenticating user : "+err.Error())
		return
	}

	responseJSON, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
