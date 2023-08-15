/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/kunxl-gg/lfx-lezgooo/config"
	"github.com/kunxl-gg/lfx-lezgooo/helpers"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// initailizing variables
		base_url := config.BaseUrl
		project_name := "home:" + config.Username
		package_name := "hello-world"
		
		// the final url
		url := base_url + "source/" + project_name + "/" + package_name + "?cmd=commit"

		resp := helpers.APICall("POST", url)

		// printing the response
		body, err := ioutil.ReadAll(resp.Body)
		helpers.CheckError(err)

		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

}
