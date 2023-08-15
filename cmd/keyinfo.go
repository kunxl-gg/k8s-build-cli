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

// keyinfoCmd represents the keyinfo command
var keyinfoCmd = &cobra.Command{
	Use:   "keyinfo",
	Short: "A brief description of your command",
	Long: `Read information about the signing keys for the specified project`,

	Run: func(cmd *cobra.Command, args []string) {
		
		// declaring the base url
		base_url := config.BaseUrl
		project_name := "home:" + config.Username

		// the base url
		url := base_url + "source/" + project_name + "/_keyinfo"

		// making a request 
		resp := helpers.APICall("GET", url)

		// printing the response status
		fmt.Println("HTTP response status: ", resp.Status)

		// printing the response body
		body, err := ioutil.ReadAll(resp.Body)
		helpers.CheckError(err)

		log.Default().Println(string(body))

	},
}

func init() {
	rootCmd.AddCommand(keyinfoCmd)
}
