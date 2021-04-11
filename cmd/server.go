/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	routesauth "github.com/applinh/kaepora/routes/auth"
	routesusers "github.com/applinh/kaepora/routes/users"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

var dbName = "kaeporadb"
var usersBucket = "users"

// Rewriter is..
func Rewriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
	})
}

func StartServer() {
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
