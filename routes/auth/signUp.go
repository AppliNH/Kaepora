package auth

import (
	"crypto/x509"
	"encoding/json"
	"log"
	"net/http"
	"primitivofr/kaepora/services/auth"
	utilserrors "primitivofr/kaepora/utils/errors"
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

	myUser, _ := auth.NewUser(data["username"].(string), data["password"].(string))

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

	keys, err := myUser.NewUserKeys()
	if err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while generating keys")
	}

	if err := keys.SaveToDB(); err != nil {
		log.Println(err)
		utilserrors.SendHTTPError(w, http.StatusInternalServerError, "Internal error occured while saving encrypted keys to db")
	}

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

	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)

}
