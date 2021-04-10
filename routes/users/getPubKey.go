package users

import (
	"crypto/x509"
	"encoding/json"
	"log"
	"net/http"

	"primitivofr/kaepora/services/auth"
	utilserrors "primitivofr/kaepora/utils/errors"

	"github.com/gorilla/mux"
)

// SignIn is the route handler to sign a user up
func GetUserPubKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]

	targetUser, err := auth.NewUser(username, "")

	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while reading incoming request")
	}

	pubKey, err := targetUser.GetPublicKey()
	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusBadRequest, "No data found for : "+username)
	}

	marshalledPublicKey := x509.MarshalPKCS1PublicKey(pubKey)

	response := map[string]interface{}{
		"publicKey": marshalledPublicKey,
	}

	responseJSON, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)

}
