/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/kunxl-gg/lfx-lezgooo/config"
	"github.com/kunxl-gg/lfx-lezgooo/helpers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `The command to list all the packages in a given project`,
	
	Run: func(cmd *cobra.Command, args []string) {

		// initial variables
		base_url := config.BaseUrl
		project_name := "home:kunxl.sol"

		// the complete url
		url := base_url + "source/" + project_name + "?deleted=0&expand=0"

		// making a get request
		resp := helpers.APICall("GET", url)

		// Status code of the response
		fmt.Println(resp.Status, http.StatusText(resp.StatusCode))

		// read the response
		body, err := ioutil.ReadAll(resp.Body)
		helpers.CheckError(err)

		log.Default().Println(string(body))

		// close the request call
		defer resp.Body.Close()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
