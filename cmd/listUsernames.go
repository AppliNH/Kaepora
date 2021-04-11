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
	"encoding/json"
	"fmt"
	"log"

	usersService "github.com/applinh/kaepora/services/users"
	"github.com/spf13/cobra"
)

// listUsernamesCmd represents the listUsernames command
var listUsernamesCmd = &cobra.Command{
	Use:   "listUsernames",
	Short: "List all usernames in the db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := usersService.GetUsernames()
		if err != nil {
			log.Fatal(err)
		} else {
			r, _ := json.Marshal(res)
			fmt.Println(string(r))
		}
	},
}

func init() {
	rootCmd.AddCommand(listUsernamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listUsernamesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listUsernamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
