package auth

import (
	"crypto/x509"
	"encoding/json"
	"log"
	"net/http"

	auth "primitivofr/kaepora/services/auth"
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

	myUser, _ := auth.NewUser(data["username"].(string), data["password"].(string))

	isAuth, err := myUser.Authenticate()
	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while authenticating user")
		return
	}

	if isAuth {
		pubKey, err := myUser.GetPublicKey()
		if err != nil {
			log.Println(err)
		}
		privKey, err := myUser.GetPrivateKey()
		if err != nil {
			log.Println(err)
		}

		marshalledPrivKey := x509.MarshalPKCS1PrivateKey(privKey)
		marshalledPublicKey := x509.MarshalPKCS1PublicKey(pubKey)

		response := map[string]interface{}{
			"publicKey":  marshalledPublicKey,
			"privateKey": marshalledPrivKey,
		}

		responseJSON, _ := json.Marshal(response)

		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
