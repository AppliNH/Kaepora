package main

import (
	"fmt"
	"log"
	"net/http"
	routesuser "primitivofr/kaepora/routes/users"
	"primitivofr/kaepora/services/user"

	"github.com/gorilla/mux"
)

var dbName = "kaeporadb"
var usersBucket = "users"

// Rewriter is..
func Rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
	})
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user.ReadAllFromDB()
	})
	r.HandleFunc("/users/signup", routesuser.SignUp).Methods("POST")
	r.HandleFunc("/users/signin", routesuser.SignIn).Methods("POST")

	log.Println("Runnin on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))

	// res, err := user.Authenticate(dbUsers, "toto", "bibo")

}
