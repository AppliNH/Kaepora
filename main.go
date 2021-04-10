package main

import (
	"fmt"
	"log"
	"net/http"
	routesauth "primitivofr/kaepora/routes/auth"
	routesusers "primitivofr/kaepora/routes/users"

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

	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	auth.ReadAllFromDB()
	// })

	r.HandleFunc("/auth/signup", routesauth.SignUp).Methods("POST")
	r.HandleFunc("/auth/signin", routesauth.SignIn).Methods("POST")

	r.HandleFunc("/users/{username}", routesusers.GetUserPubKey).Methods("GET")
	r.HandleFunc("/users", routesusers.GetUsernames).Methods("GET")

	log.Println("Runnin on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))

	// res, err := user.Authenticate(dbUsers, "toto", "bibo")

}
