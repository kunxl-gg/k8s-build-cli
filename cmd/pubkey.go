/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kunxl-gg/lfx-lezgooo/config"
	"github.com/kunxl-gg/lfx-lezgooo/helpers"
	"github.com/spf13/cobra"
)

// pubkeyCmd represents the pubkey command
var pubkeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "A brief description of your command",
	Long: `Get the project's GPG key. If the project doesn't have a default GPG key, it looks for the first one available in the namespace hierarchy, ending at the global buildservice keỷ`,
	Run: func(cmd *cobra.Command, args []string) {

			// declaring initial variables
			base_url := config.BaseUrl
			project_name := "home:" + config.Username
	
			// final url
			url := base_url + "source/" + project_name + "/_pubkey"
	
			// making the api call
			resp := helpers.APICall("GET", url)
	
			// reading the body
			body, err := ioutil.ReadAll(resp.Body)
			helpers.CheckError(err)
	
			// printing the response from the body
			fmt.Println(string(body))

	},
}

func deletePubKey() {
	log.Println("hello this is woring")
}

func init() {
	rootCmd.AddCommand(pubkeyCmd)
}
